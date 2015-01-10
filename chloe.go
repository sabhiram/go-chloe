// `chloe` is a cli binary which serves as a companion to `bower`. Its
// single purpose is to list and delete any files not required as part of
// the `bower_dependencies`.
//
// `chloe` will scan your `bower.json` file for ignore and must-preserve
// files and directories, and cull any extra junk fetched by `bower`.
// Do remember that if you delete even the `README.md` file from a bower
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
    // Set `debugLoggingEnabled` to `true` if you want debug spew
    debugLoggingEnabled = true

    // Set `traceLoggingEnabled` to `true` if you want function entry spew
    traceLoggingEnabled = true
)

var _ = ignore.CompileIgnoreFile

// Define application globals
var (
    // Trace is used for function enter exit logging
    Trace *log.Logger

    // Debug is enabled for arbitary logging
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

func getIgnoreObjectFromJSONFile(f string) *ignore.GitIgnore {
    Trace.Printf("getIgnoreObjectFromJSONFile(%s)\n", f)
    lines := []string{".git"}
    object, _ := ignore.CompileIgnoreLines(lines...)
    return object
}

// Executes the `chloe list` command
func chloeList() {
    Trace.Println("chloeList()")
    ignoreObject := getIgnoreObjectFromJSONFile(Options.File)

    // TODO: Walk script dir, ignoreObject must be valid
    if ignoreObject == nil {
        Error.Println("Ignore object is null")
    }

    workingDir, err := os.Getwd()
    if err == nil {
        visit := func(path string, fileInfo os.FileInfo, err error) error {
            relPath, _ := filepath.Rel(workingDir, path)
            if ignoreObject.IgnoresPath(relPath) {
                //Debug.Println(relPath + " is ignored")
            } else {
                Debug.Println(relPath + " is included")
            }
            return nil
        }
        err = filepath.Walk(workingDir, visit)
        if err != nil {
            Debug.Println("Error is: " + err.Error())
        }
    } else {
        Debug.Println("Error is: " + err.Error())
    }
}

// Executes the `chloe dispatch` command
func chloeDispatch() {
    Trace.Println("chloeDispatch()")
    ignoreObject := getIgnoreObjectFromJSONFile(Options.File)

    // TODO: Walk script dir, ignoreObject must be valid
    if ignoreObject == nil {
        Error.Println("Ignore object is null")
    }
}

// Application entry-point for `chloe`. Responsible for parsing
// the cli args and invoking the appropriate action
func main() {
    Trace.Println("main()")

    // Parse arguments which might get passed to `chloe`
    parser := flags.NewParser(&Options, flags.Default & ^flags.HelpFlag)
    args, error := parser.Parse()
    command := strings.Join(args, " ")

    exitCode := 0
    switch {

    // Parse Error, print usage
    case error != nil:
        Output.Printf(getAppUsageString())
        exitCode = 1

    // No args, or help requested, print usage
    case len(os.Args) == 1 || Options.Help:
        Output.Printf(getAppUsageString())

    // `--version` requested
    case Options.Version:
        Output.Println(Version)

    // `list` command invoked
    case strings.ToLower(command) == "list":
        chloeList()

    // `dispatch` command invoked
    case strings.ToLower(command) == "dispatch":
        chloeDispatch()

    // All other cases go here!
    case true:
        Output.Printf("Unknown command %s, see usage:\n", colorize.ColorString(command, "red"))
        Output.Printf(getAppUsageString())
        exitCode = 1
    }
    os.Exit(exitCode)
}
