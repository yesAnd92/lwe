package pdf

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"path/filepath"
	"testing"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

var indir = "testdata"
var outdir = "testdata/out"

func TestMergeCreateNew(t *testing.T) {
	msg := "TestMergeCreate"
	inFiles := []string{
		filepath.Join(indir, "bookletTest.pdf"),
		filepath.Join(indir, "Paclitaxel.PDF"),
		filepath.Join(indir, "blank-scan.pdf"),
		filepath.Join(indir, "bookletTest.pdf"),
	}
	outFile := filepath.Join(outdir, "out2.pdf")

	if err := api.MergeCreateFile(inFiles, outFile, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}

}

func TestValidate(t *testing.T) {
	msg := "TestValidate"
	inFiles := []string{
		filepath.Join(indir, "bookletTest.pdf"),
		filepath.Join(indir, "Paclitaxel.PDF"),
		filepath.Join(indir, "blank-scan.pdf"),
		filepath.Join(indir, "mountain.jpx"),
	}

	if err := api.ValidateFiles(inFiles, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}

}

func TestSplite(t *testing.T) {
	msg := "TestSplite"
	inFile := filepath.Join(indir, "bookletTest.pdf")

	if err := api.SplitFile(inFile, outdir, 2, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}
}

func TestExtractPagesFile(t *testing.T) {

	msg := "ExtractPagesFile"
	inFile := filepath.Join(indir, "bookletTest.pdf")

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

	msg := "ImportImagesFile"
	inFiles := []string{
		filepath.Join(indir, "github.png"),
		filepath.Join(indir, "logoSmall.png"),
		filepath.Join(indir, "lwe.jpg"),
	}
	outPdf := filepath.Join(outdir, "out.pdf")

	err := api.ImportImagesFile(inFiles, outPdf, DefaultImportConfig(), nil)
	if err != nil {
		t.Fatalf("%s ExtractPages: %v\n", msg, err)
	}
}

func DefaultImportConfig() *pdfcpu.Import {
	return &pdfcpu.Import{
		PageDim:  types.PaperSize["A4"],
		PageSize: "A4",
		Pos:      types.Full,
		Scale:    0.5,
		InpUnit:  types.POINTS,
	}
}
