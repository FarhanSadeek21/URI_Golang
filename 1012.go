package main
import ("fmt")
func main() {
	var a, b, c, triangle, circle, trapezium, square, rectangle float32
	fmt.Scanf("%f %f %f", &a, &b, &c)
	triangle = a * c * 0.5
	circle = 3.14159 * c * c
	trapezium = 0.5 * (a + b) * c
	square = b * b
	rectangle = a * b
	fmt.Printf("TRIANGULO: %.3f\n", triangle)
	fmt.Printf("CIRCULO: %.3f\n", circle)
	fmt.Printf("TRAPEZIO: %.3f\n", trapezium)
	fmt.Printf("QUADRADO: %.3f\n", square)
	fmt.Printf("RETANGULO: %.3f\n", rectangle)
}