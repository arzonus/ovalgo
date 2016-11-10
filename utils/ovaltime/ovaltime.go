// Package overload standard time package and added parse and unmarshal xml function in OVAL

package ovaltime

import (
	"encoding/xml"
	"time"
)

// List timeformats, which used for parsing data
var timeFormats = []string{
	time.RFC3339,
}

// If data don't have part of date in end, you can add suffix in end of date
// Example:
// date: 2016-11-01
// Package time can't parse this date, because it don't have time data like hours and etc
// suffix for this date - "T00:00:00Z", if we'll use RFC3339 for parsing
// 2016-11-01 => 2016-11-01T00:00:00Z

var suffixes = map[string]string{
	"Z":               time.RFC3339,
	"T00:00:00Z00:00": time.RFC3339,
	"T00:00:00Z":      time.RFC3339,
}

type Time struct {
	time.Time
}

func (c *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	var err error
	d.DecodeElement(&str, &start)

	for _, layout := range timeFormats {
		t, err := time.Parse(layout, str)
		if err == nil {
			*c = Time{t}
			return nil
		}
	}

	for suffix, layout := range suffixes {
		t, err := time.Parse(layout, str+suffix)
		if err == nil {
			*c = Time{t}
			return nil
		}
	}
	return err
}

func Parse(layout, value string) (Time, error) {
	tm, err := time.Parse(layout, value)
	return Time{tm}, err
}
