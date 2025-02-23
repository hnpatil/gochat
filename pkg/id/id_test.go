package id

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		ids  []string
		want ID
	}{
		{
			name: "multiple ids",
			ids:  []string{"1", "2", "3"},
			want: ID("ba101947-992b-d0d9-51ee3d37"),
		},
		{
			name: "multiple ids",
			ids:  []string{"2", "3", "4"},
			want: ID("85f5fc6c-6d4d-5ed3-5422d801"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.ids...)
			assert.Equal(t, tt.want, got)
		})
	}
}
