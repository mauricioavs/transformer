package embeddings

import (
	"math"

	"golang.org/x/exp/rand"
)

// generates a random matrix with dimensions (vocabSize, embedDim) with values between 0 and 1
func GenerateEmbeddings(vocabSize, embedDim int) [][]float64 {
	embeddings := make([][]float64, vocabSize) // [[], [], ..., []] with vocabSize length
	for i := range embeddings {
		embeddings[i] = make([]float64, embedDim) // [0,0,...,0] for each subarray with length embedDim
		for j := range embeddings[i] { embeddings[i][j] = rand.Float64() } // [0.1, 0.3,...,0.5] fill with random values between 0 and 1
	}
	return embeddings
}

// calculates the positional encoding matrix
func PositionalEncoding(seqLen, embedDim int) [][]float64 {
	posEnc := make([][]float64, seqLen) // [[], [], ..., []] with seqLen length
	for pos := range seqLen {
		posEnc[pos] = make([]float64, embedDim) // [0,0,...,0] for each subarray with length embedDim
		for i := range embedDim {
			divisor := math.Pow(10_000, float64(2*i)/float64(embedDim))
			argument := float64(pos) / divisor
			if i%2 == 0 {
				posEnc[pos][i] = math.Sin(argument)
			} else {
				posEnc[pos][i] = math.Cos(argument)
			}
		}
	}
	return posEnc
}

func AddEmbeddingsAndPositional(embeddings, posEnc [][]float64) [][]float64 {
	seqLen := len(embeddings)
	embedDim := len(embeddings[0])
	result := make([][]float64, seqLen)
	for i := range seqLen {
		result[i] = make([]float64, embedDim)
		for j := 0; j < embedDim; j++ {
			result[i][j] = embeddings[i][j] + posEnc[i][j]
		}
	}
	return result
}
