/****************************************************************************\

`chloe` is a simple binary which serves as a companion to `bower`. It simply
culls out any bower dependencies which are not needed for a given deployment.

\****************************************************************************/
package main

import (
    "os"
    "log"
    "strings"

    "io/ioutil"

    "github.com/sabhiram/colorize"

    "github.com/jessevdk/go-flags"
)

// Define arguments we care about
var opts struct {

    Version bool `short:"v" long:"version" description:"Print application version"`

    Help bool `short:"h" long:"help" description:"Prints this help menu"`

}

/*****************************************************************************\

Define application globals

\*****************************************************************************/
var (
    //
    // Define various log levels we wish to capture
    //

    // Trace is used for function enter exit logging
    Trace  *log.Logger

    // Debug is enabled for arbitary logging
    Debug  *log.Logger

    // Warning and error speak for themselves
    Warn   *log.Logger
    Error  *log.Logger

    // Output is any stuff we wish to print to the screen
    Output *log.Logger
)

/*****************************************************************************\

Define application constants

\*****************************************************************************/
const (
    // Enable this if we wish to dump debug and trace
    // methods. Ideally we just turn these off unless
    // we are in the process of debugging this app
    debugLoggingEnabled = true
    traceLoggingEnabled = false
)

/*****************************************************************************\

Define `init()` to setup and logging

\*****************************************************************************/
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

/*****************************************************************************\

Print app usage

\*****************************************************************************/
func printUsage() {
    Trace.Println("printUsage()")
    Output.Printf(colorize.Colorize(`Usage:

    <cyan>chloe</cyan> <command> [<options>]

Commands:

    list            lists all files deemed deletable
    dispatch        deletes any and all files marked in bower.json

Options:

    <yellow>-v --version</yellow>        prints the application version
    <yellow>-h --help</yellow>           prints this help menu

Version:

    <white>%s</white>

`), Version)
}

/*****************************************************************************\

Define `main()` application entry-point

\*****************************************************************************/
func main() {
    Trace.Println("main()")

    // Parse arguments which might get passed to `chloe`
    parser := flags.NewParser(&opts, flags.Default & ^flags.HelpFlag)
    args, error := parser.Parse()

    // If we got a parse error - print usage:
    if error != nil {
        printUsage()
        os.Exit(1)
    } else if len(os.Args) == 1 || opts.Help {
        printUsage()
        os.Exit(0)
    }

    // Handle `-version` option
    if opts.Version {
        Output.Println(Version)
        os.Exit(0)
    }

    Debug.Println(strings.Join(args, " "))
    Debug.Println("I am doing secret things...")
}
