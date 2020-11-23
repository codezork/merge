package pkg

import (
	"bytes"
	"github.com/merge/pkg/handler"
	"github.com/merge/pkg/interval"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Parse_OK(t *testing.T) {
	t.Run("parsing ok", func(t *testing.T) {
		handler := &handler.InputHandler{Verbose: true}
		reader := bytes.NewReader([]byte("[5,7][1,3]"))
		parser := NewParser(reader, handler, true)
		if err := parser.Parse(); (err != nil) {
			t.Errorf("Parse() error = %v", err)
		}
		result := []*interval.Interval{
			&interval.Interval{
				Low:  5,
				High: 7,
			},
			&interval.Interval{
				Low:  1,
				High: 3,
			},
		}
		assert.Equal(t, result, handler.Array.Inputs)
	})
}

func TestParser_Parse_Fail(t *testing.T) {
	t.Run("parsing fail", func(t *testing.T) {
		handler := &handler.InputHandler{Verbose: true}
		reader := bytes.NewReader([]byte("[5,.]"))
		parser := NewParser(reader, handler, true)
		if err := parser.Parse(); (err == nil) {
			t.Errorf("Parse() error not detected")
		}
	})
}