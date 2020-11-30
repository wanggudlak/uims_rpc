package randc_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"uims/pkg/randc"
)

func TestUUID(t *testing.T) {
	uuid := randc.UUID()
	t.Log(uuid)
	assert.Equal(t, 32, len(uuid))
}
