package eval

import (
	"bytes"
	"fmt"

	"github.com/elves/elvish/diag"
)

// CompilationError represents a compilation error and can pretty print it.
type CompilationError struct {
	Message string
	Context diag.SourceRange
}

func (ce *CompilationError) Error() string {
	return fmt.Sprintf("compilation error: %d-%d in %s: %s",
		ce.Context.Begin, ce.Context.End, ce.Context.Name, ce.Message)
}

// Range returns the range of the compilation error.
func (ce *CompilationError) Range() diag.Ranging {
	return ce.Context.Range()
}

// PPrint pretty-prints a compilation error.
func (ce *CompilationError) PPrint(indent string) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "Compilation error: \033[31;1m%s\033[m\n", ce.Message)
	buf.WriteString(ce.Context.PPrintCompact(indent + "  "))

	return buf.String()
}
