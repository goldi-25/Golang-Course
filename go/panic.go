package main 
import("fmt"
"os")
func main (){
	
	// we create panic function 
	fmt.Println("start the excution" )
	panic("something goes bas ")
	fmt.Println("end the excution" )

	// example of panic in go created 
	a,b := 1,0
	ans:= a / b 
	fmt.Println(ans) 

	// error created 
	_, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}


