package handler

import (
	"github.com/merge/pkg/helper"
	"github.com/merge/pkg/interval"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInputHandler_StartInterval(t *testing.T) {
	type fields struct {
		isLow    bool
		Array    interval.Array
		interval *interval.Interval
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Interval start", fields{
			Array:    interval.Array{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &InputHandler{
				Array:    tt.fields.Array,
				Verbose: true,
			}
			h.StartInterval()
			assert.NotNil(t, h.Interval, "Interval is nil")
			assert.True(t, h.isLow, "isLow not true")
		})
	}
}

func TestInputHandler_EndInterval(t *testing.T) {
	type fields struct {
		Array    interval.Array
		interval *interval.Interval
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"add Interval", fields{
			Array:    interval.Array{},
			interval: interval.NewInterval(),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &InputHandler{
				Array:    tt.fields.Array,
				Interval: tt.fields.interval,
			}
			h.EndInterval()
			assert.Equal(t, 1, len(h.Array.Inputs), "array not increased")
		})
	}
}

func TestInputHandler_IntervalData(t *testing.T) {
	type args struct {
		IntervalData interval.Data
	}
	h := &InputHandler{isLow: true, Interval: interval.NewInterval()}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Interval data1", args{
			[]byte("2"),
		}, false},
		{"Interval data1", args{
			[]byte("?"),
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := h.IntervalData(tt.args.IntervalData); (err != nil) != tt.wantErr {
				t.Errorf("IntervalData() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, 2, h.Interval.Low, "low value not set")
		})
	}
}

func TestInputHandler_Merge(t *testing.T) {
	t.Run("merge", func(t *testing.T) {
		h := &InputHandler{Verbose: true}
		h.StartInterval()
		h.IntervalData([]byte("5"))
		h.Splitter()
		h.IntervalData([]byte("8"))
		h.EndInterval()
		h.StartInterval()
		h.IntervalData([]byte("4"))
		h.Splitter()
		h.IntervalData([]byte("2"))
		h.EndInterval()
		h.StartInterval()
		h.IntervalData([]byte("1"))
		h.Splitter()
		h.IntervalData([]byte("3"))
		h.EndInterval()
		results := h.Merge()
		assert.Equal(t, "[1,4][5,8]", helper.Print("", results), "merge failed")
	})
}

func TestInputHandler_Splitter(t *testing.T) {
	type fields struct {
		isLow    bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"splitter", fields{isLow: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &InputHandler{
				isLow:    tt.fields.isLow,
				Verbose: true,
			}
			h.Splitter()
			assert.False(t, h.isLow, "isLow not switched")
		})
	}
}

