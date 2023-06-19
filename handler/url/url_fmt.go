package url

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func HandleUrlPathParams(uri string) {
	URL, _ := url.Parse(uri)

	fmt.Printf("Host: %s\n", URL.Host)
	fmt.Printf("Path: %s\n", URL.Path)
	fmt.Printf("%s\n", strings.Repeat("-", len(URL.Path)+6))

	paramMap := make(map[string]string, 8)
	rawQuery := URL.RawQuery
	if len(rawQuery) <= 0 {
		return
	}

	kvStr := strings.Split(rawQuery, "&")
	var keys []string
	for _, kv := range kvStr {
		if len(kv) == 0 {
			continue
		}

		kvPair := strings.Split(kv, "=")
		if len(kvPair) != 2 {
			cobra.CheckErr(fmt.Sprintf("Can't format [%s],please check url", uri))
		}
		kStr := kvPair[0]
		paramMap[kStr] = kvPair[1]
		keys = append(keys, kStr)
	}

	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) < len(keys[j])
	})

	//根据key的最大长度来格式化输出
	maxLen := len(keys[len(keys)-1])

	for _, key := range keys {
		fmt.Printf("%-"+strconv.Itoa(maxLen)+"s\t%s\n", key, paramMap[key])
	}
}
