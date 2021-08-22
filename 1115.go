package main
import "fmt"
func main(){
	var x, y int
	for true {
		fmt.Scanf("%d %d", &x, &y)
		if x == 0 || y == 0 {
			break
		} else if x > 0 {
			if y > 0{
				fmt.Println("primeiro")
			} else {
				fmt.Println("quarto")
			}
		} else {
			if y > 0 {
				fmt.Println("segundo")
			} else {
				fmt.Println("terceiro")
			}
		}
	}
}