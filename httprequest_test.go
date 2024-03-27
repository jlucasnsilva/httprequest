package httprequest

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type (
	testStruct struct {
		ID      int64     `from:"url-param=id"`
		Flag    bool      `from:"url-param=flag"`
		QueryA  string    `from:"url-query=query_a"`
		QueryT  time.Time `from:"url-query=query_t,layout=DateTime"`
		Body    testBody  `from:"request-body"`
		Ignored int       `from:"-"`
	}

	testPStruct struct {
		ID      int64     `from:"url-param=id"`
		Flag    bool      `from:"url-param=flag"`
		QueryA  string    `from:"url-query=query_a"`
		QueryT  time.Time `from:"url-query=query_t,layout=DateTime"`
		Body    *testBody `from:"request-body"`
		Ignored int       `from:"-"`
	}

	testBody struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
)

func TestOptions(t *testing.T) {
	cfg := config{}

	WithQueryFunc(defaultCfg.Query)(&cfg)
	WithURLParamFunc(defaultCfg.Param)(&cfg)
	WithUnmarshaller(defaultCfg.Unmarshal)(&cfg)

	assert.NotNil(t, cfg.Query)
	assert.NotNil(t, cfg.Param)
	assert.NotNil(t, cfg.Unmarshal)
}

func TestAs(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		now := time.Now()
		nowFmt := now.Format(time.DateTime)
		now, timeErr := time.Parse(time.DateTime, nowFmt)

		require.Nil(t, timeErr)

		expectedBody := testBody{
			ID:   834,
			Name: "Hello, world",
		}
		jsonObj, jsonErr := json.Marshal(&expectedBody)

		req, reqErr := http.NewRequest("GET", "/hello/world", nil)

		require.Nil(t, jsonErr)
		require.Nil(t, reqErr)

		expected := testStruct{
			ID:     10,
			Flag:   true,
			QueryA: "query_a_value",
			QueryT: now,
			Body:   expectedBody,
		}

		obj := testStruct{}
		err := As(
			req,
			&obj,
			WithQueryFunc(func(r *http.Request) url.Values {
				return url.Values{
					"query_a": []string{"query_a_value"},
					"query_t": []string{nowFmt},
				}
			}),
			WithURLParamFunc(func(r *http.Request, key string) string {
				if key == "id" {
					return "10"
				}
				if key == "flag" {
					return "true"
				}
				return ""
			}),
			WithUnmarshaller(func(r *http.Request, p any) error {
				return json.Unmarshal(jsonObj, p)
			}),
		)

		assert.Equal(t, expected.ID, obj.ID)
		assert.Equal(t, expected.Flag, obj.Flag)
		assert.Equal(t, expected.QueryA, obj.QueryA)
		assert.Equal(t, expected.QueryT, obj.QueryT)
		assert.Equal(t, expected.Body, obj.Body)
		assert.Nil(t, err)
	})

	t.Run("should succeed", func(t *testing.T) {
		now := time.Now()
		nowFmt := now.Format(time.DateTime)
		now, timeErr := time.Parse(time.DateTime, nowFmt)

		require.Nil(t, timeErr)

		expectedBody := &testBody{
			ID:   834,
			Name: "Hello, world",
		}
		jsonObj, jsonErr := json.Marshal(&expectedBody)

		req, reqErr := http.NewRequest("GET", "/hello/world", nil)

		require.Nil(t, jsonErr)
		require.Nil(t, reqErr)

		expected := testPStruct{
			ID:     10,
			Flag:   true,
			QueryA: "query_a_value",
			QueryT: now,
			Body:   expectedBody,
		}

		obj := testPStruct{}
		err := As(
			req,
			&obj,
			WithQueryFunc(func(r *http.Request) url.Values {
				return url.Values{
					"query_a": []string{"query_a_value"},
					"query_t": []string{nowFmt},
				}
			}),
			WithURLParamFunc(func(r *http.Request, key string) string {
				if key == "id" {
					return "10"
				}
				if key == "flag" {
					return "true"
				}
				return ""
			}),
			WithUnmarshaller(func(r *http.Request, p any) error {
				return json.Unmarshal(jsonObj, p)
			}),
		)

		assert.Equal(t, expected.ID, obj.ID)
		assert.Equal(t, expected.Flag, obj.Flag)
		assert.Equal(t, expected.QueryA, obj.QueryA)
		assert.Equal(t, expected.QueryT, obj.QueryT)
		assert.Equal(t, expected.Body, obj.Body)
		assert.Nil(t, err)
	})
}
