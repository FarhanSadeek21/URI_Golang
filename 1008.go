package main
import "fmt"
func main() {
    var hour, perHour, salary float32
    var number int32
    fmt.Scanf("%d", &number)
    fmt.Scanf("%f", &hour)
    fmt.Scanf("%f", &perHour)
    salary = hour * perHour
    fmt.Printf("NUMBER = %d\n", number)
    fmt.Printf("SALARY = U$ %.2f\n", salary)
}