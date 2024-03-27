package interface_mocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetUser(t *testing.T) {
	md := MockDatastore{
		Users: map[int]User{
			2: {
				ID:    2,
				First: "Jenny",
			},
		},
	}

	s := Service{
		ds: md,
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		s       Service
		args    args
		want    User
		wantErr bool
		err     error
	}{
		{
			name: "GetUser_OK",
			s:    s,
			args: args{
				id: 2,
			},
			want: User{
				ID:    2,
				First: "Jenny",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUser(tt.args.id)
			if tt.wantErr {
				assert.Equal(t, err, tt.err)
			}
			assert.Equal(t, got, tt.want)

		})
	}
}
