package main

import "fmt"

type Salutation struct {
    name string
    greeting string

}

func PrintLn(msg string) {
    fmt.Println(msg)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
    fmt.Println(len(greeting))
    message = greeting[1] + " " + name
    alternate =  "HEY! " + name
    return
}

func Greet(sal Salutation, do func(string)) {
    message, alternate := CreateMessage(sal.name, sal.greeting, "Yo!")
    do(message)
    do(alternate)
}

func main() {
    var s = Salutation{"Bob", "hello"}
    Greet(s, PrintLn)
}
