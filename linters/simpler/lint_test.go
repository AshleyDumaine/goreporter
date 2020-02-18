package simpler

import (
	"testing"

	"github.com/AshleyDumaine/goreporter/linters/simpler/lint/testutil"
)

func TestAll(t *testing.T) {
	testutil.TestAll(t, NewChecker(), "")
}
