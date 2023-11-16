package storage

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"

	"github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/common"
)

func encodeToGOB64(template *common.Template) (string, error) {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(template)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

func decodeFromGOB64(str string) (common.Template, error) {
	template := common.Template{}

	templateBuffer, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return template, err
	}

	buffer := bytes.Buffer{}
	buffer.Write(templateBuffer)

	decoder := gob.NewDecoder(&buffer)
	err = decoder.Decode(&template)
	if err != nil {
		return template, err
	}

	return template, nil

}
