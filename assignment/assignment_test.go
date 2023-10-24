package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExpectedStruct struct {
	Comp71 IndexStructure `json:"comp-7-s-2021.11.22"`
	Comp72 IndexStructure `json:"comp-7-s-2021.11.23"`
}

// func TestUnmarshalJSON(t *testing.T) {
// 	filePath := "assign.json"
// 	indexData, err := UnmarshalJSON(filePath)
// 	if err != nil {
// 		t.Fatalf("Error unmarshalling JSON: %v", err)
// 	}

// 	assert.True(t, indexData, "The provided file string is not correct")
// 	// var expected ExpectedStruct
// 	// var indexData2 map[string]interface{}
// 	// file, err := os.Open(filePath)
// 	// if err != nil {
// 	// 	fmt.Errorf("error opening file: %v", err)
// 	// }
// 	// defer file.Close()

// 	// content, err := io.ReadAll(file)
// 	// if err != nil {
// 	// 	fmt.Errorf("error reading file: %v", err)
// 	// }
// 	// err1 := json.Unmarshal(content, &indexData2)
// 	// if err1 != nil {
// 	// 	fmt.Println("Error found while unmarshal:", err1)
// 	// 	return
// 	// }
// 	//fmt.Println(indexData2["comp-7-s-2021.11.22"])
// 	fmt.Println(indexData.Comp71)
// 	expectedComp71RefreshInterval := "1s"
// 	expectedComp72NumberOfShards := "5"
// 	if indexData.Comp71.Settings.Index.RefreshInterval != expectedComp71RefreshInterval {
// 		t.Errorf("Expected Comp-7-s-2021.11.22 RefreshInterval %s, but got %s", expectedComp71RefreshInterval, indexData.Comp71.Settings.Index.RefreshInterval)
// 	}

// 	if indexData.Comp72.Settings.Index.NumberOfShards != expectedComp72NumberOfShards {
// 		t.Errorf("Expected Comp-7-s-2021.11.23 NumberOfShards %s, but got %s", expectedComp72NumberOfShards, indexData.Comp72.Settings.Index.NumberOfShards)
// 	}

// 	// if !reflect.DeepEqual(indexData, expected) {
// 	// 	t.Errorf("Unmarshalled struct does not match the expected struct. Expected: %v,\n Actual: %v", expected, indexData)
// 	// }

// }

func TestUnmarshalJSON(t *testing.T) {
	filePath := "assign12.json"
	isvalid := assert.FileExists(t, filePath, "dir does not exist")
	assert.True(t, isvalid, "Error:file does not exist")

	indexData, err := UnmarshalJSON(filePath)
	assert.NoError(t, err, "Error unmarshalling JSON")

	expectedComp71RefreshInterval := "1s"
	expectedComp72NumberOfShards := "5"

	assert.Equal(t, expectedComp71RefreshInterval, indexData.Comp71.Settings.Index.RefreshInterval, "Unexpected Comp-7-s-2021.11.22 RefreshInterval")
	assert.Equal(t, expectedComp72NumberOfShards, indexData.Comp72.Settings.Index.NumberOfShards, "Unexpected Comp-7-s-2021.11.23 NumberOfShards")

}
