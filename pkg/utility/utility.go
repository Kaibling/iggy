package utility

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"time"
)

func Pretty(data interface{}) string {
	a, _ := json.MarshalIndent(data, "", " ") //nolint: errchkjson

	return string(a)
}

func EncodeToBytes(p interface{}) ([]byte, error) {
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
