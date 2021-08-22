package main
import (
    "fmt"
)
func main() {
    var radius, area float64
    fmt.Scanf("%f", &radius)
    area = 3.14159 * radius * radius
    fmt.Printf("A=%.4f\n", area)       
}