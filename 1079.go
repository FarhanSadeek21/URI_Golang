package main
import "fmt"
func main() {
	var testCase int
	var n1, n2, n3, average float64
	fmt.Scanf("%d", &testCase)
	for i := 0; i < testCase; i++ {
		fmt.Scanf("%f %f %f", &n1, &n2, &n3)
		average = (n1 * 2 + n2 * 3 + n3 * 5) / (2 + 3 + 5)
		fmt.Printf("%.1f\n", average)
	}
}
