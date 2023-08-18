package main 
import "fmt"

 func main(){
	var aatish = []string{"as","ds","cd","tr"}
	fmt.Println("aatish",aatish,len(aatish))
	rangeOne :=aatish[1:2]
	fmt.Println(rangeOne)
	aatish = append(aatish,"uy")
	fmt.Println("aatish",aatish,len(aatish))
	rangeOne = append(rangeOne,"tg")
	fmt.Println(rangeOne)

	

 }
