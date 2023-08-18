package main 
import "fmt"
func main (){
	var a string = "golis is here"
	fmt.Printf("%+v %T \n", a,a)
	var b string = "efghnfds"
	fmt.Printf("%+v %T \n", a+b,a+b)
	c := []byte(a)
	fmt.Printf("%+v %T \n",c,c)
	// string has UTF-8 



}