package main
import "fmt"
func main() {
	noteArray := [7]int{100, 50, 20, 10, 5, 2, 1}
	var count, money int
	fmt.Scanf("%d", &money)
	fmt.Println(money)
	for i := 0; i < 7; i++ {
		count = money / noteArray[i]
		fmt.Printf("%d nota(s) de R$ %d,00\n", count, noteArray[i])
		money %= noteArray[i]
	}
}