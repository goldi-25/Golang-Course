package main 
import "fmt"
func main (){
	var score = []int {2,3,4,5}
	fmt.Println("score:",score,len(score))
	score = append(score,7)
	score[1] = 9
	fmt.Println("score",score,len(score))

	b:= score[2:]
	fmt.Println("b",b)
	c := score [:3]
	d:=score[:]
	e:=score[3:5]
	fmt.Println(c)
	fmt.Println(d)
fmt.Println(e)
	
}