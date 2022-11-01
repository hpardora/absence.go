package absence

import (
	"testing"
)

func TestConfig(t *testing.T) {
	path := "/tmp/absence.yaml"
	NewFromPath(path)
}
