package conversion

import (
	"encoding/json"
	"errors"
)

func Convert(source interface{}, destination interface{}) error {
	result, err := json.Marshal(source)

	if err != nil {
		return errors.New("Error in matching model")
	}

	err = json.Unmarshal([]byte(result), destination)
	if err != nil {
		return errors.New("Error in matching model")
	}

	return nil
}
