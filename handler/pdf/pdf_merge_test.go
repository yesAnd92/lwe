package pdf

import (
	"os"
	"reflect"
	"testing"
)

func TestParseMergeArg(t *testing.T) {
	type args struct {
		arg string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				arg: "testdata/*.png",
			},
			want:    []string{"testdata/github.png", "testdata/logoSmall.png", "testdata/logoVerySmall.png"},
			wantErr: false,
		},
		{
			name: "case 2",
			args: args{
				arg: "testdata/bookletTest.pdf, testdata/github.png, testdata/*.jpg",
			},
			want:    []string{"testdata/bookletTest.pdf", "testdata/github.png", "testdata/lwe.jpg"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMergeArg(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMergeArg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMergeArg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlePdfMerge(t *testing.T) {
	type args struct {
		outPdf    string
		filenames []string
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
				outPdf:    "testdata/out/out.pdf",
				filenames: []string{"testdata/lwe.jpg"},
			},
			wantErr: false},
		{
			name: "case 2",
			args: args{
				outPdf:    "testdata/out/out.pdf",
				filenames: []string{"testdata/lwe.jpg", "testdata/zineTest.pdf", "testdata/github.png"},
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				os.Remove(tt.args.outPdf)
			}()
			if err := HandlePdfMerge(tt.args.outPdf, tt.args.filenames); (err != nil) != tt.wantErr {
				t.Errorf("HandlePdfMerge() error = %v, wantErr %v", err, tt.wantErr)
			}

			//check out Pdf exist
			if _, err := os.Stat(tt.args.outPdf); os.IsNotExist(err) {
				t.Errorf("outPdf file does not exist")
			} else {
				data, _ := os.ReadFile(tt.args.outPdf)
				if len(data) == 0 {
					t.Errorf("outPdf file is empty")
				}
			}
		})
	}
}
