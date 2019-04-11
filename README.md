# timeseries [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

timeseries is a simple in-memory time series, with data being added only in
chronological order.

## Usage

```go
package main

import (
    "time"

    "github.com/zannen/timeseries"
)

func main() {
    ts := timeseries.NewLinearMemoryTimeSeries()

    d := &timeseries.Datum{
        Timestamp: time.Now(),
        Datum: []byte("Hello, world"),
    }

    ts.Add(d)
}
```
