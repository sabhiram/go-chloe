package main

import (
    "os"
    "fmt"
)


func color(s string) string {
    return "\033[36m" + s + "\033[0m"
}

func main() {
    // Print header


    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)


    const chloeHeaderString = `
Usage:

    %s <command> [<args>] [<options>]

Commands:

Options:

See 'chloe help <command>' for details on a specific command.
`
    fmt.Printf(chloeHeaderString, color("chloe"))

}
