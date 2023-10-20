package main

import (
	"testing"
)

type ExpectedStruct struct {
	Comp71 IndexStructure `json:"comp-7-s-2021.11.22"`
	Comp72 IndexStructure `json:"comp-7-s-2021.11.23"`
}

func TestUnmarshalJSON(t *testing.T) {
	filePath := "assign.json"
	indexData, err := UnmarshalJSON(filePath)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}
	expectedComp71RefreshInterval := "1s"
	expectedComp72NumberOfShards := "5"
	if indexData.Comp71.Settings.Index.RefreshInterval != expectedComp71RefreshInterval {
		t.Errorf("Expected Comp-7-s-2021.11.22 RefreshInterval %s, but got %s", expectedComp71RefreshInterval, indexData.Comp71.Settings.Index.RefreshInterval)
	}

	if indexData.Comp72.Settings.Index.NumberOfShards != expectedComp72NumberOfShards {
		t.Errorf("Expected Comp-7-s-2021.11.23 NumberOfShards %s, but got %s", expectedComp72NumberOfShards, indexData.Comp72.Settings.Index.NumberOfShards)
	}

	// var expected ExpectedStruct
	// fmt.Println(expected)

	// if !reflect.DeepEqual(indexData, expected) {
	// 	t.Errorf("Unmarshalled struct does not match the expected struct. Expected: %v,\n Actual: %v", expected, indexData)
	// }

}
