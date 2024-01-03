package yadisk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Client struct {
	token      string
	url        *url.URL
	httpClient *http.Client
}

type Opt func(*Client)

func WithHttpClient(httpClient *http.Client) Opt {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

func NewClient(token string, host string, opts ...Opt) *Client {
	c := Client{
		token: token,
		url: &url.URL{
			Scheme: "https",
			Host:   host,
			Path:   "v1",
		},
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(&c)
	}

	return &c
}

func (c *Client) doRequest(ctx context.Context, method string, path string, query url.Values, request, response any) error {
	u := *c.url
	u.Path = path
	u.RawQuery = query.Encode()

	var body io.Reader
	if request != nil {
		marshaledBytes, err := json.Marshal(request)
		if err != nil {
			return err
		}
		log.Println(string(marshaledBytes))
		body = bytes.NewReader(marshaledBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return err
	}

	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if err := checkErr(resp); err != nil {
		return err
	}

	defer resp.Body.Close()

	if response == nil {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth "+c.token)
}

func checkErr(resp *http.Response) error {
	if resp.StatusCode < 400 {
		return nil
	}

	errResp := Error{Code: resp.StatusCode}

	jdec := json.NewDecoder(resp.Body)
	if err := jdec.Decode(&errResp); err != nil {
		return err
	}

	return errResp
}

func queryFromParams(p any) (url.Values, error) {
	m := make(url.Values)
	t := reflect.TypeOf(p).Elem()
	v := reflect.ValueOf(p).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		tag, ok := f.Tag.Lookup("param")
		if !ok {
			continue
		}

		pKeys := strings.SplitN(tag, ",", 2)

		vf := v.Field(i)
		if vf.IsZero() && len(pKeys) > 1 && pKeys[1] == "required" {
			return nil, fmt.Errorf("field %s.%s required", t.Name(), f.Name)
		}
		if vf.IsZero() {
			continue
		}
		switch vf.Kind() {
		case reflect.String:
			m[pKeys[0]] = []string{vf.String()}
		case reflect.Bool:
			m[pKeys[0]] = []string{strconv.FormatBool(vf.Bool())}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			m[pKeys[0]] = []string{strconv.FormatInt(vf.Int(), 10)}
		case reflect.Slice:
			if vf.Index(0).Kind() == reflect.String {
				s := make([]string, vf.Len())
				for i := 0; i < vf.Len(); i++ {
					s[i] = vf.Index(i).String()
				}
				m[pKeys[0]] = []string{strings.Join(s, ",")}
			}
		default:
			return nil, fmt.Errorf("unsupported type")
		}
	}

	return m, nil
}
