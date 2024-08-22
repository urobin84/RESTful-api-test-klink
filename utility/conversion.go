package utility

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func StringToInt64(s string) (int64, error) {
	n, err := strconv.ParseInt(s, 10, 64);
	if err != nil {
		fmt.Println(s, "is not an integer 64.")
	}

	return n, err
}

func StringToInt16(s string) (int16, error) {
    n, err := strconv.ParseInt(s, 10, 16); 
	if err != nil {
		fmt.Println(s, "is not an integer 16.")
	}

	i16 := int16(n)
	return i16, err
}

func StringToJSONStruct(s string, v interface{}) error {
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		fmt.Println(s, "is not a json.")
	}

	return err
}

