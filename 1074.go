package main
import (
	"fmt"
)
func odd_even(number int) string {
	if number % 2 == 0 {
		return "EVEN "
	} else {
		return "ODD "
	}
}
func main(){
	var testCase, number int
	fmt.Scanf("%d", &testCase)
	for i := 0; i < testCase; i++ {
		fmt.Scanf("%d", &number)
		if number == 0 {
			fmt.Println("NULL")
		} else if number < 0 {
			number *= -1
			fmt.Printf(odd_even(number))
			fmt.Printf("NEGATIVE\n")

		} else if number > 0 {
			fmt.Printf(odd_even(number))
			fmt.Printf("POSITIVE\n")
		}
	}
}
