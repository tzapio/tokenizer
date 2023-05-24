package cl100k_base

import (
	"bufio"
	"bytes"
	"embed"
	"encoding/base64"
	"log"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/tzapio/tokenizer/codec"
)

func NewCl100kBase() *codec.Codec {
	cl100kBaseVocab := loadVocab("cl100k_base.tiktoken")
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

//go:embed cl100k_base.tiktoken
var vocabData embed.FS

func loadVocab(filename string) codec.Vocab {
	var vocab codec.Vocab = make(map[string]uint)
	data, err := vocabData.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read vocab file: %s", err)
	}
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			log.Fatalf("invalid line: %s", line)
		}

		word, err := base64.StdEncoding.DecodeString(parts[0])
		if err != nil {
			log.Fatalf("invalid word: %s", parts[0])
		}

		v, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			panic("invalid word: " + parts[1])
		}
		vocab[string(word)] = uint(v)
	}

	return vocab
}
