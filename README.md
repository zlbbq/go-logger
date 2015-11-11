# go-logger
simple logger for golang

##Install

go get github.com/zlbbq/go-logger

---

##Simple Usage
```go
package main

import "gitbub.com/zlbbq/go-logger"

func main() {
    logger.Debug("Hello, Debug")
    logger.Info("Hello, Info")
    logger.Warn("Hello, Warning")
    logger.Error("Hello, Error")
    logger.Fatal("Hello, Fata Error")
}

```

###Set global log level

```go
package main

import "gitbub.com/zlbbq/go-logger"

func main() {
    logger.SetLevel(logger.LevelInfo)
    logger.Debug("Hello, Debug")                // Debug output disappeared
    logger.Info("Hello, Info")
    logger.Warn("Hello, Warning")
    logger.Error("Hello, Error")
    logger.Fatal("Hello, Fata Error")
}

```

###Disable colorful output

```go
package main

import "gitbub.com/zlbbq/go-logger"

func main() {
    logger.SetColorful(false)                   // All log text with different levels has same color
    logger.Debug("Hello, Debug")
    logger.Info("Hello, Info")
    logger.Warn("Hello, Warning")
    logger.Error("Hello, Error")
    logger.Fatal("Hello, Fata Error")
}

```

###Redirect output stream
By default, logger writes log text to os.Stdout, but you can redirect it to any io.Writer created by yourself.

```go
package main

import (
    "os"

    "gitbub.com/zlbbq/go-logger"
)

func main() {
    f, err := os.OpenFile("a.log", os.O_CREATE | os.O_WRONLY, os.ModeAppend)
    if err == nil {
        logger.SetOutput(f)
        defer f.Close()
    }

    logger.Debug("Hello, Debug")
    logger.Info("Hello, Info")
    logger.Warn("Hello, Warning")
    logger.Error("Hello, Error")
    logger.Fatal("Hello, Fata Error")
}

```

---
##Advanced Usage
See logger documentation and samples

go get gitbub.com/zlbbq/go-logger

go doc -http=:9000

visit website: http://localhost:9000/


##License
MIT
