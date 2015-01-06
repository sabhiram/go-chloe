// This file contains version specific, and usage information for
// the `chloe` application
package main

import (
    "fmt"
    "github.com/sabhiram/colorize"
)

// Version represents the current Semantic Version of this application
const Version = "0.0.1"

// Returns the usage string for this application
func getAppUsageString() string {
    Trace.Println("getAppUsageString()")
    return fmt.Sprintf(colorize.Colorize(`Usage:

    <cyan>chloe</cyan> <command> [<options>]

Commands:

    list            lists all files deemed deletable
    dispatch        deletes any and all files marked in bower.json

Options:

    <yellow>-v --version</yellow>    prints the application version
    <yellow>-h --help</yellow>       prints this help menu

Version:

    <white>%s</white>

`), Version)
}
