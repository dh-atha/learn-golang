package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("hello world")

	// var nama bool = true
	// namaku := int8(1)
	// fmt.Println(nama)
	// fmt.Println(namaku)

	namaMap := make(map[string]string)
	namaMap["backend"] = "atha"
	namaMap["frontend"] = "narko"

	var hobi [2]string
	hobi[0] = "bola"
	hobi[1] = "basket"
	fmt.Println(len(hobi))

	var hobiSlice []string
	hobiSlice = append(hobiSlice, "bola")
	hobiSlice[0] = "basket"
	fmt.Println(hobiSlice)

	var cobaInterface interface{}
	cobaInterface = "1"
	cobaInterface = 1

	switch cobaInterface.(type) {
	case string:
		fmt.Println(len(cobaInterface.(string)))
	default:
		fmt.Println("bukan string")
	}

	nilai := 2
	if nilai == 1 {
		fmt.Println("nilai 1")
	} else {
		fmt.Println("nilai bukan 1")
	}

	namaArray := []string{"atha", "df"}
	for i := 0; i < len(namaArray); i++ {
		fmt.Println(namaArray[i])
	}

	var total int
	for total < 5 {
		total++
		printan := fmt.Sprintf("%d. halo", total)
		fmt.Println(printan)
	}

	for idx, v := range namaArray {
		fmt.Println(idx, v)
	}

	nama1 := "atha"
	nama2 := &nama1

	fmt.Println(nama1)
	*nama2 = "dhaiffathin"
	fmt.Println(nama1)

	// example json
	jsonString := `{
		"nama": "atha"
	}`

	var jsonMap = map[string]string{
		"nama": "halo",
	}
	json.Unmarshal([]byte(jsonString), &jsonMap)
	jsonData, _ := json.Marshal(jsonMap)
	fmt.Println(string(jsonData))

	fmt.Println(jsonMap)
}
