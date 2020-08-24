package json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Province string `json:"province"`
	City string `json:"city"`
	No int32  `json:"no"`
}
type IDCard struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
	Address []*Address `json:"address"`
	Remark string `json:"remark"`
}


func Json() {
	addr1 := &Address{"None", "ShangHai", 002}
	addr2 := &Address{"HeiLongJiang", "HarBin", 015}
	IdCard := IDCard{"Kphilleani", 24, []*Address{addr1, addr2}, "FE"}

	js, _ := json.Marshal(IdCard)
	fmt.Printf("JSON format %s:", js)

	file, _ := os.OpenFile("IDCARD.json", os.O_CREATE | os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(IdCard)
	if err != nil {
		log.Println("Error in encoding json")
	}
}