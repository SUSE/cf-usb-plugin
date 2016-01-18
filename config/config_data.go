package config

import (
	"encoding/json"
)

type Data struct {
	MgmtTarget string
}

func NewData() (data *Data) {
	data = new(Data)
	return
}

func (d *Data) SetConfigData() (output []byte, err error) {
	output, err = json.Marshal(d)
	if err != nil {
		return
	}

	return output, err
}

func (d *Data) ReadConfigData(input []byte) (err error) {
	err = json.Unmarshal(input, d)
	if err != nil {
		return
	}

	return nil
}
