package utility

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const HTTPClientTimeout = time.Duration(1) * time.Second

func Pretty(data interface{}) string {
	a, _ := json.MarshalIndent(data, "", " ") //nolint: errchkjson

	return string(a)
}

func PrintPretty(data interface{}) {
	a, _ := json.MarshalIndent(data, "", " ") //nolint: errchkjson

	fmt.Println(string(a)) //nolint: forbidigo
}

func EncodeToBytes(p interface{}) ([]byte, error) { //nolint: ireturn, nolintlint
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(p)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
} //nolint: ireturn, nolintlint

func DecodeToStruct[T any](s []byte) (T, error) { //nolint: ireturn, nolintlint
	var p T

	dec := gob.NewDecoder(bytes.NewReader(s))

	err := dec.Decode(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func FormatDuration(d time.Duration) string {
	switch {
	case d < time.Millisecond:
		return fmt.Sprintf("%dµs", d.Microseconds())
	case d < time.Second:
		return fmt.Sprintf("%dms", d.Milliseconds())
	case d < time.Minute:
		return fmt.Sprintf("%.1fs", d.Seconds())
	case d < time.Hour:
		return fmt.Sprintf("%.1fm", d.Minutes())
	default:
		return fmt.Sprintf("%.1fd", d.Hours())
	}
}

func Fetch(url, method string, payload *bytes.Buffer) ([]byte, error) {
	c := http.Client{Timeout: HTTPClientTimeout}
	method = strings.ToLower(method)

	var resp *http.Response

	var err error

	switch method {
	case "get":
		resp, err = c.Get(url) //nolint: bodyclose,noctx
	case "post":
		resp, err = c.Post(url, "application/json", payload) //nolint: bodyclose,noctx
	default:
		return nil, fmt.Errorf("fetch method '%s' not supportd", method) //nolint: err113
	}

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

func JSONFetch(url, method string, payload *bytes.Buffer) (any, error) {
	res, err := Fetch(url, method, payload)
	if err != nil {
		return nil, err
	}

	var result interface{}

	err = json.Unmarshal(res, &result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err) //nolint: forbidigo

		return nil, err
	}

	return result, nil
}
