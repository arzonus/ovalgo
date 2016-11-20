package ovalgo

import (
	"encoding/xml"
	"github.com/arzonus/ovalgo/models"
)

type OVALDefinitions interface{}

func New(data []byte, typeDefinitions string) (OVALDefinitions, error) {
	switch typeDefinitions {
	case "debian":
		return NewDebianDefinitions(data)
	case "ubuntu":
		return NewUbuntuDefinitions(data)
	}

	return NewDebianDefinitions(data)
}

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
