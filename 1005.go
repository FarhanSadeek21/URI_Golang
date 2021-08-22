package main
import "fmt"
func main() {
    var a, b, average float32
    fmt.Scanf("%f", &a)
    fmt.Scanf("%f", &b)
    average = (a * 3.5 + b * 7.5) / (3.5 + 7.5)
    fmt.Printf("MEDIA = %.5f\n", average)    
}