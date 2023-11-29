package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/handler/pdf"
	"os"
	"strings"
)

/**
* pdf命令相关功能
 */
var (
	merge bool = false

	pdfMergeCmd = &cobra.Command{
		Use:     `pdfm`,
		Short:   `Merge PDF or images into one PDF file`,
		Long:    `Merge multiple PDF or images(png|jpg|jpeg) into one PDF file in a gaven  `,
		Example: `lwe pdfm out.pdf in1.pdf,in2.jpg,*.png,in3.pdf ...`,
		Args:    cobra.MatchAll(),
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 2 {
				cobra.CheckErr("Please re-check syntax and try it again!")
			}

			outPdf := strings.TrimSpace(args[0])
			if len(outPdf) == 0 || !pdf.HasPdfExtension(outPdf) {
				cobra.CheckErr("Please ensure the output is a file with .pdf suffix!")
			}

			var infiles []string

			infiles, err := pdf.ParseMergeArg(args[1])
			if err != nil {
				cobra.CheckErr("Please ensure the input file is correct!")
			}

			if err, f := pdf.CheckCorrectFileExtension(infiles); !f {
				cobra.CheckErr(err)
			}

			mergeErr := pdf.HandlePdfMerge(outPdf, infiles)
			if mergeErr != nil {
				defer func() {
					//remove out.pdf when error occurs
					os.Remove(outPdf)
				}()
				cobra.CheckErr(mergeErr)
			}

		},
	}

	pdfCutCmd = &cobra.Command{
		Use:     `pdfc`,
		Short:   `Extract selected pages from pdf files`,
		Long:    `Extract selected pages from PDF files into single page PDFs,or merge PDFs into one page PDF`,
		Example: `lwe pdfc [-m] in.pdf outDir 2,3,5,7-9,15 ...`,
		Args:    cobra.MatchAll(),
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 3 {
				cobra.CheckErr("Please re-check syntax and try it again!")
			}

			inPdf := args[0]
			if len(inPdf) == 0 || !pdf.HasPdfExtension(inPdf) {
				cobra.CheckErr("Please ensure the input is a file with .pdf suffix!")
			}

			outDir := args[1]

			selectedPages, err := pdf.ParseCutArg(args[2])
			if err != nil {
				cobra.CheckErr("Please ensure the page Nums you input is correct!")
			}

			mergeErr := pdf.HandlePdfCut(inPdf, outDir, selectedPages, merge)
			if mergeErr != nil {
				cobra.CheckErr(mergeErr)
			}

		},
	}
)

func init() {

	pdfCutCmd.PersistentFlags().BoolVarP(&merge, "merge", "m", false, "merge all selected pages into one PDF,default is false")

}
