package external

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/michaelsidharta/cubic-weight/entity"
)

func TestGet(t *testing.T) {
	type args struct {
		ctx context.Context
		URL string
	}
	tests := []struct {
		name    string
		args    args
		init    func() *httptest.Server
		want    entity.APIResponse
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				URL: "",
			},
			init: func() *httptest.Server {
				resp := `{"objects":[{"category":"Gadgets","title":"10 Pack Family Car Sticker Decals","weight":120.0,"size":{"width":15.0,"length":13.0,"height":1.0}}],"next":"/api/products/2"}`
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, resp)
				}))
				return ts
			},
			want: entity.APIResponse{
				Objects: []entity.Object{
					entity.Object{
						Category: "Gadgets",
						Title:    "10 Pack Family Car Sticker Decals",
						Weight:   120.0,
						Size: entity.Size{
							Width:  15.0,
							Length: 13.0,
							Height: 1.0,
						},
					},
				},
				Next: "/api/products/2",
			},
			wantErr: false,
		}, {
			name: "error no url",
			args: args{
				ctx: context.Background(),
				URL: "",
			},
			init: func() *httptest.Server {
				return nil
			},
			want:    entity.APIResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := tt.init()
			var url string
			if server != nil {
				url = server.URL
			}
			obj := Init("")
			got, err := obj.Get(tt.args.ctx, url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
