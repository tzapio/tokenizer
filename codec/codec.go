package codec

import (
	"fmt"
	"math"

	"github.com/dlclark/regexp2"
)

type Codec struct {
	name              string
	vocabulary        Vocab
	reverseVocabulary Reverse
	splitRegexp       *regexp2.Regexp
	specialTokens     map[string]int32
}

func New(name string, vocabulary Vocab, reverse Reverse, splitRegexp *regexp2.Regexp, specialTokens map[string]int32) *Codec {
	return &Codec{
		name, vocabulary, reverse, splitRegexp, specialTokens,
	}
}
func (c *Codec) GetName() string {
	return c.name
}

func (c *Codec) Encode(input string) ([]int32, []string, error) {
	var (
		ids    []int32
		tokens []string
	)
	match, err := c.splitRegexp.FindStringMatch(input)
	if err != nil {
		return nil, nil, fmt.Errorf("error matching: %v", err)
	}

	for match != nil {
		piece := match.String()
		if id, ok := c.vocabulary[piece]; ok {
			ids = append(ids, id)
			tokens = append(tokens, piece)
		} else {
			newIds, newTokens := c.bpe([]byte(piece))
			ids = append(ids, newIds...)
			tokens = append(tokens, newTokens...)
		}
		m, err := c.splitRegexp.FindNextMatch(match)
		if err != nil {
			return nil, nil, fmt.Errorf("error matching: %v", err)
		}
		match = m
	}
	return ids, tokens, nil
}

func (c *Codec) Decode(tokens []int32) (string, error) {
	var out string
	for _, t := range tokens {
		piece, ok := c.reverseVocabulary[t]
		if !ok {
			return "", fmt.Errorf("invalid token: %d", t)
		}
		out += piece
	}
	return out, nil
}

func (c *Codec) bpe(piece []byte) ([]int32, []string) {
	type part struct {
		offset int
		rank   int32
	}

	parts := make([]part, len(piece)+1)
	for i := 0; i < len(parts); i++ {
		parts[i] = part{i, math.MaxInt32}
	}

	getRank := func(index int32, skip int32) int32 {
		if int(index)+int(skip)+2 < len(parts) {
			start := parts[index].offset
			end := parts[index+skip+2].offset
			if rank, ok := c.vocabulary[string(piece[start:end])]; ok {
				return rank
			}
		}
		return math.MaxInt32
	}

	for i := 0; i < len(parts)-2; i++ {
		parts[i].rank = getRank(int32(i), 0)
	}

	for {
		if len(parts) == 1 {
			break
		}

		minRank := int32(math.MaxInt32)
		minIndex := int32(0)
		for i, p := range parts[:len(parts)-1] {
			if p.rank < minRank {
				minRank = p.rank
				minIndex = int32(i)
			}
		}

		if minRank == math.MaxInt32 {
			break
		}

		parts[minIndex].rank = getRank(minIndex, 1)

		if minIndex > 0 {
			parts[minIndex-1].rank = getRank(minIndex-1, 1)
		}

		parts = append(parts[:minIndex+1], parts[minIndex+2:]...)
	}

	ids := make([]int32, len(parts)-1)
	tokens := make([]string, len(parts)-1)
	for i := 0; i < len(ids); i++ {
		token := string(piece[parts[i].offset:parts[i+1].offset])
		tokens[i] = token
		ids[i] = c.vocabulary[token]
	}
	return ids, tokens
}
