package p50k_base

import (
	"github.com/dlclark/regexp2"
	"github.com/tzapio/tokenizer/codec"
)

func NewP50kEdit() *codec.Codec {
	return codec.New(
		"p50k_edit",
		p50kBaseVocab,
		regexp2.MustCompile(`'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+(?!\S)|\s+`, regexp2.None),
		map[string]uint{
			"<|endoftext|>":  50256,
			"<|fim_prefix|>": 50281,
			"<|fim_middle|>": 50282,
			"<|fim_suffix|>": 50283,
		})
}
