package main
import (
    "fmt"
    //"unicode/utf8"
)
func main() {
    var input string
    fmt.Scanf("%s", &input)
    input_len := len(input)
    if input_len >= 10 {
        fmt.Printf("palavrao\n")
    } else {
        fmt.Printf("palavrinha\n")
    }
}