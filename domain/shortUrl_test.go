package domain

import (
	"github.com/nqmt/gotime/v2"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestShortUrl_ValidateUrl(t *testing.T) {
	type fields struct {
		originUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"1", fields{originUrl: "http://www.google.com"}, false},
		{"2", fields{originUrl: "https://www.google.com"}, false},
		{"3", fields{originUrl: "https://google.com"}, false},
		{"4", fields{originUrl: "https://dd.google.dd"}, false},
		{"5", fields{originUrl: "www.google.com"}, true},
		{"6", fields{originUrl: "ftp://www.google.com"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ShortUrl{
				originUrl: tt.fields.originUrl,
			}
			if err := s.ValidateUrl(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShortUrl_IsBlackList(t *testing.T) {
	type fields struct {
		originUrl string
	}
	type args struct {
		blacklists []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "1",
			fields: fields{originUrl: "found"},
			args:   args{blacklists: []string{"found"}},
			want:   true,
		},
		{
			name:   "2",
			fields: fields{originUrl: "found1"},
			args:   args{blacklists: []string{"found"}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ShortUrl{
				originUrl: tt.fields.originUrl,
			}
			if got := s.IsBlackList(tt.args.blacklists); got != tt.want {
				t.Errorf("IsBlackList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortUrl_GenShortUrl(t *testing.T) {
	freeze := time.Date(2020, 4, 18, 17, 17, 17, 123456788, time.UTC)
	gotime.Freeze(freeze)

	// 1587230237123456788_url1
	shortUrl := NewShortUrl("url1", 0)
	require.Equal(t, "b3df06", shortUrl.GenShortUrl())

	freeze = time.Date(2020, 4, 18, 17, 17, 17, 123456789, time.UTC)
	gotime.Freeze(freeze)

	// 1587230237123456789_url2
	shortUrl = NewShortUrl("url2", 0)
	require.Equal(t, "8cde9f", shortUrl.GenShortUrl())
}

func TestShortUrl_GetDefaultExpireAt(t *testing.T) {
	shortUrl := NewShortUrl("url1", 0)
	require.Equal(t, 30, shortUrl.GetDefaultExpireAt())

	shortUrl = NewShortUrl("url1", 1)
	require.Equal(t, 1, shortUrl.GetDefaultExpireAt())
}

func TestShortUrl_GetExpireAt(t *testing.T) {
	freeze := time.Date(2020, 4, 18, 17, 17, 17, 123456789, time.UTC)
	gotime.Freeze(freeze)

	shortUrl := NewShortUrl("url1", 0)
	require.Equal(t, int64(1589822237), shortUrl.GetExpireAt())
}
