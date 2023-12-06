package url

import (
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type UrlPares struct {
	Host, Path string
	keys       []string
	paramMap   map[string]string
}

func HandleUrlPathParams(uri string) (*UrlPares, error) {
	URL, parseErr := url.Parse(uri)
	if parseErr != nil {
		return nil, parseErr
	}

	//paramMap := make(map[string]string, 8)

	rawQuery := URL.RawQuery

	kvStr := strings.Split(rawQuery, "&")
	var keys []string
	var paramMap map[string]string
	for _, kv := range kvStr {
		if len(kv) == 0 {
			continue
		}

		kvPair := strings.Split(kv, "=")
		//left of '=' require not empty
		if len(kvPair) == 1 {
			return nil, errors.New(fmt.Sprintf("Can't format [%s],please check url", uri))
		}
		kStr := kvPair[0]
		//For the convenience of testing
		if paramMap == nil {
			paramMap = make(map[string]string, 8)
		}
		paramMap[kStr] = kvPair[1]
		keys = append(keys, kStr)
	}

	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) < len(keys[j])
	})

	return &UrlPares{
		Host:     URL.Host,
		Path:     URL.Path,
		keys:     keys,
		paramMap: paramMap,
	}, nil
}

func FmtPrint(pares *UrlPares) {

	fmt.Printf("Host: %s\n", pares.Host)
	fmt.Printf("Path: %s\n", pares.Path)
	fmt.Printf("%s\n", strings.Repeat("-", len(pares.Path)+6))

	//根据key的最大长度来格式化输出
	maxLen := len(pares.keys[len(pares.keys)-1])

	for _, key := range pares.keys {
		fmt.Printf("%-"+strconv.Itoa(maxLen)+"s\t%s\n", key, pares.paramMap[key])
	}
}
