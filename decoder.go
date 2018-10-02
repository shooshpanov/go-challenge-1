package drum

import (
	"io"
	"os"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
func DecodeFile(path string) (*Pattern, error) {
	//file opening
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	//decoding
	p := new(Pattern)
	err = NewDecoder(f).Decode(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

type Decoder struct {
	r io.Reader
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{r: reader}
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
type Pattern struct{}
