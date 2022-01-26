package fflogs

import (
	"net/url"
	"testing"
)

func TestUrlEncode(t *testing.T) {
	type TestCase struct {
		Question string
		Answer   string
	}

	testCases := []TestCase{
		{"123", "123"},
		{"123 abc", "123%20abc"},
		{"1-2_3", "1-2_3"},
		{"1&2?3", "1%262%3F3"},
		{"hello.世界.com", "hello.%E4%B8%96%E7%95%8C.com"},
		{"p@th", "p%40th"},
		{"한글테스트", "%ED%95%9C%EA%B8%80%ED%85%8C%EC%8A%A4%ED%8A%B8"},
		{"륜아린", "%EB%A5%9C%EC%95%84%EB%A6%B0"},
	}

	for i, testCase := range testCases {
		predict := urlEncode(testCase.Question)

		if predict != testCase.Answer {
			t.Errorf("Testcase %d is failed.\nPredict: %s\nAnswer: %s", i, predict, testCase.Answer)
		}
	}
}

func TestBuildUrl(t *testing.T) {
	type Options struct {
		Path1  string   `path:"path1"`
		Path2  string   `path:"path2"`
		Query1 *string  `query:"query1"`
		Query2 *int     `query:"query2"`
		Query3 *float32 `query:"query2"`
	}
	type TestCase struct {
		Answer string
		Path   string
		ApiKey string
		Option Options
	}

	var q1 string = "q1"
	var q2 int = 5
	var q3 float32 = 2.22

	testCases := []TestCase{
		{
			"https://www.fflogs.com:443/v1/a/b?api_key=123",
			"/{path1}/{path2}", "123", Options{Path1: "a", Path2: "b"},
		},
		{
			"https://www.fflogs.com:443/v1/a/b?api_key=42klsadjflkfds&query1=q1",
			"/{path1}/{path2}", "42klsadjflkfds", Options{Path1: "a", Path2: "b", Query1: &q1},
		},
		{
			"https://www.fflogs.com:443/v1/a/b?api_key=245sdah&query2=5",
			"/{path1}/{path2}", "245sdah", Options{Path1: "a", Path2: "b", Query2: &q2},
		},
		{
			"https://www.fflogs.com:443/v1/a/b?api_key=sadf212sfgd&query2=2.22",
			"/{path1}/{path2}", "sadf212sfgd", Options{Path1: "a", Path2: "b", Query3: &q3},
		},
		{
			"https://www.fflogs.com:443/v1/a/b?api_key=sdfasdjheb&query1=q1&query2=2.22",
			"/{path1}/{path2}", "sdfasdjheb", Options{Path1: "a", Path2: "b", Query1: &q1, Query3: &q3},
		},
		{
			"https://www.fflogs.com:443/v1/a/b?api_key=645ssdfg&query1=q1&query2=2.22",
			"/{path1}/{path2}", "645ssdfg", Options{Path1: "a", Path2: "b", Query1: &q1, Query3: &q3},
		},
		{
			"https://www.fflogs.com:443/v1/a/b?api_key=dsfg&query2=5&query2=2.22",
			"/{path1}/{path2}", "dsfg", Options{Path1: "a", Path2: "b", Query2: &q2, Query3: &q3},
		},
		{
			"https://www.fflogs.com:443/v1/a/b?api_key=sdf3w5y&query1=q1&query2=5&query2=2.22",
			"/{path1}/{path2}", "sdf3w5y", Options{Path1: "a", Path2: "b", Query1: &q1, Query2: &q2, Query3: &q3},
		},
	}

	u, _ := url.Parse(BaseUrlDefault)

	for i, testCase := range testCases {
		predict, err := buildUrl(*u, testCase.Path, &testCase.Option, testCase.ApiKey)
		if err != nil {
			panic(err)
		}

		if predict != testCase.Answer {
			t.Errorf("Testcase %d is failed.\nPredict: %s\nAnswer: %s", i, predict, testCase.Answer)
		}
	}
}
