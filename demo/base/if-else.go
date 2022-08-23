package main
import "fmt"
import (
	"os"
	"strings")

func main(){
    var str string = "hello,world!"
    path := os.Getenv("PATH")
    if strings.HasPrefix(str, "h") {
        fmt.Printf("String has h in prefix: %s\n", str)
    }
    fmt.Printf("str:%s\n",str)	
    fmt.Printf("OS path:%s\n",path)	
}
