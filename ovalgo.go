package ovalgo

import (
	"encoding/xml"
	"github.com/arzonus/ovalgo/models"
)

func NewUbuntuDefinitions(data []byte) (def models.UbuntuOVALDefinitions, err error) {
	err = xml.Unmarshal(data, &def)
	if err != nil {
		return def, err
	}

	return def, nil
}

func NewDebianDefinitions(data []byte) (def models.DebianOVALDefinitions, err error) {
	err = xml.Unmarshal(data, &def)
	if err != nil {
		return def, err
	}

	return def, nil
}
