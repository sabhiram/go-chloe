/****************************************************************************\

`colorize` is a simple package which returns an ascii colorized
string version of an input string

\****************************************************************************/
package colorize

import (
    "fmt"
    "strings"
)

/****************************************************************************\

`colorize` implements a single function colorize(...) which takes
in either a color "name" or a color ascii id.

Here is a table used for reference:

    Intensity   0       1      2       3       4       5       6       7
    Normal      Black   Red    Green   Yellow  Blue    Magenta Cyan    White
    Bright      Black   Red    Green   Yellow  Blue    Magenta Cyan    White


We also define a constant <--> color name mapping based on the above table

\****************************************************************************/
var colorToValueMap = map [string] int {
    "black":   0,
    "red":     1,
    "green":   2,
    "yellow":  3,
    "blue":    4,
    "magenta": 5,
    "cyan":    6,
    "white":   7,
}

func Colorize(input, color string) string {
    color = strings.ToLower(color)
    if colorIndex, valid := colorToValueMap[color]; valid {
        return fmt.Sprintf("\033[3%dm%s\033[0m", colorIndex, input)
    }
    return input
}
