package repository

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBlacklist_GetBlacklists_Success(t *testing.T) {
	blacklistRepo := NewBlacklistRepo("http://www.youtube.com,http://www.facebook.com")
	got := blacklistRepo.GetBlacklists()

	expect := []string{"http://www.youtube.com", "http://www.facebook.com"}
	require.Equal(t, expect, got)
}
