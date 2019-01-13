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

	return decodeConfig(config, "")
}

func decodeConfig(config map[string]interface{}, prefix string) error {
	for k, v := range config {
		name := k
		if prefix != "" {
			name = fmt.Sprintf("%s%s%s", prefix, structDelimiter, k)
		}

		switch v.(type) {
		case string:
			if err := flag.Set(name, v.(string)); err != nil {
				return err
			}
			break
		case bool:
			if err := flag.Set(name, strconv.FormatBool(v.(bool))); err != nil {
				return err
			}
			break
		case float64:
			if err := flag.Set(name, strconv.FormatFloat(v.(float64), 'G', -1, 64)); err != nil {
				return err
			}
			break
		default:
			if sub, ok := v.(map[string]interface{}); ok {
				if err := decodeConfig(sub, name); err != nil {
					return err
				}
			} else {
				return ErrUnsupportedValueType
			}
		}
	}

	return nil
}
