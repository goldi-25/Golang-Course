package main 
import ("fmt"
"strconv")
func main(){
	var a , b, c int = 1 ,2,3
	fmt.Println(a ,b,c)
	var r int 
	fmt.Println(r)
	var f = !true
	fmt.Println(f)
g := "goldi"
fmt.Println(g)



var z int 
z = 34 
fmt.Println(z)

var y int = 8
fmt.Println(y)

x := 43 
fmt.Println(x)

var u string 
fmt.Println(u)


// inside the amin function float32 is not supporting it supports outside the main function 
fmt.Printf("%v, %T \n ",z,z)

var ( name string = "goldi"
age int = 23 )
fmt.Println("name: ,age:",name ,age)

var theHTTP  string = "https://google.com"
fmt.Println("theHTTP:",theHTTP) 

//uppercase variables goballiy used 
//lowercase varibales not globally used inside the function only 

// variables conversion concept

var  p int  = 45 
fmt.Printf("%v %T \n", p,p)
var q float32 

q  = float32(p)
fmt.Printf("%v %T \n", q,q)

// strconv  package add ,and strconv.Itoa(h)

var s int = 78 
fmt.Printf("%v %T \n", s,s)
var t string 
t = strconv.Itoa(s)
fmt.Printf("%v %T \n", t,t)
}