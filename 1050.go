package main
import "fmt"
func main(){
    var code int
	fmt.Scanf("%d", &code)
    if(code == 61){
        fmt.Println("Brasilia")
    } else if(code == 71){
        fmt.Println("Salvador")
    } else if(code == 11){
        fmt.Println("Sao Paulo")
    } else if(code == 21){
        fmt.Println("Rio de Janeiro")
    } else if(code == 32){
        fmt.Println("Juiz de Fora")
    } else if(code == 19){
        fmt.Println("Campinas")
    } else if(code == 27){
        fmt.Println("Vitoria")
    } else if(code == 31){
        fmt.Println("Belo Horizonte")
    } else {
        fmt.Println("DDD nao cadastrado")
    }
}