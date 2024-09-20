package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
	

}
func zeroPtr(iptr *int)  {
	*iptr = 0 
	
}
func main() {
	a := 1
	zeroVal(a)
	fmt.Println(a)
	zeroPtr(&a)
    fmt.Println(a)
}
