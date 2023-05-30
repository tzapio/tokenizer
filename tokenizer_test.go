package tokenizer_test

import (
	"testing"

	"github.com/tzapio/tokenizer/codec/cl100k_base"
)

func TestCl100kEncoding(t *testing.T) {
	tokenizer := cl100k_base.NewCl100kBase()

	tests := []struct {
		text string
		ids  []int32
	}{
		{text: "hello world", ids: []int32{15339, 1917}},
		{text: "hello  world", ids: []int32{15339, 220, 1917}},
		{text: "hello   world", ids: []int32{15339, 256, 1917}},
		{text: "supercalifragilistic", ids: []int32{13066, 3035, 278, 333, 4193, 321, 4633}},
		{text: "We know what we are, but know not what we may be.", ids: []int32{1687, 1440, 1148, 584, 527, 11, 719, 1440, 539, 1148, 584, 1253, 387, 13}},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			ids, _, err := tokenizer.Encode(test.text)
			if err != nil {
				t.Fatalf("error encoding: %v", err)
			}

			if !sliceEqual(ids, test.ids) {
				t.Fatalf("input: %s want: %v got: %v", test.text, test.ids, ids)
			}

			text, err := tokenizer.Decode(ids)
			if err != nil {
				t.Fatalf("error decoding: %v", err)
			}

			if text != test.text {
				t.Fatalf("input: %v want: %s got: %s", test.ids, test.text, text)
			}
		})
	}
}

func sliceEqual(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}

	for i, elem := range a {
		if elem != b[i] {
			return false
		}
	}

	return true
}
