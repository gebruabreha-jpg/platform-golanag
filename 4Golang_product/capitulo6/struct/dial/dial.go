// -*- coding: utf-8 -*-
// dial.go
// -----------------------------------------------------------------------------
//
// Started on <mar 24-02-2026 06:32:28.638757516 (1771911148)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Sorting chemical elements with Dial's algorithm
package dial

import "fmt"

// Definition of the maximum ordinal
const B = 1000

// types
type Átomo struct {
	Nombre string
	UMA    float64
}

// Initialization of a bucket with B+1 entries in the range [0, B+1]
func initBucket() (bucket [][]Átomo) {

	// allocation of B+1 entries, and initialization of each bucket
	bucket = make([][]Átomo, 1+B)
	for i := 0; i <= B; i++ {
		bucket[i] = []Átomo{}
	}
	return
}

// Add the given atom to the i-th bucket according to its UMA
func insert(átomo Átomo, bucket [][]Átomo) error {

	// if this item goes beyond the capacity of the bucket then exit with an
	// error
	if int(átomo.UMA) >= len(bucket) {
		return fmt.Errorf("item out of bounds: %v > %v", átomo.UMA, B)
	}

	// otherwise insert it and exit with no error. Note that all UMAs are
	// expected to be different, no need to sort atoms within the same bucket
	i := int(átomo.UMA)
	bucket[i] = append(bucket[i], átomo)
	return nil
}

// Return all items in the bucket in ascending order just by traversing all
// buckets from 0 to B
func pop(bucket [][]Átomo) []Átomo {

	// initialize the result
	result := []Átomo{}

	for i := 0; i <= B; i++ {

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
func Sort(items []Átomo) error {

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
