// package main

// import "fmt"
// type rect struct{
// 	width, height int 
// }
// func (r *rect) area() int{
// 	return r.width * r.height
// }
// func(r rect) perim() int{
// 	r.height=50
// 	return 2*(r.width+r.height)
// }
// func main(){
// 	r:=rect{width: 10,height: 5}
// 	fmt.Println("area:",r.area())
// 	fmt.Println("perim",r.perim())
// 	rp:=&r
// 	rp.height=100
// 	rp.width=50
// 	fmt.Println("area",rp.area())
// 	fmt.Println("area",rp.perim())
// 	// functions automatically dereferences the struct object
// }
