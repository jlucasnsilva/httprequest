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

func TestOptions(t *testing.T) {
	cfg := config{}

	WithQueryFunc(defaultCfg.Query)(&cfg)
	WithURLParamFunc(defaultCfg.Param)(&cfg)
	WithUnmarshaller(defaultCfg.Unmarshal)(&cfg)

	assert.NotNil(t, cfg.Query)
	assert.NotNil(t, cfg.Param)
	assert.NotNil(t, cfg.Unmarshal)
}
