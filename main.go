// The parser for [the xml that exported by Evernote (.enex file)](https://help.evernote.com/hc/en-us/articles/209005557).
package enex

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net/url"
	"time"
)

type EncodedData struct {
	data []byte
}

func (ed *EncodedData) String() string {
	return fmt.Sprintf("{EncodedData length=%d}", len(ed.data))
}

func (ed *EncodedData) UnmarshalText(text []byte) error {
	if decoded, err := base64.StdEncoding.DecodeString(string(text)); err != nil {
		return err
	} else {
		ed.data = decoded
		return nil
	}
}

func (ed EncodedData) MarshalText() ([]byte, error) {
	s := base64.StdEncoding.EncodeToString(ed.data)
	return []byte(s), nil
}

type Recognition struct {
	XML string `xml:",cdata" json:",omitempty"`
}

type Resource struct {
	Data        EncodedData `xml:"data"`
	Type        string      `xml:"mime"`
	Name        string      `xml:"resource-attributes>file-name"`
	Attachment  string      `xml:"resource-attributes>attachment,omitempty" json:",omitempty"`
	Width       int         `xml:"width,omitempty" json:",omitempty"`
	Height      int         `xml:"height,omitempty" json:",omitempty"`
	Recognition Recognition `xml:"recognition,omitempty"`
}

type Content struct {
	XML string `xml:",cdata"`
}

type DateTime time.Time

func (dt *DateTime) UnmarshalText(text []byte) (err error) {
	var t time.Time

	for _, f := range []string{"2006-01-02T15:04:05Z", "20060102T150405Z"} {
		t, err = time.Parse(f, string(text))
		if err == nil {
			break
		}
	}

	*dt = DateTime(t)

	return
}

func (dt DateTime) MarshalText() ([]byte, error) {
	s := time.Time(dt).Format("20060102T150405Z")
	return []byte(s), nil
}

func (dt DateTime) String() string {
	return time.Time(dt).Format("2006-01-02T15:04:05Z")
}

type Note struct {
	Title      string     `xml:"title"`
	Content    Content    `xml:"content"`
	CreatedAt  DateTime   `xml:"created"`
	UpdatedAt  DateTime   `xml:"updated,omitempty" json:",omitempty"`
	Tags       []string   `xml:"tag"`
	Author     string     `xml:"note-attributes>author"`
	ReceivedAt DateTime   `xml:"note-attributes>subject-date,omitempty" json:",omitempty"`
	Source     string     `xml:"note-attributes>source,omitempty" json:",omitempty"`
	SourceURL  *url.URL   `xml:"note-attributes>source-url,omitempty" json:",omitempty"`
	Resources  []Resource `xml:"resource,omitempty" json:",omitempty"`
}

type EvernoteExportedXML struct {
	Notes      []Note   `xml:"note"`
	ExportedAt DateTime `xml:"export-date,attr"`
	ExportedBy string   `xml:"application,attr"`
	Version    string   `xml:"version,attr"`
}

/*
Parse .enex file from bytes into EvernoteExportedXML.

`data` is an xml data.

NOTE: This function is just wrapper of xml.Unmarshal. Please directly use xml package if you want customize behavior.
*/
func Parse(data []byte) (EvernoteExportedXML, error) {
	var parsed EvernoteExportedXML
	return parsed, xml.Unmarshal(data, &parsed)
}

/*
Parse .enex file from io.Reader into EvernoteExportedXML.

`r` is a reader to read .enex file.


NOTE: This function is just wrapper of xml.Decoder.Decode. Please directly use xml package if you want customize behavior.
*/
func ParseFromReader(r io.Reader) (EvernoteExportedXML, error) {
	decoder := xml.NewDecoder(r)

	var parsed EvernoteExportedXML
	return parsed, decoder.Decode(&parsed)
}
