package features

import (
	"github.com/localzet/aura/common"
)

// Feature is the interface for Aura features. All features must implement this interface.
// All existing features have an implementation in app directory. These features can be replaced by third-party ones.
type Feature interface {
	common.HasType
	common.Runnable
}
