package main

import (
	"fmt"
	"transformer/matrix"
)

func main() {
	// rand.Seed(uint64(time.Now().UnixNano()))
	// // Configuración
	// //vocabSize := 1000 // Tamaño del vocabulario
	// embedDim := 5 // Dimensión del embedding
	// seqLen := 3   // Longitud de la secuencia

	// // Generar embeddings y positional encoding
	// embeddings := emb.GenerateEmbeddings(seqLen, embedDim)
	// posEnc := emb.PositionalEncoding(seqLen, embedDim)

	// // Sumar embeddings con positional encoding
	// finalEmbeddings := embeddings + posEnc

	// // Mostrar resultados
	// fmt.Println("Embeddings:")
	// fmt.Println(embeddings)
	// fmt.Println("Positional Encoding:")
	// fmt.Println(posEnc)
	// fmt.Println("\nEmbeddings finales:")
	// fmt.Println(finalEmbeddings)
	a := matrix.Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("matrix a")
	fmt.Println(a.Rows()) // 2
	fmt.Println(a.Cols()) // 3
	fmt.Println(a.Dims()) // 2,3
	b := matrix.Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("matrix b")
	fmt.Println(b.Rows()) // 2
	fmt.Println(b.Cols()) // 3
	fmt.Println(b.Dims()) // 2,3
	fmt.Println("sum a and b")
	fmt.Println(a.Add(b))
}
