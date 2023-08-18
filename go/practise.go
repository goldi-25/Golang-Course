package main 
import "fmt"
func main (){
	var g = []int {1,2,3,4,5,6,7,8}
	fmt.Println("values:",g)
	var  a int 
fmt.Print("enter the number that we want to find ")
fmt.Scanf("%d",&a)
fmt.Println(a)
if (a == 4 )   {
	fmt.Println("number is : ",a )
} else {
	fmt.Println("not found")
}
}
