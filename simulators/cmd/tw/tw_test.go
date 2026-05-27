package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func resetState() {
	timerListMu.Lock()
	for id, v := range timerList {
		if v != nil {
			if v.ticker != nil {
				v.ticker.Stop()
			}
			if v.tmr != nil {
				v.tmr.Stop()
			}
			v.cancel()
		}
		delete(timerList, id)
	}
	timerListMu.Unlock()
	atomic.StoreUint64(&activeTimers, 0)
	httpClient = &http.Client{Transport: &mockTransport{}}
}

func createOneshot(id string, ms int64) timerData {
	now := time.Now()
	timeout := now.Add(time.Duration(ms) * time.Millisecond)

	return timerData{
		ID:        id,
		ExpireURI: "http://mock-callback",
		Type:      oneshotTimerType,
		Timeout: timerTime{
			Epoch: timeout.Unix(),
			Ms:    int64(timeout.Nanosecond() / 1000000),
		},
	}
}

func createPeriodic(id string, ms int32) timerData {
	return timerData{
		ID:        id,
		ExpireURI: "http://mock-callback",
		Type:      periodicTimerType,
		Duration:  ms,
	}
}

func TestOneShotTimer(t *testing.T) {
	resetState()
	callbackReceived := make(chan bool, 1)

	httpClient = &http.Client{
		Transport: &mockTransport{
			callback: func(_ *http.Request) {
				callbackReceived <- true
			},
		},
	}

	tw := &TwSim{}
	timerInfo := createOneshot("oneshot-1", 10)

	payloadBytes, _ := json.Marshal(timerInfo)
	req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	w := httptest.NewRecorder()

	tw.handleTimers(w, req)

	assert.Equal(t, 201, w.Code)

	select {
	case <-callbackReceived:
	case <-time.After(50 * time.Millisecond):
		t.Error("Callback not received within timeout")
	}

	select {
	case <-callbackReceived:
		t.Error("Oneshot should only fire once")
	case <-time.After(50 * time.Millisecond):
	}
}

func TestInvalidTimerType(t *testing.T) {
	resetState()

	tw := &TwSim{}
	timerInfo := createOneshot("invalid-type-timer", 10)
	timerInfo.Type = "invalidtype"

	payloadBytes, _ := json.Marshal(timerInfo)
	req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	w := httptest.NewRecorder()

	tw.handleTimers(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestOneShotTimerHitImmediately(t *testing.T) {
	resetState()
	callbackReceived := make(chan bool, 1)

	httpClient = &http.Client{
		Transport: &mockTransport{
			callback: func(_ *http.Request) {
				callbackReceived <- true
			},
		},
	}

	tw := &TwSim{}
	timerInfo := createOneshot("oneshot-immediate", 1)

	time.Sleep(10 * time.Millisecond)
	payloadBytes, _ := json.Marshal(timerInfo)
	req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	w := httptest.NewRecorder()

	tw.handleTimers(w, req)

	assert.Equal(t, 201, w.Code)

	select {
	case <-callbackReceived:
	case <-time.After(50 * time.Millisecond):
		t.Error("Callback not received within timeout")
	}
}

func TestOneShotTimerCancel(t *testing.T) {
	resetState()
	callbackReceived := make(chan bool, 1)

	httpClient = &http.Client{
		Transport: &mockTransport{
			callback: func(_ *http.Request) {
				callbackReceived <- true
			},
		},
	}

	tw := &TwSim{}
	timerInfo := createOneshot("oneshot-cancel", 100)

	payloadBytes, _ := json.Marshal(timerInfo)
	req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	w := httptest.NewRecorder()

	tw.handleTimers(w, req)

	require.Equal(t, 201, w.Code)

	var timerResponse timerData
	err := json.Unmarshal(w.Body.Bytes(), &timerResponse)
	require.NoError(t, err)

	uri := fmt.Sprintf("/v1/timers/%s", timerResponse.ID)
	deleteReq := httptest.NewRequest("DELETE", uri, http.NoBody)
	deleteW := httptest.NewRecorder()
	tw.handleTimers(deleteW, deleteReq)

	assert.Equal(t, 200, deleteW.Code)

	select {
	case <-callbackReceived:
		t.Error("Failed to cancel oneshot timer")
	case <-time.After(200 * time.Millisecond):
	}
}

func TestOneShotTimerUpdate(t *testing.T) {
	resetState()
	callbackReceived := make(chan bool, 1)

	httpClient = &http.Client{
		Transport: &mockTransport{
			callback: func(_ *http.Request) {
				callbackReceived <- true
			},
		},
	}

	tw := &TwSim{}
	timerInfo := createOneshot("oneshot-update", 100)

	payloadBytes, _ := json.Marshal(timerInfo)
	req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	w := httptest.NewRecorder()

	tw.handleTimers(w, req)

	require.Equal(t, 201, w.Code)

	var timerResponse timerData
	err := json.Unmarshal(w.Body.Bytes(), &timerResponse)
	require.NoError(t, err)

	timerInfo = createOneshot(timerResponse.ID, 200)
	timerInfo.Update = true
	payloadBytes, _ = json.Marshal(timerInfo)

	updateReq := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	updateW := httptest.NewRecorder()
	tw.handleTimers(updateW, updateReq)

	assert.Equal(t, 200, updateW.Code)

	select {
	case <-callbackReceived:
		t.Error("Timer should have been updated to hit after 200ms instead of 100ms")
	case <-time.After(100 * time.Millisecond):
	}

	select {
	case <-callbackReceived:
	case <-time.After(200 * time.Millisecond):
		t.Error("Timer should have hit after 200ms")
	}
}

func TestPeriodicTimer(t *testing.T) {
	resetState()
	var callbackCount atomic.Int32
	callbackReceived := make(chan bool, 2)

	httpClient = &http.Client{
		Transport: &mockTransport{
			callback: func(_ *http.Request) {
				callbackCount.Add(1)
				callbackReceived <- true
			},
		},
	}

	tw := &TwSim{}
	timerInfo := timerData{
		ID:        "periodic-1",
		ExpireURI: "http://mock-callback",
		Type:      periodicTimerType,
		Duration:  10,
	}

	payloadBytes, _ := json.Marshal(timerInfo)
	req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	w := httptest.NewRecorder()

	tw.handleTimers(w, req)

	require.Equal(t, 201, w.Code)

	var timerResponse timerData
	err := json.Unmarshal(w.Body.Bytes(), &timerResponse)
	require.NoError(t, err)

	for i := 0; i < 2; i++ {
		select {
		case <-callbackReceived:
		case <-time.After(1 * time.Second):
			t.Fatalf("Callback %d not received within timeout", i+1)
		}
	}

	assert.Equal(t, int32(2), callbackCount.Load())

	uri := fmt.Sprintf("/v1/timers/%s", timerResponse.ID)
	deleteReq := httptest.NewRequest("DELETE", uri, http.NoBody)
	deleteW := httptest.NewRecorder()
	tw.handleTimers(deleteW, deleteReq)
}

func TestPeriodicTimerUpdate(t *testing.T) {
	resetState()
	var callbackCount atomic.Int32
	callbackReceived := make(chan bool, 3)

	httpClient = &http.Client{
		Transport: &mockTransport{
			callback: func(_ *http.Request) {
				callbackCount.Add(1)
				callbackReceived <- true
			},
		},
	}

	tw := &TwSim{}
	timerInfo := createPeriodic("periodic-update", 50)

	payloadBytes, _ := json.Marshal(timerInfo)
	req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	w := httptest.NewRecorder()

	tw.handleTimers(w, req)

	require.Equal(t, 201, w.Code)

	var timerResponse timerData
	err := json.Unmarshal(w.Body.Bytes(), &timerResponse)
	require.NoError(t, err)

	select {
	case <-callbackReceived:
	case <-time.After(200 * time.Millisecond):
		t.Fatal("First callback not received")
	}

	timerInfo = createPeriodic(timerResponse.ID, 200)
	timerInfo.Update = true
	payloadBytes, _ = json.Marshal(timerInfo)

	updateReq := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	updateW := httptest.NewRecorder()
	tw.handleTimers(updateW, updateReq)

	assert.Equal(t, 200, updateW.Code)

	select {
	case <-callbackReceived:
		t.Error("Timer fired too early, should use new 200ms interval")
	case <-time.After(100 * time.Millisecond):
	}

	select {
	case <-callbackReceived:
	case <-time.After(200 * time.Millisecond):
		t.Error("Timer should have fired at new 200ms interval")
	}
}

func TestPeriodicTimerUpdateWithEmptyId(t *testing.T) {
	resetState()
	var callbackCount atomic.Int32
	callbackReceived := make(chan bool, 3)

	httpClient = &http.Client{
		Transport: &mockTransport{
			callback: func(_ *http.Request) {
				callbackCount.Add(1)
				callbackReceived <- true
			},
		},
	}

	tw := &TwSim{}
	timerInfo := createPeriodic("periodic-empty-update", 50)

	payloadBytes, _ := json.Marshal(timerInfo)
	req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	w := httptest.NewRecorder()

	tw.handleTimers(w, req)

	require.Equal(t, 201, w.Code)

	var timerResponse timerData
	err := json.Unmarshal(w.Body.Bytes(), &timerResponse)
	require.NoError(t, err)

	select {
	case <-callbackReceived:
	case <-time.After(200 * time.Millisecond):
		t.Fatal("First callback not received")
	}

	timerInfo = createPeriodic("", 200)
	//timerInfo.ID = timerResponse.ID //set empty id
	timerInfo.Update = true
	payloadBytes, _ = json.Marshal(timerInfo)

	updateReq := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	updateW := httptest.NewRecorder()
	tw.handleTimers(updateW, updateReq)

	assert.Equal(t, 400, updateW.Code)

	select {
	case <-callbackReceived:
	case <-time.After(100 * time.Millisecond):
		t.Error("Timer should have fired at the 50ms interval since update should fail")
	}
}

func TestPeriodicTimerCreateExistingTimer(t *testing.T) {
	resetState()
	var callbackCount atomic.Int32
	callbackReceived := make(chan bool, 3)

	httpClient = &http.Client{
		Transport: &mockTransport{
			callback: func(_ *http.Request) {
				callbackCount.Add(1)
				callbackReceived <- true
			},
		},
	}

	tw := &TwSim{}
	timerInfo := createPeriodic("periodic-existing", 50)

	payloadBytes, _ := json.Marshal(timerInfo)
	req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	w := httptest.NewRecorder()

	tw.handleTimers(w, req)

	require.Equal(t, 201, w.Code)

	var timerResponse timerData
	err := json.Unmarshal(w.Body.Bytes(), &timerResponse)
	require.NoError(t, err)

	select {
	case <-callbackReceived:
	case <-time.After(200 * time.Millisecond):
		t.Fatal("First callback not received")
	}

	timerInfo = createPeriodic(timerResponse.ID, 200)
	payloadBytes, _ = json.Marshal(timerInfo)

	updateReq := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	updateW := httptest.NewRecorder()
	tw.handleTimers(updateW, updateReq)

	assert.Equal(t, 302, updateW.Code)

	select {
	case <-callbackReceived:
	case <-time.After(100 * time.Millisecond):
		t.Error(`Timer should have fired at the 50ms interval
			since no new or updated timer should have been created with a longer timeout.`)
	}
}

func TestPeriodicTimerCancel(t *testing.T) {
	resetState()
	var callbackCount atomic.Int32
	callbackReceived := make(chan bool, 2)

	httpClient = &http.Client{
		Transport: &mockTransport{
			callback: func(_ *http.Request) {
				callbackCount.Add(1)
				callbackReceived <- true
			},
		},
	}

	tw := &TwSim{}
	timerInfo := timerData{
		ID:        "periodic-cancel",
		ExpireURI: "http://mock-callback",
		Type:      periodicTimerType,
		Duration:  100,
	}

	payloadBytes, _ := json.Marshal(timerInfo)
	req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	w := httptest.NewRecorder()

	tw.handleTimers(w, req)

	require.Equal(t, 201, w.Code)

	select {
	case <-callbackReceived:
	case <-time.After(200 * time.Millisecond):
		t.Fatal("First callback not received")
	}

	var timerResponse timerData
	err := json.Unmarshal(w.Body.Bytes(), &timerResponse)
	require.NoError(t, err)

	uri := fmt.Sprintf("/v1/timers/%s", timerResponse.ID)
	deleteReq := httptest.NewRequest("DELETE", uri, http.NoBody)
	deleteW := httptest.NewRecorder()
	tw.handleTimers(deleteW, deleteReq)

	select {
	case <-callbackReceived:
		t.Errorf("Unexpected second callback after cancellation. Expected 1 callback, got %d", callbackCount.Load())
	case <-time.After(200 * time.Millisecond):
	}
}

func TestPeriodicTimerInvalidDuration(t *testing.T) {
	resetState()
	tw := &TwSim{}
	timerInfo := createPeriodic("periodic-invalid-dur", 0)

	payloadBytes, _ := json.Marshal(timerInfo)
	req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	w := httptest.NewRecorder()

	tw.handleTimers(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestUpdatePeriodicTimerInvalidDuration(t *testing.T) {
	resetState()
	httpClient = &http.Client{Transport: &mockTransport{}}

	tw := &TwSim{}
	timerInfo := createPeriodic("periodic-update-invalid", 50)

	payloadBytes, _ := json.Marshal(timerInfo)
	req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	w := httptest.NewRecorder()

	tw.handleTimers(w, req)

	require.Equal(t, 201, w.Code)

	var timerResponse timerData
	json.Unmarshal(w.Body.Bytes(), &timerResponse)

	timerInfo = createPeriodic(timerResponse.ID, -10)
	timerInfo.Update = true
	payloadBytes, _ = json.Marshal(timerInfo)

	updateReq := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
	updateW := httptest.NewRecorder()
	tw.handleTimers(updateW, updateReq)

	assert.Equal(t, 400, updateW.Code)
}

func TestDeleteNonExistentTimer(t *testing.T) {
	resetState()
	tw := &TwSim{}

	deleteReq := httptest.NewRequest("DELETE", "/v1/timers/nonexistent-id", http.NoBody)
	w := httptest.NewRecorder()
	tw.handleTimers(w, deleteReq)

	assert.Equal(t, 200, w.Code)
}

func TestConcurrentTimerCreation(t *testing.T) {
	resetState()
	httpClient = &http.Client{Transport: &mockTransport{}}

	tw := &TwSim{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			timerInfo := createPeriodic(fmt.Sprintf("timer-%d", id), 100)

			payloadBytes, _ := json.Marshal(timerInfo)
			req := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes)))
			w := httptest.NewRecorder()

			tw.handleTimers(w, req)

			assert.Equal(t, 201, w.Code)
		}(i)
	}

	wg.Wait()

	timerListMu.RLock()
	count := len(timerList)
	timerListMu.RUnlock()

	assert.Equal(t, 10, count)
}

func TestMultiplePeriodicTimers(t *testing.T) {
	resetState()
	timer1Callbacks := 0
	timer2Callbacks := 0
	callbackReceived := make(chan string, 10)

	httpClient = &http.Client{
		Transport: &mockTransport{
			callback: func(req *http.Request) {
				callbackReceived <- req.URL.String()
			},
		},
	}

	tw := &TwSim{}

	timerInfo1 := createPeriodic("timer-1", 50)
	timerInfo1.ExpireURI = "http://callback-1"

	payloadBytes1, _ := json.Marshal(timerInfo1)
	req1 := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes1)))
	w1 := httptest.NewRecorder()
	tw.handleTimers(w1, req1)

	timerInfo2 := createPeriodic("timer-2", 50)
	timerInfo2.ExpireURI = "http://callback-2"

	payloadBytes2, _ := json.Marshal(timerInfo2)
	req2 := httptest.NewRequest("POST", "/v1/timers", strings.NewReader(string(payloadBytes2)))
	w2 := httptest.NewRecorder()
	tw.handleTimers(w2, req2)

	func() {
		timeout := time.After(200 * time.Millisecond)
		for {
			select {
			case url := <-callbackReceived:
				switch url {
				case "http://callback-1":
					timer1Callbacks++
				case "http://callback-2":
					timer2Callbacks++
				}
			case <-timeout:
				return
			}
		}
	}()

	assert.NotZero(t, timer1Callbacks, "Timer 1 did not fire")
	assert.NotZero(t, timer2Callbacks, "Timer 2 did not fire")
}

func TestHTTPMethodNotAllowed(t *testing.T) {
	resetState()
	tw := &TwSim{}

	req := httptest.NewRequest("GET", "/v1/timers", http.NoBody)
	w := httptest.NewRecorder()
	tw.handleTimers(w, req)

	assert.Equal(t, 405, w.Code)
}

func TestInvalidDeletePath(t *testing.T) {
	resetState()
	tw := &TwSim{}

	req := httptest.NewRequest("DELETE", "/v1/timers/", http.NoBody)
	w := httptest.NewRecorder()
	tw.handleTimers(w, req)

	assert.Equal(t, 400, w.Code)
}

type mockTransport struct {
	callback func(*http.Request)
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.callback != nil {
		m.callback(req)
	}

	return &http.Response{
		StatusCode: 200,
		Body:       http.NoBody,
		Header:     make(http.Header),
	}, nil
}
