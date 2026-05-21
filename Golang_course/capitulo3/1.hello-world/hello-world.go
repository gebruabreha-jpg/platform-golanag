// -*- coding: utf-8 -*-
// hello-world.go
// -----------------------------------------------------------------------------
//
// Started on <mar 17-02-2026 00:11:57.040987298 (1771283517)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Hola mundo, pero a lo grande
package main

import (
	"fmt"
	"sync"
)

func greetings() []string {
	var wg sync.WaitGroup
	wg.Add(2)
	results := make([]string, 2)

	go func() {
		defer wg.Done()
		results[0] = "Hello World!"
	}()
	go func() {
		defer wg.Done()
		results[1] = "Hallo Welt!"
	}()
	wg.Wait()
	return results
}

func main() {
	for _, g := range greetings() {
		fmt.Println(g)
	}
}

// Local Variables:
// mode:go
// fill-column:80
// End:
