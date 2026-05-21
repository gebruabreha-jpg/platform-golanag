package main

import (
	"fmt"
	"math/rand"
)

// randomGenerator generates random numbers and sends them through a channel
// Parameters:
//   - count: number of random numbers to generate
//
// Returns:
//   - <-chan int: read-only channel that receives random integers
func randomGenerator(count int) <-chan int {
	// Create an unbuffered channel for sending integers
	ch := make(chan int)

	// Launch a goroutine to generate and send random numbers
	go func() {
		// Ensure the channel is closed when the goroutine finishes
		defer close(ch)

		// Generate 'count' random numbers
		for i := 0; i < count; i++ {
			// Send random number (0-99) into the channel
			ch <- rand.Intn(100)
		}
	}()

	// Return the channel immediately (goroutine runs concurrently)
	return ch
}

func main() {
	// Create a channel that will receive 25 random numbers
	tunnel := randomGenerator(25)

	// Range over the channel until it's closed
	// This blocks and waits for each value
	for n := range tunnel {
		fmt.Println("Received from tunnel:", n)
	}
	// Loop exits when channel is closed by the goroutine
}
