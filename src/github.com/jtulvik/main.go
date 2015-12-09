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
    var o = flag.String("l", "", "search `directory` for include files")
    flag.Parse()
    fmt.Println("Selected language:")
    fmt.Println(*o)
}

type Glose struct{
    Engelsk string
    Norsk string 
}