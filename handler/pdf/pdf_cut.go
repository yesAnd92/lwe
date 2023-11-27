package pdf

import (
	"errors"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"strconv"
	"strings"
)

var ERROR_NOT_NUMBER = "the two sides of the '-' symbol are not numbers！"

func ParseCutArg(args string) ([]int, error) {

	var pages []int

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

		page, err := strconv.Atoi(arg)
		if err != nil {
			return nil, errors.New(ERROR_NOT_NUMBER)
		}
		pages = append(pages, page)

	}

	return pages, nil
}

func HandlePdfCut(inPdf, outDir string, selectedPages []string) error {

	err := api.ExtractPagesFile(inPdf, outDir, selectedPages, nil)
	if err != nil {
		return err
	}
	return nil
}

func populateSuccessivePages(l int, r int) (pages []int) {

	for i := l; i <= r; i++ {
		pages = append(pages, i)
	}
	return
}
