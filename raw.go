package fflogs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/RyuaNerin/go-fflogs/structure"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
)

var (
	ErrNotOk = errors.New("statusCode is not ok")
)

type Raw struct {
	client *Client
}

func urlEncode(str string) string {
	var sb strings.Builder

	for str != "" {
		r, sz := utf8.DecodeRuneInString(str)

		switch {
		case unicode.IsNumber(r):
			fallthrough
		case unicode.Is(unicode.Latin, r):
			fallthrough
		case r == '-' || r == '_' || r == '.':
			sb.WriteRune(r)

		default:
			for i := 0; i < sz; i++ {
				fmt.Fprintf(&sb, "%%%02X", byte(str[i]))
			}
		}

		str = str[sz:]
	}

	return sb.String()
}

func buildUrl(u url.URL, path string, opt interface{}, apiKey string) (string, error) {
	var optType reflect.Value
	var optElem reflect.Value
	var optElemType reflect.Type

	if opt != nil {
		optType = reflect.ValueOf(opt)
		optElem = optType.Elem()
		optElemType = optElem.Type()
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////
	// Path

	if opt == nil {
		u.Path += path
	} else {
		var sbPath strings.Builder
		var sbPathRaw strings.Builder

		sbPath.WriteString(u.Path)
		sbPathRaw.WriteString(u.Path)

		idx := 0
		pathStart := -1
		for idx < len(path) {
			c := path[idx]
			switch c {
			case '{':
				pathStart = idx + 1
			case '}':
				pathName := path[pathStart:idx]
				notFound := true
				for i := 0; i < optElem.NumField(); i++ {
					tag := optElemType.Field(i).Tag.Get("path")
					if tag == pathName {
						value := optElem.Field(i).Interface()
						if value != nil {
							notFound = false
							v := cast.ToString(value)
							sbPath.WriteString(v)
							sbPathRaw.WriteString(urlEncode(v))
							//sb.WriteString(cast.ToString(value))
						}
						break
					}
				}
				if notFound {
					return "", fmt.Errorf("%s cannot be empty", pathName)
				}
				pathStart = -1

			default:
				if pathStart == -1 {
					sbPath.WriteByte(c)
					sbPathRaw.WriteByte(c)
				}
			}
			idx++
		}

		u.Path = sbPath.String()
		u.RawPath = sbPathRaw.String()
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////
	// Query

	query := url.Values{}
	query.Set("api_key", apiKey)

	if opt != nil {
		for i := 0; i < optElem.NumField(); i++ {
			tag := optElemType.Field(i).Tag.Get("query")
			if tag != "" {
				value := optElem.Field(i).Interface()
				if value != nil && (reflect.ValueOf(value).Kind() != reflect.Ptr || !reflect.ValueOf(value).IsNil()) {
					query.Add(tag, cast.ToString(value))
				}
			}
		}
	}

	u.RawQuery = query.Encode()

	return u.String(), nil
}

func (c *Raw) call(context context.Context, path string, opt interface{}, respData interface{}) error {
	url, err := buildUrl(c.client.baseUrl, path, opt, c.client.apiKey)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(context)

	////////////////////////////////////////////////////////////////////////////////////////////////////
	// Response

	resp, err := c.client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	////////////////////////////////////////////////////////////////////////////////////////////////////
	// Json

	err = jsoniter.NewDecoder(resp.Body).Decode(&respData)
	if err != nil && err != io.EOF {
		return err
	}

	if r, ok := respData.(*structure.BaseResponse); ok {
		if r.Error != "" {
			return ErrNotOk
		}
	}

	return nil
}
