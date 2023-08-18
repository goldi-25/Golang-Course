// int 8 -128 to 127 
// int 16 -32768 to 32767
// int 32  -2147483648 to 2147482647
// intergers type 1) unsigned (not negative values) 2) signed (both positive and negative values )
// unit 8 0 to 255 
// unit 16 0 to 65535
// unit 32 0 to 4294967295 

package main 
import "fmt"
func main (){
var  a int = 45 
fmt.Printf("%v %T \n", a,a)  

var b uint = 45 
fmt.Printf("%v %T \n",b,b)

var c uint = 259
fmt.Printf("%v %T \n", c,c)

d := 24
e := 43 
fmt.Println("arithematic operations")
fmt.Println(d-e)
fmt.Println(d+e)
fmt.Println(d%e)
fmt.Println(d/e)
fmt.Println(d&e)
fmt.Println(d|e)
fmt.Println(d^e)
fmt.Println(d&^e)
fmt.Println(d >> e) //  / 
fmt.Println(d<<e) // *

var n complex64 = 1+2i  // or complex(1+2)
fmt.Printf("%+v %T \n", real(n) ,real(n))
fmt.Printf("%+v %T \n", imag(n) ,imag(n))

} 