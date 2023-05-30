package codec

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vocab map[string]int32
type Reverse map[int32]string

func SaveVocabToFile(filename string, vocab Vocab) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(vocab)
	if err != nil {
		log.Fatalf("Failed to serialize data: %s", err)
	}
}
func SaveReverseToFile(filename string, reverse Reverse) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(reverse)
	if err != nil {
		log.Fatalf("Failed to serialize data: %s", err)
	}
}
func LoadVocabFromByteArr(data []byte) Vocab {
	var vocab Vocab
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	err := decoder.Decode(&vocab)
	if err != nil {
		log.Fatalf("Failed to deserialize data: %s", err)
	}

	return vocab
}

func LoadReverseFromByteArr(data []byte) Reverse {
	var reverse Reverse
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	err := decoder.Decode(&reverse)
	if err != nil {
		log.Fatalf("Failed to deserialize data: %s", err)
	}

	return reverse
}

func GetReverse(vocab Vocab) Reverse {
	var reverse Reverse = make(map[int32]string)
	for k, v := range vocab {
		reverse[v] = k
	}
	return reverse
}
func LoadVocab(filename string) Vocab {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	var vocab Vocab = make(map[string]int32)

	scanner := bufio.NewScanner(file)

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

		v, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic("invalid word: " + parts[1])
		}
		vocab[string(word)] = int32(v)
	}

	return vocab
}
