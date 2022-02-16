package zserio

import (
	"bytes"
	"fmt"
)

// Marshal accepts a pointer to an instance of a zserio object and returns a
// byte array which represents a zserio stream. On error, it will return an
// empty slice.
func Marshal(m Marshaler) ([]byte, error) {
	var buf bytes.Buffer
	w := NewWriter(&buf)

	if err := m.MarshalZserio(w); err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("flush buffer: %w", err)
	}

	return buf.Bytes(), nil
}

// Unmarshal accepts a zserio byte stream and a pointer to an instance of a
// zserio object type, which is going to be populate with values deserialized
// from the bytes and returns a deserialization error.
func Unmarshal(b []byte, m Unmarshaler) error {
	r := NewReader(bytes.NewReader(b))
	if err := m.UnmarshalZserio(r); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}
