package fileserver

import "testing"

func TestServerStart(t *testing.T) {

	rootDir := "./testdata"
	ServerStart("", rootDir)

}
