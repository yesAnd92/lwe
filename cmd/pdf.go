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
	dirMode bool = false

	pdfMergeCmd = &cobra.Command{
		Use:     `pdfm`,
		Short:   `Merge PDF or images into one PDF file`,
		Long:    `Merge multiple PDF or images(png|jpg|jpeg) into one PDF file in a gaven order`,
		Example: `lwe pdfm [-m] out.pdf in1.pdf,in2.jpg,*.png,in3.pdf...`,
		Args:    cobra.MatchAll(),
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 2 {
				cobra.CheckErr("Please re-check syntax and try it again!")
			}

			outPdf := strings.TrimSpace(args[0])
			if len(outPdf) == 0 || !pdf.HasPdfExtension(outPdf) {
				cobra.CheckErr("Please ensure the output is a file with .pdf suffix!")
			}

			var infiles []string

			infiles, err := pdf.ParseMergeArg(args)
			if err != nil {
				cobra.CheckErr("Please ensure the input file is correct!")
			}

			if err, f := pdf.CheckCorrectFileExtension(infiles); f {
				cobra.CheckErr(err)
			}
			mergeErr := pdf.HandlePdfMerge(outPdf, infiles)
			if mergeErr != nil {
				go func() {
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
		Example: `lwe pdfc [-m] in.pdf outDir 2,3,5,7-9,15...`,
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
			//if os.{
			//	cobra.CheckErr("Please ensure the outDir is directory!")
			//}

			selectedPages, err := pdf.ParseCutArg(args[2])
			if err != nil {
				cobra.CheckErr("Please ensure the page Nums you input is correct!")
			}

			if f := pdf.HasPdfExtension(inPdf); f {
				cobra.CheckErr("Please ensure the inPdf you input is a PDF file!")
			}
			mergeErr := pdf.HandlePdfCut(inPdf, outDir, selectedPages)
			if mergeErr != nil {
				cobra.CheckErr(mergeErr)
			}

		},
	}
)

func init() {

	//gitCmd.PersistentFlags().BoolVarP(&detail, "detail", "d", false, "")
	//glogCmd.PersistentFlags().BoolVarP(&file, "file", "f", false, "result output to file,default value is false (meaning output to console). ")
	//glogCmd.PersistentFlags().StringVarP(&committer, "author", "a", "", "specify name of committer ")
	//glogCmd.PersistentFlags().StringVarP(&start, "start", "s", "", "specify the start of commit date. eg.'yyyy-MM-dd'")
	//glogCmd.PersistentFlags().StringVarP(&end, "end", "e", "", "specify the end of commit date. eg.'yyyy-MM-dd'")
	//glogCmd.PersistentFlags().Int16VarP(&recentN, "recentN", "n", 10, "specify the number of commit log for each git repo.")
	//
	////gcl
	//gclCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "private token")
}
