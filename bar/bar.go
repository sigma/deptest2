package bar

import "github.com/sigma/deptest/sub"

func Bar() string {
	return sub.Version
}
