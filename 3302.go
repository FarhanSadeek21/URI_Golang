package main
import "fmt"
func main() {
	var testCase int
    fmt.Scanf("%d", &testCase)
    var number int
	for i := 1; i <= testCase; i++ {
		fmt.Scanf("%d", &number)
        fmt.Printf("resposta %d: %d\n", i, number)
    }
}