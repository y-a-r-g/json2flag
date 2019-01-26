package json2flag

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//TODO: passed flags should overwrite json values

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

	err := decoder.Decode(&config)
	if err != nil {
		return err
	}

	return decodeConfig(config, "")
}

func decodeConfig(config map[string]interface{}, prefix string) error {
	for k, v := range config {
		var err error
		name := k
		if prefix != "" {
			name = fmt.Sprintf("%s%s%s", prefix, structDelimiter, k)
		}

		switch v.(type) {
		case string:
			err = flag.Set(name, v.(string))
			break
		case bool:
			err = flag.Set(name, strconv.FormatBool(v.(bool)))
			break
		case float64:
			err = flag.Set(name, strconv.FormatFloat(v.(float64), 'G', -1, 64))
			break
		default:
			sub, ok := v.(map[string]interface{})
			if !ok {
				return ErrUnsupportedValueType
			}

			err = decodeConfig(sub, name)
		}
		if err != nil && !strings.HasPrefix(err.Error(), "no such flag") {
			return err
		}
	}

	return nil
}

func WriteConfigFile(name string, perm os.FileMode) error {
	config := map[string]interface{}{}
	flag.VisitAll(func(f *flag.Flag) {
		path := strings.Split(f.Name, ".")
		node := config

		for len(path) > 1 {
			child, ok := node[path[0]]
			if !ok {
				child = map[string]interface{}{}
				node[path[0]] = child
			}
			node = child.(map[string]interface{})

			path = path[1:]
		}

		name, _ := flag.UnquoteUsage(f)
		if name == "duration" {
			node[path[0]] = f.Value.String()
		} else {
			node[path[0]] = f.Value
		}
	})

	buffer := bytes.Buffer{}
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "    ")
	err := encoder.Encode(&config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(name, buffer.Bytes(), perm)
}
