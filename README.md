# [writers] -- Writer Collection

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/writers.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/writers)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/writers/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/writers.svg)](https://github.com/spiegel-im-spiegel/writers/releases/latest)

## Usage

```go
package main

import (
    "fmt"
    "io"
    "os"
    "regexp"

    "github.com/spiegel-im-spiegel/logf"
    "github.com/spiegel-im-spiegel/writers"
)

func main() {
    file, err := os.Create("log.txt")
    if err != nil {
        fmt.Printf("%#v\n", err)
        return
    }
    defer file.Close()

    logf.SetOutput(io.MultiWriter(file, writers.Regexp(regexp.MustCompile(`\[(ERROR|FATAL)\]`), os.Stdout)))
    for i := 0; i < 6; i++ {
        logf.SetMinLevel(logf.TRACE + logf.Level(i))
        logf.Tracef("Traceing: No. %d\n", i+1)
        logf.Debugf("Debugging: No. %d\n", i+1)
        logf.Printf("Information: No. %d\n", i+1)
        logf.Warnf("Warning: No. %d\n", i+1)
        logf.Errorf("Erroring: No. %d\n", i+1)
        logf.Fatalf("Fatal Erroring: No. %d\n", i+1)
    }
}
```

```
$ go run sample.go
2020/03/28 14:44:44 [ERROR] Erroring: No. 1
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 1
2020/03/28 14:44:44 [ERROR] Erroring: No. 2
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 2
2020/03/28 14:44:44 [ERROR] Erroring: No. 3
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 3
2020/03/28 14:44:44 [ERROR] Erroring: No. 4
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 4
2020/03/28 14:44:44 [ERROR] Erroring: No. 5
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 5
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 6

$ cat log.txt
2020/03/28 14:44:44 [TRACE] Traceing: No. 1
2020/03/28 14:44:44 [DEBUG] Debugging: No. 1
2020/03/28 14:44:44 [INFO] Information: No. 1
2020/03/28 14:44:44 [WARN] Warning: No. 1
2020/03/28 14:44:44 [ERROR] Erroring: No. 1
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 1
2020/03/28 14:44:44 [DEBUG] Debugging: No. 2
2020/03/28 14:44:44 [INFO] Information: No. 2
2020/03/28 14:44:44 [WARN] Warning: No. 2
2020/03/28 14:44:44 [ERROR] Erroring: No. 2
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 2
2020/03/28 14:44:44 [INFO] Information: No. 3
2020/03/28 14:44:44 [WARN] Warning: No. 3
2020/03/28 14:44:44 [ERROR] Erroring: No. 3
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 3
2020/03/28 14:44:44 [WARN] Warning: No. 4
2020/03/28 14:44:44 [ERROR] Erroring: No. 4
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 4
2020/03/28 14:44:44 [ERROR] Erroring: No. 5
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 5
2020/03/28 14:44:44 [FATAL] Fatal Erroring: No. 6
```

[writers]: https://github.com/spiegel-im-spiegel/writers
