package main

import (
    "fmt"
    "strings"

    "github.com/sabhiram/colorize"
    "github.com/sabhiram/go-git-ignore"
)

// Returns the Usage string for this application
func getAppUsageString() string {
    Trace.Printf("getAppUsageString()\n")

    return fmt.Sprintf(colorize.Colorize(UsageString), Version)
}

// Prints the Usage for this application
func printAppUsageString() {
    Trace.Printf("printAppUsageString()\n")

    Output.Printf(getAppUsageString())
}

// Returns the application version
func getAppVersionString() string {
    Trace.Printf("getAppVersionString()\n")

    return Version
}

// Prints the application version
func printAppVersionString() {
    Trace.Printf("printAppVersionString()\n")

    Output.Printf("%s\n", getAppVersionString())
}

// Returns true if the given list of strings contains the target
// string "s"
func containsString(list []string, s string) bool {
    Trace.Printf("containsString()\n")

    for _, item := range list {
        if strings.ToLower(s) == item {
            return true
        }
    }
    return false
}

// Loads a JSON file, and fetches a GitIgnore object from
// the given lines. The object returned is a "ignore.GitIgnore"
// which is returned from the go-git-ignore package.
func getIgnoreObjectFromJSONFile(f string) *ignore.GitIgnore {
    Trace.Printf("getIgnoreObjectFromJSONFile(%s)\n", f)

    lines := []string{".git"}
    object, _ := ignore.CompileIgnoreLines(lines...)
    return object
}
