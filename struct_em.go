// package main

// import "fmt"

// type base struct {
// 	num int
// }

// func (b base) describe() string {
// 	return fmt.Sprintf("base with num=%v", b.num)
// }

// type container struct {
// 	base
// 	str string
// }

// func main() {
// 	co := container{
// 		base: base{
// 			num: 1,
// 		},
// 		str: "some name",
// 	}

// 	fmt.Printf("Co={num: %v,str: %v}",co.num,co.str)
// 	fmt.Println("also num",co.base.num)
// 	fmt.Println("describe",co.describe())  // explicit access of the field  embedded
// 	type describer interface{
// 		describe() string
// 	}

// 	var d describer=co
// 	var co_emp describer=co
// 	fmt.Println("describe emplementation by container",co_emp)
// 	fmt.Println("describer: ",d.describe())

	
// }
