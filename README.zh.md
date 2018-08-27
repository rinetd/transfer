# transfer [![GoDoc](https://godoc.org/github.com/rinetd/transfer?status.png)](https://godoc.org/github.com/rinetd/transfer)[![Build Status](https://travis-ci.org/rinetd/transfer.svg?branch=master)](https://travis-ci.org/rinetd/transfer)

* 支持全平台 Windows linux mac
* 自动检测转换文件格式
* 支持多种格式相互转换 HCL ⇄ JSON ⇄ YAML⇄TOML⇄XML⇄plist⇄pickle⇄properties ... 

### install

```
$ go get github.com/rinetd/transfer
```
### Download

```
https://github.com/rinetd/transfer/releases
```

### usage

```
usage:

	transfer [-f] [-s input.yaml] [-t output.json] /path/to/input.yaml [/path/to/output.json]

Converts from one encoding to another. Supported formats (and their file extensions):

	- JSON (.json)
	- TOML (.toml)
	- YAML (.yaml or .yml)
	- HCL (.hcl or .tf)
	- XML (.xml)
	- MSGPACK (.msgpack)
	- PLIST (.plist)
	- BSON (.bson)
	- PICKLE (.pickle)
	- PROPERTIES (.prop or .props or .properties)

```

### docker usage

```
# build the transfer image
docker build -t rientd/transfer .
```


### examples

#### yaml格式转换为json
将 ./data/main.yml 转换到 ./data/main.json
以下几种命名格式是等价的
```
$ transfer -f data/main.yaml        (default output `json` format)
$ transfer -f data/main.yaml data/main.json
$ transfer -f -t json data/main.yaml
$ transfer -f -s data/main.yaml     (default output `json` format)
$ transfer -f -s data/main.yaml -t json 
$ transfer -f -s data/main.yaml -t data/main.json 
```
```yaml
Author:
  email: rinetd@163.com
  github: rinetd
menu:
  main:
  - Identifier: categories
    Name: categories
    Pre: <i class='fa fa-category'></i>
    URL: /categories/
    Weight: -102
  - Identifier: tags
    Name: tags
    Pre: <i class='fa fa-tags'></i>
    URL: /tags/
    Weight: -101
theme: hueman

```

```json
{
	"Author": {
		"email": "rinetd@163.com",
		"github": "rinetd"
	},
	"menu": {
		"main": [
			{
				"Identifier": "categories",
				"Name": "categories",
				"Pre": "<i class='fa fa-category'></i>",
				"URL": "/categories/",
				"Weight": -102
			},
			{
				"Identifier": "tags",
				"Name": "tags",
				"Pre": "<i class='fa fa-tags'></i>",
				"URL": "/tags/",
				"Weight": -101
			}
		]
	},
	"theme": "hueman"
}
```
#### 其他格式转换: 
json -> hcl
```hcl
$ transfer main.json main.hcl
```
yaml -> toml
```toml
$ transfer main.yaml main.toml
```