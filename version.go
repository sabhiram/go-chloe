// This file contains version specific, and usage information for
// the chloe application
package main

// Version represents the current Semantic Version of this application
const Version = "0.0.1"

// List of strings which contain allowed commands
var ValidCommands = []string{"list", "command"}

// Usage string for chloe
var UsageString = `Usage:

    <cyan>chloe</cyan> <command> [<options>]

Commands:

    list            lists all files deemed deletable
    dispatch        deletes any and all files marked in bower.json

Options:

    <yellow>-f --file</yellow>       sets the input JSON file, default is "bower.json"
    <yellow>-v --version</yellow>    prints the application version
    <yellow>-h --help</yellow>       prints this help menu

Version:

    <white>%s</white>

`

