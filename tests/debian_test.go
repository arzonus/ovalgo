package test

import (
	"encoding/xml"
	"github.com/arzonus/ovalgo/models"
	"github.com/arzonus/ovalgo/utils/ovaltime"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func TestParseDebian(t *testing.T) {
	var err error

	debXML := `
  <?xml version="1.0" ?>
  <oval_definitions xmlns="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:ind-def ="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent" xmlns:linux-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#linux" xmlns:oval="http://oval.mitre.org/XMLSchema/oval-common-5" xmlns:oval-def="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:unix-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#unix" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent independent-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#linux linux-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#unix unix-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5 oval-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-common-5 oval-common-schema.xsd">
    <generator>
      <oval:product_name>Debian</oval:product_name>
      <oval:schema_version>5.3</oval:schema_version>
      <oval:timestamp>2016-11-01T11:39:07.188-04:00</oval:timestamp>
    </generator>
    <definitions>
      <definition class="vulnerability" id="oval:org.debian:def:20160402" version="1">
        <metadata>
          <title>CVE-2016-0402</title>
          <affected family="unix">
            <platform>Debian GNU/Linux 9.0</platform>
            <product>openjdk-8</product>
          </affected>
          <reference ref_id="CVE-2016-0402" ref_url="http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2016-0402" source="CVE"/>
          <description>Unspecified vulnerability in the Java SE and Java SE Embedded components in Oracle Java SE 6u105, 7u91, and 8u66 and Java SE Embedded 8u65 allows remote attackers to affect integrity via unknown vectors related to Networking.</description>
          <debian>
            <date>2016-11-01</date>
            <moreinfo></moreinfo>
          </debian>
        </metadata>
        <criteria comment="Release section" operator="AND">
          <criterion comment="Debian 9.0 is installed" test_ref="oval:org.debian.oval:tst:1"/>
          <criteria comment="Architecture section" operator="OR">
            <criteria comment="Architecture independent section" operator="AND">
              <criterion comment="all architecture" test_ref="oval:org.debian.oval:tst:2"/>
              <criterion comment="openjdk-8 DPKG is earlier than 8u72-b15-1" test_ref="oval:org.debian.oval:tst:3"/>
            </criteria>
          </criteria>
        </criteria>
      </definition>
    </definitions>
  </oval_definitions>`

	deb := new(models.DebianOVALDefinitions)

	err = xml.Unmarshal([]byte(debXML), deb)

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestLoadParseDebian(t *testing.T) {
	var err error

	path, _ := filepath.Abs("./data/debian.xml")

	debXML, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	deb := new(models.DebianOVALDefinitions)

	err = xml.Unmarshal([]byte(debXML), deb)

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestParseDebianGenerator(t *testing.T) {
	var err error

	debXML := `
  <?xml version="1.0" ?>
  <oval_definitions xmlns="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:ind-def ="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent" xmlns:linux-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#linux" xmlns:oval="http://oval.mitre.org/XMLSchema/oval-common-5" xmlns:oval-def="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:unix-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#unix" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent independent-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#linux linux-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#unix unix-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5 oval-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-common-5 oval-common-schema.xsd">
    <generator>
      <oval:product_name>Debian</oval:product_name>
      <oval:schema_version>5.3</oval:schema_version>
      <oval:timestamp>2016-11-01T11:39:07.188-04:00</oval:timestamp>
    </generator>
  </oval_definitions>`

	debTime, _ := ovaltime.Parse(time.RFC3339Nano, "2016-11-01T11:39:07.188-04:00")

	debStr := models.DebianOVALDefinitions{
		OVALDefinitions: models.OVALDefinitions{
			Generator: models.Generator{
				ProductName:   "Debian",
				SchemaVersion: "5.3",
				Timestamp:     debTime,
			},
		},
	}

	deb := new(models.DebianOVALDefinitions)

	err = xml.Unmarshal([]byte(debXML), deb)

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if !reflect.DeepEqual(deb.Generator, debStr.Generator) {
		t.Log("XML Generator :", deb.Generator)
		t.Log("Structure Generator: ", debStr.Generator)
		t.Fail()
		return
	}
}

func TestParseDebianDefinitionVulnerabilityMetadata(t *testing.T) {
	var err error

	debXML := `
  <?xml version="1.0" ?>
  <oval_definitions xmlns="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:ind-def ="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent" xmlns:linux-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#linux" xmlns:oval="http://oval.mitre.org/XMLSchema/oval-common-5" xmlns:oval-def="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:unix-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#unix" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent independent-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#linux linux-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#unix unix-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5 oval-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-common-5 oval-common-schema.xsd">
    <definitions>
    <definition class="vulnerability" id="oval:org.debian:def:20160402" version="1">
      <metadata>
        <title>CVE-2016-0402</title>
        <affected family="unix">
          <platform>Debian GNU/Linux 9.0</platform>
          <product>openjdk-8</product>
        </affected>
        <reference ref_id="CVE-2016-0402" ref_url="http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2016-0402" source="CVE"/>
        <description>Unspecified vulnerability in the Java SE and Java SE Embedded components in Oracle Java SE 6u105, 7u91, and 8u66 and Java SE Embedded 8u65 allows remote attackers to affect integrity via unknown vectors related to Networking.</description>
        <debian>
          <date>2016-11-01</date>
          <moreinfo></moreinfo>
        </debian>
      </metadata>
    </definition>
  </definitions>
  </oval_definitions>`

	debTime, _ := ovaltime.Parse(time.RFC3339, "2016-11-01T00:00:00Z")

	debStr := models.DebianOVALDefinitions{
		Definitions: []models.DebianDefinition{
			{
				Metadata: models.DebianMetadata{
					Metadata: models.Metadata{Title: "CVE-2016-0402",
						Affected: []models.Affected{
							{
								Family: "unix",
								Platforms: []string{
									"Debian GNU/Linux 9.0",
								},
								Products: []string{
									"openjdk-8",
								},
							},
						},
						References: []models.Reference{
							{
								RefID:  "CVE-2016-0402",
								RefURL: "http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2016-0402",
								Source: "CVE",
							},
						},
						Description: "Unspecified vulnerability in the Java SE and Java SE Embedded components in Oracle Java SE 6u105, 7u91, and 8u66 and Java SE Embedded 8u65 allows remote attackers to affect integrity via unknown vectors related to Networking.",
					},
					Debian: models.DebianMetadataDebian{
						Date: debTime,
					},
				},
			},
		},
	}

	deb := new(models.DebianOVALDefinitions)

	err = xml.Unmarshal([]byte(debXML), deb)

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if !reflect.DeepEqual(deb.Definitions[0].Metadata, debStr.Definitions[0].Metadata) {
		t.Log("XML Metadata :", deb.Definitions[0].Metadata)
		t.Log("Structure Metadata: ", debStr.Definitions[0].Metadata)
		t.Fail()
		return
	}
}
