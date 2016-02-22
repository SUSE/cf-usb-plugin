package config

import (
	"encoding/json"
)

//Data - config struct
type Data struct {
	MgmtTarget string
}

//NewData returns a Data object
func NewData() (data *Data) {
	data = new(Data)
	return
}

//SetConfigData returns a marshaled Data object
func (d *Data) SetConfigData() (output []byte, err error) {
	output, err = json.Marshal(d)
	if err != nil {
		return
	}

	return output, err
}

//ReadConfigData returns an unmarshaled Data object
func (d *Data) ReadConfigData(input []byte) (err error) {
	err = json.Unmarshal(input, d)
	if err != nil {
		return
	}

	return nil
}
