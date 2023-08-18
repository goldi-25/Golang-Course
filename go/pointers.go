package main 
import "fmt"
func main (){
	var a int  = 42
	var b *int = &a 
	fmt.Println(a,b)
	a = 456 
	fmt.Println(a,*b)

*b = 3456789
fmt.Println(a,*b)


c := [4] int {22,33,44,55}
d := &c [0]
e := &c [1]
fmt.Printf("%d %d %d " , c,*d,*e)
ad := []address{}
modifyAddress(&ad)
fmt.Println(ad)

}
type address struct{
	age int
	name string
	address string
}
func modifyAddress(add *[]address){
	for i:=0;i<10;i=i+1{
		*add=append(*add,address{i,"name"+string(i),"address"+string(i)})
	}
}