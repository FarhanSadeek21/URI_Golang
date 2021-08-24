package main
import "fmt"
func main(){
	var a, b, c, perimeter, area float64
	fmt.Scanf("%f %f %f", &a, &b, &c)
	if (a + b) > c && (b + c) > a && (c + a) > b {
		perimeter = a + b + c
		fmt.Printf("Perimetro = %.1f\n", perimeter)
	} else {
		area = 0.5 * (a + b) * c
		fmt.Printf("Area = %.1f\n", area)
	}
}