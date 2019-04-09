package timeseries

import (
	"errors"
	"time"
)

// TimeSeries is an ordered sequence of data
type TimeSeries interface {
	Add(d *Datum)
	GetRange(from, to time.Time) []*Datum
}

// Datum is a single piece of recordable information
type Datum struct {
	Timestamp time.Time
	Datum     []byte
}

// SearchType is the data type for searching
type SearchType int

// Constants for searching
const (
	Tstart SearchType = iota
	Tend
)

// BadChronology is an error returned when attempting to create an out-of-order
// timeseries.
var BadChronology = errors.New("Bad chronology")

// BadIndex is an error returned when attempting to access an invalid data
// location.
var BadIndex = errors.New("Bad index")
