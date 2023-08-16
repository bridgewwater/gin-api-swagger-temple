package biz_test

import (
	"github.com/bridgewwater/gin-api-swagger-temple/api/v1/model/biz"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPostForm(t *testing.T) {
	// mock gin at package test init()
	ginEngine := basicRouter
	// mock PostForm
	tests := []struct {
		name     string
		path     string
		header   map[string]string
		body     interface{}
		respCode int
		wantErr  bool
	}{
		{
			name: "sample", // testdata/TestPostForm/sample.golden
			path: basePath + "/biz/form",
			body: biz.Biz{
				Info:   "input info here",
				Id:     "id123zqqeeadg24qasd",
				Offset: 0,
				Limit:  10,
			},
			respCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do PostForm
			recorder, _ := MockFormPost(t, ginEngine, tc.path, tc.header, tc.body)
			assert.False(t, tc.wantErr)
			if tc.wantErr {
				t.Logf("want err close check case %s", t.Name())
				return
			}
			// verify PostForm
			assert.Equal(t, tc.respCode, recorder.Code)
			g.Assert(t, t.Name(), recorder.Body.Bytes())
		})
	}
}
