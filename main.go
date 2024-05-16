package main

import (
    "fmt"
    "os"
    "umlgo/parser"
    "io/ioutil"
)

func main() {
    code, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        fmt.Println("Error reading stdin:", err)
        os.Exit(1)
    }

    umlCode, err := parser.ParseAndGenerateUML(string(code))
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    // Write UML code to output file
    err = ioutil.WriteFile("output.uml", []byte(umlCode), 0644)
    if err != nil {
        fmt.Println("Error writing UML file:", err)
        os.Exit(1)
    }
}
