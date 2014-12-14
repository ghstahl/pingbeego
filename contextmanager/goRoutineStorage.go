package contextmanager

import (
 	"github.com/jtolds/gls"
)

var WellKnown = struct {
		ContextManager *gls.ContextManager
		BeegoContextKey gls.ContextKey
	}{
	gls.NewContextManager(),
	gls.GenSym(),
}
