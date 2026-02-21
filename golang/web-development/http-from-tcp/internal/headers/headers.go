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

	keyValueSeparatorIdx := bytes.Index(data[:newLineIdx], []byte(keyValueSeparator))
	if keyValueSeparatorIdx == -1 {
		return 0, false, fmt.Errorf(`invalid header: missing key value separator`)
	}
	key := string(data[:keyValueSeparatorIdx])
	value := string(data[keyValueSeparatorIdx+1 : newLineIdx])

	keyWithoutTrailingSpace := strings.TrimSpace(key)
	if key != keyWithoutTrailingSpace {
		return 0, false, fmt.Errorf("invalid header name: %s", key)
	}

	err = isValidKey(keyWithoutTrailingSpace)
	if err != nil {
		return 0, false, err
	}

	valueNormalized := strings.TrimSpace(value)
	keyNormalized := strings.ToLower(keyWithoutTrailingSpace)

	currentKeyValue, ok := h[keyNormalized]
	if ok {
		valueNormalized = strings.Join([]string{
			currentKeyValue,
			valueNormalized,
		}, ", ")
	}
	h[keyNormalized] = valueNormalized
	//println(keyNormalized, valueNormalized)

	return newLineIdx + len(crlf), false, nil
}

func (h Headers) Get(key string) (string, bool) {
	keyNormalized := strings.ToLower(key)
	value, ok := h[keyNormalized]
	return value, ok
}

func isValidKey(s string) error {
	if s == "" {
		return fmt.Errorf("invalid header: cannot be empty")
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !isFieldChar(c) {
			return fmt.Errorf("invalid header: %s", s)
		}
	}
	return nil
}

func isFieldChar(c byte) bool {
	return c == '!' || c == '#' || c == '$' || c == '%' ||
		c == '&' || c == '\'' || c == '*' || c == '+' ||
		c == '-' || c == '.' || c == '^' || c == '_' ||
		c == '`' || c == '|' || c == '~' ||
		(c >= '0' && c <= '9') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z')
}
