package main
import (
	"fmt"
	"strings"
)
func main(){
	var input_str string
	var input_int int = (int) input_str
	fmt.Scanf("%s", &input_str)
	var present bool = strings.Contains(input_str, "13")
	if present == true{
		fmt.Printf("%d es de Mala Suerte\n", input_int)
	} else {
		fmt.Printf("%d NO es de Mala Suerte\n", input_int)
	}
}