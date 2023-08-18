package main 
import "fmt"
func main (){
	if 7 % 2 == 0 {
		fmt.Println("7 is even ")
	} else {
		fmt.Println("7 is odd")
	}

	if 5 < 2 {
		fmt.Println("5 is lessthan 2 ")
	}else {
		fmt.Println("5 is not less thsan 2")
	}
	if num := -8 ; num < 0 {
		fmt.Println("number is negative")
	} else if num > 0 {
		fmt.Println("number is positive")
	} else {
		fmt.Println("number is natural")
	}



	if true {
		fmt.Println("its true ")
	}

	if false  {
		fmt.Println("its true ")
	}

	number := 234
	guess := 123456 
	if guess < number {
		fmt.Println("less than")
	}

	if guess > number {
		fmt.Println("more  than")
	}

	if guess == number {
		fmt.Println(" got it ")
	}
}