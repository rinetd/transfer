# refmt [![GoDoc](https://godoc.org/github.com/rjeczalik/refmt?status.png)](https://godoc.org/github.com/rinetd/transfer) 

[![Build Status](https://travis-ci.org/rinetd/transfer.svg?branch=master)](https://travis-ci.org/rinetd/transfer)

Reformat HCL ⇄ JSON and HCL ⇄ YAML.

### install

```
$ go get github.com/rjeczalik/refmt
```

### usage

```
usage:

	refmt [-t type] INPUT_FILE|"-" OUTPUT_FILE|"-"

Converts from one encoding to another. Supported formats (and their file extensions):

	- HCL (.hcl or .tf)
	- JSON (.json)
	- YAML (.yaml or .yml)

If INPUT_FILE's extension is not recognized or INPUT_FILE is "-" (stdin),
refmt will try to guess input format.

If OUTPUT_FILE is "-" (stdout), destination format type is required to be
passed with -t flag.

	refmt [-t type] merge ORIGINAL_FILE|"-" MIXIN_FILE|"-" OUTPUT_FILE|"-"

Merges the object defined in ORIGINAL_FILE with the object from MIXIN_FILE, writing
the resulting object to the OUTPUT_FILE.

The ORIGINAL_FILE, MIXIN_FILE and OUTPUT_FILE can have different encodings.

If ORIGINAL_FILE's extension is not recognized or ORIGINAL_FILE is "-" (stdin),
refmt will try to guess original format.

If ORIGINAL_FILE does not exist or is empty, refmt is going to use empty
object instead.

If MIXIN_FILE's extension is not recognized or MIXIN_FILE is "-" (stdin),
refmt will try to guess mixin format.

If OUTPUT_FILE is "-" (stdout), destination format type is required to be
passed with -t flag.
```

### docker usage

```
# build the refmt image
docker build -t refmt .

# convert test.yml to hcl
cat test.yml | docker run -i --rm refmt -t hcl - -
```


### examples

```
$ refmt -t yaml main.yaml -
```
```yaml
provider:
  aws:
    access_key: ${var.aws_access_key}
    secret_key: ${var.aws_secret_key}
resource:
  aws_instance:
    aws-instance:
      instance_type: t2.nano
      user_data: echo "hello world!" >> /tmp/helloworld.txt
```
```
$ refmt main.yaml main.json
```
```json
{
        "provider": {
                "aws": {
                        "access_key": "${var.aws_access_key}",
                        "secret_key": "${var.aws_secret_key}"
                }
        },
        "resource": {
                "aws_instance": {
                        "aws-instance": {
                                "instance_type": "t2.nano",
                                "user_data": "echo \"hello world!\" >> /tmp/helloworld.txt"
                        }
                }
        }
}
```
```hcl
$ refmt main.json main.hcl
```
```
provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
}

resource "aws_instance" "aws-instance" {
  instance_type = "t2.nano"
  user_data = "echo \"hello world!\" >> /tmp/helloworld.txt"
}
```

#### pretty reformat in-place

```
$ refmt main.tf.json main.tf.json
```

### merge configurations

```yaml
$ cat ~/.kube/config
apiVersion: v1
clusters:
- cluster:
    certificate-authority: /Volumes/rjk.io/.kube/ca.pem
    server: https://178.0.20.1
  name: default-cluster
contexts:
- context:
    cluster: default-cluster
    user: default-admin
  name: default-system
current-context: default-system
kind: Config
preferences: {}
users:
- name: default-admin
  user:
    client-certificate: /Volumes/rjk.io/.kube/admin.pem
    client-key: /Volumes/rjk.io/.kube/admin-key.pem
```
```yaml
$ cat >>another-cluster <<EOF
apiVersion: v1
clusters:
- cluster:
    certificate-authority: /Volumes/rjk.io/.kube/another-ca.pem
    server: https://10.0.0.1
  name: another-cluster
contexts:
- context:
    cluster: another-cluster
    user: another-admin
  name: another-system
current-context: another-system
kind: Config
preferences: {}
users:
- name: another-admin
  user:
    client-certificate: /Volumes/rjk.io/.kube/another-admin.pem
    client-key: /Volumes/rjk.io/.kube/another-admin-key.pem
EOF
```
```bash
$ refmt merge -t yaml ~/.kube/config ./another-cluster -
```
```yaml
apiVersion: v1
clusters:
- cluster:
    certificate-authority: /Volumes/rjk.io/.kube/ca.pem
    server: https://178.0.20.1
  name: default-cluster
- cluster:
    certificate-authority: /Volumes/rjk.io/.kube/another-ca.pem
    server: https://10.0.0.1
  name: another-cluster
contexts:
- context:
    cluster: default-cluster
    user: default-admin
  name: default-system
- context:
    cluster: another-cluster
    user: another-admin
  name: another-system
current-context: another-system
kind: Config
preferences: {}
users:
- name: default-admin
  user:
    client-certificate: /Volumes/rjk.io/.kube/admin.pem
    client-key: /Volumes/rjk.io/.kube/admin-key.pem
- name: another-admin
  user:
    client-certificate: /Volumes/rjk.io/.kube/another-admin.pem
    client-key: /Volumes/rjk.io/.kube/another-admin-key.pem
```

### todo

- inline docs
- fix hcl marshaling:
  - fix excessive newlines
  - fix excessive quotes
