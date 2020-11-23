package handler

import (
	"github.com/merge/pkg/interval"
)

type Handler interface {
	StartInterval()
	EndInterval()
	IntervalData(interval.Data) error
	Splitter()
}

type VoidHandler struct{}

func (h VoidHandler) StartInterval() {}
func (h VoidHandler) EndInterval()   {}
func (h VoidHandler) IntervalData() error {return nil}
func (h VoidHandler) Splitter()  {}
