// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package lockss

import (
	"io/ioutil"
	"log"
	"os"
)

// logger is a package logger that can be enabled from client code to allow
// logging output from this package when desired/needed for troubleshooting
var logger *log.Logger

const logTimeLayout string = "15:04:05"

func init() {

	// Disable logging output by default unless client code explicitly
	// requests it
	logger = log.New(os.Stderr, "[go-lockss/lockss] ", 0)
	logger.SetOutput(ioutil.Discard)

}

// EnableLogging enables logging output from this package. Output is muted by
// default unless explicitly requested (by calling this function).
func EnableLogging() {
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logger.SetOutput(os.Stderr)
}

// DisableLogging reapplies default package-level logging settings of muting
// all logging output.
func DisableLogging() {
	logger.SetFlags(0)
	logger.SetOutput(ioutil.Discard)
}
