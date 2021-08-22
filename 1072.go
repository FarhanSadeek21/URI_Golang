package main
import "fmt"
func main() {
	var testCase int
	var number, in, out int
	in = 0
	out = 0
	fmt.Scanf("%d", &testCase)
	for i := 0; i < testCase; i++ {
		fmt.Scanf("%d", &number)
		if number >= 10 && number <= 20 {
			in += 1
		} else {
			out += 1
		}
	}
	fmt.Printf("%d in\n", in)
	fmt.Printf("%d out\n", out)
}
