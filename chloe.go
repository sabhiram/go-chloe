/****************************************************************************\

`chloe` is a simple binary which serves as a companion to `bower`. It simply
culls out any bower dependencies which are not needed for a given deployment.

\****************************************************************************/
package main

import (
    "os"
    "log"
    "flag"

    "io/ioutil"

    "github.com/sabhiram/colorize"
)

// Define arguments we care about
type CLIArgs struct {
    version bool
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

    //
    // Define arguments
    //
    Args CLIArgs
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

Setup custom logging

\*****************************************************************************/
func setupLogging() {
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

Define `init()` to setup cli arguments and logging

\*****************************************************************************/
func init() {
    setupLogging()

    // Setup flags we expect to parse
    flag.BoolVar(&Args.version, "version", false, "Prints the version of the application")

    // Override the `flag.Usage()` to have a pretty custom one for `chloe`
    flag.Usage = func() {
        Output.Printf(colorize.Colorize(`Usage:

    <cyan>chloe</cyan> <command> [<options>]

Commands:

    list            lists all files deemed deletable
    dispatch        deletes any and all files marked in bower.json

Options:

    <yellow>-version</yellow>        prints the application version

Version:
    %s

`), Version)
    }
}

/*****************************************************************************\

Define `main()` application entry-point

\*****************************************************************************/
func main() {
    Trace.Println("main()")

    // Parse arguments which might get passed to `chloe`
    flag.Parse()

    // If we got no arguments - print usage
    if len(os.Args) == 1 {
        flag.Usage()
        os.Exit(1)
    }

    // Handle `-version` option
    if Args.version {
        Output.Println(Version)
        os.Exit(0)
    }

    // Do chloe stuff!
    Debug.Println("I am doing secret things...")
}
