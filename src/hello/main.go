package main

import (
	"fmt"
	"github.com/kamisuzuri/lrn_package/src/tr"
	//"somepackage"
	// "github.com/kamisuzuri/lrn_package/src/internal"
)

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	ff := AddUpper()
	fmt.Println(ff(1))
	fmt.Println(ff(2))
	fmt.Println(ff(3))
	//   disp.SomeFunction()
	//	itnldisp.SomeFunction()
	tr.SomeFunction()

}
