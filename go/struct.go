package main 
import "fmt"
type employee struct{
	name string
	age int 
	no int 
}
type family struct{
	name string 
	age int 
	numbers [] int 
}
func main (){
	Eemployee:= employee{ name:"goldi",
age:12,
no:123456}
fmt.Println(Eemployee)


Efamily := family{
	name : "aatish",
	age : 22 ,
	numbers : []int {
		12345,
		54321,
		123456,
		5432112345,
		5432123,
		43212345,
		54321234,
	},
}
fmt.Println(Efamily)
fmt.Println(Efamily)
	}