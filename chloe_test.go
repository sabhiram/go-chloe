// Implement test cases for the "chloe" application
package main

import (
    "testing"
    "strings"

    "github.com/stretchr/testify/assert"
)

// Test application usage string printing
func TestChloeUsage(test *testing.T) {

}

// Test Helper functions
func TestContainsString(test *testing.T) {
    sampleStrings  := []string{"a", "b", "c"}
    assert.Equal(test, false, containsString(sampleStrings, "d"))
    assert.Equal(test, false, containsString(sampleStrings, "aa"))
    assert.Equal(test, false, containsString(sampleStrings, "abc"))

    assert.Equal(test, true,  containsString(sampleStrings, "a"))
    assert.Equal(test, true,  containsString(sampleStrings, "b"))
    assert.Equal(test, true,  containsString(sampleStrings, "c"))
}

func TestGetAllOptions(test *testing.T) {
    commands, options := getAllOptions()

    commandLines := strings.Split(commands, "\n")
    assert.Equal(test, 2, len(commandLines) - 1, "only two commands are supported")

    optionLines := strings.Split(options, "\n")
    assert.Equal(test, 4, len(optionLines) - 1, "only four options are supported")
}

func TestIsValidCommand(test *testing.T) {
    for _, item := range ValidCommands {
        assert.Equal(test, true, isValidCommand(item.name), "valid commands should be valid")
    }
    assert.Equal(test, false, isValidCommand("invalid_command"), "invalid commands should be rejected")
}
