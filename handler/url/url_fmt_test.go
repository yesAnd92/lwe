package url

import (
	"reflect"
	"testing"
)

func TestHandleUrlPathParams(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		args    args
		want    *UrlPares
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				"http://api.demo.com/api/user/getList?a=v1&b=v2&c=1685673384000",
			},
			want: &UrlPares{
				Host: "api.demo.com",
				Path: "/api/user/getList",
				keys: []string{"a", "b", "c"},
				paramMap: map[string]string{
					"a": "v1",
					"b": "v2",
					"c": "1685673384000",
				},
			},
			wantErr: false,
		},
		{
			name: "case 2",
			args: args{
				"/api/user/getList?a=v1&b=v2&c=1685673384000",
			},
			want: &UrlPares{
				Path: "/api/user/getList",
				keys: []string{"a", "b", "c"},
				paramMap: map[string]string{
					"a": "v1",
					"b": "v2",
					"c": "1685673384000",
				},
			},
			wantErr: false,
		},
		{
			name: "case 3",
			args: args{
				"/api/user/getList",
			},
			want: &UrlPares{
				Path: "/api/user/getList",
			},
			wantErr: false,
		},
		{
			name: "case 4",
			args: args{
				"?a=111",
			},
			want: &UrlPares{
				keys: []string{"a"},
				paramMap: map[string]string{
					"a": "111",
				},
			},
			wantErr: false,
		},
		{
			name: "bad case 5",
			args: args{
				"&a=&&",
			},
			want: &UrlPares{
				Path: "&a=&&",
			},
			wantErr: false,
		},

		{
			name: "bad case 6",
			args: args{
				"111",
			},
			want: &UrlPares{
				Path: "111",
			},
			wantErr: false,
		},

		{
			name: "bad case 7",
			args: args{
				"http://api.demo.com/api/user/getList?platfor=&a",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleUrlPathParams(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleUrlPathParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleUrlPathParams() got = %v, want %v", got, tt.want)
			}
		})
	}
}
