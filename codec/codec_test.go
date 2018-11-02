package codec

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/franela/goblin"
)

func Test_typ(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("文件类型监测", func() {

		g.It(" Should type json", func() {
			g.Assert(Typ("yaml")).Equal(FileTypeYAML)
			g.Assert(Typ("json")).Equal(FileTypeJSON)
			g.Assert(Typ("a.json")).Equal(FileTypeJSON)
			g.Assert(Typ("./a.json")).Equal(FileTypeJSON)
			g.Assert(Typ("/data/a.json")).Equal(FileTypeJSON)
		})
		g.It(" Should type json", func() {
			g.Assert(CheckType("yaml")).Equal(FileTypeYAML)
			g.Assert(CheckType("json")).Equal(FileTypeJSON)
			g.Assert(CheckType("a.json")).Equal(FileTypeJSON)
			g.Assert(CheckType("./a.json")).Equal(FileTypeJSON)
			g.Assert(CheckType("/data/a.json")).Equal(FileTypeJSON)
		})
	})
}

var input = []byte(`{
	"Author": {
		"email": "rinetd@163.com",
		"github": true,
		"age": 10,
		"rss": "rss.xml"
	}}`)

func TestUnmarshal(t *testing.T) {
	var data interface{}
	json.Unmarshal(input, &data)
	fmt.Println(data)
	fmt.Println(xml.Marshal(data))

}

func Test_marshal(t *testing.T) {
	var data interface{}
	json.Unmarshal(input, &data)
	fmt.Println(data)
	d, err := Marshal(data, FileTypeHCL)
	if err != nil {

	}
	fmt.Println(string(d))
}
