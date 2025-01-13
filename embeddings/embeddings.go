package embeddings

import (
	"math"
	"time"

	"golang.org/x/exp/rand"
)

// generates a random matrix with dimensions (vocabSize, embedDim) with values between 0 and 1
func GenerateEmbeddings(vocabSize, embedDim int) [][]float64 {
	rand.Seed(uint64(time.Now().UnixNano()))
	embeddings := make([][]float64, vocabSize) // [[], [], ..., []] with vocabSize length
	for i := range embeddings {
		embeddings[i] = make([]float64, embedDim) // [0,0,...,0] for each subarray with length embedDim
		for j := range embeddings[i] {
			embeddings[i][j] = rand.Float64() // [0.1, 0.3,...,0.5] fill with random values between 0 and 1
		}

	}
	return embeddings
}

// calculates the positional encoding matrix
func PositionalEncoding(seqLen, embedDim int) [][]float64 {
	posEnc := make([][]float64, seqLen) // [[], [], ..., []] with seqLen length
	for pos := 0; pos < seqLen; pos++ {
		posEnc[pos] = make([]float64, embedDim) // [0,0,...,0] for each subarray with length embedDim
		for i := 0; i < embedDim; i++ {
			divisor := math.Pow(10_000, float64(2*i)/float64(embedDim))
			if i%2 == 0 {
				posEnc[pos][i] = math.Sin(float64(pos) / divisor)
			} else {
				posEnc[pos][i] = math.Cos(float64(pos) / divisor)
			}
		}
	}
	return posEnc
}

func AddEmbeddingsAndPositional(embeddings, posEnc [][]float64) [][]float64 {
	seqLen := len(embeddings)
	embedDim := len(embeddings[0])
	result := make([][]float64, seqLen)
	for i := 0; i < seqLen; i++ {
		result[i] = make([]float64, embedDim)
		for j := 0; j < embedDim; j++ {
			result[i][j] = embeddings[i][j] + posEnc[i][j]
		}
	}
	return result
}
