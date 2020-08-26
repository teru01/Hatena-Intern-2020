package fetcher

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TC struct {
	in       string
	out      string
	endpoint string
}

func TestExtractTitle(t *testing.T) {
	testCases := []TC{
		TC{
			in: `<title>  hello</title>
`,
			out: "  hello",
		},
		TC{
			in: `
<head>
	<title>こんにちは世界</title>
</head>
`,
			out: "こんにちは世界",
		},
		TC{
			in: `
<HEAD>
<TITLE>こんにちは世界</TITLE>
</HEAD>
`,
			out: "こんにちは世界",
		},
	}

	errTestCases := []TC{
		TC{
			in: `<html></html>
`,
			out: "",
		},
		TC{
			in: `
<head>
<title></title>
</head>
`,
			out: "",
		},
		TC{
			in: `
<head>
<title>  </title>
</head>
`,
			out: "hello",
		},
	}
	for _, testCase := range testCases {
		title, err := extractTitle(strings.NewReader(testCase.in))
		assert.NoError(t, err)
		assert.Equal(t, testCase.out, title)
	}

	for _, testCase := range errTestCases {
		_, err := extractTitle(strings.NewReader(testCase.in))
		assert.Error(t, err)
	}
}

func TestFetchTitle(t *testing.T) {
	testCases := []TC{
		TC{
			in: `<html>
<head>
<title>hello, world</title>
</head>
<body>
title fetcher
hello world
</body>
</html>
`,
			out:      "hello, world",
			endpoint: "/success",
		},
	}

	mux := http.NewServeMux()
	for _, testCase := range testCases {
		mux.HandleFunc(testCase.endpoint, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, testCase.in)
		})
	}
	ts := httptest.NewServer(mux)
	defer ts.Close()

	for _, testCase := range testCases {
		title, err := Fetch(context.Background(), ts.URL+testCase.endpoint, ts.Client())
		assert.NoError(t, err)
		assert.Equal(t, testCase.out, title)
	}
}

func TestFetchTitleTimeout(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 3010)
	})
	ts := httptest.NewServer(mux)
	defer ts.Close()

	_, err := Fetch(context.Background(), ts.URL+"/timeout", ts.Client())
	assert.Error(t, err)
}
