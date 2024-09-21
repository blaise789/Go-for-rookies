package main
import "fmt"
func addNums[T int |float64]  (a ,b T) T{
	return a+b


}

func SlicesIndex[S ~[]E,E comparable](s S,v E) int{
	for i := range s{
		if(s[i]==v){
			return i
		}
	}
	return -1

}
type List[T any] struct{
	head,tail *element[T]
}
type element[T any] struct{
  next *element[T]
  val T

}
func (lst *List[T]) push(v T) {
	if  lst.tail==nil{
		lst.head=&element[T]{val: v}
		lst.tail=lst.head

	}else{
		lst.tail.next=&element[T]{val: v}
		lst.tail=lst.tail.next
		
	}
}
func (lst *List[T]) AllElements() []T{
	var elems []T
	for e:=lst.head;e !=nil;e=e.next{
		elems=append(elems, e.val)
	}
	return elems
}
func main(){
	fmt.Println(addNums(4,9.5))

	var s =[] string{"foo","bar","zoo"}
	
	var b= [] int {1,2,3,4}
	fmt.Println(s)
	fmt.Println("index of zoo",SlicesIndex(s,"b"))
	fmt.Println("index of zoo",SlicesIndex(b,3))
_=SlicesIndex[[] string,string](s,"zoo")
lst:=List[int] {}
fmt.Println(lst)
lst.push(10)
lst.push(20)
lst.push(23)
fmt.Println(lst.head)
fmt.Println(lst.tail)
fmt.Println("List:",lst.AllElements())

}

