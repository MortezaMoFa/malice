// Code generated by go-bindata.
// sources:
// plugins/plugins.toml
// plugins/templates/python/.dockerignore
// plugins/templates/python/.gitignore.template
// plugins/templates/python/CHANGELOG.md
// plugins/templates/python/Dockerfile
// plugins/templates/python/LICENSE.md
// plugins/templates/python/README-short.txt
// plugins/templates/python/README.md
// plugins/templates/python/circle.yml
// plugins/templates/python/plugin.toml
// plugins/templates/python/requirements.txt
// plugins/templates/python/scan.py
// DO NOT EDIT!

package plugins

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _pluginsPluginsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\xdf\x53\xdb\xb8\x13\x7f\xcf\x5f\xb1\x63\x5e\xbe\xdf\x1b\x42\x42\x9a\x90\xb6\x33\x3c\xe4\x20\xb4\xf4\xa0\x30\x49\xca\x4d\xaf\xc3\x50\x45\x5e\xc7\x3a\x6c\xcb\x23\xc9\x81\xf0\x70\x7f\xfb\x8d\x64\x27\xce\x0f\xd1\x5a\xa9\x8f\x07\x9c\x68\xf5\xd9\xdd\xcf\x47\x2b\x6d\xe4\x03\x38\xe3\xe9\x42\xb0\x59\xa8\xe0\x7f\xf4\xff\xd0\x69\x1f\xbf\x81\xa6\x7e\xbc\x85\x69\x44\xe8\xa3\xe2\x29\x7c\xe2\x32\xcc\x08\x5c\x13\x96\xe0\x21\x0c\xa2\x08\x46\x1a\x20\x61\x84\x12\xc5\x1c\xfd\xa3\xc6\x01\x8c\x11\xe1\xea\xf2\x6c\xf8\x79\x3c\x84\x80\x0b\x88\x18\xc5\x44\x22\xb0\x24\xe0\x22\x26\x8a\xf1\xe4\xa8\xd1\x38\xa8\xe7\xaf\x71\x00\xb7\x57\x5f\x3e\x5c\x7e\x86\x33\x9e\x04\x6c\x96\x09\x13\x00\xdc\xfd\xd4\x94\x4f\x43\x31\x15\x21\x9c\x82\x77\x4d\x34\x73\xb8\x8d\xb2\x19\x4b\x36\xd3\x93\x5e\xa3\xf1\xed\x5b\x6a\x2c\xf7\xf7\x0d\x00\x4c\xc8\x34\x42\x1f\x4e\x41\x89\x0c\x1b\x00\x09\x89\x8d\x93\x44\x8a\xc8\x6b\x00\xf8\x28\xa9\x60\xa9\xe1\x76\x0a\xde\xe7\xf1\xe8\x0a\xce\x89\x22\x53\x22\x11\x3e\x12\x19\xc2\x18\x89\xa0\xa1\x9e\x4b\x89\xc2\x19\x17\x0b\x3d\x91\x25\x0a\x8d\x03\x16\x93\x99\xf1\x18\x9b\xb4\x5a\xda\xf1\x7b\x19\x92\x63\x6d\x14\x98\x72\xc9\x54\x81\x09\x95\x4a\xe5\xfb\x56\x6b\xc6\x54\x98\x4d\x8f\x28\x8f\x5b\x39\xa8\x99\x67\x2c\x0d\xf8\x68\xc6\x94\xc6\x4e\x33\x16\xe9\xc4\x03\x12\x49\x9d\x39\x8d\xf5\x37\x2f\xe2\xfc\x31\x4b\xf5\x84\x98\xe5\x54\x42\x22\x4d\x7a\xfa\xa9\x16\x29\x4a\x38\x85\x6f\xe0\x99\x1c\xe0\xbe\xaa\x20\x73\x26\x32\xa9\xb8\x22\x36\x59\xee\xb4\x71\xa2\x8d\xd0\x84\x80\x45\x28\x41\x52\x92\x00\x49\x7c\x13\x16\xca\xac\x2a\x89\xb4\x19\xcc\x55\xa5\x12\xfd\x9a\x56\x59\x1a\x71\xb2\x1a\x00\xbd\x77\xd4\x72\x50\x71\xf8\xae\xb9\x7f\x07\x16\xc0\x82\x67\xf0\x44\x12\xa5\x47\x0b\xbb\x24\x71\x1a\xa1\x1e\x58\x8b\x43\x79\xdc\x00\x20\x29\x7b\x44\x93\x64\xa7\xf7\xe6\x5d\xef\xf8\xc4\xef\xf6\x8f\xfd\xfe\x14\xa7\x27\xd3\xce\x5b\xd2\xef\xb4\xfd\x3e\xe9\xf4\xda\x9d\x2e\xfa\xb4\x1d\xf4\x7b\xef\xda\xfe\x9b\x6e\x2f\xa0\xfd\x6e\xbf\x7b\xfc\xf6\xa4\xdb\x6b\xb7\x3b\x84\x76\xfb\x53\x0f\x0e\x60\x12\x32\x09\x4c\x02\x01\x85\x52\xc1\x23\x2e\x0e\x21\x8d\x50\x17\x9e\xc0\x34\x22\x14\xe1\x89\xa9\x50\xe7\x28\x80\x3f\x25\x7b\xd4\x40\xec\xf7\xbc\xc3\xa2\x14\xf2\x67\xa7\x77\xa2\x8b\x42\x17\xc2\x5c\xcf\xf1\xae\x07\xfa\x4c\x79\xb8\x9b\x3c\x0c\x6e\x2f\xf5\xa4\x62\x60\x72\x79\x3d\xbc\xf9\x32\xf1\x5e\x2f\xa0\xa5\xd8\xcb\x0a\x32\x4a\x2d\x13\xd9\x2a\xa0\x83\x95\x11\x9a\x7b\x55\xcc\x86\x73\xd7\x82\x59\x81\xff\xdb\x7a\x29\xc3\xd0\x45\x2c\xb2\xa2\x68\x32\x89\x42\xa7\xa9\xe3\x16\xd5\xe3\xfd\xc2\x76\xde\x59\xb9\xc9\xc7\x87\x2f\xe3\xe1\x68\x7d\xe9\x3e\x3e\xfc\x31\xfc\xfa\x83\x95\xdb\xda\xfa\x32\x24\x3e\x7f\x6a\x9a\x4e\x23\x2c\x8b\x37\x36\xf6\xb1\x31\xef\xb9\x7c\x3b\x21\x5c\x97\x70\xc3\x41\x5d\x47\xe4\xfa\xf6\xf8\xc1\x49\xb9\x53\xe8\x48\xe2\xa6\x59\x62\x8b\x58\x13\x24\xf1\x99\xb6\xed\x5b\xe8\x1b\xce\x9d\x2b\x7d\x85\x7e\x4d\xa3\x7a\x24\xd9\xaa\x20\xdd\x12\xf4\xcf\x10\x8b\x1e\x52\xfa\x88\x69\x6b\x32\x62\xe7\x2d\x7c\x66\x81\xe2\x3c\xda\x16\x23\x46\x45\x7c\xa2\x88\x4d\x8f\x75\xd7\xae\x6a\x2c\xb1\x3f\xd3\xe2\xb7\xca\x3f\x1b\x16\x44\x10\x0b\xc9\xaf\x83\xd1\x00\xc6\x94\x24\xdb\xcc\xc8\xdc\xc6\x69\xe9\xc5\x95\x8f\xc6\xd5\xc7\x85\xcc\x89\x54\x16\x32\x03\x3d\x0e\x83\x44\x31\xd3\xf6\xab\x51\x5a\x39\x73\xe5\x64\x80\x75\x92\x9a\xd9\x28\xdd\x7d\x70\x27\x34\xdb\x8f\xce\xec\x17\xc8\x6c\x9f\x34\x64\xce\xac\xf5\x36\xd0\xe3\xee\x8c\xd8\x7e\x65\x67\x80\x15\x58\x99\xfe\x26\x05\xd5\x5f\xff\x69\x1d\x15\x61\x37\xbc\xb4\xc2\x29\xfa\xf3\xa3\x47\x5c\x14\xd3\x7d\xa9\xf4\xf4\x16\x4f\xd5\xee\x8c\x8a\x6b\x3e\x65\xca\xc7\x00\x13\xdf\xda\xbd\x7e\x2f\xad\xae\x8a\x6d\x39\x76\xd5\x6d\x0d\x5e\x5f\x81\xd3\x88\xc4\x79\xb6\x5b\x3c\xcf\x22\x12\x0f\xee\xaa\x11\x2b\x9d\xb8\x72\xca\x91\x35\xd2\xe1\x31\xf7\x6d\x7d\xe3\xcc\x18\x5c\x57\xac\x74\xe7\x4c\xcc\x20\xeb\x23\x86\xb2\x68\x07\x5b\xbc\x50\xb7\x09\x57\x5a\x2b\x67\xae\xac\x0c\xb0\x3e\x52\x41\x2a\xb8\xad\x65\x5c\x34\x6f\x47\x37\x13\x57\x56\x2b\x6f\xce\x7d\x5d\x03\x6b\x64\x25\x91\x66\x02\xad\xbc\xc6\xc6\xe4\xcc\xac\xf4\xe8\xcc\x2d\x87\xd6\xd8\x44\x62\x4a\x02\xb4\xb1\xbb\xa6\x83\x00\x9d\xb9\x95\xee\x5c\xa9\xe5\xc8\xfa\xd6\x4d\xf2\x34\xe4\xd2\x76\x61\x31\x06\x57\x62\xa5\x3b\xe7\xcb\x89\x41\xd6\x47\xec\x89\x25\x3e\x7f\x92\xcd\x1f\x74\xb5\x3f\xf3\x29\x70\xbe\x67\x6b\xb3\x85\x70\xa5\xbd\xed\xa3\x3e\x01\x5e\x78\x62\x65\xfd\x97\x1e\x77\x64\xba\xf2\xe5\x4a\xcf\x00\xeb\xe3\x94\xbe\xd6\x11\x6e\x87\xd0\x04\x7d\x2d\x32\xef\x10\x04\xd3\xe9\xa7\x5c\x28\xed\x07\xf0\x19\x69\x66\x3e\xee\xd0\xc5\x67\xb4\xf1\x4d\xf7\x6e\x16\x69\xa5\x6e\x41\xd2\x34\x62\xd4\xbc\x6e\x6d\x3d\x37\x7d\x2e\x75\x8a\x6b\x6f\x34\xf2\xe8\x55\x8f\xde\x88\x4b\xdb\x0e\xbe\x60\x02\x87\x0b\x84\x2b\x32\x95\x70\x33\x0d\x32\xa9\x99\xfb\x30\x56\x82\x25\x33\x18\xf3\xa8\x78\x8f\x50\x41\x90\x55\x0c\xe7\xa3\x58\x03\xf7\x94\xa3\xea\xe1\xcc\x83\x80\x51\xdb\xe1\x7c\x63\x0c\xbb\x95\x71\x73\x35\x6c\x8d\x26\x17\xe0\x73\x9a\xc5\x98\xa8\x9d\xb2\x58\x1a\x6c\x52\x94\xd1\x5c\xb5\xc8\x91\x35\xee\x06\x3f\xb0\x6d\x85\xf3\x8b\x5d\xc6\x7a\x70\x2f\xb6\x45\x0c\xe7\x6d\xe0\x07\x2e\x8b\x5e\x44\xa9\x54\xfd\xdb\xab\xff\xb7\xbe\xfd\x1a\x01\x2c\x62\x7c\x5a\x19\x77\x35\xf9\x34\x86\xdc\xe4\x24\xc8\x66\x38\x57\x5d\x4a\xb4\x8b\x3c\xeb\x31\x2b\x5f\x7b\x05\x0d\xd9\xdc\xb6\x2b\x06\xb9\x65\x4d\x90\x2c\x29\x66\x43\xf1\xdc\x6d\x0b\xa5\xb7\x9d\x2b\x71\x69\x72\xbe\x14\xe7\xd0\x9f\x4a\x51\x46\x78\x61\xe9\xda\x8b\x36\x45\x84\x77\x08\xde\xec\x45\xff\x4f\xfb\x2f\x1e\xdc\x37\xfe\x0d\x00\x00\xff\xff\x23\x12\x19\xef\x2c\x1c\x00\x00")

func pluginsPluginsTomlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsPluginsToml,
		"plugins/plugins.toml",
	)
}

func pluginsPluginsToml() (*asset, error) {
	bytes, err := pluginsPluginsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/plugins.toml", size: 7212, mode: os.FileMode(420), modTime: time.Unix(1535991814, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonDockerignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xf0\x4c\xcf\xcb\x2f\x4a\x55\xd0\x4b\xcf\x2c\x51\x48\xcb\xcf\x49\x49\x2d\xe2\x02\xb1\xb5\xb8\xc0\x54\x26\x58\x96\x2b\xc8\xd5\xd1\xc5\xd7\x55\x2f\x37\x85\xcb\xc7\xd3\xd9\xd5\x2f\x18\xcc\x74\xf6\x70\xf4\x73\x77\xf5\xf1\x77\x07\x71\x92\x4a\x33\x73\x52\xb8\x8a\x52\x73\x52\x13\x8b\x53\xb9\x7c\x13\xb3\x53\xd3\x32\x73\x52\xb9\x92\x33\x8b\x92\x73\x52\xf5\x2a\x73\x73\x00\x01\x00\x00\xff\xff\xba\xc0\x2a\xa8\x6a\x00\x00\x00")

func pluginsTemplatesPythonDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonDockerignore,
		"plugins/templates/python/.dockerignore",
	)
}

func pluginsTemplatesPythonDockerignore() (*asset, error) {
	bytes, err := pluginsTemplatesPythonDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/.dockerignore", size: 106, mode: os.FileMode(420), modTime: time.Unix(1504646603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonGitignoreTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x52\xb1\x6e\x1b\x31\x0c\xdd\xf9\x15\x04\xb2\x19\xb1\xb4\x14\x1d\x3a\xa6\x5e\x0a\x64\x08\x90\x66\x2a\x0a\x43\x96\x78\x77\x4c\x75\xa2\x20\xd2\x8e\xaf\x5f\x5f\x48\x4e\x9b\x2e\x77\x8f\x7c\xe2\xd3\xa3\x48\x77\x78\x3e\x3e\x9b\x34\x02\xb8\xc3\x87\xcd\x68\x1f\x65\xad\x9c\x29\xa1\x47\xa9\xc6\x2b\xff\x1e\xf8\xf0\xf8\x88\x13\x67\x52\x38\x1e\xeb\x16\x43\x5c\xe8\x78\xf4\xb0\x73\x75\xfb\x11\x25\xfd\xec\xf5\x5f\x91\xae\x46\x45\x59\x8a\xc2\xce\xa9\xf4\xe4\x81\xd5\x1a\x9f\xce\xc6\x52\xd0\x63\x0d\xf1\x57\x98\xb9\xcc\xe0\x9e\x36\x5b\xa4\x00\x95\x8b\x87\xd3\x99\x73\xf2\x90\xe8\x42\x59\xea\x9e\xe6\x59\x3d\x24\x56\xf3\x90\xe4\xad\x64\x09\x49\x3d\xdc\xd2\xee\xf6\xcb\x7c\x1a\x9f\xcf\x9f\x3c\xd4\xd0\x4c\x3d\xe8\xad\xe2\x12\x5a\x77\x46\xf3\xbc\xe7\x32\x89\x07\xc7\x45\x2d\xe4\x4c\xc9\xc5\x69\xbe\x51\xdd\xdb\xd3\xf6\xed\x9d\x68\x70\x87\xf8\xa2\xe7\x90\xf3\x86\xb6\x90\xd2\xad\x5b\x0c\x8d\xf0\xad\xb1\x19\x15\x3c\x6d\x18\xb0\x0e\xd7\xa8\xb1\x71\x35\x9c\x9a\xac\x18\xd0\x68\xad\x39\x18\x75\x95\x13\x4d\xd2\xe8\x7f\x6d\x1c\xdd\x69\xd7\x45\xba\xd2\x3d\xaa\x60\x50\x34\x41\x2e\xaf\x14\x0d\x53\x30\xf2\x62\x0b\x35\xec\x86\x15\xb9\x74\xd2\x1c\xec\xdc\x1a\x0a\x4f\xa4\xd6\x1f\xb4\x52\xec\xb6\x3f\x84\xb3\xcc\x0a\x95\xeb\x3e\xcb\xec\xec\x6a\x03\x27\xca\x64\xb4\xb7\x85\x75\x9f\xb8\x51\x34\x69\xdb\x60\xe1\x0e\x5f\x0a\x1b\x1a\xa9\xa1\xc7\x28\x17\x6a\x61\x26\x6c\x54\xa5\x99\xc2\x62\x6b\x8e\x72\xf1\xe0\x4c\xae\x1e\xdc\xdf\x03\x1f\xc8\xed\xc0\x8d\xe1\x43\x11\xa5\xae\xa3\xee\xba\x66\xf8\xc7\xf7\x60\x77\x3f\xc2\x7e\xdd\xf7\x16\x8a\xe6\x60\xef\x2b\xb1\x4a\xdf\x18\x19\x4e\x0e\xaf\xa1\xcc\x82\x6a\xe7\x69\xfa\x02\x3b\x97\x65\xcc\xe4\xb9\x2e\x5c\xae\x98\x24\x9e\x57\x2a\x36\x4a\x21\x49\x54\x7f\x7c\x5f\x92\x31\xb8\x87\x8e\xa9\x81\x85\x36\x93\xf9\x3f\x01\x00\x00\xff\xff\x27\xc3\x36\xdb\xc8\x02\x00\x00")

func pluginsTemplatesPythonGitignoreTemplateBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonGitignoreTemplate,
		"plugins/templates/python/.gitignore.template",
	)
}

func pluginsTemplatesPythonGitignoreTemplate() (*asset, error) {
	bytes, err := pluginsTemplatesPythonGitignoreTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/.gitignore.template", size: 712, mode: os.FileMode(493), modTime: time.Unix(1504646603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonChangelogMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x56\x30\xd0\x33\xd4\x33\x50\xd0\x55\x70\xcb\x2c\x2a\x2e\x51\x08\x4a\xcd\x49\x4d\x2c\x4e\xe5\xd2\x52\x70\x2d\x4b\x2d\xaa\x54\x48\x4b\x4d\x2c\x29\x2d\x4a\x55\x48\x4c\x49\x49\x4d\x81\x8b\x26\x95\xa6\x2b\xa4\x65\x56\xa4\xa6\x70\x01\x02\x00\x00\xff\xff\x82\x96\x39\xb5\x41\x00\x00\x00")

func pluginsTemplatesPythonChangelogMdBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonChangelogMd,
		"plugins/templates/python/CHANGELOG.md",
	)
}

func pluginsTemplatesPythonChangelogMd() (*asset, error) {
	bytes, err := pluginsTemplatesPythonChangelogMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/CHANGELOG.md", size: 65, mode: os.FileMode(493), modTime: time.Unix(1504646603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x41\xab\xe2\x30\x14\x85\xf7\xf9\x15\x97\x2e\x5c\x0c\xa4\xdd\x0f\xb8\xe8\xb4\x1d\x0c\x8e\x6d\x89\xce\x0c\xe2\x48\x89\xed\xb5\x06\xd3\x34\x93\xa6\xa8\x14\xff\xfb\x50\xe7\x55\x1e\x3c\x79\xbb\xe4\xdc\xef\x72\xcf\x39\xdf\x79\xb6\x82\x46\x28\x59\x62\x20\x94\x91\x1a\xbf\x3a\xa9\x25\x21\xab\x90\xa5\x9b\x90\xa5\x09\x87\x61\x80\xd2\xa2\x70\xad\x85\xfb\x9d\x90\x28\xcb\xb7\xe0\x43\xe0\x1a\x13\x74\xb6\x0c\x86\x01\x8c\xea\x6b\xa9\x0b\x2d\x1a\x1c\x11\xfe\x33\x05\x61\xce\x54\xea\xce\x09\xa5\xc0\xdc\xdc\xa9\xd5\x1f\x64\xea\xc0\x3f\xf4\x52\x55\xb4\x42\xd3\x41\x2d\x1d\x34\x68\xcb\xde\x4a\x31\xee\x50\x23\x0d\xfc\x21\x00\xb3\x19\x74\xe8\x80\x5e\xa7\x5f\x59\x7d\x72\x7d\x82\xf0\x6a\x5a\xeb\x20\x67\x79\x91\x66\x45\x14\x46\x8b\xa4\x88\x19\x9f\xb7\xc7\xe3\x0b\x24\x66\xeb\xf0\xdb\x8f\xa4\x18\xdf\xbf\x12\xbe\x66\x59\x5a\x44\x8b\x24\x5a\xce\x5b\x3d\xe1\xa3\x9f\xa7\x77\xda\x9b\xda\x8a\x0a\x1f\xea\xe5\x84\xa8\x5e\x62\x16\x2c\xfe\xed\xa5\xc5\x06\xb5\xeb\x7c\x77\x75\x4f\xec\x51\xca\x98\xac\x37\xbe\xb9\x3d\x57\xde\xc6\xb6\x01\x6a\x8f\xff\x73\x7e\x99\x44\x61\xce\x50\xe1\x78\xdd\xf4\xb6\xc6\xf7\xf5\x11\xf2\x3b\xe3\xcb\x98\x71\x08\x1a\xa1\x2e\xc2\x22\x21\x49\xba\xe1\xdb\x3c\x63\xe9\x06\x76\x5e\x70\x90\x3a\xe8\x4a\xa1\xbd\x3d\x21\xd1\x2a\x86\x9d\x47\xe9\x09\x95\xf1\xf6\xff\x02\x00\x00\xff\xff\x22\x74\x40\xdc\x05\x02\x00\x00")

func pluginsTemplatesPythonDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonDockerfile,
		"plugins/templates/python/Dockerfile",
	)
}

func pluginsTemplatesPythonDockerfile() (*asset, error) {
	bytes, err := pluginsTemplatesPythonDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/Dockerfile", size: 517, mode: os.FileMode(420), modTime: time.Unix(1504646603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonLicenseMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x51\x4f\x8f\xe3\x26\x14\xbf\xfb\x53\xfc\x34\xa7\x5d\xc9\x9a\xde\x7b\x63\x62\x1c\xa3\x3a\x10\x61\xb2\x69\x4e\x15\xb1\x49\x4c\xe5\x40\x04\xa4\xa3\x68\xb5\xdf\xbd\x82\x64\x3b\x3b\x3d\x59\xe6\xbd\xdf\xdf\xb7\xf2\xd7\x7b\xb0\xe7\x39\xe1\xcb\xf8\x15\xdf\xbf\x63\xbc\x85\x60\x5c\xfa\xeb\x6e\x74\xc0\x8f\x1f\xe5\x29\x18\x9d\x7c\xfe\xab\xaa\xad\x09\x17\x1b\xa3\xf5\x0e\x36\x62\x36\xc1\x1c\xef\x38\x07\xed\x92\x99\x6a\x9c\x82\x31\xf0\x27\x8c\xb3\x0e\x67\x53\x23\x79\x68\x77\xc7\xd5\x84\xe8\x1d\xfc\x31\x69\xeb\xac\x3b\x57\x1a\xa3\xbf\xde\xf3\x66\x9a\x6d\x44\xf4\xa7\xf4\xae\x83\x81\x76\x13\x74\x8c\x7e\xb4\x3a\x99\x09\x93\x1f\x6f\x17\xe3\x92\x4e\x59\xef\x64\x17\x13\xf1\x25\xcd\xa6\x7a\x19\x9e\x88\x97\xaf\x45\x64\x32\x7a\x81\x75\x48\xb3\xc1\xcf\x11\xde\x6d\x9a\xfd\x2d\x21\x98\x98\x82\x1d\x33\x47\x0d\xeb\xc6\xe5\x36\x65\x0f\x3f\xc7\x8b\xbd\xd8\xa7\x42\x86\x97\x32\x62\x26\xbd\x45\x53\x17\x9f\x35\x2e\x7e\xb2\xa7\xfc\x35\x25\xd6\xf5\x76\x5c\x6c\x9c\xeb\x6a\xb2\x99\xfa\x78\x4b\xa6\x46\xcc\x8f\xa3\x71\x19\xa5\xdd\xf4\x9b\x0f\x88\x66\x59\x32\x83\x35\xf1\x91\xf5\xc3\x5d\xd9\x41\xf2\xd5\x35\x17\x9a\x9e\x15\x15\xdd\xf7\xd9\x5f\x3e\x27\xb1\x11\xa7\x5b\x70\x36\xce\x66\x2a\x71\x3d\xa2\x2f\x8a\x7f\x9b\x31\x65\x96\xbc\x7e\xf2\xcb\xe2\xdf\xad\x3b\x63\xf4\x6e\xb2\x39\x51\xfc\xbd\xaa\xd4\x6c\xa0\x8f\xfe\x1f\x53\xb2\x3c\x6e\xed\x7c\xb2\xe3\xa3\xee\x72\x80\xeb\xc7\x55\x9f\xa3\x38\xeb\x65\xc1\xd1\x54\x8f\xc2\xcc\x94\xeb\xd5\xbf\xc4\x09\x59\x3e\x26\xed\x92\xd5\x0b\xae\x3e\x14\xbd\xff\xc7\x7c\xad\x2a\xd5\x51\x0c\xa2\x55\x7b\x22\x29\xd8\x80\xad\x14\xdf\x58\x43\x1b\xbc\x90\x01\x6c\x78\xa9\xb1\x67\xaa\x13\x3b\x85\x3d\x91\x92\x70\x75\x80\x68\x41\xf8\x01\x7f\x30\xde\xd4\x15\xfd\x73\x2b\xe9\x30\x40\x48\xb0\xcd\xb6\x67\xb4\xa9\xc1\xf8\xaa\xdf\x35\x8c\xaf\xf1\xb6\x53\xe0\x42\xa1\x67\x1b\xa6\x68\x03\x25\x90\x05\x9f\x54\x8c\x0e\x10\x6d\xb5\xa1\x72\xd5\x11\xae\xc8\x1b\xeb\x99\x3a\xd4\x68\x99\xe2\x99\xb3\x15\x12\x04\x5b\x22\x15\x5b\xed\x7a\x22\xb1\xdd\xc9\xad\x18\x28\x08\x6f\x2a\x2e\x38\xe3\xad\x64\x7c\x4d\x37\x94\xab\x57\x30\x0e\x2e\x40\xbf\x51\xae\x30\x74\xa4\xef\x8b\x14\xd9\xa9\x4e\xc8\xe2\x6f\x25\xb6\x07\xc9\xd6\x9d\x42\x27\xfa\x86\xca\x01\x6f\xb4\xea\x19\x79\xeb\xe9\x43\x8a\x1f\xb0\xea\x09\xdb\xd4\x68\xc8\x86\xac\x69\x41\x09\xd5\x51\x89\xbc\xf6\x74\xb7\xef\x68\x79\x62\x1c\x84\x83\xac\x14\x13\xbc\x12\x2d\x56\x82\x2b\x49\x56\xaa\x86\x12\x52\xfd\x07\xdd\xb3\x81\xd6\x20\x92\x0d\xb9\x90\x56\x8a\x4d\x8d\x5c\xa7\x68\x4b\x67\x3c\xe3\x38\x7d\xb0\xe4\xaa\xf1\xe9\x22\x42\x96\xff\xdd\x40\x3f\xbc\x34\x94\xf4\x8c\xaf\x87\x0c\xfe\x75\xf9\xb5\xfa\x37\x00\x00\xff\xff\xe4\x3f\x0e\xa2\x2f\x04\x00\x00")

func pluginsTemplatesPythonLicenseMdBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonLicenseMd,
		"plugins/templates/python/LICENSE.md",
	)
}

func pluginsTemplatesPythonLicenseMd() (*asset, error) {
	bytes, err := pluginsTemplatesPythonLicenseMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/LICENSE.md", size: 1071, mode: os.FileMode(493), modTime: time.Unix(1504646603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonReadmeShortTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\x28\xc8\x29\x4d\xcf\xcc\x8b\x4f\x49\x2d\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\x53\xa8\xad\x05\x04\x00\x00\xff\xff\xf7\x36\xbe\xfa\x18\x00\x00\x00")

func pluginsTemplatesPythonReadmeShortTxtBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonReadmeShortTxt,
		"plugins/templates/python/README-short.txt",
	)
}

func pluginsTemplatesPythonReadmeShortTxt() (*asset, error) {
	bytes, err := pluginsTemplatesPythonReadmeShortTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/README-short.txt", size: 24, mode: os.FileMode(420), modTime: time.Unix(1504646603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcb\xb1\x4e\xc5\x20\x14\x80\xe1\x9d\xa7\x38\xc6\x45\x07\xa0\xb4\x29\xbd\xb8\xf9\x00\x26\xee\xc6\x18\x38\x1c\x5a\x92\x5e\x68\x0a\x0c\xa6\xe1\xdd\x4d\x8c\x83\xeb\xff\xe7\x7b\x84\xeb\x82\x63\x6f\x6b\x4c\x5f\xc9\xde\x09\x7a\x67\xec\xcd\xee\x11\xe9\xdf\xf1\x54\xf0\x8c\x47\x8d\x39\x41\xef\xf0\xfe\x5b\x05\x63\x0f\x1f\xaf\x50\xf0\x24\x4a\x65\xcb\x15\x72\x80\xef\xdc\xce\x3f\xf5\xf9\xb4\xd5\x7a\x94\x17\x29\x83\xc0\x3d\x37\x2f\xd6\x58\xb7\xe6\x04\xe6\xbb\xb4\xa5\x50\x2d\x52\x1b\xa5\x8d\x1c\x47\x33\x8c\xf3\x20\x71\x9a\xfd\x4d\x2f\x96\xdb\x41\x2d\x5c\x29\x9a\xf8\x4d\x3b\xe2\xe8\x17\x9c\x5d\x98\x42\x30\x4e\xac\x31\x3c\xb3\x9f\x00\x00\x00\xff\xff\x6d\x30\xd9\xc7\xb7\x00\x00\x00")

func pluginsTemplatesPythonReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonReadmeMd,
		"plugins/templates/python/README.md",
	)
}

func pluginsTemplatesPythonReadmeMd() (*asset, error) {
	bytes, err := pluginsTemplatesPythonReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/README.md", size: 183, mode: os.FileMode(493), modTime: time.Unix(1504646603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonCircleYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xcd\x6e\xc2\x30\x10\x84\xef\x79\x8a\x15\xed\xa1\x3d\x84\xdc\x41\xea\x8b\x20\x84\x16\x7b\x48\x56\xf8\x27\x5a\x3b\x91\x2a\x44\x9f\xbd\x22\x38\x14\x24\xca\xcd\x1a\xef\x7c\x33\x1a\xcf\xa6\x93\x80\x55\x45\x94\xa0\xa3\x18\xa4\xcb\x9b\xa8\x26\x1b\xcd\x11\x5a\x55\x16\x3d\x82\x45\x30\x72\xfd\x33\x6c\x3a\xec\xac\x28\x4c\x8e\x2a\x7f\x86\xc5\x4f\x73\xf5\x2c\x2a\xa2\x38\x42\x55\x2c\x1e\x69\x24\xe1\x10\x8b\x22\x07\xda\x6c\xa8\x06\xcd\xb6\x46\x3c\xb7\x58\x66\x56\xda\x6e\xd7\x94\x3b\x84\xd9\xe6\x22\x5b\xaa\x6b\x09\xfd\x90\x9f\xdc\xaf\xe9\x20\x8f\x39\xfb\x41\x9c\xa5\x3a\x93\x11\x35\x0e\x46\x9a\xd3\x89\x7a\x37\xb4\x12\x76\x81\x3d\xe8\x7c\xa6\x65\xb1\xf8\xa3\x15\xa5\xba\xbf\x81\xd7\x33\x25\xf1\x88\x57\x84\xaf\x27\x5d\xaa\x2a\x23\xe5\xd5\xff\x13\xe8\x10\x5e\x31\x2f\xee\x69\x74\x17\xbf\x3d\xc2\x44\x7a\x23\xcf\x29\x43\xaf\x6f\xa2\xbd\x72\x30\xdd\xaa\xa8\x45\x34\xd1\x7b\x0e\x36\xcd\x47\x97\x48\x8c\xec\xe8\xfd\xe3\x2e\xb9\x75\x62\xa1\x8e\xf7\xa9\xe9\xb9\xc5\xb4\x13\xf4\x56\xa8\x36\xde\x7e\x56\x44\x0a\x07\x4e\xa5\xfa\x1c\x57\xc4\x49\xbb\x4f\x2b\x33\xf2\x11\xb7\x93\xdf\x00\x00\x00\xff\xff\x5b\x80\xf9\x05\x58\x02\x00\x00")

func pluginsTemplatesPythonCircleYmlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonCircleYml,
		"plugins/templates/python/circle.yml",
	)
}

func pluginsTemplatesPythonCircleYml() (*asset, error) {
	bytes, err := pluginsTemplatesPythonCircleYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/circle.yml", size: 600, mode: os.FileMode(420), modTime: time.Unix(1504646603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonPluginToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\xcb\x6a\xeb\x30\x10\x86\xf7\x7e\x8a\x41\xd9\x1e\x0e\xb4\x90\xee\xb2\x28\xa1\x94\x40\x5a\xba\xe9\xca\x04\xa3\xc4\x13\x67\x88\x6e\x48\xb2\x4b\x08\x7e\xf7\xa2\x8b\xd3\xd8\xe9\xa6\x10\x6f\x0c\xdf\x7c\x1a\xfd\x9a\x99\xdd\xe7\x2b\x66\xf0\xb1\xfe\x7c\x5d\xbd\xc3\x52\xab\x3d\x35\xad\xe5\x9e\xb4\x82\xbf\xf7\xb9\x53\x9e\xa2\x2c\x8d\x68\x1b\x52\x9b\x4d\x01\xa0\xb8\x44\x58\x00\x3b\x9f\x21\xd1\x2a\x92\xbe\x67\x05\x40\x87\xd6\x85\xac\xa3\xfa\x00\x93\x52\xa3\xdb\x59\x32\xfe\x46\xbb\x2e\x24\x75\xc7\x3d\x36\xda\x9e\xc6\xde\x85\x26\xe9\x88\xa7\x2f\x6d\x6b\x07\x0b\x28\xaf\xac\x0b\xee\x7b\x16\x62\x93\xe4\xcd\x24\x77\x42\xa9\x8b\x45\xa3\x1d\xf9\x9b\xcb\xae\x78\x12\x05\xed\x50\xb9\x49\xa7\x01\x26\x05\x55\x43\x0a\x5d\x56\x24\x0f\xd5\x2a\xc3\xca\xb5\xc6\x68\xeb\xb1\xce\x72\x1e\xed\x7f\x43\x46\x04\x21\x64\x05\xf0\x7a\x7c\x81\x21\x83\xae\xf2\x3a\x1f\x02\xd8\x5b\x2d\x7f\x53\x22\x4f\xd2\xb6\x25\x51\xc3\x02\xf6\x5c\x38\x2c\x00\xb8\xa1\x23\xc6\xe7\xc5\xd9\xca\x7a\x32\x56\x39\x44\x92\x34\xdd\x70\x24\xa9\x78\xe0\xee\xe0\x4f\x26\x3e\xaf\x04\x26\xeb\x39\xfb\x07\xcc\x1d\xf8\x43\xfe\x3f\xce\x9f\x18\x6c\xe2\x18\xba\xb8\x93\xb7\xe7\xf5\x6a\xf9\x52\xfd\x74\x43\xd5\x55\x1d\xb7\xc3\x66\x50\xf1\xad\xc0\x90\xc6\xdb\x16\xbf\x03\x00\x00\xff\xff\xab\x02\x2d\x63\x45\x03\x00\x00")

func pluginsTemplatesPythonPluginTomlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonPluginToml,
		"plugins/templates/python/plugin.toml",
	)
}

func pluginsTemplatesPythonPluginToml() (*asset, error) {
	bytes, err := pluginsTemplatesPythonPluginTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/plugin.toml", size: 837, mode: os.FileMode(420), modTime: time.Unix(1504646603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonRequirementsTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x28\xc8\x2c\x50\xc8\xcc\x2b\x2e\x49\xcc\xc9\x51\xd0\x2d\x52\x28\x4a\x2d\x2c\xcd\x2c\x4a\xcd\x4d\xcd\x2b\x29\xd6\x2b\xa9\x28\x01\x04\x00\x00\xff\xff\x5d\x24\x95\x7b\x21\x00\x00\x00")

func pluginsTemplatesPythonRequirementsTxtBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonRequirementsTxt,
		"plugins/templates/python/requirements.txt",
	)
}

func pluginsTemplatesPythonRequirementsTxt() (*asset, error) {
	bytes, err := pluginsTemplatesPythonRequirementsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/requirements.txt", size: 33, mode: os.FileMode(420), modTime: time.Unix(1504646603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonScanPy = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\x4f\x8b\xdb\x3e\x10\xbd\xeb\x53\xbc\x9f\x96\xc5\xc9\x6f\x63\xa7\x2d\x6c\x29\x06\x1d\x0a\xd9\x42\x0f\x0b\x3d\xf4\x6e\x14\x7b\x62\x0b\x62\x49\x1d\xc9\xd9\x1a\x93\x7e\xf6\x22\xc5\xed\x16\x4a\x7d\x11\xf3\xe6\xbd\x79\xf3\xc7\x77\xff\xed\xa7\xc0\xfb\xa3\xb1\x7b\xb2\x17\xf8\x39\x0e\xce\x8a\x3b\x94\xff\x97\x68\x5d\x67\x6c\x5f\x63\x8a\xa7\xf2\x43\x42\x84\x94\x52\x84\x56\xdb\xca\xcf\xe2\xc7\xfa\x89\x65\x41\x47\xa1\x65\xe3\xa3\x71\x16\xd7\x6b\x25\xea\xd6\xf9\x99\x4d\x3f\xc4\x1a\x9b\x76\x8b\x77\x6f\xde\xbe\xc7\x71\xc6\xb2\xa0\x65\xd2\xd1\xf1\x8d\x76\x36\x2d\xd9\x40\x35\x9e\x3f\x7f\xcd\xc5\x85\x19\xbd\xe3\x08\xcd\xbd\xd7\x1c\x48\x08\xd1\xd1\x09\xa3\x36\x76\xb3\xad\x05\x00\x64\x9c\xa1\x7e\x73\xaa\x8f\xdc\x4f\x23\xd9\xf8\x25\x67\x36\x9e\x5d\xaf\x8a\x65\x81\x3f\x4f\xbd\xb1\x8d\xd5\x23\xe1\x7a\x2d\xb6\x7f\xc8\x2b\xdd\x75\x8d\x5e\x75\x1b\x59\x5e\xe4\x0e\xb2\x2c\x2f\xc4\x47\x17\x48\xee\x30\xd0\xd9\x2b\x79\x30\xc1\x9f\xf5\x8c\x15\x87\x9b\xa2\x9f\x22\x46\x0a\x41\xf7\x89\xa6\xdb\x34\xb3\x92\x21\x3a\xa6\x26\xf2\x94\x40\xa6\x6f\x93\x61\xea\xd4\x27\x7d\x0e\xf4\x6f\xdb\x62\xd0\x61\x28\x76\x18\x29\xea\x8b\x66\x55\x3c\x1f\x1e\x8b\x1d\xe2\xec\x49\x85\xc8\x3b\x58\xcd\x7d\x50\xc5\x43\xb1\xf6\x53\x68\x8c\xdd\x23\x92\x0c\xd1\x21\x90\xe6\x76\xc0\xc9\x71\xb5\x0e\x97\xf8\x50\xbf\xcc\xf2\x93\xec\xc2\x66\x2b\x72\xde\x9c\x32\xa5\x4a\x15\x6e\xdb\xcc\xbd\xb1\xb1\x11\xf2\xe0\x10\xdc\x48\x71\x30\xb6\x87\x7e\xa1\x14\xe0\xc5\xc4\x21\x1b\xd6\x90\x78\x78\x55\x67\x31\x53\x9c\xd8\x0a\x21\xcc\x09\x4d\xde\x73\xd3\x40\x29\xc8\xa6\x49\x17\x6b\x1a\x79\x33\x89\x3c\xbf\xba\xdd\x6e\x99\x43\xfa\xde\x92\x8f\x78\xca\x4f\xfa\x77\x74\x00\xfd\xd5\xd7\x13\xb3\xe3\x1a\xf7\x41\xe2\x1e\x24\x7e\x06\x00\x00\xff\xff\xe2\x82\x12\x89\xb2\x02\x00\x00")

func pluginsTemplatesPythonScanPyBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonScanPy,
		"plugins/templates/python/scan.py",
	)
}

func pluginsTemplatesPythonScanPy() (*asset, error) {
	bytes, err := pluginsTemplatesPythonScanPyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/scan.py", size: 690, mode: os.FileMode(420), modTime: time.Unix(1504646603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"plugins/plugins.toml": pluginsPluginsToml,
	"plugins/templates/python/.dockerignore": pluginsTemplatesPythonDockerignore,
	"plugins/templates/python/.gitignore.template": pluginsTemplatesPythonGitignoreTemplate,
	"plugins/templates/python/CHANGELOG.md": pluginsTemplatesPythonChangelogMd,
	"plugins/templates/python/Dockerfile": pluginsTemplatesPythonDockerfile,
	"plugins/templates/python/LICENSE.md": pluginsTemplatesPythonLicenseMd,
	"plugins/templates/python/README-short.txt": pluginsTemplatesPythonReadmeShortTxt,
	"plugins/templates/python/README.md": pluginsTemplatesPythonReadmeMd,
	"plugins/templates/python/circle.yml": pluginsTemplatesPythonCircleYml,
	"plugins/templates/python/plugin.toml": pluginsTemplatesPythonPluginToml,
	"plugins/templates/python/requirements.txt": pluginsTemplatesPythonRequirementsTxt,
	"plugins/templates/python/scan.py": pluginsTemplatesPythonScanPy,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"plugins": &bintree{nil, map[string]*bintree{
		"plugins.toml": &bintree{pluginsPluginsToml, map[string]*bintree{}},
		"templates": &bintree{nil, map[string]*bintree{
			"python": &bintree{nil, map[string]*bintree{
				".dockerignore": &bintree{pluginsTemplatesPythonDockerignore, map[string]*bintree{}},
				".gitignore.template": &bintree{pluginsTemplatesPythonGitignoreTemplate, map[string]*bintree{}},
				"CHANGELOG.md": &bintree{pluginsTemplatesPythonChangelogMd, map[string]*bintree{}},
				"Dockerfile": &bintree{pluginsTemplatesPythonDockerfile, map[string]*bintree{}},
				"LICENSE.md": &bintree{pluginsTemplatesPythonLicenseMd, map[string]*bintree{}},
				"README-short.txt": &bintree{pluginsTemplatesPythonReadmeShortTxt, map[string]*bintree{}},
				"README.md": &bintree{pluginsTemplatesPythonReadmeMd, map[string]*bintree{}},
				"circle.yml": &bintree{pluginsTemplatesPythonCircleYml, map[string]*bintree{}},
				"plugin.toml": &bintree{pluginsTemplatesPythonPluginToml, map[string]*bintree{}},
				"requirements.txt": &bintree{pluginsTemplatesPythonRequirementsTxt, map[string]*bintree{}},
				"scan.py": &bintree{pluginsTemplatesPythonScanPy, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

