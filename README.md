A debug module for Go
=====================

Package `debug` implements the missing debug printing function for Go.

The debug messages printed include filename, line number, package and name
of the calling function, and they are printed in color so they stand out
in the console terminal. Colorization is auto-disabled when output is
redirected to a file or piped to another program.

Despite being just a souped-up printf, I wouldn't want to do without
this package. Take it, as a convenience to ease your life in software dev.

## Usage

Step 1. Import the package.

```go
import "github.com/walterdejong/debug"
```

Step 2. Enable debug mode. Additionally, you may configure the colors
and/or the output file handle.

```go
    debug.Enabled = true

    // change colorization (if we want to)
    debug.Info = debug.DarkCyan
    debug.Message = debug.Cyan

    // change output to stdout (default is stderr)
    debug.SetOutput(os.Stdout)
```

Step 3. Log messages. `Debug()` works just like `fmt.Printf()`.

```go
    debug.Debug("the value of x == %#v", x)
```

This produces output like:

```
% file.go:42 pkg.MyFunction() the value of x == []int{1, 2, 3}
```

See also the provided [example program](example/example.go).

## Copyright and License

Copyright (c) 2025 by Walter de Jong <walter@heiho.net>

This software is freely available under terms described in the MIT license.
