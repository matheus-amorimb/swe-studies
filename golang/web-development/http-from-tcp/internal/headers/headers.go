package headers

import (
	"bytes"
	"fmt"
	"strings"
)

type Headers map[string]string

const crlf = "\r\n"
const keyValueSeparator = ":"

func NewHeaders() Headers {
	return make(Headers)
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	newLineIdx := bytes.Index(data, []byte(crlf))
	if newLineIdx == -1 {
		return 0, false, nil
	}

	//The blank line separates the headers from the body
	if newLineIdx == 0 {
		return len(crlf), true, nil
	}

	keyValueSeparatorIdx := bytes.Index(data, []byte(keyValueSeparator))
	key := string(data[:keyValueSeparatorIdx])
	value := string(data[keyValueSeparatorIdx+1 : newLineIdx])

	keyNormalized := strings.TrimRight(key, " ")
	if key != keyNormalized {
		return 0, false, fmt.Errorf("invalid header name: %s", key)
	}

	valueNormalized := strings.TrimSpace(value)
	keyNormalized = strings.TrimSpace(keyNormalized)
	fmt.Println(keyNormalized)
	fmt.Println(valueNormalized)
	h[keyNormalized] = valueNormalized

	return newLineIdx + len(crlf), false, nil
}
