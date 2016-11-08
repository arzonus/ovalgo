package rfctime

import (
	"encoding/xml"
	"time"
)

type Time struct {
	time.Time
}

const layout = time.RFC3339

func (c *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var strtime string
	d.DecodeElement(&strtime, &start)

	strtime = strtime + "Z"

	parse, err := time.Parse(layout, strtime)
	if err != nil {
		return err
	}
	*c = Time{parse}
	return nil
}
