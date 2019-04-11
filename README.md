# timeseries [![CircleCI][circleci-img]][circleci] [![Coverage Status][codecov-img]][codecov]

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

[circleci-img]: https://circleci.com/gh/zannen/timeseries/tree/master.svg?style=svg
[circleci]: https://circleci.com/gh/zannen/timeseries/tree/master
[codecov-img]: https://codecov.io/gh/zannen/timeseries/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/zannen/timeseries
