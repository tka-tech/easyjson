// generated by gotemplate

package opt

import (
	"fmt"

	"github.com/tka-tech/easyjson/jlexer"
	"github.com/tka-tech/easyjson/jwriter"
)

// template type Optional(A)

// A 'gotemplate'-based type for providing optional semantics without using pointers.
type Uint32 struct {
	V       uint32
	Defined bool
}

// Creates an optional type with a given value.
func OUint32(v uint32) Uint32 {
	return Uint32{V: v, Defined: true}
}

// Get returns the value or given default in the case the value is undefined.
func (v Uint32) Get(deflt uint32) uint32 {
	if !v.Defined {
		return deflt
	}
	return v.V
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Uint32) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Defined {
		w.Uint32(v.V)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Uint32) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Uint32{}
	} else {
		v.V = l.Uint32()
		v.Defined = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v *Uint32) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Uint32) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}

// IsDefined returns whether the value is defined, a function is required so that it can
// be used in an interface.
func (v Uint32) IsDefined() bool {
	return v.Defined
}

// String implements a stringer interface using fmt.Sprint for the value.
func (v Uint32) String() string {
	if !v.Defined {
		return "<undefined>"
	}
	return fmt.Sprint(v.V)
}
