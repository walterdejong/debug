/*
	debug.go	WJ125

	Copyright (c) 2025 Walter de Jong <walter@heiho.net>

	Permission is hereby granted, free of charge, to any person obtaining a copy of
	this software and associated documentation files (the "Software"), to deal in
	the Software without restriction, including without limitation the rights to
	use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
	of the Software, and to permit persons to whom the Software is furnished to do
	so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
*/

// Package debug implements debug printing functionality.
package debug

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/term"
)

// debug color codes.
const (
	Reset       string = "\x1b[0m"
	DarkGray           = "\x1b[30;1m"
	Red         string = "\x1b[31;1m"
	Green       string = "\x1b[32;1m"
	Yellow      string = "\x1b[33;1m"
	Blue        string = "\x1b[34;1m"
	Magenta     string = "\x1b[35;1m"
	Cyan        string = "\x1b[36;1m"
	White       string = "\x1b[37;1m"
	Black              = "\x1b[30m"
	DarkRed     string = "\x1b[31m"
	DarkGreen   string = "\x1b[32m"
	DarkYellow  string = "\x1b[33m"
	DarkBlue    string = "\x1b[34m"
	DarkMagenta string = "\x1b[35m"
	DarkCyan    string = "\x1b[36m"
	Gray        string = "\x1b[37m"
)

var (
	Enabled      bool                 // set this to true to enable debug output
	Colorize     bool                 // will be set to true if on a tty terminal
	ColorInfo             = DarkGreen // color of the informational lead
	ColorMessage          = Reset     // color of the message. Default is no color
	ColorReset            = Reset     // reset the color back to normal after printing the message
	output       *os.File = os.Stderr // by default log to stderr
)

func init() {
	// colorize if we're on a tty
	Colorize = term.IsTerminal(int(output.Fd()))
}

func SetOutput(f *os.File) {
	output = f

	// colorize if we're on a tty
	Colorize = term.IsTerminal(int(output.Fd()))
}

// Debug prints debug message (if debug output is enabled).
// A newline is automatically appended.
func Debug(format string, a ...any) {
	if !Enabled {
		return
	}

	var b strings.Builder

	if Colorize {
		fmt.Fprint(&b, ColorInfo)
	}

	pc, fullpath, lineno, ok := runtime.Caller(1)
	if !ok {
		fmt.Fprint(&b, "% ??")
	} else {
		funcname := runtime.FuncForPC(pc).Name()
		filename := filepath.Base(fullpath)
		fmt.Fprintf(&b, "%% %s:%d %s()", filename, lineno, funcname)
	}
	if Colorize {
		fmt.Fprint(&b, ColorMessage)
	}
	fmt.Fprint(&b, " ")
	fmt.Fprintf(&b, format, a...)

	if Colorize {
		fmt.Fprint(&b, ColorReset)
	}

	fmt.Fprintln(output, b.String())
}

// EOB
