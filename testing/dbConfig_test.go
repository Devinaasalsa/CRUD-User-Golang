package testing

import (
	"myapp-me/config"
	"testing"
)

func TestConnect(t *testing.T) {
	config.DatabaseInit()
}
