package main

import (
	"fmt"
	emb "transformer/embeddings"
)

func main() {
	//vocabSize := 2
	embedDim := 3
	seqLen := 5
	fmt.Println(emb.PositionalEncoding(seqLen, embedDim))
	//fmt.Println(GenerateEmbeddings(vocabSize, embedDim))
}
