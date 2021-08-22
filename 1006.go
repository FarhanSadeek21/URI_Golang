package main
import "fmt"
func main() {
    var a, b, c, average float32
    fmt.Scanf("%f", &a)
    fmt.Scanf("%f", &b)
    fmt.Scanf("%f", &c)
    average = (a * 2 + b * 3 + c * 5) / 10
    fmt.Printf("MEDIA = %.1f\n", average)    
}