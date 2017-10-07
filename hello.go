package main

import (
    "fmt"
    stringutil "github.com/jshields/go_stringutil"
)

func main() {
    fmt.Printf("Hello Go World!\n")
    fmt.Printf(stringutil.Reverse("that gum you like is coming back in style") + "\n")
}
