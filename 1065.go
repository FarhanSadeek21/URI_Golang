package main
import "fmt"
func main() {
	var count int32
	count = 0
	var number int32
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &number)
		if number % 2 == 0 {
			count += 1
		}
	}
	fmt.Printf("%d valores pares\n", count)
}