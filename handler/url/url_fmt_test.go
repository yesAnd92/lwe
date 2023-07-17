package url

import "testing"

func Test_handleUrlPathParams(t *testing.T) {

	cases := []string{
		"http://api.demo.com/api/user/getList?platform=ios&signature=bd7e1fd4e65e8199fd817006e709fb33&currentTimeMillis=1685673384000&pageNum=1",
		"http://api.demo.com/api/user/getList?platform=&a",
		"/api/user/getList?aabb=23",
		"/api/user/getList",
		"?a=111",
		"111",
		"&a=&&",
		"http://api.demo.com/api/user/getList?&platform=a",
	}

	badCase := []string{
		"http://api.demo.com/api/user/getList?platform=&a",
	}

	for _, uri := range cases {

		HandleUrlPathParams(uri)
	}

	for _, uri := range badCase {

		HandleUrlPathParams(uri)
	}
}
