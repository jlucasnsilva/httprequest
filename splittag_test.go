package httprequest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	splitTagTest struct {
		Label          string
		Meta           string
		ExpectedKind   string
		ExpectedSource string
		ExpectedMeta   map[string]string
		ExpectedError  error
	}
)

func TestSplitTag(t *testing.T) {
	testTable := []splitTagTest{
		{
			Label:          "should fail because it is a invalid tag",
			ExpectedKind:   "",
			ExpectedSource: "",
			ExpectedError:  ErrInvalidParamTag,
		},
		{
			Label:          "should fail because it is a blank kind",
			ExpectedKind:   "        ",
			ExpectedSource: "",
			ExpectedError:  ErrInvalidParamTag,
		},
		{
			Label:        "should succeed with request-body kind",
			ExpectedKind: requestBodyTag,
		},
		{
			Label:         "should fail with invalid KV tag",
			ExpectedKind:  urlParamTag,
			ExpectedError: ErrInvalidParamTagKeyValue,
		},
		{
			Label:          "should succeed with valid KV tag",
			ExpectedKind:   urlParamTag,
			ExpectedSource: "field",
		},
		{
			Label:          "should succeed with valid KV tag",
			ExpectedKind:   urlQueryTag,
			ExpectedSource: "query",
		},
		{
			Label:          "should succeed with valid KV tag and valid meta",
			ExpectedKind:   urlParamTag,
			ExpectedSource: "field",
			ExpectedMeta: map[string]string{
				"hello": "world",
				"where": "somewhere",
			},
		},
		{
			Label:          "should fail with valid KV tag and invalid meta",
			ExpectedError:  ErrInvalidParamTagKeyValue,
			ExpectedKind:   urlParamTag,
			ExpectedSource: "field",
			ExpectedMeta: map[string]string{
				"hello": "",
				"where": "somewhere",
			},
		},
		{
			Label:          "should fail with valid KV tag and blank meta",
			ExpectedError:  ErrInvalidParamTagKeyValue,
			ExpectedKind:   urlParamTag,
			ExpectedSource: "field",
			ExpectedMeta: map[string]string{
				"hello": "           ",
				"where": "somewhere",
			},
		},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.Label, func(t *testing.T) {
			tag := makeTag(test.ExpectedKind, test.ExpectedSource, test.ExpectedMeta)
			kind, source, meta, err := splitTag(tag)

			if test.ExpectedError != nil {
				assert.Equal(t, test.ExpectedError, err)
			} else {
				assert.Equal(t, test.ExpectedKind, kind)
				assert.Equal(t, test.ExpectedSource, source)
				assert.Equal(t, test.ExpectedMeta, meta)
				assert.Nil(t, err)
			}
		})
	}
}

func makeTag(kind, source string, meta map[string]string) string {
	m := ""
	parts := make([]string, 0, len(meta))
	for k, v := range meta {
		parts = append(parts, k+"="+v)
	}
	if len(parts) > 0 {
		m = "," + strings.Join(parts, ",")
	}

	switch {
	case kind != "" && source == "":
		return kind + m
	case kind != "" && source != "":
		return kind + "=" + source + m
	default:
		return ""
	}
}
