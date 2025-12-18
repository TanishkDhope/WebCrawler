package main

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/boltdb/bolt"
)

type Posting struct {
	DF   int            //Document Frequency
	Docs map[string]int //Inverted Doc Index
}

type Document struct {
	TotalTokens int
}

func addDocument(db *bolt.DB, base string, totalTokens int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("document"))

		if err != nil {
			return err
		}

		data, _ := json.Marshal(Document{
			TotalTokens: totalTokens,
		})

		return b.Put([]byte(base), data)

	})

}

func addToIndex(db *bolt.DB, word, docID string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("index"))

		if err != nil {
			return err
		}

		var posting Posting

		v := b.Get([]byte(word))
		if v != nil {
			json.Unmarshal(v, &posting)
		} else {
			posting = Posting{
				DF:   0,
				Docs: make(map[string]int),
			}
		}

		if _, ok := posting.Docs[docID]; !ok {
			posting.DF++
		}

		posting.Docs[docID]++

		data, _ := json.Marshal(posting)
		return b.Put([]byte(word), data)
	})

}

func search(db *bolt.DB, query string, docs float64) map[string]int {
	results := make(map[string]int)
	scores := make(map[string]float64)
	tokens := tokenize(query)

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("index"))
		buck := tx.Bucket([]byte("document"))

		if b == nil {
			return nil
		}

		for _, token := range tokens {
			v := b.Get([]byte(token))
			if v == nil {
				continue
			}

			var posting Posting
			json.Unmarshal(v, &posting)

			//tfidf score= tf*idf
			//tf = (freq of word in document // total words in document)
			//idf = log(no of docs// (1 + no of docs containing term))
			//For Each token in Query add their scores

			idf := math.Log(docs / float64(1+posting.DF))

			for doc, freq := range posting.Docs {
				//freq for each token in each document

				//s1 calculate the tf score
				var document Document

				json.Unmarshal(buck.Get([]byte(doc)), &document)

				if document.TotalTokens == 0 {
					continue
				}

				totalTokens := document.TotalTokens

				tf := float64(freq) / float64(totalTokens)

				scores[doc] += tf * idf
				results[doc] += freq
			}
		}
		return nil
	})
	fmt.Println(scores)
	return results
}
