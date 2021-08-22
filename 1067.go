package main
import "fmt"
func main() {
	var limit int
	fmt.Scanf("%d", &limit)
	if limit % 2 == 0 {
		limit -= 1
	}
	for i := 1; i <= limit; i += 2 {
		fmt.Println(i)
	}
}