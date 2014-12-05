package reflect

import (
	"encoding/xml"
	"io"
	"os"
)

func LoadModelFile(path string) (*Model, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadModel(file)
}

func ReadModel(r io.Reader) (*Model, error) {
	decoder := xml.NewDecoder(r)
	model := NewModel()
	err := decoder.Decode(model)
	if err != nil {
		return nil, err
	}
	model.Initialize()
	return model, nil
}
