package main
import ("fmt")
func main() {
	var radius float64
	fmt.Scanf("%f", &radius)
	var volume float64 = 4 / 3.0 * 3.14159 * radius * radius * radius
	fmt.Printf("VOLUME = %.3f\n", volume)
}