package main
import(
    "fmt"
)
func main() {
    var a, b, c, d, result int32
    fmt.Scanf("%d", &a)
    fmt.Scanf("%d", &b)
    fmt.Scanf("%d", &c)
    fmt.Scanf("%d", &d)
    result = (a * b - c * d)
    fmt.Printf("DIFERENCA = %d\n", result)
}