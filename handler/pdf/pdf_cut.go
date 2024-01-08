package pdf

import (
	"errors"
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/yesAnd92/lwe/utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var ERROR_NOT_NUMBER = "the two sides of the '-' symbol are not numbers！"

func ParseCutArg(args string) ([]string, error) {

	var pages []string

	//multiple page are separated by ","
	argArr := strings.Split(strings.ReplaceAll(args, " ", ""), ",")

	for _, arg := range argArr {
		if len(arg) == 0 {
			continue
		}
		//parse range pages by "-"
		if strings.Contains(arg, "-") {
			var l, r int
			var el, er error
			s := strings.Split(arg, "-")
			l, el = strconv.Atoi(s[0])
			r, er = strconv.Atoi(s[1])
			if el != nil || er != nil {
				return nil, errors.New(ERROR_NOT_NUMBER)
			}
			pages = append(pages, populateSuccessivePages(l, r)...)
			continue
		}

		if !isNumeric(arg) {
			return nil, errors.New(ERROR_NOT_NUMBER)
		}
		pages = append(pages, arg)

	}

	return pages, nil
}

func HandlePdfCut(inPdf, outDir string, selectedPages []string, merge bool) error {

	utils.MkdirIfNotExist(outDir)

	err := api.ExtractPagesFile(inPdf, outDir, selectedPages, nil)
	if err != nil {
		return err
	}

	if merge {

		originName := strings.TrimSuffix(filepath.Base(inPdf), ".pdf")

		//find all generated pdf by rule('{originName}_page_*.pdf')
		//but it may contain the  pdf which has existed under dir before
		matches, matchErr := filepath.Glob(filepath.Join(outDir, fmt.Sprintf("%s_page_*.pdf", originName)))
		if matchErr != nil {
			return matchErr
		}

		defer func() {
			for _, p := range matches {
				os.Remove(p)
			}
		}()

		outMergeFile := filepath.Join(outDir, fmt.Sprintf("%s_selected.pdf", originName))

		//merge single pdf but this ignores the order that user input
		mergeErr := HandlePdfMerge(outMergeFile, matches)
		if mergeErr != nil {
			return err
		}
	}

	fmt.Println("PDF cut result >> " + utils.ToAbsPath(outDir))

	return nil
}

func populateSuccessivePages(l int, r int) (pages []string) {

	for i := l; i <= r; i++ {
		pages = append(pages, strconv.Itoa(i))
	}
	return
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
