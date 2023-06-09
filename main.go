package main

import (
    "os"
    "os/user"
    "fmt"
    "github.com/gray-adeyi/monkey/repl"
)

func main(){
    user, err := user.Current()
    if err != nil{
        panic(err)
    }
    fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
    fmt.Println("Feel free to type in commands")
    repl.Start(os.Stdin, os.Stdout)
}
