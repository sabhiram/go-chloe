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
    "bufio"
    "io/ioutil"
    "path/filepath"

    "github.com/sabhiram/colorize"
    "github.com/sabhiram/go-git-ignore"

    "github.com/jessevdk/go-flags"
)

// Define application constants
const (
    // Set "debugLoggingEnabled" to "true" if you want debug spew
    debugLoggingEnabled = true // false

    // Set "traceLoggingEnabled" to "true" if you want function entry spew
    traceLoggingEnabled = true // false

    // Set "timestampEnable" to "true" if you want timestamp output w/ all logs (except the Output logger)
    timestampEnabled = true // false
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
        Version     bool   `short:"v" long:"version"`
        Help        bool   `short:"h" long:"help"`
        File        string `short:"f" long:"file" default:"bower.json"`
        ForceDelete bool   `short:"y" long:"force"`
    }
)

// Sets up any application logging, and any other startup-y
// things we might need to do when this package is used (first-time)
func init() {
    var timestamp   = 0
    var debugWriter = ioutil.Discard
    var traceWriter = ioutil.Discard

    if timestampEnabled    { timestamp = log.Ldate | log.Ltime }
    if debugLoggingEnabled { debugWriter = os.Stdout }
    if traceLoggingEnabled { traceWriter = os.Stdout }

    Trace  = log.New(traceWriter, colorize.ColorString("TRACE: ", "magenta"), timestamp)
    Debug  = log.New(debugWriter, colorize.ColorString("DEBUG: ", "green"),   timestamp)
    Warn   = log.New(os.Stdout,   colorize.ColorString("WARN:  ", "yellow"),  timestamp)
    Error  = log.New(os.Stderr,   colorize.ColorString("ERROR: ", "red"),     timestamp)
    Output = log.New(os.Stdout,   "",                                         0)
}

// Executes the "chloe dispatch" command and its subset ("chloe list")
func chloeDispatch(command string) int {
    Trace.Printf("chloeDispatch()\n")

    var workingDir   string
    var files        []string
    var err          error
    var ignoreObject *ignore.GitIgnore

    // Build an ignore object from the input file
    if err == nil {
        ignoreObject, err = getIgnoreObjectFromJSONFile(Options.File)
    }

    // Fetch the current working dir where "chloe" was run from
    if err == nil {
        workingDir, err = os.Getwd()
    }

    // Fetch files we might want to delete using the "workingDir" as the base
    if err == nil {
        // Define function to aggregate matched paths into the "files" slice
        aggregateMatchedFilesFn := func(path string, fileInfo os.FileInfo, err error) error {
            relPath, _ := filepath.Rel(workingDir, path)
            if ignoreObject.MatchesPath(relPath) {
                files = append(files, relPath)
            }
            return nil
        }
        err = filepath.Walk(workingDir, aggregateMatchedFilesFn)
    }

    // List and delete files
    if err == nil && len(files) > 0 {
        Output.Printf("Found %d extra files:\n", len(files))
        for _, file := range files {
            Output.Printf(" - %s\n", file)
        }

        // Only attempt to delete if we are running a dispatch command
        if command == "dispatch" {
            deletePaths := Options.ForceDelete
            if !Options.ForceDelete {
                var input string
                reader := bufio.NewReader(os.Stdin)

                Output.Printf("Purge %d files? [ Yes | No ]: ", len(files))
                input, err = reader.ReadString('\n')
                input = strings.ToLower(strings.Trim(input, "\n"))

                deletePaths = false
                if containsString([]string{"t", "y", "true", "yes", "1"}, input) {
                    deletePaths = true
                }
            }

            // Actually walk and delete files
            if deletePaths {
                for _, file := range files {
                    fullPath, _ := filepath.Abs(file)
                    err = os.Remove(fullPath)

                    if err != nil { break }
                }

                if err == nil {
                    Output.Printf("Deleted %d files!\n", len(files))
                }
            }
        }
    } else if err == nil {
        Output.Printf("Found no files to cleanup\n")
    }

    // Handle error condition
    if err != nil {
        Error.Printf("%s\n", err.Error())
        return 1
    }
    return 0
}

// Runs the appropriate chloe command
func runCommand(command string) int {
    Trace.Printf("runCommand()\n")

    switch {
    case command == "list" || command == "dispatch":
        return chloeDispatch(command)
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
    command := strings.ToLower(strings.Join(args, " "))

    exitCode := 0
    switch {

    // Parse Error, print usage
    case error != nil:
        printAppUsageString()
        exitCode = 1

    // No arguments, or help requested, print usage
    case len(os.Args) == 1 || Options.Help:
        printAppUsageString()

    // "--version" requested
    case Options.Version:
        printAppVersionString()

    // "list" command invoked
    case isValidCommand(command):
        exitCode = runCommand(command)

    // All other cases go here!
    case true:
        Output.Printf("Unknown command %s, see usage:\n", colorize.ColorString(command, "red"))
        printAppUsageString()
        exitCode = 1
    }
    os.Exit(exitCode)
}
