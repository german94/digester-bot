package user_usage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUserUsage(t *testing.T) {
	uu := New(2)
	require.False(t, uu.HasReachedLimit())

	uu.Increment("req 1")
	require.False(t, uu.HasReachedLimit())
	uu.Increment("req 2")
	require.True(t, uu.HasReachedLimit())

	uu.requests[0].At = uu.requests[0].At.Add(-30 * time.Hour)
	require.False(t, uu.HasReachedLimit())

	uu.Increment("req 3")
	require.True(t, uu.HasReachedLimit())
}
