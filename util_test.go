package mq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	got := RandStringBytes(6)
	assert.Equal(t, 6, len(got))

}
