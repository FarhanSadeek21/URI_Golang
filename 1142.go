package main
import "fmt"
func main() {
	var limit int
	fmt.Scanf("%d", &limit)
	number := 1
	for i := 0; i < limit; i++ {
		for j := number; j < (number + 3); j++{
			fmt.Printf("%d ", j)
		}
		number += 4
		fmt.Printf("PUM\n")
	}
}