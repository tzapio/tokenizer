package tokenizer

// Package tokenizer provides functions for encoding and decoding text using
// different tokenization schemes.
//
// Encoding Formats
//
// The following encoding formats are supported:
// - Cl100kBase
// - R50kBase
// - P50kBase
// - P50kEdit
//
// Alternatively you can request a tokenizer using OpenAI's model name, the
// following OpenAI models are supported:
// - GPT4
// - GPT35Turbo
// - TextEmbeddingAda002
// - TextDavinci003
// - TextDavinci002
// - CodeDavinci002
// - CodeDavinci001
// - CodeCushman002
// - CodeCushman001
// - DavinciCodex
// - CushmanCodex
// - TextDavinci001
// - TextCurie001
// - TextBabbage001
// - TextAda001
// - Davinci
// - Curie
// - Babbage
// - Ada
// - TextSimilarityDavinci001
// - TextSimilarityCurie001
// - TextSimilarityBabbage001
// - TextSimilarityAda001
// - TextSearchDavinciDoc001
// - TextSearchCurieDoc001
// - TextSearchAdaDoc001
// - TextSearchBabbageDoc001
// - CodeSearchBabbageCode001
// - CodeSearchAdaCode001
// - TextDavinciEdit001
// - CodeDavinciEdit001
//
// Usage Example
//
// Here is an example of how to encode a string using the `ForModel` function:
//
//	package main
//
//	import (
//		"fmt"
//		"github.com/tiktoken-go/tokenizer"
//	)
//
//	func main() {
//		enc, err := tokenizer.Get(tokenizer.Cl100kBase)
//		if err != nil {
//			panic("oh oh")
//		}
//
//		// this should print a list of token ids
//		ids, token, _ := enc.Encode("supercalifragilistic")
//		fmt.Println(ids)
//
//		// this should print the original string back
//		text, _ := enc.Decode(ids)
//		fmt.Println(text)
//}

import (
	"errors"

	"github.com/tzapio/tokenizer/codec/cl100k_base"
	"github.com/tzapio/tokenizer/codec/p50k_base"
	"github.com/tzapio/tokenizer/codec/r50k_base"
)

var (
	ErrModelNotSupported    = errors.New("model not supported")
	ErrEncodingNotSupported = errors.New("encoding not supported")
)

type Codec interface {
	GetName() string
	Encode(string) ([]uint, []string, error)
	Decode([]uint) (string, error)
}

type Model string

const (
	GPT4                     Model = "gpt-4"
	GPT35Turbo               Model = "gpt-3.5-turbo"
	TextEmbeddingAda002      Model = "text-embedding-ada-002"
	TextDavinci003           Model = "text-davinci-003"
	TextDavinci002           Model = "text-davinci-002"
	CodeDavinci002           Model = "code-davinci-002"
	CodeDavinci001           Model = "code-davinci-001"
	CodeCushman002           Model = "code-cushman-002"
	CodeCushman001           Model = "code-cushman-001"
	DavinciCodex             Model = "davinci-codex"
	CushmanCodex             Model = "cushman-codex"
	TextDavinci001           Model = "text-davinci-001"
	TextCurie001             Model = "text-curie-001"
	TextBabbage001           Model = "text-babbage-001"
	TextAda001               Model = "text-ada-001"
	Davinci                  Model = "davinci"
	Curie                    Model = "curie"
	Babbage                  Model = "babbage"
	Ada                      Model = "ada"
	TextSimilarityDavinci001 Model = "text-similarity-davinci-001"
	TextSimilarityCurie001   Model = "text-similarity-curie-001"
	TextSimilarityBabbage001 Model = "text-similarity-babbage-001"
	TextSimilarityAda001     Model = "text-similarity-ada-001"
	TextSearchDavinciDoc001  Model = "text-search-davinci-doc-001"
	TextSearchCurieDoc001    Model = "text-search-curie-doc-001"
	TextSearchAdaDoc001      Model = "text-search-ada-doc-001"
	TextSearchBabbageDoc001  Model = "text-search-babbage-doc-001"
	CodeSearchBabbageCode001 Model = "code-search-babbage-code-001"
	CodeSearchAdaCode001     Model = "code-search-ada-code-001"
	TextDavinciEdit001       Model = "text-davinci-edit-001"
	CodeDavinciEdit001       Model = "code-davinci-edit-001"
	GPT2                     Model = "gpt2"
)

type Encoding string

const (
	GPT2Enc    Encoding = "gpt2"
	R50kBase   Encoding = "r50k_base"
	P50kBase   Encoding = "p50k_base"
	P50kEdit   Encoding = "p50k_edit"
	Cl100kBase Encoding = "cl100k_base"
)

// Get returns a new instance of a Codec implementation based on the specified
// encoding format. The returned Codec instance can be used to encode (tokenize)
// and decode (reassemble) text. If the specified encoding is not supported,
// an error is returned.
func Get(encoding Encoding) (Codec, error) {
	switch encoding {
	case Cl100kBase:
		return cl100k_base.NewCl100kBase(), nil
	case R50kBase:
		return r50k_base.NewR50kBase(), nil
	case P50kBase:
		return p50k_base.NewP50kBase(), nil
	case P50kEdit:
		return p50k_base.NewP50kEdit(), nil
	default:
		return nil, ErrEncodingNotSupported
	}
}

// ForModel returns a new instance of a Codec implementation based on the
// specified OpenAI model. If the specified model is not supported, an error
// is returned.
func ForModel(model Model) (Codec, error) {
	switch model {
	case GPT4, GPT35Turbo, TextEmbeddingAda002:
		return Get(Cl100kBase)

	case TextDavinci003, TextDavinci002, CodeDavinci001,
		CodeDavinci002, CodeCushman002, CodeCushman001,
		DavinciCodex, CushmanCodex:
		return Get(P50kBase)

	case TextDavinci001, TextCurie001, TextBabbage001, TextAda001, Davinci,
		Curie, Babbage, Ada, TextSimilarityDavinci001, TextSimilarityCurie001,
		TextSimilarityBabbage001, TextSimilarityAda001, TextSearchDavinciDoc001,
		TextSearchCurieDoc001, TextSearchAdaDoc001, TextSearchBabbageDoc001,
		CodeSearchBabbageCode001, CodeSearchAdaCode001:
		return Get(R50kBase)

	case TextDavinciEdit001, CodeDavinciEdit001:
		return Get(P50kEdit)

	case GPT2:
		return Get(GPT2Enc)
	default:
		return nil, ErrModelNotSupported
	}
}
