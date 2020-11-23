package handler

import (
	"github.com/merge/pkg/interval"
	"strconv"
)

type InputHandler struct {
	isLow    bool
	Array    interval.Array
	Interval *interval.Interval
	Verbose  bool
}

func (h *InputHandler) StartInterval() {
	h.isLow = true
	h.Interval = interval.NewInterval()
}

func (h *InputHandler) EndInterval() {
	h.Array.Inputs = append(h.Array.Inputs, h.Interval)
}

func (h *InputHandler) IntervalData(IntervalData interval.Data) error {
	value, err := strconv.Atoi(string(IntervalData))
	if err != nil {
		return err
	}
	if h.isLow {
		h.Interval.Low = value
	} else {
		// exchange low/high if in wrong order
		if value < h.Interval.Low {
			h.Interval.High = h.Interval.Low
			h.Interval.Low = value
		} else {
			h.Interval.High = value
		}
	}
	return nil
}

func (h *InputHandler) Splitter() {
	h.isLow = false
}
