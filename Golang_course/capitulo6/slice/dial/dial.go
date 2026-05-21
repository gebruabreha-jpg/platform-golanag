// -*- coding: utf-8 -*-
// dial.go
// -----------------------------------------------------------------------------
//
// Started on <jue 19-02-2026 17:08:40.730964721 (1771517320)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Sorting integers with Dial's algorithm
package dial

import "fmt"

// Definition of the maximum ordinal
const upperBound = 1000

// Initialization of a bucket with upperBound+1 entries in the range [0,
// upperBound)
func initBucket() (bucket [][]int) {

	// allocation of upperBound+1 entries, and initialization of each bucket
	bucket = make([][]int, 1+upperBound)
	for i := range upperBound {
		bucket[i] = []int{}
	}
	return
}

// Add number i to the i-th bucket
func insert(i int, bucket [][]int) error {

	// if this item goes beyond the length of the bucket then exit with an
	// error
	if i >= len(bucket) {
		return fmt.Errorf("item out of bounds: %v > %v", i, upperBound)
	}

	// otherwise insert it and exit with no error
	bucket[i] = append(bucket[i], i)
	return nil
}

// Return all items in the bucket in ascending order just by traversing all
// buckets from 0 to upperBound
func pop(bucket [][]int) []int {

	// initialize the result
	result := []int{}

	for i := 0; i <= upperBound; i++ {

		// if this bucket is not empty
		if len(bucket[i]) > 0 {
			result = append(result, bucket[i]...)
		}
	}

	// and return the result
	return result
}

// sort all numbers in the given slice using buckets. It returns nil if no error
// happened
func Sort(items []int) error {

	// first, create a bucket as a slice of slices and initialize it
	bckt := initBucket()

	// now, insert all items in the bucket
	for _, v := range items {
		if err := insert(v, bckt); err != nil {
			return err
		}
	}

	// modify the given slice
	copy(items, pop(bckt))
	return nil
}

// Local Variables:
// mode:go
// fill-column:80
// End:
