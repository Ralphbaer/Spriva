package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("./gtfs-files/agency.txt")
	var dataSlice = make([]Agency, 0)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		keyVal := strings.Split(line, "=")
		dataSlice = append(dataSlice, Element{Name: keyVal[0], Value: keyVal[1]})
	}
	bts, err := json.Marshal(dataSlice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", bts)
	fmt.Print(reflect.TypeOf(string(dat)))
}
