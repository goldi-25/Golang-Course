package main 
import "fmt"
 func main (){
	var g [5] int 
	fmt.Println("g:",g)
	g[3]=25
	fmt.Println("set:",g)
	fmt.Println("get:",g[3])
	fmt.Println("len:",len(g))
	  r:= [5] int {1,2,3,4,6}
	 fmt.Println("r:",r)
	 var twod [2][4]int 
	 for i:=0 ;i<2;i++{
		for j:=0;j<4;j++ {
			twod[i][j]=i+j
		} }
fmt.Println("2D",twod)

var gol = [3]int{1,2,3}
fmt.Println("gold:",gol,"length: ",len(gol))
fruit := [4]string{"apple","mango","guaua" ,"orange"}
fmt.Println("fruit:",fruit,"lenth:",len(fruit))

grades := [3]int{23,45,67}
fmt.Printf("grades: %v  \n ",grades)
goldi := [...] int {11,22,33,55}
fmt.Printf("goldi: %v \n ", goldi)
  var student [2]string
  fmt.Printf("student:%v\n", student)
  student[0] = "aatish"
  student[1]= "goldi" 
  fmt.Printf("student:%v\n", student)

  //identity matrix 
  var threeD [3][3]int 
  for i:=0;i<3;i++{
	for j := 0 ;j<3 ;j++{
		threeD [i][j] = i+j
	}}
  fmt.Println("threeD:",threeD)
 }
