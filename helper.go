package main

import (
    "fmt"
    "strings"

    "github.com/sabhiram/colorize"
)

// Returns the Usage string for this application
func getAppUsageString() string {
    Trace.Println("getAppUsageString()")
    return fmt.Sprintf(colorize.Colorize(UsageString), Version)
}

// Prints the Usage for this application
func printAppUsageString() {
    Trace.Println("printAppUsageString")
    Output.Printf(getAppUsageString())
}

// Returns true if the given list of strings contains the target
// string "s"
func containsString(list []string, s string) bool {
    for _, item := range list {
        if strings.ToLower(s) == item {
            return true
        }
    }
    return false
}
