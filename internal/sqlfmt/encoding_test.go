package sqlfmt

import (
	"encoding/binary"
	"testing"
	"unicode/utf16"
)

func TestDetectSQLEncodingIssue_UTF8Plain(t *testing.T) {
	raw := []byte("CREATE TABLE dbo.TableA (id INT NOT NULL);\n")
	if msg, bad := DetectSQLEncodingIssue(raw); bad {
		t.Fatalf("unexpected encoding issue: %s", msg)
	}
}

func TestDetectSQLEncodingIssue_UTF8BOM(t *testing.T) {
	raw := append([]byte{0xEF, 0xBB, 0xBF}, []byte("CREATE OR ALTER VIEW dbo.ViewA AS SELECT 1 AS id;\n")...)
	if msg, bad := DetectSQLEncodingIssue(raw); bad {
		t.Fatalf("unexpected encoding issue: %s", msg)
	}
}

func TestDetectSQLEncodingIssue_InvalidUTF8(t *testing.T) {
	raw := []byte{'-', '-', ' ', 0x96, '\n'}
	if _, bad := DetectSQLEncodingIssue(raw); !bad {
		t.Fatalf("expected invalid UTF-8 issue")
	}
}

func TestDetectSQLEncodingIssue_UTF16BOM(t *testing.T) {
	raw := []byte{0xFF, 0xFE, 'C', 0x00, 'R', 0x00, 'E', 0x00, 'A', 0x00, 'T', 0x00, 'E', 0x00}
	if _, bad := DetectSQLEncodingIssue(raw); !bad {
		t.Fatalf("expected UTF-16 BOM issue")
	}
}

func TestDetectSQLEncodingIssue_UTF16WithoutBOMPattern(t *testing.T) {
	raw := []byte{'C', 0x00, 'R', 0x00, 'E', 0x00, 'A', 0x00, 'T', 0x00, 'E', 0x00}
	if _, bad := DetectSQLEncodingIssue(raw); !bad {
		t.Fatalf("expected suspected UTF-16 issue")
	}
}

func TestNormalizeSQLToUTF8(t *testing.T) {
	text := "CREATE PROCEDURE dbo.Utf AS SELECT N'東京😀';\n"
	tests := []struct {
		name   string
		raw    []byte
		source string
	}{
		{name: "utf8", raw: []byte(text), source: "UTF-8"},
		{name: "utf8 bom", raw: append(append([]byte{}, utf8BOM...), []byte(text)...), source: "UTF-8 BOM"},
		{name: "utf16 le bom", raw: encodeUTF16ForTest(text, binary.LittleEndian, true), source: "UTF-16LE BOM"},
		{name: "utf16 be bom", raw: encodeUTF16ForTest(text, binary.BigEndian, true), source: "UTF-16BE BOM"},
		{name: "utf16 le no bom", raw: encodeUTF16ForTest("CREATE TABLE dbo.Utf16 (id INT);\n", binary.LittleEndian, false), source: "UTF-16LE (no BOM)"},
		{name: "utf32 le bom", raw: encodeUTF32ForTest(text, binary.LittleEndian), source: "UTF-32LE BOM"},
		{name: "utf32 be bom", raw: encodeUTF32ForTest(text, binary.BigEndian), source: "UTF-32BE BOM"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, change, err := NormalizeSQLToUTF8(tt.raw)
			if err != nil {
				t.Fatalf("NormalizeSQLToUTF8() error = %v", err)
			}
			want := text
			if tt.name == "utf16 le no bom" {
				want = "CREATE TABLE dbo.Utf16 (id INT);\n"
			}
			if string(got) != want {
				t.Fatalf("normalized text = %q, want %q", got, want)
			}
			if change.SourceEncoding != tt.source {
				t.Fatalf("source encoding = %q, want %q", change.SourceEncoding, tt.source)
			}
			if change.Changed != (tt.source != "UTF-8") {
				t.Fatalf("changed = %v for %s", change.Changed, tt.source)
			}
		})
	}
}

func TestNormalizeSQLToUTF8RejectsAmbiguousOrMalformedSource(t *testing.T) {
	for _, raw := range [][]byte{
		{'-', '-', ' ', 0x96, '\n'},
		{0xFF, 0xFE, 0x00},
		{0x00, 0x00, 0xFE, 0xFF, 0x00, 0x11, 0x00, 0x00},
	} {
		if _, _, err := NormalizeSQLToUTF8(raw); err == nil {
			t.Fatalf("NormalizeSQLToUTF8(% X) unexpectedly succeeded", raw)
		}
	}
}

func encodeUTF16ForTest(text string, order binary.ByteOrder, bom bool) []byte {
	words := utf16.Encode([]rune(text))
	raw := make([]byte, 0, len(words)*2+2)
	if bom {
		if order == binary.LittleEndian {
			raw = append(raw, utf16LEBOM...)
		} else {
			raw = append(raw, utf16BEBOM...)
		}
	}
	for _, word := range words {
		buf := make([]byte, 2)
		order.PutUint16(buf, word)
		raw = append(raw, buf...)
	}
	return raw
}

func encodeUTF32ForTest(text string, order binary.ByteOrder) []byte {
	raw := make([]byte, 0, len(text)*4+4)
	if order == binary.LittleEndian {
		raw = append(raw, utf32LEBOM...)
	} else {
		raw = append(raw, utf32BEBOM...)
	}
	for _, r := range text {
		buf := make([]byte, 4)
		order.PutUint32(buf, uint32(r))
		raw = append(raw, buf...)
	}
	return raw
}
