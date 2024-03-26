package httprequest

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetValueBool(t *testing.T) {
	t.Run("setValue bool should fail", func(t *testing.T) {
		flag := true
		param := "error"
		_, expected := strconv.ParseBool(param)
		err := setValue(reflect.ValueOf(&flag).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue bool succeed with value false", func(t *testing.T) {
		flag := true
		param := "false"
		err := setValue(reflect.ValueOf(&flag).Elem(), param, nil)

		assert.Nil(t, err)
		assert.False(t, flag)
	})

	t.Run("setValue bool succeed with value true", func(t *testing.T) {
		flag := false
		param := "true"
		err := setValue(reflect.ValueOf(&flag).Elem(), param, nil)

		assert.Nil(t, err)
		assert.True(t, flag)
	})

	type boolStruct struct {
		Flag bool
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := boolStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Flag")

		err := setValue(f, "true", nil)

		assert.Nil(t, err)
		assert.True(t, st.Flag)
	})
}

func TestSetValueInt(t *testing.T) {
	t.Run("setValue int should fail", func(t *testing.T) {
		i := 0
		param := "error"
		_, expected := strconv.ParseInt(param, 10, 64)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue int succeed", func(t *testing.T) {
		i := 0
		expected := 10
		param := strconv.FormatInt(int64(expected), 10)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, i)
	})

	type intStruct struct {
		Int int
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := intStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Int")
		expected := 10
		param := strconv.FormatInt(int64(expected), 10)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, st.Int)
	})
}

func TestSetValueInt8(t *testing.T) {
	t.Run("setValue int8 should fail", func(t *testing.T) {
		i := int8(0)
		param := "error"
		_, expected := strconv.ParseInt(param, 10, 8)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue int8 succeed", func(t *testing.T) {
		i := int8(0)
		expected := int8(10)
		param := strconv.FormatInt(int64(expected), 10)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, i)
	})

	type intStruct struct {
		Int int8
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := intStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Int")
		expected := int8(10)
		param := strconv.FormatInt(int64(expected), 10)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, st.Int)
	})
}

func TestSetValueInt16(t *testing.T) {
	t.Run("setValue int16 should fail", func(t *testing.T) {
		i := int16(0)
		param := "error"
		_, expected := strconv.ParseInt(param, 10, 8)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue int16 succeed", func(t *testing.T) {
		i := int16(0)
		expected := int16(10)
		param := strconv.FormatInt(int64(expected), 10)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, i)
	})

	type intStruct struct {
		Int int16
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := intStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Int")
		expected := int16(10)
		param := strconv.FormatInt(int64(expected), 10)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, st.Int)
	})
}

func TestSetValueInt32(t *testing.T) {
	t.Run("setValue int32 should fail", func(t *testing.T) {
		i := int32(0)
		param := "error"
		_, expected := strconv.ParseInt(param, 10, 8)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue int32 succeed", func(t *testing.T) {
		i := int32(0)
		expected := int32(10)
		param := strconv.FormatInt(int64(expected), 10)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, i)
	})

	type intStruct struct {
		Int int32
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := intStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Int")
		expected := int32(10)
		param := strconv.FormatInt(int64(expected), 10)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, st.Int)
	})
}

func TestSetValueInt64(t *testing.T) {
	t.Run("setValue int64 should fail", func(t *testing.T) {
		i := int64(0)
		param := "error"
		_, expected := strconv.ParseInt(param, 10, 8)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue int64 succeed", func(t *testing.T) {
		i := int64(0)
		expected := int64(10)
		param := strconv.FormatInt(int64(expected), 10)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, i)
	})

	type intStruct struct {
		Int int64
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := intStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Int")
		expected := int64(10)
		param := strconv.FormatInt(int64(expected), 10)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, st.Int)
	})
}

func TestSetValueUint(t *testing.T) {
	t.Run("setValue uint should fail", func(t *testing.T) {
		i := uint(0)
		param := "error"
		_, expected := strconv.ParseUint(param, 10, 64)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue uint succeed", func(t *testing.T) {
		i := uint(0)
		expected := uint(10)
		param := strconv.FormatInt(int64(expected), 10)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, i)
	})

	type intStruct struct {
		Int uint
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := intStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Int")
		expected := uint(10)
		param := strconv.FormatUint(uint64(expected), 10)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, st.Int)
	})
}

func TestSetValueUint8(t *testing.T) {
	t.Run("setValue uint8 should fail", func(t *testing.T) {
		i := uint8(0)
		param := "error"
		_, expected := strconv.ParseUint(param, 10, 8)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue uint8 succeed", func(t *testing.T) {
		i := uint8(0)
		expected := uint8(10)
		param := strconv.FormatUint(uint64(expected), 10)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, i)
	})

	type intStruct struct {
		Int uint8
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := intStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Int")
		expected := uint8(10)
		param := strconv.FormatUint(uint64(expected), 10)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, st.Int)
	})
}

func TestSetValueUint16(t *testing.T) {
	t.Run("setValue uint16 should fail", func(t *testing.T) {
		i := uint16(0)
		param := "error"
		_, expected := strconv.ParseUint(param, 10, 8)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue uint16 succeed", func(t *testing.T) {
		i := uint16(0)
		expected := uint16(10)
		param := strconv.FormatUint(uint64(expected), 10)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, i)
	})

	type intStruct struct {
		Int uint16
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := intStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Int")
		expected := uint16(10)
		param := strconv.FormatUint(uint64(expected), 10)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, st.Int)
	})
}

func TestSetValueUint32(t *testing.T) {
	t.Run("setValue uint32 should fail", func(t *testing.T) {
		i := uint32(0)
		param := "error"
		_, expected := strconv.ParseUint(param, 10, 8)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue uint32 succeed", func(t *testing.T) {
		i := uint32(0)
		expected := uint32(10)
		param := strconv.FormatUint(uint64(expected), 10)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, i)
	})

	type intStruct struct {
		Int uint32
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := intStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Int")
		expected := uint32(10)
		param := strconv.FormatUint(uint64(expected), 10)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, st.Int)
	})
}

func TestSetValueUint64(t *testing.T) {
	t.Run("setValue uint64 should fail", func(t *testing.T) {
		i := uint64(0)
		param := "error"
		_, expected := strconv.ParseUint(param, 10, 8)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue uint64 succeed", func(t *testing.T) {
		i := uint64(0)
		expected := uint64(10)
		param := strconv.FormatUint(uint64(expected), 10)
		err := setValue(reflect.ValueOf(&i).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, i)
	})

	type intStruct struct {
		Int uint64
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := intStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Int")
		expected := uint64(10)
		param := strconv.FormatUint(uint64(expected), 10)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, st.Int)
	})
}

func TestSetValueFloat32(t *testing.T) {
	t.Run("setValue float32 should fail", func(t *testing.T) {
		f := float32(0)
		param := "error"
		_, expected := strconv.ParseFloat(param, 32)
		err := setValue(reflect.ValueOf(&f).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue float32 succeed", func(t *testing.T) {
		f := float32(0)
		expected := float32(123.4567)
		param := strconv.FormatFloat(float64(expected), 'f', 4, 32)
		err := setValue(reflect.ValueOf(&f).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, f)
	})

	type floatStruct struct {
		Float float32
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := floatStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Float")
		expected := float32(123.4567)
		param := strconv.FormatFloat(float64(expected), 'f', 4, 32)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, st.Float)
	})
}

func TestSetValueFloat64(t *testing.T) {
	t.Run("setValue float64 should fail", func(t *testing.T) {
		f := 0.0
		param := "error"
		_, expected := strconv.ParseFloat(param, 64)
		err := setValue(reflect.ValueOf(&f).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue float64 succeed", func(t *testing.T) {
		f := 0.0
		expected := 123.456790123
		param := strconv.FormatFloat(expected, 'f', 9, 64)
		err := setValue(reflect.ValueOf(&f).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, f)
	})

	type floatStruct struct {
		Float float64
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := floatStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Float")
		expected := 123.456790123
		param := strconv.FormatFloat(expected, 'f', 9, 64)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected, st.Float)
	})
}

func TestSetValueTime(t *testing.T) {
	t.Run("setValue time should fail", func(t *testing.T) {
		var tm time.Time
		param := "error"
		_, expected := time.Parse(time.RFC3339, param)
		err := setValue(reflect.ValueOf(&tm).Elem(), param, nil)

		assert.Equal(t, expected, err)
	})

	t.Run("setValue time succeed", func(t *testing.T) {
		var tm time.Time
		expected := time.Now()
		param := expected.Format(time.RFC3339)
		err := setValue(reflect.ValueOf(&tm).Elem(), param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected.Format(time.RFC3339), tm.Format(time.RFC3339))
	})

	t.Run("setValue time using meta succeed", func(t *testing.T) {
		metas := []map[string]string{
			{timeLayoutMeta: "Layout"},
			{timeLayoutMeta: "ANSIC"},
			{timeLayoutMeta: "UnixDate"},
			{timeLayoutMeta: "RubyDate"},
			{timeLayoutMeta: "RFC822"},
			{timeLayoutMeta: "RFC822Z"},
			{timeLayoutMeta: "RFC850"},
			{timeLayoutMeta: "RFC1123"},
			{timeLayoutMeta: "RFC1123Z"},
			{timeLayoutMeta: "RFC3339"},
			{timeLayoutMeta: "RFC3339Nano"},
			{timeLayoutMeta: "Kitchen"},
			{timeLayoutMeta: "Stamp"},
			{timeLayoutMeta: "StampMilli"},
			{timeLayoutMeta: "StampMicro"},
			{timeLayoutMeta: "StampNano"},
			{timeLayoutMeta: "DateTime"},
			{timeLayoutMeta: "DateOnly"},
			{timeLayoutMeta: "TimeOnly"},
		}

		expected := time.Now()
		for _, meta := range metas {
			t.Run("meta: "+meta[timeLayoutMeta], func(t *testing.T) {
				var tm time.Time

				layout := timeLayout(meta)
				param := expected.Format(layout)
				err := setValue(reflect.ValueOf(&tm).Elem(), param, meta)

				assert.Nil(t, err)
				assert.Equal(t, expected.Format(layout), tm.Format(layout))
			})
		}
	})

	type timeStruct struct {
		Time time.Time
	}

	t.Run("test using a struct - should succeed", func(t *testing.T) {
		st := timeStruct{}
		v := reflect.ValueOf(&st).Elem()
		f := v.FieldByName("Time")
		expected := time.Now()
		param := expected.Format(time.RFC3339)
		err := setValue(f, param, nil)

		assert.Nil(t, err)
		assert.Equal(t, expected.Format(time.RFC3339), st.Time.Format(time.RFC3339))
	})
}
