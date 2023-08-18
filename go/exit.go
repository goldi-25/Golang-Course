package main 
import ("os"
"fmt")
func main (){
	defer fmt.Println("hellon")
	os.Exit(5)
}