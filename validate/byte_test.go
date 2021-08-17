package validate

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsEmptyByte(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Ok",
			args: args{
				data: []byte{},
			},
			want: true,
		},
		{
			name: "Empty",
			args: args{
				data: []byte("Lorem Ipsum"),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEmptyByte(tt.args.data)
			require.Equal(t, got, tt.want)
		})
	}
}
