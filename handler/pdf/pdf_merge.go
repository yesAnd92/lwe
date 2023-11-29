package pdf

import (
	"errors"
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"path/filepath"
	"strings"
)

var DEFALUT_IMPORT_CONFIG *pdfcpu.Import

func init() {
	DEFALUT_IMPORT_CONFIG = &pdfcpu.Import{
		PageDim:  types.PaperSize["A4"],
		PageSize: "A4",
		Pos:      types.Full,
		Scale:    0.5,
		InpUnit:  types.POINTS,
	}
}

func HandlePdfMerge(outPdf string, filenames []string) error {

	size := len(filenames)
	for i := 0; i < size; {
		file := filenames[i]
		var tmpImgFiles []string

		//same type file merge in one operation
		if HasImgExtension(file) {
			tmpImgFiles = append(tmpImgFiles, file)
			for j := i + 1; j < size; j++ {
				if HasImgExtension(filenames[j]) {
					tmpImgFiles = append(tmpImgFiles, filenames[j])

					continue
				}
				break
			}
			i += len(tmpImgFiles)

			err := api.ImportImagesFile(tmpImgFiles, outPdf, DEFALUT_IMPORT_CONFIG, nil)
			if err != nil {
				return err
			}
		}

		var tmpPdfFiles []string
		if HasPdfExtension(file) {
			tmpPdfFiles = append(tmpPdfFiles, file)
			for j := i + 1; j < size; j++ {
				if HasPdfExtension(filenames[j]) {
					tmpPdfFiles = append(tmpPdfFiles, filenames[j])

					continue
				}
				break
			}
			i += len(tmpPdfFiles)

			if err := api.MergeAppendFile(tmpPdfFiles, outPdf, nil); err != nil {
				return err
			}
		}

	}
	return nil
}

func ParseMergeArg(arg string) ([]string, error) {

	var mergeParams []string

	//multiple PDF file are separated by ","
	argArr := strings.Split(arg, ",")
	for i := 0; i < len(argArr); i++ {
		in := argArr[i]
		//support "*" match files
		if strings.Contains(in, "*") {
			matches, err := filepath.Glob(in)
			if err != nil {
				return nil, err
			}

			for _, match := range matches {
				mergeParams = append(mergeParams, match)
			}

			continue
		}

		mergeParams = append(mergeParams, in)

	}

	return mergeParams, nil
}

func HasPdfExtension(filename string) bool {

	return strings.HasSuffix(strings.ToLower(filename), ".pdf")
}

func HasImgExtension(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	return types.MemberOf(ext, []string{".png", ".webp", ".tif", ".tiff", ".jpg", ".jpeg"})
}

func CheckCorrectFileExtension(filenames []string) (error, bool) {

	for _, file := range filenames {
		if HasPdfExtension(file) || HasImgExtension(file) {
			continue
		}

		return errors.New(fmt.Sprintf("Not support this file : %s ", file)), false
	}

	return nil, true
}
