package timeseries

import (
	"time"
)

// LinearMemoryTimeSeries is an in-memory TimeSeries
type LinearMemoryTimeSeries struct {
	data []*Datum
}

// NewLinearMemoryTimeSeries creates a new empty LinearMemoryTimeSeries
func NewLinearMemoryTimeSeries() *LinearMemoryTimeSeries {
	ts := LinearMemoryTimeSeries{data: make([]*Datum, 0, 2048)}
	return &ts
}

// Add adds a Datum to a LinearMemoryTimeSeries
func (ts *LinearMemoryTimeSeries) Add(d *Datum) error {
	// O(1)
	l := len(ts.data)
	if l > 0 {
		last := ts.data[l-1]
		if d.Timestamp.Before(last.Timestamp) {
			return BadChronology
		}
	}
	ts.data = append(ts.data, d)
	return nil
}

// Get returns the specified Datum
func (ts *LinearMemoryTimeSeries) Get(i int) (*Datum, error) {
	if i < 0 || i >= len(ts.data) {
		return nil, BadIndex
	}
	return ts.data[i], nil
}

// AddN adds multiple Datums to a LinearMemoryTimeSeries
func (ts *LinearMemoryTimeSeries) AddN(d ...*Datum) {
	// O(1)
	ts.data = append(ts.data, d...)
}

// Len returns the length of the internal data
func (ts *LinearMemoryTimeSeries) Len() int {
	return len(ts.data)
}

func (ts *LinearMemoryTimeSeries) binarySearch(t time.Time, st SearchType) int {
	// O(log(N))
	var start, mid, end int
	end = len(ts.data) - 1
	for start <= end {
		mid = (start + end) / 2
		tmid := ts.data[mid].Timestamp
		if t.After(tmid) {
			start = mid + 1
		} else if t.Before(tmid) {
			end = mid - 1
		} else {
			if st == Tend {
				return mid + 1
			}
			return mid
		}
	}
	return start
}

// GetRange returns a slice of data from the LinearMemoryTimeSeries
func (ts *LinearMemoryTimeSeries) GetRange(from, to time.Time) []*Datum {
	// O(log(N))
	indexFrom := ts.binarySearch(from, Tstart)
	indexTo := ts.binarySearch(to, Tend)
	return ts.data[indexFrom:indexTo]
}
