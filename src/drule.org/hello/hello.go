package main

import "fmt"

type Salutation struct {
    name string
    greeting string

}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
    fmt.Println(len(greeting))
    message = greeting[1] + " " + name
    alternate =  "HEY! " + name
    return
}

func Greet(sal Salutation) {
    message, alternate := CreateMessage(sal.name, sal.greeting, "Yo!")
    fmt.Println(message)
    fmt.Println(alternate)

}

func main() {
    var s = Salutation{"Bob", "hello"}
    Greet(s)
}
