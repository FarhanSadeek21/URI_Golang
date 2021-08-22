package main
import "fmt"
func main(){
	var testCase int
	var a, b, result float64
	fmt.Scanf("%d", &testCase)
	for i:= 0; i < testCase; i++{
		fmt.Scanf("%f %f", &a, &b)
		if b == 0 {
			fmt.Println("divisao impossivel")
		} else {
		result = a / b
		fmt.Printf("%.1f\n", result)
		}
	}
}