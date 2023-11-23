package pdf

import (
	"path/filepath"
	"strings"
)

func HandlePdfMerge(filenames []string) error {

	return nil
}

func parseMergeArg(args []string) ([]string, error) {

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
