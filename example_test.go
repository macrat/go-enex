package enex_test

import (
	"encoding/json"
	"fmt"

	"github.com/macrat/go-enex"
)

func Example() {
	enexData := []byte(`
		<?xml version="1.0" encoding="UTF-8"?>
		<!DOCTYPE en-export SYSTEM "http://xml.evernote.com/pub/evernote-export2.dtd">

		<en-export export-date="20190101T130203Z" application="Evernote/Windows" version="6.x">
			<note>
				<title>test note</title>
				<content><![CDATA[<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE en-note SYSTEM "http://xml.evernote.com/pub/enml2.dtd"><en-note>hello world<br /></en-note>]]></content>
				<created>20190102T140304Z</created>
				<tag>shidax</tag>
				<note-attributes>
					<subject-date>20190103T150405Z</subject-date>
					<author>Jhon Due</author>
					<source>web.clip</source>
				</note-attributes>
				<resource>
					<data encoding="base64">aGVsbG8gd29ybGQ=</data>
					<mime>text/plain</mime>
					<resource-attributes>
						<file-name>test attached file</file-name>
						<attachment>true</attachment>
					</resource-attributes>
				</resource>
			</note>
		</en-export>
	`)

	parsed, err := enex.Parse(enexData)
	if err != nil {
		panic(err.Error())
	}

	j, err := json.MarshalIndent(parsed, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(j))

	// Output:
	// {
	//   "Notes": [
	//     {
	//       "Title": "test note",
	//       "Content": {
	//         "XML": "\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\u003c!DOCTYPE en-note SYSTEM \"http://xml.evernote.com/pub/enml2.dtd\"\u003e\u003cen-note\u003ehello world\u003cbr /\u003e\u003c/en-note\u003e"
	//       },
	//       "CreatedAt": "20190102T140304Z",
	//       "UpdatedAt": "00010101T000000Z",
	//       "Tags": [
	//         "shidax"
	//       ],
	//       "Author": "Jhon Due",
	//       "ReceivedAt": "20190103T150405Z",
	//       "Source": "web.clip",
	//       "Resources": [
	//         {
	//           "Data": "aGVsbG8gd29ybGQ=",
	//           "Type": "text/plain",
	//           "Name": "test attached file",
	//           "Attachment": "true",
	//           "Recognition": {}
	//         }
	//       ]
	//     }
	//   ],
	//   "ExportedAt": "20190101T130203Z",
	//   "ExportedBy": "Evernote/Windows",
	//   "Version": "6.x"
	// }
}
