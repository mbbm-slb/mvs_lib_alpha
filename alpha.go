package mvslibalpha

import (
	"fmt"
	"github.com/mbbm-slb/mvs_lib_gamma"
)

func Alpha() string {
	return fmt.Sprintf("Alpha forwarding: %s", mvslibgamma.Gamma())
}