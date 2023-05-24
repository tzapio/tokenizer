package r50k_base

import (
	"github.com/dlclark/regexp2"
	"github.com/tzapio/tokenizer/codec"
)

func NewR50kBase() *codec.Codec {
	return codec.New(
		"r50k_base",
		r50kBaseVocab,
		regexp2.MustCompile(`'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+(?!\S)|\s+`, regexp2.None),
		map[string]uint{
			"<|endoftext|>": 50256,
		})
}
