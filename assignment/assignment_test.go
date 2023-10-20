package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	// jsonData := []byte(`{
	// 	"comp-7-s-2021.11.22": {
	// 		"settings": {
	// 			"index": {
	// 				"refresh_interval": "1s",
	// 				"number_of_shards": "5",
	// 				"provided_name": "comp-7-s-2021.11.22",
	// 				"creation_date": "1637661228822",
	// 				"analysis": {
	// 					"normalizer": {
	// 						"case_insensitive": {
	// 							"filter": ["lowercase", "asciifolding"],
	// 							"type": "custom",
	// 							"char_filter": []
	// 						}
	// 					},
	// 					"analyzer": {
	// 						"autocomplete": {
	// 							"filter": ["lowercase"],
	// 							"tokenizer": "autocomplete_tokenizer"
	// 						},
	// 						"autocomplete_version_numbers": {
	// 							"filter": ["lowercase"],
	// 							"tokenizer": "autocomplete_version_number_tokenizer"
	// 						}
	// 					},
	// 					"tokenizer": {
	// 						"autocomplete_version_number_tokenizer": {
	// 							"token_chars": ["letter", "digit", "punctuation"],
	// 							"min_gram": "2",
	// 							"type": "edge_ngram",
	// 							"max_gram": "20"
	// 						},
	// 						"autocomplete_tokenizer": {
	// 							"token_chars": ["letter", "digit"],
	// 							"min_gram": "2",
	// 							"type": "edge_ngram",
	// 							"max_gram": "20"
	// 						}
	// 					}
	// 				},
	// 				"number_of_replicas": "1",
	// 				"uuid": "w72cfW0GR8uCG8snqUI7Dg",
	// 				"version": {
	// 					"created": "6081499"
	// 				}
	// 			}
	// 		}
	// 	}
	// }`)

	file, err := os.Open("assign.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var indexData IndexFormat

	err1 := json.Unmarshal(content, &indexData)
	if err1 != nil {
		t.Errorf("Error unmarshalling JSON data: %v", err1)
	}

	expectedRefreshInterval := "1s"
	if indexData.Comp71.Settings.Index.RefreshInterval != expectedRefreshInterval {
		t.Errorf("Expected RefreshInterval %s, but got %s", expectedRefreshInterval, indexData.Comp71.Settings.Index.RefreshInterval)
	}

	expectedNumberOfShards := "5"
	if indexData.Comp71.Settings.Index.NumberOfShards != expectedNumberOfShards {
		t.Errorf("Expected NumberOfShards %s, but got %s", expectedNumberOfShards, indexData.Comp71.Settings.Index.NumberOfShards)
	}

}
