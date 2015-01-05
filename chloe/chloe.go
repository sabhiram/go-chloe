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
    // Handle header print if no args present
    // TODO: Handle bad arguments / commands / options
    if 0 == len(os.Args[1:]) {
        fmt.Printf(chloeHeaderString, colorize.Colorize("chloe", "cyan"))
    }


}
