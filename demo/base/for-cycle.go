package main
import "fmt"
import (
	"os"
)

func main(){
    var str string = "hello,world! go is a beautiful language"
    path := os.Getenv("PATH")
    fmt.Printf("str:%s\n",str)  
    fmt.Printf("OS path:%s\n",path) 

    fmt.Println("*********for cycle demo**************")
    var num1 int = 7
    for i:=0; i<num1;i++{
        fmt.Printf("This is the %d iteration\n", i)
    }

    for pos, char := range str{
        fmt.Printf("character %c starts at byte position %d\n", char, pos)
    }

    var j int = 1;    
    for{
        j++
        if j==6{
            break
        }
        fmt.Printf("Test for ture and break %d\n", j);
    }

}
