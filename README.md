# Tokenizer

This is a pure go port of OpenAI's tokenizer.

## Usage

```go
package main

import (
    "fmt"
    "github.com/tiktoken-go/tokenizer"
)

func main() {
    enc, err := tokenizer.Get(tokenizer.Cl100kBase)
    if err != nil {
        panic("oh oh")
    }

    // this should print a list of token ids
    ids, _, _ := enc.Encode("supercalifragilistic")
    fmt.Println(ids)

    // this should print the original string back
    text, _ := enc.Decode(ids)
    fmt.Println(text)
}
```

Alternatively you can use the included command-line tool

```sh
> tokenizer -h

Usage of tokenizer:
  -decode string
        tokens to decode
  -encode string
        text to encode
  -token string
        text to calculate token

> tokenizer -encode supercalifragilistic
```

## Todo

- ❌ port code
- ✅ cl100k_base encoding
- ❌ r50k_base encoding
- ❌ p50k_base encoding
- ❌ p50k_edit encoding
- ✅ tests
- ❌ handle special tokens
- ❌ gpt-2 model

## Caveats

This library embeds OpenAI's vocabularies—which are not small (~4Mb)— as go
maps. This is different than what the way python version of tiktoken works, 
which downloads the dictionaries and puts them in a cache folder.

However, since the dictionaries are compiled during the go build process
the performance and start-up times should be better than downloading and loading
them at runtime.


