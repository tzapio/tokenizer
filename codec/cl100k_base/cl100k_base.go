package cl100k_base

import (
	_ "embed"

	"github.com/dlclark/regexp2"
	"github.com/tzapio/tokenizer/codec"
)

//go:generate go run ../../internal/binizer/main.go
func NewCl100kBase() *codec.Codec {
	cl100kBaseVocab := codec.LoadVocabFromByteArr(vocabData)
	cl100kBaseReverse := codec.LoadReverseFromByteArr(reverseData)
	return codec.New(
		"cl100k_base",
		cl100kBaseVocab,
		cl100kBaseReverse,
		regexp2.MustCompile(`(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}{1,3}| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+`, regexp2.None),
		map[string]int32{
			"<|endoftext|>":   100257,
			"<|fim_prefix|>":  100258,
			"<|fim_middle|>":  100259,
			"<|fim_suffix|>":  100260,
			"<|endofprompt|>": 100276,
		})

}

//go:embed vocab.gob
var vocabData []byte

//go:embed reverse.gob
var reverseData []byte
