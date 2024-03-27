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
	tagName        = "from"
	timeLayoutMeta = "layout"
)

const (
	urlParamTag    = "url-param"
	urlQueryTag    = "url-query"
	requestBodyTag = "request-body"
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

func As(req *http.Request, obj any, opts ...Option) error {
	var (
		values url.Values

		cfg         = defaultCfg
		decodedBody = false
	)

	for _, opt := range opts {
		opt(&cfg)
	}

	values = cfg.Query(req)
	v := reflect.ValueOf(obj).Elem()
	for i, f := range reflect.VisibleFields(v.Type()) {
		tag := f.Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}

		kind, source, meta, err := splitTag(tag)
		if err != nil {
			return err
		}

		switch kind {
		case urlParamTag:
			setValue(v.FieldByName(f.Name), cfg.Param(req, source), meta)
		case urlQueryTag:
			setValue(v.FieldByName(f.Name), values.Get(source), meta)
		case requestBodyTag:
			if decodedBody {
				panic("Cannot decode the body twice")
			}

			decodedBody = true
			rvalue := reflect.ValueOf(req)
			target := v.FieldByIndex([]int{i})
			decode := reflect.ValueOf(cfg.Unmarshal)
			if target.Kind() == reflect.Pointer {
				if target.IsNil() {
					typ := target.Type().Elem()
					target.Set(reflect.New(typ))
				}
			}

			in := []reflect.Value{rvalue, target}
			if target.Kind() != reflect.Pointer {
				in[1] = target.Addr()
			}

			ret := decode.Call(in)
			if len(ret) > 0 && !ret[0].IsNil() {
				return ret[0].Interface().(error)
			}
		default:
			panic("Invalid kind: " + kind)
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
	case reflect.String:
		f.SetString(param)
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

func splitTag(tag string) (kind, source string, meta map[string]string, err error) {
	parts := strings.Split(tag, ",")
	if len(parts) < 1 || parts[0] == "" {
		return "", "", nil, ErrInvalidParamTag
	}

	kv := strings.TrimSpace(parts[0])
	if kv == "" || kv == "=" {
		return "", "", nil, ErrInvalidParamTag
	}

	if kv == requestBodyTag {
		return kv, "", nil, nil
	}

	kind, source, err = splitKV(kv)
	if len(parts) < 2 {
		return kind, source, nil, err
	}

	m := make(map[string]string)
	for _, p := range parts[1:] {
		key, value, err := splitKV(p)
		if err != nil {
			return "", "", nil, err
		}

		k := strings.TrimSpace(key)
		v := strings.TrimSpace(value)
		if k == "" || v == "" {
			return "", "", nil, ErrInvalidParamTagKeyValue
		}
		m[k] = v
	}
	return kind, source, m, nil
}

func splitKV(kv string) (string, string, error) {
	parts := strings.Split(kv, "=")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", ErrInvalidParamTagKeyValue
	}
	return parts[0], parts[1], nil
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
