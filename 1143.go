package main
import "fmt"
func main(){
	var limit, square, cube int
	fmt.Scanf("%d", &limit)
	for i := 1; i <= limit; i++ {
		square = i * i
		cube = square * i
		fmt.Printf("%d %d %d\n", i, square, cube)
	}
}