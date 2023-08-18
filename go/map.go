package main 
import "fmt"
func main (){
m := make(map[string]int) //syntax

m["k1"] = 2 //key pair/value
m["k2"] = 4 //key pair/value


fmt.Println("map:",m)

v1 := m["k2"]
v3 :=m["k3"]
fmt.Println("v1",v1)
fmt.Println("v3",v3)

delete(m,"k2")
fmt.Println("map:",m)

_, gol := m["k2"]
    fmt.Println("gol:", gol)

	n:=map[string]int{"go":1,"ldi":3}
	fmt.Println("n:",n)

	_, sang := m["k8"]
fmt.Println("sang",sang)

 goldi := map [string] int {
	"as":1,
 }

 fmt.Println("godli",goldi)
 r := map [[12]int]string{}
fmt.Println("r:",r)

}
