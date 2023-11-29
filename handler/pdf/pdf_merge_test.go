package pdf

import (
	"testing"
)

func Test_parseMergeArg(t *testing.T) {
	type args struct {
		args string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				args: "testdata/bookletTest.pdf, testdata/github.png, testdata/*.jpg",
			},
		},
		{
			name: "case 2",
			args: args{
				args: "testdata/*.pdf",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re, err := ParseMergeArg(tt.args.args)
			if err != nil {
				t.Fail()
			}
			if len(re) != 3 {
				t.Fail()
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
		{name: "case 1",
			args: args{
				outPdf:    "",
				filenames: []string{"a.pdf", "b.jpg", "c.jpg", "d.jpg", "e.pdf", "f.pdf"},
			}},
		{name: "case 2",
			args: args{
				outPdf:    "",
				filenames: []string{"a.pdf", "b.jpg"},
			}},
		{name: "case 3",
			args: args{
				outPdf:    "",
				filenames: []string{"a.pdf"},
			}},

		{name: "case 4",
			args: args{
				outPdf:    "testdata/out/out.pdf",
				filenames: []string{"testdata/lwe.jpg", "testdata/zineTest.pdf", "testdata/github.png"},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := HandlePdfMerge(tt.args.outPdf, tt.args.filenames); (err != nil) != tt.wantErr {
				t.Errorf("HandlePdfMerge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
