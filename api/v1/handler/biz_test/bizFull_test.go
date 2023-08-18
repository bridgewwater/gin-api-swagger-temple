package biz_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetHeadFull(t *testing.T) {
	// mock gin at package test init()
	ginEngine := basicRouter
	// mock GetHeadFull
	tests := []struct {
		name     string
		path     string
		header   map[string]string
		respCode int
		wantErr  bool
	}{
		{
			name: "sample",
			path: basePath + "/biz/header_full",
			header: map[string]string{
				"BIZ_FOO": "foo",
				"BIZ_BAR": "bar",
			},
			respCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do GetHeadFull
			recorder, _ := MockRequestGet(t, ginEngine, tc.path, tc.header)

			// verify GetHeadFull
			assert.False(t, tc.wantErr)
			if tc.wantErr {
				t.Logf("want err close check case %s", t.Name())
				return
			}
			// verify GetPath
			assert.Equal(t, tc.respCode, recorder.Code)
			t.Logf("resp body: %s", recorder.Body.String())
		})
	}
}

func TestGetQueryFull(t *testing.T) {
	// mock gin at package test init()
	ginEngine := basicRouter

	type query struct {
		Foo string `form:"foo" json:"foo" binding:"required"`
		Bar string `form:"bar" json:"bar" binding:"required"`
		Baz string `form:"baz" json:"baz" binding:"required"`
	}

	// mock GetQueryFull
	tests := []struct {
		name     string
		path     string
		header   map[string]string
		query    interface{}
		respCode int
		wantErr  bool
	}{
		{
			name: "sample",
			path: basePath + "/biz/query_full",
			query: query{
				Foo: "foo",
				Bar: "bar",
				Baz: "baz",
			},
			respCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// do GetQueryFull
			recorder, _ := MockJsonQueryGet(t, ginEngine, tc.path, tc.header, tc.query)

			// verify GetQueryFull
			assert.False(t, tc.wantErr)
			if tc.wantErr {
				t.Logf("want err close check case %s", t.Name())
				return
			}
			assert.Equal(t, tc.respCode, recorder.Code)
			t.Logf("resp body: %s", recorder.Body.String())
		})
	}
}

func TestPostFormFull(t *testing.T) {
	// mock gin at package test init()
	ginEngine := basicRouter
	type query struct {
		Foo string `form:"foo" json:"foo" binding:"required"`
		Bar string `form:"bar" json:"bar" binding:"required"`
		Baz string `form:"baz" json:"baz" binding:"required"`
	}
	// mock PostFormFull
	tests := []struct {
		name     string
		path     string
		header   map[string]string
		body     interface{}
		respCode int
		wantErr  bool
	}{
		{
			name:     "sample",
			path:     basePath + "/biz/form_full",
			body:     query{Foo: "foo", Bar: "bar", Baz: "baz"},
			respCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// do PostFormFull
			recorder, _ := MockFormPost(t, ginEngine, tc.path, tc.header, tc.body)

			// verify PostFormFull
			assert.False(t, tc.wantErr)
			if tc.wantErr {
				t.Logf("want err close check case %s", t.Name())
				return
			}
			assert.Equal(t, tc.respCode, recorder.Code)
			t.Logf("resp body: %s", recorder.Body.String())
		})
	}
}
