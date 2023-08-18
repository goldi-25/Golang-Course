package main 
import "fmt"
func test (){
	fmt.Println("test")
}
func test1 (x int ){
	fmt.Println(x)
}

func add(x int,y int){
	fmt.Println(x+y)
}

func sub(g ,o int)int{
	return g - o
}

func calu(x,y int)(int ,int,int){
	return x-y,x+y,x%y
}


//func calu(x,y int)(z1 int ,z2 int,z3 int){
	// return 

	//func calu(x,y int)(z1,z2,z3 int){
	// 	z1 = x-y z2 = x+y  z3 = x%y return 
func main (){
	test()
	test1(7)
	add(55,66)
	ans:= sub(22,66)
	fmt.Println("ans:",ans)
	ans1,ans2,ans3:=calu(33,11)
fmt.Println(ans1,ans2,ans3)
	
}