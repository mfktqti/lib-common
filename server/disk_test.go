package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisk(t *testing.T) {
	s, err := GetServerInfo()
	if err != nil {
		t.Fatal(err)
	}
	assert.LessOrEqual(t, s.Disk.UsedPercent, 80)
}
