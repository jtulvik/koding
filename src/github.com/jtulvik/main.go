// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

import (
    "fmt"
    "flag"
) 


func main() {
    fmt.Println("hello world")
    g := Glose {Engelsk: "car", Norsk: "bil"}
    fmt.Println(g.Engelsk)   
    var språk = flag.String("s", "", "Velg språk Engelsk eller Norsk")
    flag.Parse()
    fmt.Println("Selected language:")
    fmt.Println(*språk)
    if (*språk == "Norsk"){
        fmt.Println ("fant Norsk")
    }
}

type Glose struct{
    Engelsk string
    Norsk string 
}