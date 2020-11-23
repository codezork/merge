package pkg

import (
	"fmt"
	"github.com/merge/pkg/handler"
	"github.com/merge/pkg/helper"
	"github.com/merge/pkg/interval"
	"io"
)

type Parser struct {
	*interval.Decoder
	handler handler.Handler
	verbose bool
}

//Create a New Parser
func NewParser(reader io.Reader, handler handler.Handler, verbose bool) *Parser {
	decoder := interval.NewDecoder(reader)
	return &Parser{decoder, handler, verbose}
}

func (p *Parser) Parse() (err error) {
	defer helper.Elapsed("parse", p.verbose)()
	for {
		token, err := p.Token()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			panic(err)
		}
		switch token.(type) {
		case interval.StartInterval:
			p.handler.StartInterval()
		case interval.EndInterval:
			p.handler.EndInterval()
		case interval.Data:
			data := token.(interval.Data)
			err := p.handler.IntervalData(data)
			if err != nil {
				return err
			}
		case interval.Splitter:
			p.handler.Splitter()
		default:
			return fmt.Errorf("unknown token %s.", token)
		}
	}
	return nil
}

