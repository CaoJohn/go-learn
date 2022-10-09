package main
import "fmt"
import (
	"os"
	"strings"
)

func main(){
    var str string = "hello,world!"
    path := os.Getenv("PATH")
    if strings.HasPrefix(str, "h") {
        fmt.Printf("String has h in prefix: %s\n", str)
    }else if strings.HasPrefix(str, "x"){
        fmt.Printf("String has x in prefix: %s\n", str)
    }else{
        fmt.Printf("String has not been matched in prefix: %s\n", str)
    }
    fmt.Printf("str:%s\n",str)	
    fmt.Printf("OS path:%s\n",path)	
}
