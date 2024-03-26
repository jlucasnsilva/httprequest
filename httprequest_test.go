package httprequest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	boolStruct struct {
		Param bool `url-param:"param"`
		Query bool `url-query:"query"`
	}

	intStruct struct {
		Int     int
		Int8    int8
		Int16   int16
		Int32   int32
		Int64   int64
		Uint    uint
		Uint8   uint8
		Uint16  uint16
		Uint32  uint32
		Uint64  uint64
		Float32 float32
		Float64 float64
	}
)

func TestSplitTag(t *testing.T) {
	t.Run("should fail because of invalid tag", func(t *testing.T) {
		_, _, err := splitTag("")
		assert.Equal(t, ErrInvalidParamTag, err)
	})
	t.Run("should fail because of blank tag", func(t *testing.T) {
		_, _, err := splitTag("       ")
		assert.Equal(t, ErrInvalidParamTag, err)
	})

	t.Run("should return correct tag", func(t *testing.T) {
		expected := "ninja"
		tag, meta, err := splitTag(expected)
		assert.Nil(t, err)
		assert.Nil(t, meta)
		assert.Equal(t, expected, tag)
	})

	t.Run("should fail because of invalid meta", func(t *testing.T) {
		expectedTag := "ninja"
		invalidMeta := "samurai"
		tag, meta, err := splitTag(expectedTag + "," + invalidMeta)
		assert.NotNil(t, err)
		assert.Equal(t, ErrInvalidParamTagKeyValue, err)
		assert.Nil(t, meta)
		assert.Equal(t, "", tag)
	})

	t.Run("should fail because of invalid meta", func(t *testing.T) {
		expectedTag := "ninja"
		invalidMeta := " = "
		tag, meta, err := splitTag(expectedTag + "," + invalidMeta)
		assert.NotNil(t, err)
		assert.Equal(t, ErrInvalidParamTagKeyValue, err)
		assert.Nil(t, meta)
		assert.Equal(t, "", tag)
	})

	t.Run("should succeed", func(t *testing.T) {
		expectedTag := "ninja"
		expectedMeta := map[string]string{
			"stealthy": "true",
			"sword":    "ninja-sword",
		}
		s := expectedTag + "," +
			"sword" + "=" + expectedMeta["sword"] + "," +
			"stealthy" + "=" + expectedMeta["stealthy"]

		tag, meta, err := splitTag(s)
		assert.Nil(t, err)
		assert.Equal(t, expectedTag, tag)
		assert.Equal(t, expectedMeta, meta)
	})
}

func TestOptions(t *testing.T) {
	cfg := config{}

	WithQueryFunc(defaultCfg.Query)(&cfg)
	WithURLParamFunc(defaultCfg.Param)(&cfg)
	WithUnmarshaller(defaultCfg.Unmarshal)(&cfg)

	assert.Equal(t, defaultCfg, cfg)
}

// func TestBoolStruct(t *testing.T) {
// 	t.Run("setValue bool should succeed", func(t *testing.T) {
// 		query := func(t *http.Request) url.Values {
// 			return url.Values{
// 				"query": []string{""},
// 			}
// 		}

// 		s := boolStruct{}
// 		req := &http.Request{}
// 		err := As(req, &s, WithQuery(query))

// 		assert.Nil(t, err)
// 	})
// }
