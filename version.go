// This file contains version specific, and usage information for
// the chloe application
package main

// Version represents the current Semantic Version of this application
const Version = "0.0.1"

// List of strings which contain allowed commands
var ValidCommands = [] struct {
    command, description string
} {
    { `list`,     `lists all files which are deletable`                                  },
    { `dispatch`, `deletes any files which are redundant as indicated by the input file` },
}

// List of options which chloe supports
var ValidOptions = [] struct {
    short, long, description string
} {
    { `f`, `file`,    `sets the input JSON file, default is "bower.json"` },
    { `v`, `version`, `prints the application version`                    },
    { `h`, `help`,    `prints this help menu`                             },
}

// Usage string for chloe
var UsageString = `Usage:

    <cyan>chloe</cyan> <command> [<options>]

Commands:

    list            lists all files deemed deletable
    dispatch        deletes any and all files marked in bower.json

Options:

    <yellow>-f --file</yellow>
    <yellow>-v --version</yellow>
    <yellow>-h --help</yellow>       prints this help menu

Version:

    <white>%s</white>

`

