// Implement test cases for the "chloe" application
package main

import (
    "os"
    "fmt"
    "testing"
    "strings"

    "path/filepath"
    "io/ioutil"

    "github.com/stretchr/testify/assert"
)

const (
    TEST_DIR = "tmp_test_dir"
)

// Helper function to setup a test fixture dir and write to
// it a file with the name "fname" and content "content"
func writeFileToTestDir(fname, content string) {
    testDirPath := "." + string(filepath.Separator) + TEST_DIR
    testFilePath := testDirPath + string(filepath.Separator) + fname

    _ = os.MkdirAll(testDirPath, 0755)
    _ = ioutil.WriteFile(testFilePath, []byte(content), os.ModePerm)
}

func cleanupTestDir() {
    _ = os.RemoveAll(fmt.Sprintf(".%s%s", string(filepath.Separator), TEST_DIR))
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

func TestGetAppUsage(test *testing.T) {
    usageStr := getAppUsageString()
    assert.Contains(test, usageStr, "<command> [<options>]", "usage must contain sample usage str")
}

func TestGetAppVersion(test *testing.T) {
    version := getAppVersionString()
    assert.Contains(test, version, Version, "version must match internal global var")
}

func TestGetIgnoreObjectFromFile(test *testing.T) {
    writeFileToTestDir("test.json", `
{
    "chloe": [
        "node_modules",
        "**/*.out"
    ]
}
`)
    defer cleanupTestDir()

    object, err := getIgnoreObjectFromJSONFile("tmp_test_dir/test.json")
    assert.Nil(test, err, "err should be nil")

    assert.Equal(test, false, object.MatchesPath("foo"), "foo should not match")

    assert.Equal(test, true,  object.MatchesPath("foobar/blah.out"),  "foobar/blah.out should match")
    assert.Equal(test, true,  object.MatchesPath("node_modules/foo"), "node_modules/foo should match")
}
