package main
import "fmt"
func main(){
	var a, b int
	for true {
		fmt.Scanf("%d %d", &a, &b)
		if a == b {
			break
		} else if a > b {
			fmt.Printf("Decrescente\n")
		} else if a < b {
			fmt.Printf("Crescente\n")
		}
	}
}