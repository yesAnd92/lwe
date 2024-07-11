package fileserver

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/utils"
	"net/http"
	"strconv"
	"sync"
)

var (
	defaultIP         = "127.0.0.1"
	defaultPort       = 9527
	MaxPort           = 9999
	fileAccessCounter = make(map[string]int)
	lock              sync.Mutex
)

// add visit count
func addAccessCount(filePath string) int {
	lock.Lock()
	defer lock.Unlock()

	count := 0
	if _, ok := fileAccessCounter[filePath]; ok {
		count = fileAccessCounter[filePath]
	}
	count++
	fileAccessCounter[filePath] = count
	return count
}

func ServerStart(port, rootDir string) {

	if len(port) == 0 {
		port = strconv.Itoa(havaAvailablePort())
	}

	rootDir = utils.ToAbsPath(rootDir)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		filePath := r.URL.Path

		//file access statistic
		count := addAccessCount(filePath)

		fmt.Printf("%s  - %d visit\n", filePath, count)

		http.ServeFile(w, r, rootDir+filePath)

	})

	// listen local file server
	fmt.Printf("%s/ ==>  http://%s/\n", rootDir, fmt.Sprintf("%s:%s", defaultIP, port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		cobra.CheckErr(err)
	}
}

func havaAvailablePort() int {

	for defaultPort <= MaxPort && !utils.PortAvailable(defaultPort) {
		defaultPort++
	}
	return defaultPort
}
