package main

import (
	"fmt"
	// "slices"
)
func main(){
/*	fmt.Println("1+1=",1+1)
	fmt.Print("1+1=",1+1)
    fmt.Println("go"+"lang")
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	var a="initial"
	fmt.Println(a)
	var b,c int=1,2
	fmt.Println(b,c)
	var d=true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	f:="apple"
	fmt.Println(f)
	*/
	// i:=1
	// for i <=3 {
	//     fmt.Println(i)
	// 	i++
	// }
	// for j:=0; j<=3;j++{
	// 	fmt.Println(j)
	// }
	// for i := range 3{
	// 	fmt.Println(i)
	// }
/*	for {
		fmt.Print("loop")
		break
	}
	for n:=range 6{
		if n%2==0{
			continue

		}
		fmt.Println(n)
	}
	if num:=0;num<0{
		fmt.Println("negative number")
	} else if num==0{
		fmt.Println("number is zero")
	}else{
	fmt.Println("Positive number")	
	}
	fmt.Print("hello")
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println(time.Now().Weekday())
	default:
		fmt.Println("It's weekday",time.Now().Weekday())
	}
	t:=time.Now()
	switch{
	case t.Hour()<12:
		fmt.Println("It is before noon")
	
	default:
		fmt.Println("it is afternoon")


	
}  
whatAmi:=func(i interface{}){
	switch t:=i.(type){
	case bool:
	
		fmt.Println(t)
		fmt.Println("I'm a bool")

	case int:
		fmt.Println("I am an int")
    default:
		fmt.Printf("Don't know type %T",t)		
}  
	}
	whatAmi(true)
	arr:=[4] int {1,3,5}
	fmt.Print(arr)
    a:=[...]int { 1,2,5,8}
	fmt.Print(a)
// two dimensions array
var twoD [2][3] int
for i:=0;i<2;i++{
	for j:=0;j<3;j++{
		twoD[i][j]=i+j
	}
}
fmt.Println("2d",twoD)
twoD=[2][3] int{
	{1,2,3},

}
fmt.Print(twoD)

// slices
var s []string
fmt.Println("uninit",s,len(s))
s=make([]string,3)
fmt.Println("emp",s,"len",len(s),"capacity",cap(s))

s[0],s[1],s[2]="a","b","c"
// fmt.Println(s)
// fmt.Println(s[2])
s=append(s, "d")
s=append(s, "e","f")
c:=make([] string,len(s))
copy(c,s)
// fmt.Print(c)
l:=s[2:5]

fmt.Println("sl1",l)
t1:=[] string{"a","b","c"}
t2:=[] string{"a","b","d"}
if slices.Equal(t2,t1){
	fmt.Println("t1 and t2 are equal")
} */
// twoD:=[5][5] int{}
// for i:=0;i<5;i++{
// 	for j:=0;j<=i;j++{
//       twoD[i][j]=i+j
// 	}
// }
// fmt.Print(twoD)
m:=make(map[string] int)
m["k1"]=45
m["k2"]=40
fmt.Println("map:",m)
v1:=m["k1"]
fmt.Println("v1:",v1)
v3:=m["k3"]
fmt.Println(v3)


} 
	
