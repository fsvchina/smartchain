package authz

import (
	"testing"
	"time"


	"github.com/stretchr/testify/require"
)

func expecError(r *require.Assertions, expected string, received error) {
	if expected == "" {
		r.NoError(received)
	} else {
		r.Error(received)
		r.Contains(received.Error(), expected)
	}
}

func TestNewGrant(t *testing.T) {

	a := NewGenericAuthorization("some-type")
	var tcs = []struct {
		title     string
		a         Authorization
		blockTime time.Time
		expire    time.Time
		err       string
	}{


		{"good expire time (1)", a, time.Unix(10, 0), time.Unix(10, 1), ""},
		{"good expire time (2)", a, time.Unix(10, 0), time.Unix(11, 0), ""},
	}

	for _, tc := range tcs {
		t.Run(tc.title, func(t *testing.T) {

			_, err := NewGrant(tc.a, tc.expire)
			expecError(require.New(t), tc.err, err)
		})
	}

}
