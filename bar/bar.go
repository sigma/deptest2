package bar

import (
	"github.com/sigma/deptest/sub"
	_ "rsc.io/quote"
)

func Bar() string {
	return sub.Version
}
