# Rag-FAQBot
FAQBot made with custom RAG pipeline implemented manually using raw GO &amp; maths using TF-IDF algorithm and Cosine Similarity for search.
# Features
> Custom RAG pipeline implemented purley in Go with out thridparty libraries or bindings
> No API keys no dependancies, runs entrirly offline
> Use Tf-IDF for smarter matching 

# Getting Started
1. Clone the repo: git clone https://github.com/tanush94i/Rag-FAQBot 
2. Add FAQs to the faq.txt file (separate each questions with @@@)
3. Run the Bot: go run main.go 
4. Ask the the query.

# How it works:
1. Reads & splits the FAQ file
2. Builds TF-IDF vectors for search
3. Uses cosine similarity to match queries
4. Returns the best answer (or says " Sorry, I couldn't find a relevant answer." )

# License
MIT – Use it, tweak it, build on it! 

