# About Simulator

Simulates the Timer Wheel (TW) microservice. It manages periodic and oneshot
timers on behalf of the SUT, sending HTTPS callbacks when timers expire.

## Functionality

### REST API

The simulator exposes an HTTPS server on port **8443**. Only `POST` and
`DELETE` are accepted; any other method returns `405 Method Not Allowed`.

#### POST /v1/timers — Create or update a timer

Request body (JSON):

```json
{
    "id":         "string (required)",
    "type":       "periodic | oneshot",
    "geored":     false,
    "timeout":    { "epoch": 1700000000, "ms": 500 },
    "duration":   5000,
    "expire_uri": "https://sut-host:port/timers",
    "expire_arg": "base64-encoded-string",
    "update":     false
}
```

**Create behaviour (`update: false`):**

| Condition                        | Response         |
| :------------------------------- | :--------------- |
| Timer with given `id` not found  | `201 Created`    |
| Timer with given `id` exists     | `302 Found`      |
| Invalid `type`                   | `400 Bad Request`|
| Missing `id`                     | `400 Bad Request`|
| `periodic` with `duration <= 0`  | `400 Bad Request`|

On `201 Created`, the timer is started in a background goroutine and
`rxQuestionsTw` is incremented by 1.

**Update behaviour (`update: true`):**

| Condition                        | Response                    |
| :------------------------------- | :-------------------------- |
| Timer with given `id` not found  | `404 Not Found`             |
| Timer found but nil (internal)   | `500 Internal Server Error` |
| `periodic` with `duration <= 0`  | `400 Bad Request`           |
| Success                          | `200 OK`                    |

For `oneshot` timers, the underlying `time.Timer` is reset to the new timeout.
For `periodic` timers, the underlying `time.Ticker` is reset to the new
duration.

#### DELETE /v1/timers/{id} — Delete a timer

Cancels and removes the timer identified by `{id}` from the path.

| Condition              | Response         |
| :--------------------- | :--------------- |
| `{id}` missing         | `400 Bad Request`|
| Always otherwise       | `200 OK`         |

The response body is the JSON of the request body as received (may be empty).
If the timer was found and removed, `rxQuestionsTw` is decremented by 1.

### Timer Expiry Callback

When a timer fires, the simulator sends a `POST` request over HTTPS to the URL
specified in `expire_uri`. The request body is the full JSON of the
`timerInfoType` struct as originally received at creation time.

The HTTP client used for callbacks is initialised with mutual TLS, loading
certificates from `emu/tls/server/` and the CA from
`emu/tls/root/cacertbundle.pem`.

### Statistics

`GetStats` returns the following counter:

`activeTimers`: count of active timers. Incremented on `201 Created`,
decremented on a successful `DELETE`.
