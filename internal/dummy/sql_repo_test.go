package dummy

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"

	"ws-dummy-go/internal/dummy/domain"
)

var (
	testPostgresPool *pgxpool.Pool
)

func Test_usersSQLRepo_Insert(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.UserID
		wantErr bool
	}{
		{
			name:    "Positive: Insert user",
			args:    args{name: "testname3843"},
			want:    domain.UserID("1"),
			wantErr: false,
		},
	}

	r := usersSQLRepo{
		pool: testPostgresPool,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert := assert.New(t)

			got, err := r.Insert(context.Background(), tt.args.name)

			assert.Equal(tt.want, got)
			assert.Equal(tt.wantErr, err != nil, err)
		})
	}
}

func Test_usersSQLRepo_Update(t *testing.T) {
	type args struct {
		id   domain.UserID
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Positive: Update user",
			args:    args{id: domain.UserID("1"), name: "testname123"},
			wantErr: false,
		},
	}

	r := usersSQLRepo{
		pool: testPostgresPool,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert := assert.New(t)

			err := r.Update(context.Background(), tt.args.id, tt.args.name)

			assert.Equal(tt.wantErr, err != nil, err)
		})
	}
}
