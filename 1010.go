package main
import "fmt"
func main() {
	var code1, code2, unit1, unit2, price1, price2, total float32
	fmt.Scanf("%f", &code1)
	fmt.Scanf("%f", &unit1)
	fmt.Scanf("%f", &price1)
	fmt.Scanf("%f", &code2)
	fmt.Scanf("%f", &unit2)
	fmt.Scanf("%f", &price2)
	total = price1 * unit1 + price2 * unit2
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}