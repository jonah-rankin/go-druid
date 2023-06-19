package granularity

import (
	"github.com/jonah-rankin/go-druid/builder/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSimple(t *testing.T) {
	expected := []byte(`"all"`)
	s := NewSimple()
	s.SetGranularity("all")

	built, err := Load(expected)
	assert.Nil(t, err)

	testutil.Compare(t, expected, s, built)
}
