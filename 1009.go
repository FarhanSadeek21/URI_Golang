package main
import (
    "fmt"
)
func main() {
    var name string
    var init_sal, tot_sal, sales float64
    fmt.Scanf("%s", &name)
    fmt.Scanf("%f", &init_sal)
    fmt.Scanf("%f", &sales)
    tot_sal = init_sal + sales * 0.15
    fmt.Printf("TOTAL = R$ %.2f\n", tot_sal)
}