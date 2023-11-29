package cmd

import (
	"fmt"
	"testing"
)

func TestPdfMergeCmd(t *testing.T) {

	output, err := executeCommand(pdfMergeCmd, "pdfm", "../handler/pdf/testdata/out/ooo.pdf", "../handler/pdf/testdata/*.pdf")
	if err != nil {
		t.Fail()
	}
	fmt.Println(output.String())
}

func TestPdfCutCmd(t *testing.T) {

	output, err := executeCommand(pdfCutCmd, "pdfc", "../handler/pdf/testdata/bookletTest.pdf", "../handler/pdf/testdata/out/", "1,3,4,,7,9-11")
	if err != nil {
		t.Fail()
	}
	fmt.Println(output.String())
}
