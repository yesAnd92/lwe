package env

import (
	"testing"
)

func TestMacEVnVariable_findEnvInfo(t *testing.T) {
	mac := &MacEVnVariable{}

	EnvCat(mac.FindEnvInfo())
}
