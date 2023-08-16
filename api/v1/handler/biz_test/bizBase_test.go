package biz_test

import (
	"fmt"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetPath(t *testing.T) {
	// mock gin at package test init()
	ginEngine := basicRouter
	// mock GetPath
	tests := []struct {
		name     string
		path     string
		header   map[string]string
		respCode int
		wantErr  bool
	}{
		{
			name:     "sample 123",
			path:     basePath + "/biz/path/123",
			respCode: http.StatusOK,
		},
		{
			name:     "sample 567",
			path:     basePath + "/biz/path/567",
			respCode: http.StatusOK,
		},
		{
			name:     "StatusNotFound",
			path:     basePath + "/biz/path/",
			respCode: http.StatusNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do GetPath
			recorder, _ := MockRequestGet(t, ginEngine, tc.path, tc.header)
			assert.False(t, tc.wantErr)
			if tc.wantErr {
				t.Logf("want err close check case %s", t.Name())
				return
			}
			// verify GetPath
			assert.Equal(t, tc.respCode, recorder.Code)
			g.Assert(t, t.Name(), recorder.Body.Bytes())
		})
	}
}

func TestGetQuery(t *testing.T) {
	// mock gin at package test init()
	ginEngine := basicRouter

	type query struct {
		Offset string `form:"offset" json:"offset"`
		Limit  string `form:"limit" json:"limit"`
	}

	// mock GetQuery
	tests := []struct {
		name     string
		path     string
		query    any
		header   map[string]string
		respCode int
		wantErr  bool
	}{
		{
			name: "sample", // testdata/TestGetQuery/sample.golden
			path: basePath + "/biz/query",
			query: query{
				Offset: "0",
				Limit:  "10",
			},
			respCode: http.StatusOK,
		},
		{
			name: "fail offset", // testdata/TestGetQuery/sample.golden
			path: basePath + "/biz/query",
			query: query{
				Offset: "a",
				Limit:  "10",
			},
			respCode: http.StatusBadRequest,
		},
		{
			name: "fail limit", // testdata/TestGetQuery/sample.golden
			path: basePath + "/biz/query",
			query: query{
				Offset: "0",
				Limit:  "abc",
			},
			respCode: http.StatusBadRequest,
		},
		{
			name: "fail not exist url", // testdata/TestGetQuery/sample.golden
			path: basePath + "/biz/query/",
			query: query{
				Offset: "0",
				Limit:  "10",
			},
			respCode: http.StatusMovedPermanently,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do GetQuery
			realUrl := fmt.Sprintf("%s?%s", tc.path, MockQueryStrFrom(tc.query))
			recorder, _ := MockRequestGet(t, ginEngine, realUrl, nil)
			assert.False(t, tc.wantErr)
			if tc.wantErr {
				t.Logf("want err close check case %s", t.Name())
				return
			}
			// verify GetQuery
			assert.Equal(t, tc.respCode, recorder.Code)
			g.Assert(t, t.Name(), recorder.Body.Bytes())
		})
	}
}

func TestGetString(t *testing.T) {
	recorder, _ := MockRequestGet(t, basicRouter, basePath+"/biz/string", nil)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "this is biz message", recorder.Body.String())
}
