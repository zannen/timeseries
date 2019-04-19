package timeseries_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/zannen/timeseries"

	"github.com/stretchr/testify/assert"
)

func TestLinearMemoryTimeSeriesAdd(t *testing.T) {
	ts := atestLinearMemoryTimeSeries()
	assert.Equal(t, 31, ts.Len())
	assert.Equal(t, time.Date(2019, time.December, 1, 12, 0, 0, 0, time.UTC), getOrPanic(ts, 0).Timestamp)
	assert.Equal(t, time.Date(2019, time.December, 31, 12, 0, 0, 0, time.UTC), getOrPanic(ts, 30).Timestamp)
}

func TestLinearMemoryTimeSeriesAddN(t *testing.T) {
	ts := timeseries.NewLinearMemoryTimeSeries()
	dt := time.Date(2019, time.December, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 32; i += 4 {
		d := make([]*timeseries.Datum, 4)
		for j := 0; j < 4; j++ {
			d[j] = &timeseries.Datum{
				Timestamp: dt,
				Datum:     []byte(fmt.Sprintf("item #%03d", i+j)),
			}
			dt = dt.Add(time.Hour * 24)
		}
		ts.AddN(d[0], d[1], d[2], d[3])
	}
	assert.Equal(t, 32, ts.Len())
	assert.Equal(t, time.Date(2019, time.December, 1, 12, 0, 0, 0, time.UTC), getOrPanic(ts, 0).Timestamp)
	assert.Equal(t, time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC), getOrPanic(ts, 31).Timestamp)
}

func TestLinearMemoryTimeSeriesBadChronology(t *testing.T) {
	ts := timeseries.NewLinearMemoryTimeSeries()
	d1 := &timeseries.Datum{
		Timestamp: time.Date(2019, time.December, 1, 12, 0, 0, 0, time.UTC),
		Datum:     []byte("an item"),
	}
	ts.Add(d1)
	d2 := &timeseries.Datum{
		Timestamp: time.Date(2019, time.December, 1, 11, 59, 59, 0, time.UTC),
		Datum:     []byte("an item"),
	}
	err := ts.Add(d2)
	assert.Error(t, err, "expected an error when adding a datum before the last datum")
}

func TestLinearMemoryTimeSeriesBadIndex(t *testing.T) {
	ts := timeseries.NewLinearMemoryTimeSeries()
	d, err := ts.Get(-1)
	assert.Nil(t, d)
	assert.Error(t, err, "expected an error when using an invalid index (-1)")
	d1 := &timeseries.Datum{
		Timestamp: time.Date(2019, time.December, 1, 12, 0, 0, 0, time.UTC),
		Datum:     []byte("an item"),
	}
	ts.Add(d1)
	d, err = ts.Get(1)
	assert.Nil(t, d)
	assert.Error(t, err, "expected an error when using an invalid index (1)")
}

func TestLinearMemoryTimeSeriesGetRange(t *testing.T) {
	ts := atestLinearMemoryTimeSeries()

	// the first (index==0) item in ts
	exactStart := time.Date(2019, time.December, 1, 12, 0, 0, 0, time.UTC)
	// the last (index==30) item in ts
	exactEnd := time.Date(2019, time.December, 31, 12, 0, 0, 0, time.UTC)

	wayBeforeRange := exactStart.Add(time.Hour * 24 * -15)
	wayAfterRange := exactEnd.Add(time.Hour * 24 * 15)

	justBeforeStart := exactStart.Add(time.Hour * -1)
	justAfterStart := exactStart.Add(time.Hour)

	justBeforeEnd := exactEnd.Add(time.Hour * -1)
	justAfterEnd := exactEnd.Add(time.Hour)

	tests := []struct {
		from        time.Time
		to          time.Time
		shouldlen   int
		shouldfirst *timeseries.Datum
		shouldlast  *timeseries.Datum
	}{
		{wayBeforeRange, wayBeforeRange, 0, &timeseries.Datum{}, &timeseries.Datum{}},
		{wayBeforeRange, justBeforeStart, 0, &timeseries.Datum{}, &timeseries.Datum{}},
		{wayBeforeRange, exactStart, 1, getOrPanic(ts, 0), getOrPanic(ts, 0)},
		{wayBeforeRange, justAfterStart, 1, getOrPanic(ts, 0), getOrPanic(ts, 0)},
		{wayBeforeRange, justBeforeEnd, 30, getOrPanic(ts, 0), getOrPanic(ts, 29)},
		{wayBeforeRange, exactEnd, 31, getOrPanic(ts, 0), getOrPanic(ts, 30)},
		{wayBeforeRange, justAfterEnd, 31, getOrPanic(ts, 0), getOrPanic(ts, 30)},

		{wayBeforeRange, wayAfterRange, 31, getOrPanic(ts, 0), getOrPanic(ts, 30)},

		{justBeforeStart, wayAfterRange, 31, getOrPanic(ts, 0), getOrPanic(ts, 30)},
		{exactStart, wayAfterRange, 31, getOrPanic(ts, 0), getOrPanic(ts, 30)},
		{justAfterStart, wayAfterRange, 30, getOrPanic(ts, 1), getOrPanic(ts, 30)},
		{justBeforeEnd, wayAfterRange, 1, getOrPanic(ts, 30), getOrPanic(ts, 30)},
		{exactEnd, wayAfterRange, 1, getOrPanic(ts, 30), getOrPanic(ts, 30)},
		{justAfterEnd, wayAfterRange, 0, &timeseries.Datum{}, &timeseries.Datum{}},
		{wayAfterRange, wayAfterRange, 0, &timeseries.Datum{}, &timeseries.Datum{}},
	}

	for _, tt := range tests {
		actual := ts.GetRange(tt.from, tt.to)
		assert.Equal(t, tt.shouldlen, len(actual))
		if tt.shouldlen > 0 && tt.shouldlen == len(actual) {
			assert.Equal(t, tt.shouldfirst, actual[0])
			assert.Equal(t, tt.shouldlast, actual[len(actual)-1])
		}
	}
}

func atestLinearMemoryTimeSeries() *timeseries.LinearMemoryTimeSeries {
	ts := timeseries.NewLinearMemoryTimeSeries()
	dt := time.Date(2019, time.December, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 31; i++ {
		ts.Add(&timeseries.Datum{
			Timestamp: dt,
			Datum:     []byte(fmt.Sprintf("item #%03d", i)),
		})
		dt = dt.Add(time.Hour * 24)
	}
	return ts
}

func getOrPanic(ts *timeseries.LinearMemoryTimeSeries, i int) *timeseries.Datum {
	d, err := ts.Get(i)
	if err != nil {
		panic(err)
	}
	return d
}

func printLinearMemoryTimeSeries(ts *timeseries.LinearMemoryTimeSeries) {
	l := ts.Len()
	for i := 0; i < l; i++ {
		fmt.Printf("%3d: At %v\n", i, getOrPanic(ts, i).Timestamp)
	}
}

func BenchmarkLinearMemoryTimeSeriesAdd(b *testing.B) {
	d := &timeseries.Datum{
		Timestamp: time.Date(2019, time.December, 1, 12, 0, 0, 0, time.UTC),
		Datum:     make([]byte, 10),
	}
	ts := timeseries.NewLinearMemoryTimeSeries()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ts.Add(d)
		if ts.Len() > 10485760 {
			ts = timeseries.NewLinearMemoryTimeSeries()
		}
	}
}

func BenchmarkLinearMemoryTimeSeriesAdd8(b *testing.B) {
	d := &timeseries.Datum{
		Timestamp: time.Date(2019, time.December, 1, 12, 0, 0, 0, time.UTC),
		Datum:     make([]byte, 10),
	}
	ts := timeseries.NewLinearMemoryTimeSeries()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ts.Add(d)
		ts.Add(d)
		ts.Add(d)
		ts.Add(d)
		ts.Add(d)
		ts.Add(d)
		ts.Add(d)
		ts.Add(d)
		if ts.Len() > 10485760 {
			ts = timeseries.NewLinearMemoryTimeSeries()
		}
	}
}

func BenchmarkLinearMemoryTimeSeriesAddN8(b *testing.B) {
	d := &timeseries.Datum{
		Timestamp: time.Date(2019, time.December, 1, 12, 0, 0, 0, time.UTC),
		Datum:     make([]byte, 10),
	}
	ts := timeseries.NewLinearMemoryTimeSeries()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ts.AddN(d, d, d, d, d, d, d, d)
		if ts.Len() > 10485760 {
			ts = timeseries.NewLinearMemoryTimeSeries()
		}
	}
}

func BenchmarkLinearMemoryTimeSeriesGetRange(b *testing.B) {
	ts := atestLinearMemoryTimeSeries()
	wayBeforeRange := time.Date(2019, time.November, 15, 12, 0, 0, 0, time.UTC)
	wayAfterRange := time.Date(2020, time.January, 15, 12, 0, 0, 0, time.UTC)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ts.GetRange(wayBeforeRange, wayAfterRange)
	}
}
