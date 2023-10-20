package main

import (
	"reflect"
	"testing"
)

type ExpectedStruct struct {
	Comp71 IndexStructure `json:"comp-7-s-2021.11.22"`
	Comp72 IndexStructure `json:"comp-7-s-2021.11.23"`
}

func TestUnmarshalJSON(t *testing.T) {
	filePath := "test_data.json"
	indexData, err := UnmarshalJSON(filePath)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}

	expectedIndexData := IndexFormat{
		Comp71: IndexStructure{
			Settings: Settings{
				Index: Index{
					RefreshInterval: "1s",
					NumberOfShards:  "3",
				},
			},
		},
		Comp72: IndexStructure{
			Settings: Settings{
				Index: Index{
					RefreshInterval: "2s",
					NumberOfShards:  "5",
				},
			},
		},
	}

	// Perform deep equality check between expected struct and unmarshalled struct
	if !reflect.DeepEqual(indexData, expectedIndexData) {
		t.Errorf("Unmarshalled struct does not match the expected struct. Expected: %+v, Actual: %+v", expectedIndexData, indexData)
	}

}
