package gitcmd

import (
	"fmt"
	"testing"
)

func TestGetCommit(t *testing.T) {

	commitLogs, err := GetCommitLog("yesand", 20)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, log := range *commitLogs {
		fmt.Printf("-----------Hash:%s-----------\n", log.CommitHash)
		fmt.Printf("@%s  %s\n"+
			"commit msg: %s\n\n", log.Username, log.CommitAt, log.CommitMsg)
	}
}

func TestGetChangedFile(t *testing.T) {
	filenames, err := GetChangedFile("6f635d7")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, filename := range filenames {
		fmt.Println(filename)
	}
}
