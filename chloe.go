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

    "github.com/sabhiram/colorize"
    "github.com/sabhiram/go-git-ignore"

    "github.com/jessevdk/go-flags"
)

// Define application constants
const (
    // Set "debugLoggingEnabled" to "true" if you want debug spew
    debugLoggingEnabled = false // true

    // Set "traceLoggingEnabled" to "true" if you want function entry spew
    traceLoggingEnabled = false // true
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
        Version     bool   `short:"v" long:"version" description:"Print application version"`
        Help        bool   `short:"h" long:"help" description:"Prints this help menu"`
        File        string `short:"f" long:"file" description:"Set the file to be read. Default bower.json" default:"bower.json"`
        ForceDelete bool   `short:"y" long:"force" description:"Delete files without prompting"`
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

    Trace  = log.New(traceWriter, colorize.ColorString("TRACE: ", "magenta"), log.Ldate|log.Ltime)
    Debug  = log.New(debugWriter, colorize.ColorString("DEBUG: ", "green"),   log.Ldate|log.Ltime)
    Warn   = log.New(os.Stdout,   colorize.ColorString("WARN:  ", "yellow"),  log.Ldate|log.Ltime)
    Error  = log.New(os.Stderr,   colorize.ColorString("ERROR: ", "red"),     log.Ldate|log.Ltime)
    Output = log.New(os.Stdout,   "",                                         0)
}

// Executes the "chloe list" command
func chloeList() int {
    Trace.Printf("chloeList()\n")

    var workingDir  string
    var files       []string

    ignoreObject, err := getIgnoreObjectFromJSONFile(Options.File)

    if err == nil {
        workingDir, err = os.Getwd()
    }

    // Fetch files we might want to delete
    if err == nil {
        files, err = getDeletableFilesInPath(workingDir, ignoreObject)
    }

    // List files
    if err == nil && len(files) > 0 {
        for _, file := range files {
            Output.Printf("%s\n", file)
        }
        Output.Printf("Found %d un-needed files. Run 'chloe dispatch' to remove them.\n", len(files))
    } else if err == nil {
        Output.Printf("No un-needed files to list\n")
    }

    // Handle error condition
    if err != nil {
        Debug.Printf("Error is: %s\n", err.Error())
        return 1
    }
    return 0
}

// Executes the "chloe dispatch" command
func chloeDispatch() int {
    Trace.Printf("chloeDispatch()\n")

    var workingDir  string
    var files       []string

    ignoreObject, err := getIgnoreObjectFromJSONFile(Options.File)

    if err == nil {
        workingDir, err = os.Getwd()
    }

    // Fetch files we might want to delete
    if err == nil {
        files, err = getDeletableFilesInPath(workingDir, ignoreObject)
    }

    // List and delete files
    if err == nil && len(files) > 0 {
        Output.Printf("Found %d files to delete:\n", len(files))
        for _, file := range files {
            Output.Printf(" - %s\n", file)
        }

        deletePaths := Options.ForceDelete
        if !Options.ForceDelete {
            var input string
            reader := bufio.NewReader(os.Stdin)

            Output.Printf("Delete %d files? [True|False]: ", len(files))
            input, err = reader.ReadString('\n')
            input = strings.ToLower(strings.Trim(input, "\n"))

            deletePaths = false
            if containsString([]string{"t", "y", "true", "yes", "1"}, input) {
                deletePaths = true
            }
            Debug.Printf("Got value for deletePaths: %t\n", deletePaths)
        }

        if deletePaths {
            err = removeFiles(files)
        }
    } else if err == nil {
        Output.Printf("No un-needed files to delete\n")
    }

    // Handle error condition
    if err != nil {
        Debug.Printf("Error is: %s\n", err.Error())
        return 1
    }
    return 0
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
