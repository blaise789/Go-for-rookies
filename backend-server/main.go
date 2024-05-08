package main
import (
	"fmt"
	"log"
	"net/http"
) 
func formHandler(w http.ResponseWriter,r *http.Request){
	if err:=r.ParseForm(); err !=nil{
		fmt.Fprintf(w,"error while parsing the form")
	}
	fmt.Print("form submitted successfully")
	
	name:=r.FormValue("name")
	email:=r.FormValue("email")
	fmt.Print(email)
	fmt.Fprintf(w,"your name is %s \n",name)
	fmt.Fprintf(w,"your email is %s \n",email)

	return


}
func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path !="/hello"{
	 	http.Error(w,"page not found",http.StatusNotFound)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)

	}
	 fmt.Fprintf(w,"hello")
}
func main(){
	fileServer :=http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)
	fmt.Printf("started server at port 8080")
	if err:=http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal(err)
	}
	
}