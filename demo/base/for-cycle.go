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

    fmt.Println("*********for cycle demo**************")
    var num1 int = 7
    for i:=0; i<num1;i++{
        fmt.Printf("This is the %d iteration\n", i)
    }
}
