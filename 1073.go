package main
import "fmt"
func main(){
	var limit, square int
	fmt.Scanf("%d", &limit)
	if limit % 2 != 0 {
		limit -= 1
	}
	for i := 2; i <= limit; i += 2 {
		square = i * i
		fmt.Printf("%d^2 = %d\n", i, square)
	}
}