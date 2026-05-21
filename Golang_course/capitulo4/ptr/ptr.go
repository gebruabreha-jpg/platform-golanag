// -*- coding: utf-8 -*-
// ptr.go
// -----------------------------------------------------------------------------
//
// Started on <vie 06-02-2026 11:10:58.743906865 (1770372658)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Description
package main

import "fmt"

func main() {

	var a int = 5
	var cadena string = "sin(Θ)=cos(π/2-θ)"
	fmt.Printf("a %T: %v\n", a, a)
	fmt.Printf("cadena %T: %v\n", cadena, cadena)

	var ar = &a
	var cadenar = &cadena
	fmt.Printf("ar %T: %v\n", ar, ar)
	fmt.Printf("cadenar %T: %v\n", cadenar, cadenar)

	var ad = *ar
	var cadenad = *cadenar
	fmt.Printf("ad %T: %v\n", ad, ad)
	fmt.Printf("cadenad %T: %v\n", cadenad, cadenad)

	var aref *int = &a // Ok
	fmt.Println(aref)  // 0xc000012300

	// var bref *bool = &false   // cannot take address of false
	// var cref *int = a << 3    // cannot use a << 3 as *int value in variable declaration
	// var dref *int = &(a << 3) // cannot take address of a << 3

	var e int = a << 3 // Ok
	var eref *int = &e // Ok
	fmt.Println(eref)  // 0xc000012308

	// var g = &*a // cannot indirect a

	var h = &*eref         // Ok
	fmt.Println(h == eref) // true

	var i = *&e         // Ok
	fmt.Println(i == e) // true
}

// Local Variables:
// mode:go
// fill-column:80
// End:
