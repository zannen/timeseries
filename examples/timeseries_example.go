package main

import (
	"fmt"
	"time"

	"github.com/zannen/timeseries"
)

func exampleTimeSeries() *timeseries.LinearMemoryTimeSeries {
	ts := timeseries.NewLinearMemoryTimeSeries()
	dt := time.Date(2019, time.January, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 31; i++ {
		ts.Add(&timeseries.Datum{
			Timestamp: dt,
			Datum:     []byte(fmt.Sprintf("item #%03d", i)),
		})
		dt = dt.Add(time.Hour * 24)
	}
	return ts
}

func main() {
	var ts timeseries.TimeSeries

	ts = exampleTimeSeries()

	start := time.Date(2019, time.January, 10, 9, 0, 0, 0, time.UTC)
	end := time.Date(2019, time.January, 20, 12, 0, 0, 0, time.UTC)

	fmt.Printf("Getting data:\n- starting on or after: %v\n- ending strictly before: %v\n", start, end)
	datarange := ts.GetRange(start, end)

	l := len(datarange)
	for i := 0; i < l; i++ {
		item := datarange[i]
		fmt.Printf("%v:\t%s\n", item.Timestamp, string(item.Datum))
	}

}
