package parser

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/xerrors"
)

// Flag describes conditional parameter.
type Flag struct {
	// Name of the parameter.
	Name string
	// Index represent bit index.
	Index int
}

func (f Flag) MarshalText() (text []byte, err error) {
	return []byte(f.String()), nil
}

func (f *Flag) UnmarshalText(text []byte) (err error) {
	return f.Parse(string(text))
}

func (f Flag) String() string {
	return fmt.Sprintf("%s.%d", f.Name, f.Index)
}

func (f *Flag) Parse(s string) error {
	pos := strings.Index(s, ".")
	if pos < 1 {
		return xerrors.New("bad flag")
	}
	f.Name = s[:pos]
	idx, err := strconv.Atoi(s[pos+1:])
	if err != nil {
		return xerrors.New("bad index")
	}
	f.Index = idx
	return nil
}