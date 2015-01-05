package main

import (
    "os"
    "fmt"

    "github.com/sabhiram/go-chloe/colorize"
)

/*****************************************************************************\

Define a header string for this command line application

\*****************************************************************************/
const chloeHeaderString = `
Usage:

    %s <command> [<args>] [<options>]

Commands:

Options:

See 'chloe help <command>' for details on a specific command.
`

/*****************************************************************************\

Define `main()` application entry-point

\*****************************************************************************/
func main() {
    fmt.Println(os.Args)

    // Print header
    fmt.Printf(chloeHeaderString, colorize.Colorize("chloe", "cyan"))
}
