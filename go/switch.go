package main 
import ("fmt"
"time" )
func main (){
	goldi := 5
	switch goldi { // we add multiples cases at a time in a oneline 
	case 1 : 
	fmt.Println("goldi 1")
	case 2 : 
	fmt.Println("goldi 2")
	case 3 : 
	fmt.Println("goldi 3")
	case 4 : 
	fmt.Println("goldi 4")
	case 5: 
	fmt.Println("goldi 15")
	default :
   fmt.Println("not as case")
	}
   t := time.Now()
   switch {
   case t.Hour() < 12:
	   fmt.Println("It's before noon")
   default:
	   fmt.Println("It's after noon")

}

}