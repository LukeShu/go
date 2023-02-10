// errorcheck -0 -m

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p

func f(...*int) {} // ERROR "can inline f$"

func g() {
	defer f()                   // ERROR "func literal does not escape"
	defer f(new(int))           // ERROR "func literal does not escape" "... argument does not escape$" "new\(int\) does not escape$"
	defer f(new(int), new(int)) // ERROR "func literal does not escape" "... argument does not escape$" "new\(int\) does not escape$"

	defer f(nil...)                        // ERROR "func literal does not escape"
	defer f([]*int{}...)                   // ERROR "func literal does not escape" "\[\]\*int{} does not escape$"
	defer f([]*int{new(int)}...)           // ERROR "func literal does not escape" "\[\]\*int{...} does not escape$" "new\(int\) does not escape$"
	defer f([]*int{new(int), new(int)}...) // ERROR "func literal does not escape" "\[\]\*int{...} does not escape$" "new\(int\) does not escape$"

	go f()                   // ERROR "func literal escapes to heap"
	go f(new(int))           // ERROR "func literal escapes to heap" "... argument does not escape$" "new\(int\) escapes to heap$"
	go f(new(int), new(int)) // ERROR "func literal escapes to heap" "... argument does not escape$" "new\(int\) escapes to heap$"

	go f(nil...)                        // ERROR "func literal escapes to heap"
	go f([]*int{}...)                   // ERROR "func literal escapes to heap" "\[\]\*int{} does not escape$"
	go f([]*int{new(int)}...)           // ERROR "func literal escapes to heap" "\[\]\*int{...} does not escape$" "new\(int\) escapes to heap$"
	go f([]*int{new(int), new(int)}...) // ERROR "func literal escapes to heap" "\[\]\*int{...} does not escape$" "new\(int\) escapes to heap$"

	for {
		defer f()                   // ERROR "func literal escapes to heap"
		defer f(new(int))           // ERROR "func literal escapes to heap" "... argument does not escape$" "new\(int\) escapes to heap$"
		defer f(new(int), new(int)) // ERROR "func literal escapes to heap" "... argument does not escape$" "new\(int\) escapes to heap$"

		defer f(nil...)                        // ERROR "func literal escapes to heap"
		defer f([]*int{}...)                   // ERROR "func literal escapes to heap" "\[\]\*int{} does not escape$"
		defer f([]*int{new(int)}...)           // ERROR "func literal escapes to heap" "\[\]\*int{...} does not escape$" "new\(int\) escapes to heap$"
		defer f([]*int{new(int), new(int)}...) // ERROR "func literal escapes to heap" "\[\]\*int{...} does not escape$" "new\(int\) escapes to heap$"

		go f()                   // ERROR "func literal escapes to heap"
		go f(new(int))           // ERROR "func literal escapes to heap" "... argument does not escape$" "new\(int\) escapes to heap$"
		go f(new(int), new(int)) // ERROR "func literal escapes to heap" "... argument does not escape$" "new\(int\) escapes to heap$"

		go f(nil...)                        // ERROR "func literal escapes to heap"
		go f([]*int{}...)                   // ERROR "func literal escapes to heap" "\[\]\*int{} does not escape$"
		go f([]*int{new(int)}...)           // ERROR "func literal escapes to heap" "\[\]\*int{...} does not escape$" "new\(int\) escapes to heap$"
		go f([]*int{new(int), new(int)}...) // ERROR "func literal escapes to heap" "\[\]\*int{...} does not escape$" "new\(int\) escapes to heap$"
	}
}
