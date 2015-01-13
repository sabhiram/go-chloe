package main

import (
    "fmt"
    "os"
    "strings"
    "path/filepath"

    "github.com/sabhiram/colorize"
    "github.com/sabhiram/go-git-ignore"
)

// Returns the Usage string for this application
func getAppUsageString() string {
    Trace.Printf("getAppUsageString()\n")

    commands, options := getAllOptions()
    return colorize.Colorize(fmt.Sprintf(UsageString, commands, options, getAppVersionString()))
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

// Loads a JSON file, and fetches a GitIgnore object from
// the given lines. The object returned is a "ignore.GitIgnore"
// which is returned from the go-git-ignore package.
func getIgnoreObjectFromJSONFile(f string) (*ignore.GitIgnore, error) {
    Trace.Printf("getIgnoreObjectFromJSONFile(%s)\n", f)

    lines := []string{"LICE*", "*.go"}
    object, err := ignore.CompileIgnoreLines(lines...)
    return object, err
}

// Returns true if the ValidCommands struct contains an entry with the
// input string "s"
func isValidCommand(s string) bool {
    Trace.Printf("isValidCommand()\n")

    for _, item := range ValidCommands {
        if strings.ToLower(s) == item.name {
            return true
        }
    }
    return false
}

// Returns true if the list "strings" contains the "target" string
func containsString(strings []string, target string) bool {
    Trace.Printf("containsString()\n")

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

// Walks a given basePath and returns all files which are ignored
func getDeletableFilesInPath(basePath string, ignoreObject *ignore.GitIgnore) ([]string, error) {
    Trace.Printf("getDeletableFilesInPath()\n")

    returnObject := []string{}
    listFilesFn  := func(path string, fileInfo os.FileInfo, err error) error {
        relPath, _ := filepath.Rel(basePath, path)
        if ignoreObject.MatchesPath(relPath) {
            returnObject = append(returnObject, relPath)
        }
        return nil
    }

    return returnObject, filepath.Walk(basePath, listFilesFn)
}

func removeFiles(files []string) error {
    for _, file := range files {
        Debug.Printf("rm -rf %s\n", file)
    }
    return nil
}


