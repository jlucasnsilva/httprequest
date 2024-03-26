package httprequest

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type (
	testStruct struct {
		Param bool     `from:"url-param=param"`
		Query bool     `from:"url-query=query"`
		Body  testBody `from:"request-body"`
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
	obj := testStruct{}
	err := As(
		nil,
		&obj,
		WithQueryFunc(func(r *http.Request) url.Values {
			return url.Values{
				"query_a": []string{"query_a_value"},
				"query_t": []string{time.Now().Format(time.RFC3339)},
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
	)
}
