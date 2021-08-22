package main
import "fmt"
func main() {
	var limit int
	for true {
		fmt.Scanf("%d", &limit)
		if limit == 0 {
			break
		} else {
			for i := 1; i <= limit; i++ {
				fmt.Printf("%d ", i)
			}
			fmt.Printf("\n")
		}
	}
}