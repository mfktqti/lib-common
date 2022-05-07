package sms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAccount(t *testing.T) {
	account, err := GetAccount("e55f980fd79102d98a387b13eacd589d")
	if err != nil {
		t.Fatal(err)
	}
	require.NotNil(t, account)
}
