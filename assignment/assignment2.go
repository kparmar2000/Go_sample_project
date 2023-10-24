package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// type Index struct {
// 	settings Settings `json:"settings"`
// }
// type Settings struct {
// 	Index struct {
// 		RefreshInterval  string   `json:"refresh_interval"`
// 		NumberOfShards   string   `json:"number_of_shards"`
// 		ProvidedName     string   `json:"provided_name"`
// 		CreationDate     string   `json:"creation_date"`
// 		analysis         Analysis `json:"analysis"`
// 		NumberOfReplicas string   `json:"number_of_replicas"`
// 		UUID             string   `json:"uuid"`
// 		Version          struct {
// 			Created string `json:"created"`
// 		} `json:"version"`
// 	} `json:"index"`
// }
// type Analysis struct {
// 	normalizer Normalizer `json:"normalizer"`
// 	analyzer   Analyzer   `json:"analyzer"`
// 	tokenizer  Tokenizer  `json:"tokenizer"`
// }
// type Normalizer struct {
// 	CaseInsensitive struct {
// 		Filter     []string `json:"filter"`
// 		Type       string   `json:"type"`
// 		CharFilter []string `json:"char_filter"`
// 	} `json:"case_insensitive"`
// }
// type Analyzer struct {
// 	autocomplete               Autocomplete               `json:"autocomplete"`
// 	autocompleteVersionNumbers AutocompleteVersionNumbers `json:"autocomplete_version_numbers"`
// }
// type Autocomplete struct {
// 	Filter    []string `json:"filter"`
// 	Tokenizer string   `json:"tokenizer"`
// }
// type AutocompleteTokenizer struct {
// 	TokenChars []string `json:"token_chars"`
// 	MinGram    string   `json:"min_gram"`
// 	Type       string   `json:"type"`
// 	MaxGram    string   `json:"max_gram"`
// }
// type AutocompleteVersionNumberTokenizer struct {
// 	TokenChars []string `json:"token_chars"`
// 	MinGram    string   `json:"min_gram"`
// 	Type       string   `json:"type"`
// 	MaxGram    string   `json:"max_gram"`
// }
// type AutocompleteVersionNumbers struct {
// 	Filter    []string `json:"filter"`
// 	Tokenizer string   `json:"tokenizer"`
// }
// type Tokenizer struct {
// 	autocompleteVersionNumberTokenizer AutocompleteVersionNumberTokenizer `json:"autocomplete_version_number_tokenizer"`
// 	autocompleteTokenizer              AutocompleteTokenizer              `json:"autocomplete_tokenizer"`
// }

//	type IndexFormat struct {
//		Comp71 Index `json:"comp-7-s-2021.11.22"`
//		Comp72 Index `json:"comp-7-s-2021.11.23"`
//	}
type AutocompleteTokenizer struct {
	TokenChars []string `json:"token_chars"`
	MinGram    string   `json:"min_gram"`
	Type       string   `json:"type"`
	MaxGram    string   `json:"max_gram"`
}

type Analyzer struct {
	Autocomplete               Autocomplete `json:"autocomplete"`
	AutocompleteVersionNumbers Autocomplete `json:"autocomplete_version_numbers"`
}

type Autocomplete struct {
	Filter    []string `json:"filter"`
	Tokenizer string   `json:"tokenizer"`
}

type Normalizer struct {
	CaseInsensitive CaseInsensitive `json:"case_insensitive"`
}
type CaseInsensitive struct {
	Filter     []string `json:"filter"`
	Type       string   `json:"type"`
	CharFilter []string `json:"char_filter"`
}

type Analysis struct {
	Normalizer Normalizer `json:"normalizer"`
	Analyzer   Analyzer   `json:"analyzer"`
	Tokenizer  Tokenizer  `json:"tokenizer"`
}
type Tokenizer struct {
	AutocompleteVersionNumberTokenizer AutocompleteTokenizer `json:"autocomplete_version_number_tokenizer"`
	AutocompleteTokenizer              AutocompleteTokenizer `json:"autocomplete_tokenizer"`
}

type Version struct {
	Created string `json:"created"`
}

type Index struct {
	RefreshInterval  string   `json:"refresh_interval"`
	NumberOfShards   string   `json:"number_of_shards"`
	ProvidedName     string   `json:"provided_name"`
	CreationDate     string   `json:"creation_date"`
	Analysis         Analysis `json:"analysis"`
	NumberOfReplicas string   `json:"number_of_replicas"`
	UUID             string   `json:"uuid"`
	Version          Version  `json:"version"`
}

type Settings struct {
	Index Index `json:"index"`
}

type IndexStructure struct {
	Settings Settings `json:"settings"`
}

type IndexFormat struct {
	Comp71 IndexStructure `json:"comp-7-s-2021.11.22"`
	Comp72 IndexStructure `json:"comp-7-s-2021.11.23"`
}

func main() {

	indexData, err := UnmarshalJSON("assign.json")
	if err != nil {
		log.Fatalf("Error in UnMarshalJson: %v", err)
	}

	fmt.Println("Comp-7-s-2021.11.22:")
	fmt.Println("Refresh Interval:", indexData.Comp71.Settings.Index.RefreshInterval)
	fmt.Println("Number of Shards:", indexData.Comp71.Settings.Index.NumberOfShards)

	fmt.Println("Comp-7-s-2021.11.23:")
	fmt.Println("Refresh Interval:", indexData.Comp72.Settings.Index.RefreshInterval)
	fmt.Println("Number of Shards:", indexData.Comp72.Settings.Index.NumberOfShards)
	// jsonString, err := json.MarshalIndent(indexData, "", " ")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	//fmt.Println(string(jsonString))
}
func UnmarshalJSON(filePath string) (IndexFormat, error) {

	content, err := JsonData(filePath)
	if err != nil {
		return IndexFormat{}, fmt.Errorf("error reading file: %v", err)
	}

	var jsonData IndexFormat
	//fmt.Println(jsonData)
	err = json.Unmarshal([]byte(content), &jsonData)
	if err != nil {
		return IndexFormat{}, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return jsonData, nil
}
func JsonData(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "error in opening file.", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "error in reading file.", fmt.Errorf("error reading file: %v", err)
	}
	fmt.Print(string(content))
	return string(content), nil
}
