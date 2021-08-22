package main
import "fmt"
func main(){
	var position, max int
	var numArray [100]int
	for i:= 0; i < 100; i++ {
		fmt.Scanf("%d", &numArray[i])
	}
	max = numArray[0]
	for i:= 0; i < 100; i++ {
		if max < numArray[i] {
			max = numArray[i]
			position = i + 1
		}	
	}
	fmt.Println(max)
	fmt.Println(position)
}