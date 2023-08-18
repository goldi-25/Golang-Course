package main 
import "fmt"
func main (){
	var n bool = false 
	fmt.Printf("%v %T \n",n,n)
	var a bool = true 
	fmt.Printf("%v %T \n",a,a)
	 b := 1 == 1 // == is termed as equal operators 
	 c := 2 ==4 
	 fmt.Printf("%+v %T \n", b,b)
	 fmt.Printf("%+v %T\n", c,c)

	 var d bool // if we cant inititase th evalue to variables it result in false values 
	 fmt.Printf("%+v %T\n", d,d)

}