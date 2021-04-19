package domain

import "testing"

func TestShortCode_Validate(t *testing.T) {
	tests := []struct {
		name    string
		s       ShortCode
		wantErr bool
	}{
		{"1", "123456", false},
		{"2", "12345a", false},
		{"3", "+2345a", true},
		{"4", "1234567", true},
		{"5", "1", true},
		{"6", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
