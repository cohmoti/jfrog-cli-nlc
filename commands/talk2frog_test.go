package commands

import (
	"github.com/jfrog/jfrog-client-go/utils/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	log.SetLogger(log.NewLogger(log.DEBUG, nil))
}

func TestSimpleDo(t *testing.T) {
	conf := &doConfiguration{
		nlCommand: "Audit the Go project at the current directory using the watch1 watch defined in Xray",
	}
	assert.Equal(t, "jfrog xr ago --watches \"watch1\"", doTranslate(conf))
}
