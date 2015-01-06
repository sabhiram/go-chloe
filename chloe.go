/****************************************************************************\

`chloe` is a simple binary which serves as a companion to `bower`. It simply
culls out any bower dependencies which are not needed for a given deployment.

\****************************************************************************/
package main

import (
	"log"
	"os"
	"strings"

	"io/ioutil"

	"github.com/sabhiram/colorize"

	"github.com/jessevdk/go-flags"
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

Define application globals

\*****************************************************************************/
var (
	//
	// Define various log levels we wish to capture
	//

	// Trace is used for function enter exit logging
	Trace *log.Logger

	// Debug is enabled for arbitary logging
	Debug *log.Logger

	// Warning and error speak for themselves
	Warn  *log.Logger
	Error *log.Logger

	// Output is any stuff we wish to print to the screen
	Output *log.Logger

	//
	// Define holders arguments we wish to parse
	//
	Options struct {
		Version bool `short:"v" long:"version" description:"Print application version"`
		Help    bool `short:"h" long:"help" description:"Prints this help menu"`
	}
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

    <yellow>-v --version</yellow>    prints the application version
    <yellow>-h --help</yellow>       prints this help menu

Version:

    <white>%s</white>

`), Version)
}

func chloeList() {
	Trace.Println("chloeList()")
}
func chloeDispatch() {
	Trace.Println("chloeDispatch()")
}

/*****************************************************************************\

Define `main()` application entry-point

\*****************************************************************************/
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
		printUsage()
		exitCode = 1

		// No args, or help requested, print usage
	case len(os.Args) == 1 || Options.Help:
		printUsage()

		// `--version` requested
	case Options.Version:
		Output.Println(Version)

		// `list` command invoked
	case strings.ToLower(command) == "list":
		chloeDispatch()

		// `dispatch` command invoked
	case strings.ToLower(command) == "dispatch":
		chloeList()

	// All other cases go here!
	case true:
		Output.Printf("Unknown command %s, see usage:\n", colorize.ColorString(command, "red"))
		printUsage()
		exitCode = 1
	}
	os.Exit(exitCode)
}
