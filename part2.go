// package main

// import (
// 	"fmt"
// )

// //  closures
// // import "fmt"
// // func counter()func() int{
// //  count:=0

// // 	return func() int {
// // count++
// // return count
// // 	}
// // }
// // func main(){
// // increment:=counter()
// // fmt.Println(increment())

// // fmt.Println(increment())

// // }

// // # recursion #
// // func fact(n int) int {
// // 	// fmt.Println(n)

// // 	if n == 0 {
// // 		return 1

// // 	}
// // 	return n * fact(n-1)

// // }
// // func fib(n int) int {
// // 	if n == 0 {
// // 		return 0
// // 	} else if n == 1 {
// // 		return 1
// // 	}
// // 	return fib(n-2) + fib(n-1)
// // }
// // func main() {
// // 	fmt.Println(fact(6))
// // 	var fib func(n int) int
// // 	fib = func(n int) int {
// // 		if n < 2 {
// // 			return n
// // 		}

// // 		return fib(n-1) + fib(n-2)
// // 	}

// // 	fmt.Println(fib(10))
// // }
// func main(){
// // nums:=[]int {1,2,3,4}
// // var total int
// // // for _,num := range nums{
// // // 	total+=num
// // // }
// // fmt.Println(total)
// // for i,num := range nums{
// // 	fmt.Println(i,num)
// // }
// // var fruits =map[string] string{} 
//  fruits:=map[string] string{"a":"mango","b":"banana","c":"citron"}
//  for _,v:= range  fruits{
// 	fmt.Println(v)
//  }
//  for key:= range fruits{
// 	fmt.Println(key)
//  }
//  for key,val:= range "hello"{

// 	fmt.Println(key,val)
//  }
// }
