package main

import (
	"encoding/json"
	"github.com/boltdb/bolt"
)
type Posting map[string]int

func addToIndex(db *bolt.DB, word, docID string)error{
	return db.Update(func(tx *bolt.Tx)error{
		b,err:=tx.CreateBucketIfNotExists([]byte("index"))

		if err!=nil{
			return err
		}

		var posting Posting

		v:=b.Get([]byte(word))
		if v!=nil{
			json.Unmarshal(v, &posting)
		}else{
			posting=make(Posting)
		}

		posting[docID]++

		data,_:=json.Marshal(posting)
		return b.Put([]byte(word),data)
	})

}


func search(db *bolt.DB, query string) map[string]int {
    results := make(map[string]int)
    tokens := tokenize(query)

    db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("index"))
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

            for doc, score := range posting {
                results[doc] += score
            }
        }
        return nil
    })

    return results
}