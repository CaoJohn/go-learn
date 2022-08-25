package main
import "fmt"
import (
	"os"
)

func main(){
    var str string = "hello,world!"
    path := os.Getenv("PATH")
    fmt.Printf("str:%s\n",str)  
    fmt.Printf("OS path:%s\n",path) 

    switchVar1 := "A"
    switch switchVar1{
    case "A":
        fmt.Printf("switchVar1 is A\n")  
    case "B":
        fmt.Printf("switchVar1 is B\n")
    default:
        fmt.Printf("default\n")
    }
    fmt.Println("**********demo*************")
    switchVar2 := "B"
    switch switchVar2{
    case "A":
        fmt.Printf("switchVar2 is A\n")  
    case "B":
        fmt.Printf("switchVar2 is B\n")
        fallthrough
    default:
        fmt.Printf("default, support fallthrough\n")
    }
    // use switch as if-else
    fmt.Println("*********use switch as if-else**************")
    var num1 int = 7
    switch{
    case num1 < 0 :
        fmt.Println("Number is negative")
    case num1>0 && num1<10:
        fmt.Println("Number is between 0 and 10")
    default:
        fmt.Println("Number is bigger than 9")
    }
}
