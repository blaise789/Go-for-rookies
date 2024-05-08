package main

import (
	"fmt"
	// "maps"
	// "time"
	// "math"https://pkg.go.dev/golang.org/x/tools/internal/typesinternal#BrokenImport
	// "constants"
)

// const s string="constants"

func main(){
	


// fmt.Println("Hello world")
// fmt.Println("go"+"lang")
// fmt.Println("1+1",1+1,)
// fmt.Println(true && false)
// fmt.Println(true||false)
// fmt.Println(!true)
// variables 
// var a="initial"
// fmt.Println(a)
// var b,c int=1,2
// fmt.Println(b,c)
// var e int
// fmt.Println(e)
// f :="apple"
// fmt.Println(f)

// fmt.Println(s)
// const n=5000000
// const d=3e20/n
// fmt.Println(int64(d))
// fmt.Println(math.Sin(n))
// i:=1
// for i<=3{
// 	fmt.Println(i)
// 	i=i+1
// }
// for j:=1;j<3;j++{
// 	fmt.Println(j)
// }
// nums:=[]int{2,3,4}
// for i := range nums{
// 	fmt.Println("range", nums[i])
// }
// for{
// 	fmt.Println("loop")
// 	break

// }
// if 7%2==0{
// 	fmt.Println("7 is even")

// } else{
// 	fmt.Println("7 is odd")
	
// }
	
// if 8%2 == 0 || 7%2 == 0 {
// 	fmt.Println("either 8 or 7 are even")
// }
// if num :=9; num<0{
// 	fmt.Println(num,"is negative")
// } else if num<10{
// 	fmt.Println(num,"has 1 digit")
// } else{
// 	fmt.Println(num,"has multiple digits")
// }
// i:=2
// switch i{
// case 1:
// 	fmt.Println("one")

// case 2:
// 	fmt.Println("two")
// case 3:
// 	fmt.Println("three")	
	
// // }
// switch time.Now().Weekday(){
// case time.Saturday,time.Sunday:
// 	fmt.Println(time.Now())
// 	fmt.Println("its the weekend")
// default:
// 	fmt.Println(time.Now())
// 	fmt.Println("it's a weekday")
// }
// t :=time.Now()
// switch {
// case t.Hour()<12:
// 	fmt.Println("it's before noon")
// default:
// 	fmt.Println("it's after noon")	
// }
// whatAmi:=func(i interface{}){
// 	switch t:=i.(type){
		
// 	case bool:
// 		fmt.Println("I'm a bool")
// case int:
// 	fmt.Println("I am an int")
// default:
// 	fmt.Printf("Don't know type %T\n",t)
// 	}

	
// }
// whatAmi(true)
//// arrays
// 
// var s []string;
// fmt.Println(s==nil,len(s)==0)
// s=make([]string,3 )
// fmt.Println(s,len(s))
// s[0]="a"
// s[1]="b"
// s[2]="c"
// s = append(s, "d")
// s = append(s, "e", "f")
// fmt.Println(s)
// fmt.Println(s[0:2])



m:=make(map[string]int)
m["k1"]=7 
m["k2"]=8
m["k3"]=9
m["k4"]=10


fmt.Println("Map:",m)
delete(m,"k2")
clear(m)
fmt.Println("Map:",m)
// n:=map[string] int{"age":12,"year":2024,"height":200,"weight":90}
// if(maps.Equal(n,m)){
// fmt.Println("the maps are equal")
// // fmt.Print(maps)
// }




}
