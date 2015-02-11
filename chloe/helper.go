package main

import (
    "fmt"
    "io/ioutil"

    "strings"
    "encoding/json"

    "github.com/sabhiram/go-colorize"
    "github.com/sabhiram/go-git-ignore"
)

// Returns the Usage string for this application
func getAppUsageString() string {
    Trace.Printf("getAppUsageString()\n")

    commands, options := getAllOptions()
    return colorize.Colorize(fmt.Sprintf(UsageString, commands, options, getAppVersionString()))
}

// Returns the application version
func getAppVersionString() string {
    Trace.Printf("getAppVersionString()\n")

    return Version
}

// Loads a JSON file, and fetches a GitIgnore object from
// the given lines. The object returned is a "ignore.GitIgnore"
// which is returned from the go-git-ignore package.
func getIgnoreObjectFromJSONFile(f string) (*ignore.GitIgnore, error) {
    Trace.Printf("getIgnoreObjectFromJSONFile(%s)\n", f)

    var err         error
    var object      *ignore.GitIgnore
    var jsonContent []byte

    // Define tag so we can effectively marshal json -> our patterns struct
    var chloeJSON = struct {
        Patterns []string `json:"chloe"`
    } {}

    jsonContent, err = ioutil.ReadFile(f)

    if err == nil {
        err = json.Unmarshal(jsonContent, &chloeJSON)
    }

    if err == nil {
        object, err = ignore.CompileIgnoreLines(chloeJSON.Patterns...)
    }

    return object, err
}

// Returns true if the ValidCommands struct contains an entry with the
// input string "s"
func isValidCommand(s string) bool {
    Trace.Printf("isValidCommand(%s)\n", s)

    for _, item := range ValidCommands {
        if strings.ToLower(s) == item.name {
            return true
        }
    }
    return false
}

// Returns true if the list "strings" contains the "target" string
func containsString(strings []string, target string) bool {
    Trace.Printf("containsString(%s)\n", target)

    for _, item := range strings {
        if item == target {
            return true
        }
    }
    return false
}

// Returns a tuple of commands and options which we support
func getAllOptions() (string, string) {
    Trace.Printf("getAllOptions()\n")

    commands, options := "", ""
    for _, c := range ValidCommands {
        commands += fmt.Sprintf("    %-16s %s\n", c.name, c.description)
    }
    for _, o := range ValidOptions {
        options += fmt.Sprintf("    <yellow>-%s --%-8s</yellow>    %s\n", o.short, o.long, o.description)
    }
    return commands, options
}
