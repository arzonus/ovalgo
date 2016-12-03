package ovalgo

import (
	"github.com/arzonus/ovalgo/models"
	"github.com/arzonus/ovalgo/utils/ovaltime"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

// Debian tests
func TestNewDebianDefinitionsFullXML(t *testing.T) {
	path, _ := filepath.Abs("./data/debian.xml")

	debXML, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	_, err = NewDebianDefinitions([]byte(debXML))

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestNewDebianDefinitionsCheckGenerator(t *testing.T) {
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

	_, err = NewDebianDefinitions([]byte(debXML))

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestNewDebianDefinitionsCheckWrongGenerator(t *testing.T) {
	var err error

	debXML := `
  <?xml version="1.0" ?>
  <oval_definitions xmlns="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:ind-def ="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent" xmlns:linux-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#linux" xmlns:oval="http://oval.mitre.org/XMLSchema/oval-common-5" xmlns:oval-def="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:unix-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#unix" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent independent-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#linux linux-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#unix unix-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5 oval-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-common-5 oval-common-schema.xsd">
    <generator>
      <oval:product_name>Debian</oval:product_nam>
      <oval:schema_version>5.3</oval:schema_version>
      <oval:timestamp>2016:00</oval:timestamp>
    </generator>
  </oval_definitions>`

	_, err = NewDebianDefinitions([]byte(debXML))

	if err == nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestNewDebianDefinitionsCheckVulnerabilityMetadata(t *testing.T) {
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

	deb, err := NewDebianDefinitions([]byte(debXML))

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

// Ubuntu tests
func TestNewUbuntuDefinitionsFullXML(t *testing.T) {
	path, _ := filepath.Abs("./data/ubuntu.xenial.xml")

	debXML, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	_, err = NewUbuntuDefinitions([]byte(debXML))

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestNewUbuntuDefinitionsCheckGenerator(t *testing.T) {
	var err error

	debXML := `
  <?xml version="1.0" ?>
  <oval_definitions xmlns="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:ind-def ="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent" xmlns:linux-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#linux" xmlns:oval="http://oval.mitre.org/XMLSchema/oval-common-5" xmlns:oval-def="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:unix-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#unix" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent independent-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#linux linux-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#unix unix-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5 oval-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-common-5 oval-common-schema.xsd">
    <generator>
        <oval:product_name>Canonical CVE OVAL Generator</oval:product_name>
        <oval:product_version>1</oval:product_version>
        <oval:schema_version>5.11.1</oval:schema_version>
        <oval:timestamp>2016-11-01T15:23:45</oval:timestamp>
    </generator>
  </oval_definitions>`

	_, err = NewUbuntuDefinitions([]byte(debXML))

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestNewUbuntuDefinitionsCheckWrongGenerator(t *testing.T) {
	var err error

	debXML := `
  <?xml version="1.0" ?>
  <oval_definitions xmlns="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:ind-def ="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent" xmlns:linux-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#linux" xmlns:oval="http://oval.mitre.org/XMLSchema/oval-common-5" xmlns:oval-def="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:unix-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#unix" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent independent-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#linux linux-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#unix unix-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5 oval-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-common-5 oval-common-schema.xsd">
    <generator>
        <oval:product_name>Canonical CVE OVAL Generator</oval:product_name>
        <oval:product_version>1</oval:product_versio>
        <oval:schema_version>5.11.1</oval:schema_version>
        <oval:timestamp>2016-11-01T15:23:45</oval:timestamp>
    </generator>
  </oval_definitions>`

	_, err = NewUbuntuDefinitions([]byte(debXML))

	if err == nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestNewUbuntuDefinitionsCheckVulnerabilityMetadata(t *testing.T) {
	var err error

	debXML := `
  <?xml version="1.0" ?>
  <oval_definitions xmlns="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:ind-def ="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent" xmlns:linux-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#linux" xmlns:oval="http://oval.mitre.org/XMLSchema/oval-common-5" xmlns:oval-def="http://oval.mitre.org/XMLSchema/oval-definitions-5" xmlns:unix-def="http://oval.mitre.org/XMLSchema/oval-definitions-5#unix" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://oval.mitre.org/XMLSchema/oval-definitions-5#independent independent-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#linux linux-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5#unix unix-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-definitions-5 oval-definitions-schema.xsd http://oval.mitre.org/XMLSchema/oval-common-5 oval-common-schema.xsd">
    <definitions>
    <definition class="vulnerability" id="oval:com.ubuntu.xenial:def:20022439000" version="1">
            <metadata>
                <title>CVE-2002-2439 on Ubuntu 16.04 LTS (xenial) - low.</title>
                <description>operator new[] sometimes returns pointers to heap blocks which are too small.  When a new array is allocated, the C++ run-time has to calculate its size.  The product may exceed the maximum value which can be stored in a machine register.  This error is ignored, and the truncated value is used for the heap allocation. This may lead to heap overflows and therefore security bugs. (See http://cert.uni-stuttgart.de/advisories/calloc.php for further references.)</description>
                <affected family="unix">
                    <platform>Ubuntu 16.04 LTS</platform>
                </affected>
                <reference source="CVE" ref_id="CVE-2002-2439" ref_url="https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2002-2439" />
                <advisory>
                    <severity>Low</severity>
                    <rights>Copyright (C) 2015 Canonical Ltd.</rights>
                    <public_date>2015-02-23</public_date>
                    <bug>http://gcc.gnu.org/bugzilla/show_bug.cgi?id=19351</bug>
                    <bug>https://bugzilla.redhat.com/show_bug.cgi?id=850911</bug>
                    <ref>http://people.canonical.com/~ubuntu-security/cve/2002/CVE-2002-2439.html</ref>
                </advisory>
            </metadata>
            <oval:notes>
                <oval:note>sbeattie&gt; fixed upstream in gcc 4.8.0 sbeattie&gt; backporting fixes may be problematic for ABI issues</oval:note>
            </oval:notes>
            <criteria>
                <extend_definition definition_ref="oval:com.ubuntu.xenial:def:100" comment="Ubuntu 16.04 LTS (xenial) is installed." applicability_check="true" />
                <criteria operator="OR">
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:10" comment="The vulnerability of the 'gcc-4.7' package in xenial is not known (status: 'needs-triage'). It is pending evaluation." />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:10" comment="The vulnerability of the 'gcc-4.7-armel-cross' package in xenial is not known (status: 'needs-triage'). It is pending evaluation." />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:10" comment="The vulnerability of the 'gcc-4.7-armhf-cross' package in xenial is not known (status: 'needs-triage'). It is pending evaluation." />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:20022439030" comment="While related to the CVE in some way, the 'gcc-4.8' package in xenial is not affected (note: '4.8.2-19ubuntu1')." negate = "true" />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:20022439040" comment="While related to the CVE in some way, the 'gcc-4.8-arm64-cross' package in xenial is not affected." negate = "true" />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:20022439050" comment="While related to the CVE in some way, the 'gcc-4.8-armhf-cross' package in xenial is not affected." negate = "true" />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:20022439060" comment="While related to the CVE in some way, the 'gcc-4.8-powerpc-cross' package in xenial is not affected." negate = "true" />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:20022439070" comment="While related to the CVE in some way, the 'gcc-4.8-ppc64el-cross' package in xenial is not affected." negate = "true" />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:20022439080" comment="While related to the CVE in some way, the 'gcc-4.9' package in xenial is not affected (note: '4.9.1-8ubuntu1')." negate = "true" />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:10" comment="The vulnerability of the 'gcc-arm-linux-androideabi' package in xenial is not known (status: 'needs-triage'). It is pending evaluation." />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:10" comment="The vulnerability of the 'gcc-arm-none-eabi' package in xenial is not known (status: 'needs-triage'). It is pending evaluation." />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:20022439110" comment="While related to the CVE in some way, the 'gcc-avr' package in xenial is not affected." negate = "true" />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:10" comment="The vulnerability of the 'gcc-h8300-hms' package in xenial is not known (status: 'needs-triage'). It is pending evaluation." />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:10" comment="The vulnerability of the 'gcc-i686-linux-android' package in xenial is not known (status: 'needs-triage'). It is pending evaluation." />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:10" comment="The vulnerability of the 'gcc-mingw-w64' package in xenial is not known (status: 'needs-triage'). It is pending evaluation." />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:10" comment="The vulnerability of the 'gcc-msp430' package in xenial is not known (status: 'needs-triage'). It is pending evaluation." />
                    <criterion test_ref="oval:com.ubuntu.xenial:tst:20022439160" comment="While related to the CVE in some way, the 'gcc-snapshot' package in xenial is not affected (note: '20140405-0ubuntu1')." negate = "true" />
                </criteria>
            </criteria>
        </definition>
  </definitions>
  </oval_definitions>`

	_, err = NewDebianDefinitions([]byte(debXML))

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}

// New tests
func TestNewFullUbuntuXML(t *testing.T) {
	path, _ := filepath.Abs("./data/ubuntu.xenial.xml")

	debXML, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	def, err := New([]byte(debXML), "ubuntu")

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	switch tp := def.(type) {
	case models.UbuntuOVALDefinitions:
		t.Log("It's ubuntu definition: ", tp.Generator.ProductName)
	case models.DebianOVALDefinitions:
		t.Log("It's debian definition: ", tp.Generator.ProductName)
	}
}

func TestNewFullDebianXML(t *testing.T) {
	path, _ := filepath.Abs("./data/debian.xml")

	debXML, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	def, err := New([]byte(debXML), "debian")

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	switch tp := def.(type) {
	case models.UbuntuOVALDefinitions:
		t.Log("It's ubuntu definition: ", tp.Generator.ProductName)
	case models.DebianOVALDefinitions:
		t.Log("It's debian definition: ", tp.TextFileContentTests)
	}
}

// Check Tests
