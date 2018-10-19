package gen_test

import (
	"testing"

	"github.com/rerorero/gopter/gen"
)

func TestBool(t *testing.T) {
	commonGeneratorTest(t, "bool", gen.Bool(), func(value interface{}) bool {
		_, ok := value.(bool)
		return ok
	})
}
