package utility

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
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
}

func DecodeToStruct[T any](s []byte) (T, error) { //nolint: ireturn
	var p T

	dec := gob.NewDecoder(bytes.NewReader(s))

	err := dec.Decode(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}
