package sqlfmt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
)

var (
	utf16LEBOM = []byte{0xFF, 0xFE}
	utf16BEBOM = []byte{0xFE, 0xFF}
	utf32LEBOM = []byte{0xFF, 0xFE, 0x00, 0x00}
	utf32BEBOM = []byte{0x00, 0x00, 0xFE, 0xFF}
	utf8BOM    = []byte{0xEF, 0xBB, 0xBF}
)

// EncodingNormalization describes a lossless source-encoding normalization.
// SQL source files are stored as UTF-8 without a BOM after normalization.
type EncodingNormalization struct {
	Changed        bool
	SourceEncoding string
}

// NormalizeSQLToUTF8 converts supported Unicode SQL encodings to UTF-8.
//
// It only converts byte streams whose encoding can be identified safely:
// UTF-8 (optionally with BOM), UTF-16 and UTF-32 with BOM, plus the narrow
// UTF-16-without-BOM pattern already detected by DetectSQLEncodingIssue.
// Arbitrary non-UTF-8 legacy code pages are intentionally rejected: guessing
// one could silently alter SQL identifiers or string literals.
func NormalizeSQLToUTF8(raw []byte) ([]byte, EncodingNormalization, error) {
	if len(raw) == 0 {
		return raw, EncodingNormalization{SourceEncoding: "UTF-8"}, nil
	}

	// UTF-32 LE has the UTF-16 LE BOM as a prefix, so test UTF-32 first.
	switch {
	case bytes.HasPrefix(raw, utf32LEBOM):
		text, err := decodeUTF32(raw[len(utf32LEBOM):], binary.LittleEndian)
		if err != nil {
			return nil, EncodingNormalization{}, fmt.Errorf("decode UTF-32LE SQL source: %w", err)
		}
		return []byte(text), EncodingNormalization{Changed: true, SourceEncoding: "UTF-32LE BOM"}, nil
	case bytes.HasPrefix(raw, utf32BEBOM):
		text, err := decodeUTF32(raw[len(utf32BEBOM):], binary.BigEndian)
		if err != nil {
			return nil, EncodingNormalization{}, fmt.Errorf("decode UTF-32BE SQL source: %w", err)
		}
		return []byte(text), EncodingNormalization{Changed: true, SourceEncoding: "UTF-32BE BOM"}, nil
	case bytes.HasPrefix(raw, utf16LEBOM):
		text, err := decodeUTF16(raw[len(utf16LEBOM):], binary.LittleEndian)
		if err != nil {
			return nil, EncodingNormalization{}, fmt.Errorf("decode UTF-16LE SQL source: %w", err)
		}
		return []byte(text), EncodingNormalization{Changed: true, SourceEncoding: "UTF-16LE BOM"}, nil
	case bytes.HasPrefix(raw, utf16BEBOM):
		text, err := decodeUTF16(raw[len(utf16BEBOM):], binary.BigEndian)
		if err != nil {
			return nil, EncodingNormalization{}, fmt.Errorf("decode UTF-16BE SQL source: %w", err)
		}
		return []byte(text), EncodingNormalization{Changed: true, SourceEncoding: "UTF-16BE BOM"}, nil
	case bytes.HasPrefix(raw, utf8BOM):
		text := raw[len(utf8BOM):]
		if !utf8.Valid(text) {
			return nil, EncodingNormalization{}, fmt.Errorf("UTF-8 BOM SQL source contains invalid UTF-8")
		}
		return append([]byte(nil), text...), EncodingNormalization{Changed: true, SourceEncoding: "UTF-8 BOM"}, nil
	}

	if littleEndian, ok := utf16WithoutBOMEndian(raw); ok {
		var order binary.ByteOrder = binary.BigEndian
		name := "UTF-16BE (no BOM)"
		if littleEndian {
			order = binary.LittleEndian
			name = "UTF-16LE (no BOM)"
		}
		text, err := decodeUTF16(raw, order)
		if err != nil {
			return nil, EncodingNormalization{}, fmt.Errorf("decode %s SQL source: %w", name, err)
		}
		return []byte(text), EncodingNormalization{Changed: true, SourceEncoding: name}, nil
	}

	if !utf8.Valid(raw) {
		return nil, EncodingNormalization{}, fmt.Errorf("SQL source is not valid UTF-8 and its encoding cannot be safely identified")
	}
	return raw, EncodingNormalization{SourceEncoding: "UTF-8"}, nil
}

func decodeUTF16(raw []byte, order binary.ByteOrder) (string, error) {
	if len(raw)%2 != 0 {
		return "", fmt.Errorf("odd byte length %d", len(raw))
	}
	words := make([]uint16, len(raw)/2)
	for i := range words {
		words[i] = order.Uint16(raw[i*2:])
	}
	for i, word := range words {
		if word >= 0xD800 && word <= 0xDBFF {
			if i+1 >= len(words) || words[i+1] < 0xDC00 || words[i+1] > 0xDFFF {
				return "", fmt.Errorf("unpaired high surrogate at UTF-16 unit %d", i)
			}
		}
		if word >= 0xDC00 && word <= 0xDFFF && (i == 0 || words[i-1] < 0xD800 || words[i-1] > 0xDBFF) {
			return "", fmt.Errorf("unpaired low surrogate at UTF-16 unit %d", i)
		}
	}
	return string(utf16.Decode(words)), nil
}

func decodeUTF32(raw []byte, order binary.ByteOrder) (string, error) {
	if len(raw)%4 != 0 {
		return "", fmt.Errorf("byte length %d is not divisible by 4", len(raw))
	}
	runes := make([]rune, len(raw)/4)
	for i := range runes {
		value := order.Uint32(raw[i*4:])
		if value > utf8.MaxRune || (value >= 0xD800 && value <= 0xDFFF) {
			return "", fmt.Errorf("invalid Unicode scalar U+%04X at UTF-32 unit %d", value, i)
		}
		runes[i] = rune(value)
	}
	return string(runes), nil
}

// DetectSQLEncodingIssue validates SQL file text encoding policy.
//
// Supported:
// - UTF-8 (with or without BOM)
//
// Rejected:
// - UTF-16/UTF-32 with BOM
// - likely UTF-16 (no BOM) based on NUL-byte pattern
// - invalid UTF-8 byte sequences
func DetectSQLEncodingIssue(raw []byte) (string, bool) {
	if len(raw) == 0 {
		return "", false
	}

	if bytes.HasPrefix(raw, utf16LEBOM) || bytes.HasPrefix(raw, utf16BEBOM) ||
		bytes.HasPrefix(raw, utf32LEBOM) || bytes.HasPrefix(raw, utf32BEBOM) {
		return "SQL file encoding must be UTF-8 (UTF-16/UTF-32 BOM detected)", true
	}

	trimmed := raw
	if bytes.HasPrefix(trimmed, utf8BOM) {
		trimmed = trimmed[len(utf8BOM):]
	}

	if likelyUTF16WithoutBOM(trimmed) {
		return "SQL file encoding must be UTF-8 (suspected UTF-16 content without BOM)", true
	}

	if !utf8.Valid(trimmed) {
		return "SQL file encoding must be valid UTF-8", true
	}

	return "", false
}

func likelyUTF16WithoutBOM(raw []byte) bool {
	_, ok := utf16WithoutBOMEndian(raw)
	return ok
}

func utf16WithoutBOMEndian(raw []byte) (littleEndian bool, ok bool) {
	if len(raw) < 8 {
		return false, false
	}

	nul := 0
	evenNul := 0
	oddNul := 0
	for i, b := range raw {
		if b != 0 {
			continue
		}
		nul++
		if i%2 == 0 {
			evenNul++
		} else {
			oddNul++
		}
	}
	if nul == 0 {
		return false, false
	}

	// UTF-16 text typically has NUL bytes in alternating positions.
	if float64(nul)/float64(len(raw)) < 0.20 {
		return false, false
	}
	if oddNul*100/nul >= 85 {
		return true, true
	}
	if evenNul*100/nul >= 85 {
		return false, true
	}
	return false, false
}
