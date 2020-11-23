package interval

// an array of intervals
type Array struct {
	// input intervals
	Inputs []*Interval
}

// a single interval
type Interval struct {
	// the begin of the interval
	Low int
	// the end of teh interval
	High int
}

func NewInterval() *Interval {
	return &Interval{}
}
