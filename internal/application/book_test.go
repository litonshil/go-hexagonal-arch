package app_test

import (
	app "hexagonal-arch/internal/application"
	db "hexagonal-arch/internal/framework/mysql"
	"hexagonal-arch/internal/ports/types"
	"testing"

	. "github.com/stretchr/testify/require"
)

func TestAdapter_CreateBook(t *testing.T) {
	dbAdapter, err := db.NewAdapterForTest(t)
	Nil(t, err)

	type args struct {
		bookReq types.BookReqBody
	}
	tests := []struct {
		name      string
		args      args
		want      types.BookResp
		wantError bool
	}{
		{
			name: "create book",
			args: args{
				bookReq: types.BookReqBody{
					Code:  "B-1",
					Title: "Book 1",
				},
			},
			want: types.BookResp{
				BookReqBody: types.BookReqBody{
					Code: "B-1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := app.NewApplication(dbAdapter)
			got, err := a.CreateBook(tt.args.bookReq)
			if tt.wantError {
				NotNil(t, err)
				return
			}
			Nil(t, err)
			NotZero(t, got.ID)
			Equal(t, tt.want.Code, got.Code)
		})
	}
}
