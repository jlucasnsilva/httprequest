package httprequest

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type (
	config struct {
		Param     func(*http.Request, string) string
		Unmarshal func(*http.Request, any) error
		Query     func(*http.Request) url.Values
	}

	Option func(*config)
)

const (
	urlParamTag    = "url-param"
	urlQueryTag    = "url-query"
	requestBodyTag = "request-body"
	timeLayoutMeta = "layout"
)

var (
	ErrInvalidParamTag         = errors.New("invalid param tag")
	ErrInvalidParamTagKeyValue = errors.New("invalid param tag key-value pair")
	ErrInvalidURLValueTag      = errors.New("invalid url value tag")
)

var timeType = reflect.TypeOf(time.Time{})

var defaultCfg = config{
	Unmarshal: func(r *http.Request, v any) error {
		return json.NewDecoder(r.Body).Decode(v)
	},
	Param: func(r *http.Request, key string) string {
		return r.PathValue(key)
	},
	Query: func(r *http.Request) url.Values {
		return r.URL.Query()
	},
}

func As[T any](req *http.Request, obj *T, opts ...Option) error {
	var (
		values      = req.URL.Query()
		cfg         = defaultCfg
		decodedBody = false
	)

	for _, opt := range opts {
		opt(&cfg)
	}

	v := reflect.ValueOf(obj).Elem()
	for i, f := range reflect.VisibleFields(v.Type()) {
		if param := f.Tag.Get(urlParamTag); param != "" {
			key, meta, err := splitTag(param)
			if err != nil {
				return err
			}
			setValue(v.FieldByName(f.Name), cfg.Param(req, key), meta)
		} else if valueTag := f.Tag.Get(urlQueryTag); valueTag != "" {
			key, meta, err := splitTag(valueTag)
			if err != nil {
				return err
			}
			// TODO handle array
			setValue(v.FieldByName(f.Name), values.Get(key), meta)
		} else if body := f.Tag.Get(requestBodyTag); body != "" {
			if decodedBody {
				panic("Cannot decode the body twice")
			}

			decodedBody = true
			rvalue := reflect.ValueOf(req)
			target := v.FieldByIndex([]int{i})
			decode := reflect.ValueOf(cfg.Unmarshal)
			in := []reflect.Value{rvalue, target}
			ret := decode.Call(in)
			if len(ret) > 0 {
				return ret[0].Interface().(error)
			}
		}
	}
	return nil
}

func WithURLParamFunc(getter func(*http.Request, string) string) Option {
	return func(cfg *config) {
		cfg.Param = getter
	}
}

func WithUnmarshaller(unmarshal func(*http.Request, any) error) Option {
	return func(cfg *config) {
		cfg.Unmarshal = unmarshal
	}
}

func WithQueryFunc(q func(*http.Request) url.Values) Option {
	return func(cfg *config) {
		cfg.Query = q
	}
}

func setValue(f reflect.Value, param string, meta map[string]string) error {
	switch f.Kind() {
	case reflect.Bool:
		if v, err := strconv.ParseBool(param); err != nil {
			return err
		} else {
			f.SetBool(v)
		}
	case reflect.Int:
		if v, err := strconv.ParseInt(param, 10, 64); err != nil {
			return err
		} else {
			f.SetInt(v)
		}
	case reflect.Int8:
		if v, err := strconv.ParseInt(param, 10, 8); err != nil {
			return err
		} else {
			f.SetInt(v)
		}
	case reflect.Int16:
		if v, err := strconv.ParseInt(param, 10, 16); err != nil {
			return err
		} else {
			f.SetInt(v)
		}
	case reflect.Int32:
		if v, err := strconv.ParseInt(param, 10, 32); err != nil {
			return err
		} else {
			f.SetInt(v)
		}
	case reflect.Int64:
		if v, err := strconv.ParseInt(param, 10, 64); err != nil {
			return err
		} else {
			f.SetInt(v)
		}
	case reflect.Uint:
		if v, err := strconv.ParseUint(param, 10, 64); err != nil {
			return err
		} else {
			f.SetUint(v)
		}
	case reflect.Uint8:
		if v, err := strconv.ParseUint(param, 10, 8); err != nil {
			return err
		} else {
			f.SetUint(v)
		}
	case reflect.Uint16:
		if v, err := strconv.ParseUint(param, 10, 16); err != nil {
			return err
		} else {
			f.SetUint(v)
		}
	case reflect.Uint32:
		if v, err := strconv.ParseUint(param, 10, 32); err != nil {
			return err
		} else {
			f.SetUint(v)
		}
	case reflect.Uint64:
		if v, err := strconv.ParseUint(param, 10, 64); err != nil {
			return err
		} else {
			f.SetUint(v)
		}
	case reflect.Float32:
		if v, err := strconv.ParseFloat(param, 32); err != nil {
			return err
		} else {
			f.SetFloat(v)
		}
	case reflect.Float64:
		if v, err := strconv.ParseFloat(param, 64); err != nil {
			return err
		} else {
			f.SetFloat(v)
		}
	}

	switch f.Type() {
	case timeType:
		layout := timeLayout(meta)
		t, err := time.Parse(layout, param)
		if err != nil {
			return err
		}
		f.Set(reflect.ValueOf(t))
	}

	return nil
}

func splitTag(tag string) (string, map[string]string, error) {
	parts := strings.Split(tag, ",")
	if len(parts) < 1 || parts[0] == "" {
		return "", nil, ErrInvalidParamTag
	}

	name := strings.TrimSpace(parts[0])
	if name == "" {
		return "", nil, ErrInvalidParamTag
	}

	if len(parts) < 2 {
		return name, nil, nil
	}

	m := make(map[string]string)
	for _, p := range parts[1:] {
		kv := strings.Split(p, "=")
		if len(kv) != 2 {
			return "", nil, ErrInvalidParamTagKeyValue
		}

		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])
		if k == "" || v == "" {
			return "", nil, ErrInvalidParamTagKeyValue
		}
		m[k] = v
	}
	return name, m, nil
}

func timeLayout(m map[string]string) string {
	switch layout := m[timeLayoutMeta]; layout {
	case "Layout":
		return time.Layout
	case "ANSIC":
		return time.ANSIC
	case "UnixDate":
		return time.UnixDate
	case "RubyDate":
		return time.RubyDate
	case "RFC822":
		return time.RFC822
	case "RFC822Z":
		return time.RFC822Z
	case "RFC850":
		return time.RFC850
	case "RFC1123":
		return time.RFC1123
	case "RFC1123Z":
		return time.RFC1123Z
	case "RFC3339":
		return time.RFC3339
	case "RFC3339Nano":
		return time.RFC3339Nano
	case "Kitchen":
		return time.Kitchen
	case "Stamp":
		return time.Stamp
	case "StampMilli":
		return time.StampMilli
	case "StampMicro":
		return time.StampMicro
	case "StampNano":
		return time.StampNano
	case "DateTime":
		return time.DateTime
	case "DateOnly":
		return time.DateOnly
	case "TimeOnly":
		return time.TimeOnly
	default:
		return time.RFC3339
	}
}
