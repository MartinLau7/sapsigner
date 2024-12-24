package trace

import (
	"io"
	"os"
)

var (
	Output = io.Writer(os.Stderr)
)
