package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Document struct {
	Text   string
	Vector map[string]float64
}

var documents []Document
var vocabulary map[string]float64

// read and segment the FAQ file
func segmentFAQ(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	segments := strings.Split(string(data), "@@@")
	for i := range segments {
		segments[i] = strings.TrimSpace(segments[i])
	}
	return segments
}

// tokenize the text
func tokenize(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

// make vocabulary and calculate TF-IDF vectors
func buildVectors(docs []string) {
	vocabulary = make(map[string]float64)
	docFreq := make(map[string]int)
	totalDocs := len(docs)

	for _, doc := range docs {
		wordSet := make(map[string]bool)
		words := tokenize(doc)
		for _, word := range words {
			wordSet[word] = true
		}
		for word := range wordSet {
			docFreq[word]++
		}
	}

	for word, count := range docFreq {
		vocabulary[word] = math.Log(float64(totalDocs) / float64(count+1))
	}

	for _, doc := range docs {
		vec := make(map[string]float64)
		words := tokenize(doc)
		termFreq := make(map[string]float64)
		for _, word := range words {
			termFreq[word]++
		}
		for word, freq := range termFreq {
			vec[word] = freq * vocabulary[word]
		}
		documents = append(documents, Document{Text: doc, Vector: vec})
	}
}

// cosine similarity between two vectors
func cosineSimilarity(vec1, vec2 map[string]float64) float64 {
	dotProduct := 0.0
	normA := 0.0
	normB := 0.0

	for word, weight := range vec1 {
		normA += weight * weight
		if val, ok := vec2[word]; ok {
			dotProduct += weight * val
		}
	}
	for _, weight := range vec2 {
		normB += weight * weight
	}
	if normA == 0 || normB == 0 {
		return 0
	}
	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}

// search for the FAQ
func searchFAQ(query string) string {
	queryVec := make(map[string]float64)
	words := tokenize(query)
	for _, word := range words {
		if idf, ok := vocabulary[word]; ok {
			queryVec[word]++
			queryVec[word] *= idf
		}
	}

	type result struct {
		doc     string
		similar float64
	}
	var results []result

	for _, doc := range documents {
		sim := cosineSimilarity(queryVec, doc.Vector)
		results = append(results, result{doc: doc.Text, similar: sim})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].similar > results[j].similar
	})

	if len(results) > 0 && results[0].similar > 0 {
		return results[0].doc
	}
	return "Sorry, I couldn't find a relevant answer."
}

func main() {
	filename := "faq.txt"
	segments := segmentFAQ(filename)
	buildVectors(segments)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ask me a question (type 'exit' to quit):")
	for {
		fmt.Print("You: ")
		scanner.Scan()
		query := scanner.Text()
		if query == "exit" {
			break
		}
		answer := searchFAQ(query)
		fmt.Println("Bot:", answer)
	}
}
