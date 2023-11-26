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
		Short:   `Get all git repository commit log under the given dir `,
		Long:    `Get all git repository commit log under the given dir ,and  specify author，date etc. supported!`,
		Example: `lwe pdfm [-m] out.pdf in1.pdf,in2.jpg,in3.pdf...`,
		Args:    cobra.MatchAll(),
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 2 {
				cobra.CheckErr("Please re-check syntax and try it again!")
			}

			outPdf := strings.TrimSpace(args[0])
			if len(outPdf) == 0 || !pdf.HasPdfExtension(outPdf) {
				cobra.CheckErr("Please ensure the output is a file with a. pdf suffix!")
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
