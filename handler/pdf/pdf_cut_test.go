package pdf

import (
	"reflect"
	"testing"
)

func TestParseCutArg(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "case 1",
			args:    "1, 2, 4-9,20",
			want:    []string{"1", "2", "4", "5", "6", "7", "8", "9", "20"},
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
			want:    []string{"1"},
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
		merger        bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				inPdf:         "testdata/zineTest.pdf",
				outDir:        "testdata/out/",
				selectedPages: []string{"1", "3", "4", "6", "20"},
			},
			wantErr: false,
		},
		{
			name: "no exist outDir",
			args: args{
				inPdf:         "testdata/zineTest.pdf",
				outDir:        "testdata/out/1/",
				selectedPages: []string{"1", "1", "1"},
			},
			wantErr: true,
		},
		{
			name: " wrong outDir",
			args: args{
				inPdf:         "testdata/zineTest.pdf",
				outDir:        "testdata/out/out.pdf",
				selectedPages: []string{"1", "1", "1"},
			},
			wantErr: true,
		},
		{
			name: "merge all pdf",
			args: args{
				inPdf:         "testdata/zineTest.pdf",
				outDir:        "testdata/out/",
				selectedPages: []string{"1", "3", "5"},
				merger:        true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := HandlePdfCut(tt.args.inPdf, tt.args.outDir, tt.args.selectedPages, tt.args.merger); (err != nil) != tt.wantErr {
				t.Errorf("HandlePdfCut() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
