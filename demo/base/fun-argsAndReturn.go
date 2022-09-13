package main
import "fmt"
import (
	"os"
)

// https://learnku.com/docs/the-way-to-go/chapter-description/3598

func main(){
    var str string = "go is a beautiful language"
    path := os.Getenv("PATH")
    fmt.Printf("str:%s\n",str)  
    fmt.Printf("OS path:%s\n",path) 

    fmt.Println("*********for cycle demo**************")
    var num1 int = 4
    for i:=0; i<num1;i++{
        fmt.Printf("This is the %d iteration\n", i)
    }

    for pos, char := range str{
        fmt.Printf("character %c starts at byte position %d\n", char, pos)
    }
    gfun("GFUNTEST");

}


func gfun(mstr string){
    fmt.Printf("gfun 的入参是 %s \n", mstr)
    var j int = 1;    
    for{
        j++
        if j==6{
            break
        }
        fmt.Printf("Test for ture and break %d\n", j);
    }

func gfun(mstr string){
    fmt.Printf("gfun 的入参是 %s \n", mstr)
    var j int = 1;    
    for{
        j++
        if j==6{
            break
        }
        fmt.Printf("Test for ture and break %d\n", j);
    }
}
