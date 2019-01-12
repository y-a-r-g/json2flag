package json2flag

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
)

var ErrUnsupportedValueType = fmt.Errorf("unsupported config value type")

func ReadConfigFile(fileName string) error {
	if fileName == "" {
		fileName = "config.json"
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return ReadConfigData(data)
}

func ReadConfigString(json string) error {
	return ReadConfigData([]byte(json))
}

func ReadConfigData(data []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	config := map[string]interface{}{}

	if err := decoder.Decode(&config); err != nil {
		return err
	}

	for k, v := range config {
		switch v.(type) {
		case string:
			if err := flag.Set(k, v.(string)); err != nil {
				return err
			}
			break
		case bool:
			if err := flag.Set(k, strconv.FormatBool(v.(bool))); err != nil {
				return err
			}
			break
		case float64:
			if err := flag.Set(k, strconv.FormatFloat(v.(float64), 'G', -1, 64)); err != nil {
				return err
			}
			break
		default:
			return ErrUnsupportedValueType
		}
	}

	return nil
}
