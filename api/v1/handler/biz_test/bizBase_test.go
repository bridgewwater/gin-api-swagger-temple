package biz_test

import (
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

func TestGetString(t *testing.T) {
	recorder, _ := MockRequestGet(t, basicRouter, basePath+"/biz/string", nil)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "this is biz message", recorder.Body.String())
}
