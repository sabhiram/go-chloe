// This file contains version specific, and usage information for
// the chloe application
package main

// Version represents the current Semantic Version of this application
const Version = "0.0.1"

// List of strings which contain allowed commands
var ValidCommands = [] struct {
    name, description string
} {
    { `list`,     `lists all files which are deletable`                                  },
    { `dispatch`, `deletes any files which are deemed redundant` },
}

// List of options which chloe supports
var ValidOptions = [] struct {
    short, long, description string
} {
    { `i`, `input`,   `sets the input JSON file, default is "bower.json"` },
    { `f`, `force`,   `force delete without prompting user`               },
    { `v`, `version`, `prints the application version`                    },
    { `h`, `help`,    `prints this help menu`                             },
}

// Usage string for chloe
var UsageString = `Usage:

    <cyan>chloe</cyan> <command> [<options>]

Commands:

%s

Options:

%s

Version:

    <white>%s</white>

`

