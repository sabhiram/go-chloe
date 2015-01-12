// "chloe" is a cli binary which serves as a companion to "bower". Its
// single purpose is to list and delete any files not required as part of
// the "bower_dependencies".
//
// "chloe" will scan your "bower.json" file for ignore and must-preserve
// files and directories, and cull any extra junk fetched by "bower".
// Do remember that if you delete even the "README.md" file from a bower
// package - it will prompt bower to re-fetch it on the next update.
package main

import (
    "log"
    "os"
    "strings"
    "io/ioutil"
    "path/filepath"

    "github.com/sabhiram/colorize"
    "github.com/sabhiram/go-git-ignore"

    "github.com/jessevdk/go-flags"
)

// Define application constants
const (
    // Set "debugLoggingEnabled" to "true" if you want debug spew
    debugLoggingEnabled = true

    // Set "traceLoggingEnabled" to "true" if you want function entry spew
    traceLoggingEnabled = true
)

var _ = ignore.CompileIgnoreFile

// Define application globals
var (
    // Trace is used for function enter exit logging
    Trace *log.Logger

    // Debug is enabled for arbitrary logging
    Debug *log.Logger

    // Warning and error speak for themselves
    Warn  *log.Logger
    Error *log.Logger

    // Output is any stuff we wish to print to the screen
    Output *log.Logger

    // Define holders for the cli arguments we wish to parse
    Options struct {
        Version bool   `short:"v" long:"version" description:"Print application version"`
        Help    bool   `short:"h" long:"help" description:"Prints this help menu"`
        File    string `short:"f" long:"file" description:"Set the file to be read. Default bower.json" default:"bower.json"`
    }
)

// Sets up any application logging, and any other startup-y
// things we might need to do when this package is used (first-time)
func init() {
    var debugWriter = ioutil.Discard
    if debugLoggingEnabled {
        debugWriter = os.Stdout
    }

    var traceWriter = ioutil.Discard
    if traceLoggingEnabled {
        traceWriter = os.Stdout
    }

    Trace = log.New(traceWriter,
        colorize.ColorString("TRACE: ", "magenta"),
        log.Ldate|log.Ltime)

    Debug = log.New(debugWriter,
        colorize.ColorString("DEBUG: ", "green"),
        log.Ldate|log.Ltime)

    Warn = log.New(os.Stdout,
        colorize.ColorString("WARN:  ", "yellow"),
        log.Ldate|log.Ltime)

    Error = log.New(os.Stderr,
        colorize.ColorString("ERROR: ", "red"),
        log.Ldate|log.Ltime)

    Output = log.New(os.Stdout, "", 0)
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

// Executes the "chloe list" command
func chloeList() int {
    Trace.Printf("chloeList()\n")
    ignoreObject := getIgnoreObjectFromJSONFile(Options.File)

    // TODO: Walk script directory, ignoreObject must be valid
    if ignoreObject == nil {
        Error.Printf("Ignore object is null\n")
    }

    workingDir, err := os.Getwd()
    if err == nil {
        visit := func(path string, fileInfo os.FileInfo, err error) error {
            relPath, _ := filepath.Rel(workingDir, path)
            if ignoreObject.IgnoresPath(relPath) {
                //Debug.Printf(relPath + " is ignored\n")
            } else {
                Debug.Printf(relPath + " is included\n")
            }
            return nil
        }
        err = filepath.Walk(workingDir, visit)
    }

    if err != nil {
        Debug.Printf("Error is: %s\n", err.Error())
    }

    return 0
}

// Executes the "chloe dispatch" command
func chloeDispatch() int {
    Trace.Printf("chloeDispatch()\n")
    ignoreObject := getIgnoreObjectFromJSONFile(Options.File)

    // TODO: Walk script dir, ignoreObject must be valid
    if ignoreObject == nil {
        Error.Printf("Ignore object is null\n")
    }

    return 1
}

// Runs the appropriate chloe command
func runCommand(command string) int {
    Trace.Printf("runCommand()\n")

    switch {
    case command == "list":
        return chloeList()
    case command == "dispatch":
        return chloeDispatch()
    }
    panic(command + " is not a valid command, this code should not be hit!")
    return 1
}

// Application entry-point for "chloe". Responsible for parsing
// the cli arguments and invoking the appropriate action
func main() {
    Trace.Printf("main()\n")

    // Parse arguments which might get passed to "chloe"
    parser := flags.NewParser(&Options, flags.Default & ^flags.HelpFlag)
    args, error := parser.Parse()
    command := strings.Join(args, " ")

    exitCode := 0
    switch {

    // Parse Error, print usage
    case error != nil:
        Output.Printf(getAppUsageString())
        exitCode = 1

    // No arguments, or help requested, print usage
    case len(os.Args) == 1 || Options.Help:
        printAppUsageString()

    // "--version" requested
    case Options.Version:
        Output.Printf("%s\n", Version)

    // "list" command invoked
    case containsString(ValidCommands, command):
        exitCode = runCommand(command)

    // All other cases go here!
    case true:
        Output.Printf("Unknown command %s, see usage:\n", colorize.ColorString(command, "red"))
        Output.Printf(getAppUsageString())
        exitCode = 1
    }
    os.Exit(exitCode)
}
