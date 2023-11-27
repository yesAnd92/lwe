package pdf

import (
	"reflect"
	"testing"
)

func TestParseCutArg(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		want    []int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "case 1",
			args:    "1, 2, 4-9,20",
			want:    []int{1, 2, 4, 5, 6, 7, 8, 9, 20},
			wantErr: false,
		},

		{
			name:    "case 2",
			args:    "1, 2, a-9,20",
			want:    nil,
			wantErr: true,
		},

		{
			name:    "case 3",
			args:    "1, ,,,,",
			want:    []int{1},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCutArg(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCutArg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCutArg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlePdfCut(t *testing.T) {
	type args struct {
		inPdf         string
		outDir        string
		selectedPages []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := HandlePdfCut(tt.args.inPdf, tt.args.outDir, tt.args.selectedPages); (err != nil) != tt.wantErr {
				t.Errorf("HandlePdfCut() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
