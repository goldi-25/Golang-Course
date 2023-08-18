package main 
import "fmt"
func main (){
	a := []int {1,2,3,4,5,6,7,8,9}
	index := binarySearch(a,6)
	fmt.Println(index)
}
func binarySearch(a [] int ,value int )int { 
	high := len(a) -1 
	low := 0 
	for low <= high {
		mid := (high+ low )/2 
		 if value == a [mid]{
			return mid 
		 }else if value < a [mid]{
			high = mid -1 
		 } else if value > a [mid]{
			low = mid +1 
		 }
		}
	return - 1 
}
