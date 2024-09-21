package main
import "fmt"
func addNums[T int |float64]  (a ,b T) T{
	return a+b
}
func main(){

	fmt.Println(addNums(4,9.5))

}
