package main

import (
	"log"

	"github.com/tzapio/tokenizer/codec"
)

func main() {
	// Creating a Vocab map
	log.Println("reading tiktoken")
	vocab := codec.LoadVocab("cl100k_base.tiktoken")
	log.Println("making reverse")
	reverse := codec.GetReverse(vocab)
	// Saving the Vocab map to a file
	codec.SaveVocabToFile("vocab.gob", vocab)
	codec.SaveReverseToFile("reverse.gob", reverse)
}
