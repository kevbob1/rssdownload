package main

import "fmt"

type Salutation struct {
    name string
    greeting string

}

func PrintLn(msg string) {
    fmt.Println(msg)
}

func CreateMessage(name string, greeting string) (message string) {
    fmt.Println(len(greeting))
    message = greeting + " " + GetPrefix(name) + name
    return
}

func GetPrefix(name string) (prefix string) {
    switch {
        case name == "Bob": prefix = "Mr "
        case name == "Joe", name == "Amy", len(name) == 10: prefix = "Dr "
        case name == "Mary": prefix = "Mrs "
        default: prefix = "Dude "
    }

    return
}


func Greet(sal Salutation, do func(string)) {
    message := CreateMessage(sal.name, sal.greeting)
    do(message)
}

func main() {
    var s = Salutation{"Amyxxxxxxx", "hello"}
    Greet(s, PrintLn)
    a := []int{3,2,1,5,6,7,8,8}
    for _, aa := range a {
        fmt.Println(aa)
    }
    
}
