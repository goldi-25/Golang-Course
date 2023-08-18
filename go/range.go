package main 
import "fmt"
func main (){
	//sllice
	 g := [] int {1,2,3,4,5,6}
	for i := range g {
		fmt.Println("ranhe:",g ,"as", g[i])
	}

	//array 
	number := [4] int {1,2,5,4}
	for index , item := range number {
		fmt.Printf("number[%d] = %d\n ", index ,item )
	} 

	//string 
	goldis := "golang"
	for i ,j := range goldis {
		fmt.Printf("%d = %c \n",i,j)
	} 

	//maps 
	numofHome := map[string]float32{"ds":21,"wdf":23}

	for subject , home := range numofHome {
		fmt.Println(subject, ":" , home)
	}

	for i,g:= range "abcd" {
		fmt.Println(i,g)
	}
}


