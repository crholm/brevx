package iconv

import (
	"fmt"
	"github.com/crholm/brevx/envelope"
	"io"

	ico "gopkg.in/iconv.v1"
)

// iconv enables using GNU iconv for converting 7bit to UTF-8.
// iconv supports a larger range of encodings.
// It's a cgo package, the build system needs have Gnu library headers available.
// when importing, place an underscore _ in front to import for side-effects

func init() {
	envelope.Dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if cd, err := ico.Open("UTF-8", charset); err == nil {
			r := ico.NewReader(cd, input, 32)
			return r, nil
		}
		return nil, fmt.Errorf("unhandled charset %q", charset)
	}

}
