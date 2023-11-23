package pdf

import (
	"errors"
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"path/filepath"
	"strings"
)

func HandlePdfMerge(outPdf string, filenames []string) error {

	size := len(filenames)
	for i := 0; i < size; {
		file := filenames[i]
		var tmpImgFiles []string

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
			fmt.Println(" merge img:", tmpImgFiles)
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
			fmt.Println(" merge pdf:", tmpPdfFiles)
		}

	}

	return nil
}

func ParseMergeArg(args []string) ([]string, error) {

	var mergeParams []string

	for i := 1; i < len(args); i++ {
		arg := args[i]
		//support "*" match files
		if strings.Contains(arg, "*") {
			matches, err := filepath.Glob(arg)
			if err != nil {
				return nil, err
			}

			for _, match := range matches {
				mergeParams = append(mergeParams, match)
			}

			continue
		}

		mergeParams = append(mergeParams, arg)

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
