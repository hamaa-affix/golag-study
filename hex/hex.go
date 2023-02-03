package hex

import (
	"fmt"
)

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}
