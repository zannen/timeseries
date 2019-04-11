# timeseries [![Build Status][Build Status]][ci] [![Coverage Status][cov-img]][cov]

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

[ci]: https://circleci.com/gh/zannen/timeseries
[cov-img]: https://codecov.io/gh/zannen/timeseries/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/zannen/timeseries
