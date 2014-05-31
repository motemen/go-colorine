colorine [![GoDoc](https://godoc.org/github.com/motemen/go-colorine?status.png)](https://godoc.org/github.com/motemen/go-colorine) [![Build Status](https://api.travis-ci.org/motemen/go-colorine.svg?branch=master)](https://travis-ci.org/motemen/go-colorine)
========

A simple colorized console logger for golang

Example
-------

```go
import "github.com/motemen/go-colorine"

logger := colorine.NewLogger(
    colorine.Prefixes{
        // Using colorine's predefined TextStyles
        "create": colorine.Info,
        "exist":  colorine.Notice,
        // Or specify your own
        "debug": colorine.TextStyle{colorine.White, colorine.None},
        "error": colorine.TextStyle{colorine.BrightRed, colorine.White},
    },
    // The default prefix color
    colorine.TextStyle{colorine.Green, colorine.None},
)

logger.Log("create", "path/to/file")
logger.Log("exist", "path/to/another/file")
logger.Log("error", "something went wrong!")
```


Screenshot
----------

![screenshot](screenshot.png)
