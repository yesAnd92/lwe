package sync

import (
	"reflect"
	"testing"
)

func Test_findAllFile(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{name: "", args: args{dir: "E:"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllFile(tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
