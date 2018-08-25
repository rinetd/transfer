package codec

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	plist "github.com/DHowett/go-plist"
	"github.com/clbanning/mxj"
	"github.com/ghodss/yaml"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
	"github.com/hydrogen18/stalecucumber"
	"github.com/rinetd/jy/utils"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

// FileType represents file type FileType used in Smartling API.
type FileType string

const (
	//FileTypeUnknown represents we don't now (yet)
	FileTypeUnknown    = ""
	FileTypeYAML       = "yaml"
	FileTypeTOML       = "toml"
	FileTypeHCL        = "hcl"
	FileTypeJSON       = "json"
	FileTypeXML        = "xml"
	FileTypeMsgpack    = "msgpack"
	FileTypePlist      = "plist"
	FileTypeBson       = "bson"
	FileTypePickle     = "pickle"
	FileTypeProperties = "properties"
	// FIXME: handle other plist FileTypes: binary, openstep, gnustep
	// FIXME: SDL

)

var (
	extensions = map[string]string{
		"":           FileTypeUnknown,
		"yml":        FileTypeYAML,
		"yaml":       FileTypeYAML,
		"toml":       FileTypeTOML,
		"hcl":        FileTypeHCL,
		"tf":         FileTypeHCL,
		"json":       FileTypeJSON,
		"msgpack":    FileTypeMsgpack,
		"plist":      FileTypePlist,
		"bson":       FileTypeBson,
		"xml":        FileTypeXML,
		"pickle":     FileTypePickle,
		"properties": FileTypeProperties,
		"prop":       FileTypeProperties,
	}
)

// DetectFileType get fileType
func CheckType(path string) string {
	// ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(path), "."))
	ext := filepath.Base(path)
	if i := strings.LastIndex(ext, "."); i != -1 {
		ext = ext[i+1:]
	}
	fmt.Println(ext)
	return extensions[ext]
}
func Typ(file string) string {
	ext := filepath.Base(file)
	if i := strings.LastIndex(ext, "."); i != -1 {
		ext = ext[i+1:]
	}
	switch ext = strings.ToLower(ext); ext {
	case "yml", "yaml":
		return FileTypeYAML
	case "tf", "hcl":
		return FileTypeHCL
	case "json":
		return FileTypeJSON
	case "toml":
		return FileTypeTOML
	case "msgpack":
		return FileTypeMsgpack
	case "plist":
		return FileTypePlist
	case "bson":
		return FileTypeBson
	case "xml":
		return FileTypeXML
	case "pickle":
		return FileTypePickle
	case "prop", "props", "properties":
		return FileTypeProperties
	default:
		return FileTypeUnknown
	}
}
func Unmarshal(input []byte, t string) (interface{}, error) {
	var data interface{}
	var err error
	switch t {
	case FileTypeJSON:
		json.Unmarshal(input, &data)
		// decoder := json.NewDecoder(bytes.NewReader(input))
		// decoder.UseNumber()
		// err = decoder.Decode(&data)
		// FIXME: convert numbers to int64
	case FileTypeTOML:
		_, err = toml.Decode(string(input), &data)
		// FIXME: use effective bytes to string instead whole copy
	case FileTypeHCL:
		err = hcl.Unmarshal(input, &data)
		utils.FixHCL(data)
	case FileTypeXML:
		err = xml.Unmarshal(input, &data)
	case FileTypeMsgpack:
	case FileTypePickle:
		buf := new(bytes.Buffer)
		buf.Write(input)
		err = stalecucumber.UnpackInto(&data).From(stalecucumber.Unpickle(buf))
	case FileTypeBson:
		// err = bson.Unmarshal(input, &data)
	case FileTypePlist:
		input := bytes.NewReader(input)
		decoder := plist.NewDecoder(input)
		err = decoder.Decode(data)
	case FileTypeYAML:
		err = yaml.Unmarshal(input, &data)
		if err != nil {
			return nil, err
		}
		// if err == nil {
		// 	data, err = utils.ConvertMapsToStringMaps(data)
		// }
		return utils.FixYAML(data), nil
	default:
		err = fmt.Errorf("unsupported input FileType")
	}

	return data, err
}

func jsonMarshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Marshal Convert struct to  []byte
func Marshal(data interface{}, outputFileType string) ([]byte, error) {
	var result []byte
	var err error

	switch outputFileType {
	case FileTypeJSON:
		result, err = jsonMarshal(&data)
		// result, err = json.Marshal(&data)
		// result, err = json.MarshalIndent(&data, "", "  ")
	case FileTypeXML:
		// fmt.Println("marshal xml")
		// result, err = xml.MarshalIndent(&data, "", "    ")
		jsonValue, err := jsonMarshal(data)
		if err != nil {
			return nil, err
		}
		mv, err := mxj.NewMapJson(jsonValue)
		if err != nil {
			return nil, err
		}
		result, err = mv.XmlIndent("", "  ")
	case FileTypeYAML:
		result, err = yaml.Marshal(&data)
	case FileTypeHCL:
		p, err := jsonMarshal(data)
		if err != nil {
			return nil, err
		}
		nd, err := hcl.Parse(string(p))
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		if err := printer.Fprint(&buf, nd); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
		// err = hcl.Parse(&obj, config)
		// hcl.Marshal(input, &data)
	case FileTypeMsgpack:
		result, err = msgpack.Marshal(&data)
	case FileTypeBson:
		// result, err = bson.Marshal(&data)
	case FileTypePickle:
		buf := new(bytes.Buffer)
		pickler := stalecucumber.NewPickler(buf)
		_, err = pickler.Pickle(result)
		result = buf.Bytes()
	case FileTypePlist:
		// result, err = plist.Marshal(&data, plist.XMLformat)
		output := new(bytes.Buffer)
		encoder := plist.NewEncoder(output)
		err = encoder.Encode(data)
		result = output.Bytes()
	case FileTypeTOML:
		buf := new(bytes.Buffer)
		err = toml.NewEncoder(buf).Encode(data)
		result = buf.Bytes()
	default:
		err = fmt.Errorf("unsupported output FileType")
	}

	return result, err
}
