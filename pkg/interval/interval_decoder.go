package interval

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"unicode/utf8"
)

// A SyntaxError represents a syntax error in the input stream.
type SyntaxError struct {
	Msg  string
	Line int
}

func (e *SyntaxError) Error() string {
	return "Interval syntax error on line " + strconv.Itoa(e.Line) + ": " + e.Msg
}

// A Token is an interface holding one of the token types:
// StartInterval, EndInterval, IntervalData, Splitter.
type Token interface{}

// A StartInterval represents an Interval start element.
type StartInterval struct {
}

// An EndInterval represents an Interval end element.
type EndInterval struct {
}

// A Data represents Interval character data (raw text)
type Data []byte

// A Splitter represents a character datas' split element
type Splitter struct {
}

// A TokenReader is anything that can decode a stream of interval tokens
// When Token encounters an error or end-of-file condition after successfully
// reading a token, it returns the token. It may return the (non-nil) error from
// the same call or return the error (and a nil token) from a subsequent call.
// An instance of this general case is that a TokenReader returning a non-nil
// token at the end of the token stream may return either io.EOF or a nil error.
// The next Read should return nil, io.EOF.
type TokenReader interface {
	Token() (Token, error)
}

// A Decoder represents an interval parser reading a particular input stream.
// The parser assumes that its input is encoded in UTF-8.
type Decoder struct {
	r              io.ByteReader
	t              TokenReader
	buf            bytes.Buffer
	saved          *bytes.Buffer
	nextByte       int
	err            error
	line           int
	offset         int64
	unmarshalDepth int
}

func NewDecoder(r io.Reader) *Decoder {
	d := &Decoder{
		nextByte: -1,
		line:     1,
	}
	d.switchToReader(r)
	return d
}

func (d *Decoder) Token() (Token, error) {
	var t Token
	var err error
	if t, err = d.rawToken(); err != nil {
		return t, err
	}
	switch t1 := t.(type) {
	case StartInterval:
		t = t1
	case EndInterval:
		t = t1
	}
	return t, err
}

func (d *Decoder) switchToReader(r io.Reader) {
	// Get efficient byte at a time reader.
	// Assume that if reader has its own
	// ReadByte, it's efficient enough.
	// Otherwise, use bufio.
	if rb, ok := r.(io.ByteReader); ok {
		d.r = rb
	} else {
		d.r = bufio.NewReader(r)
	}
}

// Creates a SyntaxError with the current line number.
func (d *Decoder) syntaxError(msg string) error {
	return &SyntaxError{Msg: msg, Line: d.line}
}

func (d *Decoder) rawToken() (Token, error) {
	if d.t != nil {
		return d.t.Token()
	}
	if d.err != nil {
		return nil, d.err
	}
	b, ok := d.getc()
	if !ok {
		return nil, d.err
	}

	if b != '[' && b != ']' && b != ',' {
		// Text section.
		d.ungetc(b)
		data := d.text()
		if data == nil {
			return nil, d.err
		}
		return Data(data), nil
	}

	switch b {
	case ']':
		// End element
		d.space()
		return EndInterval{}, nil

	case ',':
		// Separator
		d.space()
		return Splitter{}, nil
	}
	return StartInterval{}, nil
}

// Skip spaces if any
func (d *Decoder) space() {
	for {
		b, ok := d.getc()
		if !ok {
			return
		}
		switch b {
		case ' ', '\r', '\n', '\t':
		default:
			d.ungetc(b)
			return
		}
	}
}

// Read a single byte.
// If there is no byte to read, return ok==false
// and leave the error in d.err.
// Maintain line number.
func (d *Decoder) getc() (b byte, ok bool) {
	if d.err != nil {
		return 0, false
	}
	if d.nextByte >= 0 {
		b = byte(d.nextByte)
		d.nextByte = -1
	} else {
		b, d.err = d.r.ReadByte()
		if d.err != nil {
			return 0, false
		}
		if d.saved != nil {
			d.saved.WriteByte(b)
		}
	}
	if b == '\n' {
		d.line++
	}
	d.offset++
	return b, true
}

// Unread a single byte.
func (d *Decoder) ungetc(b byte) {
	if b == '\n' {
		d.line--
	}
	d.nextByte = int(b)
	d.offset--
}

func (d *Decoder) text() []byte {
	var b1 byte
	var trunc int
	d.buf.Reset()
Input:
	for {
		b, ok := d.getc()
		if !ok {
			break Input
		}
		if b == ']' || b == ',' {
			d.ungetc(b)
			break Input
		}

		// We must rewrite unescaped \r and \r\n into \n.
		if b == '\r' {
			d.buf.WriteByte('\n')
		} else if b1 == '\r' && b == '\n' {
			// Skip \r\n--we already wrote \n.
		} else {
			d.buf.WriteByte(b)
		}

		b1 = b
	}
	data := d.buf.Bytes()
	data = data[0 : len(data)-trunc]

	// Inspect each rune for being a disallowed character.
	buf := data
	for len(buf) > 0 {
		r, size := utf8.DecodeRune(buf)
		if r == utf8.RuneError && size == 1 {
			d.err = d.syntaxError("invalid UTF-8")
			return nil
		}
		buf = buf[size:]
	}

	return data
}