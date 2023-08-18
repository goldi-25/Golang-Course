package main 
import "fmt"
func main (){
i := 1 
for i <= 5 {
	fmt.Println(i)
	i = i + 1
}	
for {
	fmt.Println("aatishgupta")
	break
 }
for r := 3 ; r <= 6 ;r++{
	fmt.Println(r)
}

for n := 0 ;n <10 ;n++{
	if n%2 == 0{
		continue
	}
	fmt.Println(n)
}

for i := 0 ; i < 5 ; i++{
	fmt.Println("rtgh",i)
}

for i := 0 ; i <34;i++ {
	fmt.Println(i)
}
 
}