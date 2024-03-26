package httprequest

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
