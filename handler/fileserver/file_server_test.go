package fileserver

import "testing"

func TestServerStart(t *testing.T) {

	//go test skip this case
	t.Skip()
	rootDir := "../../testdata"
	ServerStart("", rootDir)

}
