package main

import "fmt"
type person struct{
	name string 
	age int 
}
func newPerson (name string) *person{
	p:=person{name:name}
	p.age=42
	return &p
}
type student struct{
	studentId int
	studentName string
}

func main(){
	st1:=student{1,"blaise"}
	st1.studentId=2
	fmt.Println(st1)
	st2:=student{studentId:2,studentName:"Karim" }
	st3:=&student{studentId: 3,studentName: "Kevant"}
	st4:=&student{studentId:4}
	st5:=&student{studentName: "blaise"}
	st5.studentId=1
	fmt.Println(st5)
 st6:=student{studentName:"test"}
	fmt.Println(st1,st2,st3,st4,st5,st6)
// 	teacher:=struct{
// 		teachers_name string
// 		teachers_id int
// 	}{
// 		"kevin",
// 		1}
// fmt.Println(teacher)
// fmt.Println(person{"Bob",20})
// fmt.Println(person{name: "Alice", age: 30})

//     fmt.Println(person{name: "Fred"})

//     fmt.Println(&person{name: "Ann", age: 40})

// fmt.Println(newPerson("Jon"))
// s:=person{name: "Sean",age:50}
// sp:=&s
// sp.age=51
// fmt.Println(sp.age)
// fmt.Println(s.name)
// dog:=struct{
// 	name string
// 	isGood bool
// }{
// 	"Rex",
// 	true,
// }
// fmt.Println(dog)
}
