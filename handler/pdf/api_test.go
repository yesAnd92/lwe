package pdf

import (
	"path/filepath"
	"testing"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func TestMergeCreateNew(t *testing.T) {
	dir := "/Users/wangyj/Desktop/pdf_test"
	msg := "TestMergeCreate"
	inFiles := []string{
		filepath.Join(dir, "1.pdf"),
		filepath.Join(dir, "2.pdf"),
		filepath.Join(dir, "s.pdf"),
		filepath.Join(dir, "a.pdf"),
	}
	outFile := filepath.Join(dir, "out2.pdf")

	// Merge inFiles by concatenation in the order specified and write the result to outFile.
	// outFile will be overwritten.

	// Bookmarks for the merged document will be created/preserved per default (see config.yaml)

	if err := api.MergeCreateFile(inFiles, outFile, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}

	//if err := api.ValidateFile(outFile, conf); err != nil {
	//	t.Fatalf("%s: %v\n", msg, err)
	//}
}

func TestValidate(t *testing.T) {
	dir := "/Users/wangyj/Desktop/pdf_test"
	msg := "TestMergeCreate"
	inFiles := []string{
		filepath.Join(dir, "1.pdf"),
		filepath.Join(dir, "2.pdf"),
		filepath.Join(dir, "s.pdf"),
		filepath.Join(dir, "a.pdf"),
	}

	if err := api.ValidateFiles(inFiles, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}

}

func TestSplite(t *testing.T) {
	dir := "/Users/wangyj/Desktop/pdf_test"
	msg := "TestMergeCreate"
	inFile := filepath.Join(dir, "out2.pdf")
	outdir := filepath.Join(dir, "out")

	if err := api.SplitFile(inFile, outdir, 2, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}
}
func TestSplitLowLevel(t *testing.T) {
	dir := "/Users/wangyj/Desktop/pdf_test"

	msg := "TestSplitLowLevel"
	inFile := filepath.Join(dir, "out2.pdf")
	outdir := filepath.Join(dir, "out")

	// Extract a page span.
	selectedPages := []string{
		"1", "3",
	}
	err := api.ExtractPagesFile(inFile, outdir, selectedPages, nil)
	if err != nil {
		t.Fatalf("%s ExtractPages: %v\n", msg, err)
	}
}

func TestImportImagesFile(t *testing.T) {
	dir := "/Users/wangyj/Desktop/pdf_test"

	msg := "ImportImagesFile"
	inFiles := []string{
		filepath.Join(dir, "a.jpg"),
		filepath.Join(dir, "收入1.jpg"),
		filepath.Join(dir, "收入2.jpg"),
	}
	outPdf := filepath.Join(dir, "out3.pdf")

	err := api.ImportImagesFile(inFiles, outPdf, nil, nil)
	if err != nil {
		t.Fatalf("%s ExtractPages: %v\n", msg, err)
	}
}
