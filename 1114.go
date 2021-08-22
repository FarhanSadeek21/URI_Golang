package main
import "fmt"
func main(){
	password := 2002
	var input int
	for true {
		fmt.Scanf("%d", &input)
		if password == input {
			fmt.Println("Acesso Permitido")
			break
		} else {
			fmt.Println("Senha Invalida")
		}
	}
}