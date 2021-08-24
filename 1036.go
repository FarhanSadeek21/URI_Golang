package main
import (
	"fmt"
	"math"
)
func main(){
	var a, b, c, d, r1, r2 float64
	fmt.Scanf("%f %f %f", &a, &b, &c)
	d = (b * b) - (4 * a * c)
	if a == 0 || d < 0 {
		fmt.Println("Impossivel calcular")
	} else {		
		r1 = (-1 * b + math.Sqrt(d)) / (2 * a)
		r2 = (-1 * b - math.Sqrt(d)) / (2 * a)
		fmt.Printf("R1 = %.5f\n", r1)
		fmt.Printf("R2 = %.5f\n", r2)
	}
}