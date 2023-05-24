package cl100k_base

import (
	"github.com/dlclark/regexp2"
	"github.com/tzapio/tokenizer/codec"
)

func NewCl100kBase() *codec.Codec {
	return codec.New(
		"cl100k_base",
		cl100kBaseVocab,
		regexp2.MustCompile(`(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}{1,3}| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+`, regexp2.None),
		map[string]uint{
			"<|endoftext|>":   100257,
			"<|fim_prefix|>":  100258,
			"<|fim_middle|>":  100259,
			"<|fim_suffix|>":  100260,
			"<|endofprompt|>": 100276,
		})

}
