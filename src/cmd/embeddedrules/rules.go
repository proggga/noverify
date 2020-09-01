// Package embeddedrules Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// embeddedrules/rules.php
package embeddedrules

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _embeddedrulesRulesPhp = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x98\xd1\x6f\xe2\xb8\x13\xc7\xdf\x91\xf8\x1f\x06\x94\x52\x60\x81\xfc\xaa\xdf\x5b\x5b\xda\xbd\x7b\x3b\x69\x75\x95\xee\x1e\x4f\x2d\x38\xc9\x84\x78\x1b\xec\x9c\xed\x2c\x70\x85\xff\xfd\xe4\x38\x84\x40\x0c\x24\x5b\x5d\xa5\x52\x18\xdb\xf3\x99\x99\xef\x24\x19\xfa\xf8\x9c\x44\x49\xbb\xd5\x6e\xb9\xc3\x61\xbb\x05\x43\xf8\xca\x38\x65\x32\x41\x5f\x51\xce\xe0\x97\x6f\xdf\x8c\x35\xa6\x4c\xa1\x00\xf3\x13\x50\x49\xbc\x18\xf5\x8a\x7b\x74\xd6\xe7\xcb\x25\x32\x05\x7f\x60\xc2\x85\x02\x85\x82\x11\xb1\x01\x5c\x27\x02\xa5\xa4\x9c\x49\x50\x11\x51\xe0\x13\x06\x1e\x82\xa4\xcb\x24\xa6\x21\xc5\x60\x62\xce\x7b\x18\x72\x81\x00\xce\x1a\x9e\xf5\xcb\x3d\x38\x1b\xb3\x42\x42\x83\xd7\x2b\x07\x6b\x48\xd7\x59\x40\x65\xab\xdb\x6e\x85\x29\x33\xe1\xe7\x01\xfc\x69\x38\x9b\xfe\x00\x3e\xda\x2d\x00\x13\x2f\x68\x0f\x4b\xb2\xf1\x10\x7c\x9e\xc6\x01\x08\x4c\x62\xe2\x23\xa8\x08\x8b\xd0\x57\x54\x45\xf0\x3d\x95\x0a\x1c\x9f\xb3\x60\x7f\x4c\x6d\x12\x04\x8f\xf3\xb8\x6c\x76\xf5\x9f\xec\x33\x3c\x83\x12\x29\xc2\x3d\x84\x24\x96\xf8\xa0\x8b\x74\x01\xbb\x12\x54\x21\x10\x09\xf3\xbe\xf6\x39\xc8\x7c\xcc\x8f\x58\xd0\xf9\x0f\x68\xfb\xb2\x15\x28\x5d\xcf\x52\x2d\x33\x5b\x92\x0a\x04\x67\x5d\x86\x96\xd4\x69\x06\x7b\xb6\xc1\x9e\x0b\x58\xe6\x9d\x4a\x89\xaa\xef\xac\x07\x47\x90\xdd\xa5\x3e\x4b\xb8\x42\xa6\x28\x89\x81\x27\x28\x48\x26\x7d\x22\xd0\xc7\x00\x99\x8f\xda\x63\x8a\xb2\xda\x62\x3d\x70\x96\x44\xbe\xc3\x74\x0a\xff\x7b\x00\xd7\xd5\x6f\x22\x22\x21\xa2\x8b\x08\x45\xd9\x83\x8a\x08\x83\xde\x71\x27\xf6\x0f\x1e\x06\x99\x8b\xd3\xe6\x3b\x9c\x2f\xfa\xce\x85\xdf\xb9\xc2\x7b\x58\xa1\x6e\x36\x1d\xfa\xbc\x14\x47\x67\xaa\xcb\xa3\x8b\x45\x2a\xe1\x03\xfe\x40\x96\xfb\xa0\x21\xd0\xe2\x1a\xf2\x49\xba\x88\x94\xe9\xd3\xb9\x6e\x15\x3f\x42\xff\xfd\x25\x99\xe7\x57\x9a\xfe\x24\xc1\xe3\x2a\x32\xc5\x31\xed\xe3\xba\xa0\xf7\x4a\xc8\x1a\x0e\xa8\x04\xc6\x15\x10\x58\x70\x1e\xec\xf7\x41\xc8\x05\x78\x54\xad\xa8\xc4\x43\x61\x07\x13\x73\x3e\xf7\x62\x14\xa0\x6c\x01\x73\xca\x7e\x90\x98\x06\xc6\xf1\x48\x5f\xf4\xe8\x2b\x0c\x80\xa5\x4b\x0f\x05\x84\x3c\x65\x41\x76\xcd\xcc\x81\xca\xfc\xb8\xa6\x66\x71\x46\x18\x27\x61\x1a\x8f\xc0\x43\x9f\xa4\xd2\x5c\x84\x82\x73\x05\x3c\xcc\xde\x27\x82\x7b\x31\x2e\x75\xa8\x87\xda\xe4\xb1\xc0\x6f\x65\x34\x10\x81\x40\x40\xa0\x4c\xe3\xfc\x38\x51\x93\xa2\x4f\xe1\xeb\x8a\x08\xa6\x23\xbe\x22\x77\xde\x8f\x84\x6d\x66\xf8\xf7\xcc\xa3\x8a\xb0\xe0\xde\x08\x09\xe0\xcc\xf4\x71\x67\xa6\xc5\x9b\x3d\x14\xb6\x5e\xb1\x90\xd9\x76\x15\x66\xa7\x36\x93\x59\xa1\x9d\x73\xd0\xce\x05\xe8\xb4\x49\xa6\xff\xb7\xa6\x7a\x3e\xd7\x8b\xc9\x36\xc9\xd6\x06\xee\x9c\x05\x77\x8e\xc0\x0d\xb5\xdd\x56\xb4\xe5\xa2\x2a\xed\xf6\x88\xba\xfd\x9c\xb4\xdb\xaa\xb4\xc7\xcc\xce\x39\xe6\x27\x94\xdd\x56\x95\x3d\x4d\xf4\x7c\xa6\x3f\x2f\xec\xd6\x22\xec\x69\xb2\xe7\xb3\x2d\x73\x2f\xde\xf6\x89\x94\x74\xc1\xb4\xa5\xfe\x58\x31\xd5\x2f\x5f\xcc\x33\xe5\x64\xae\xf8\x32\x3d\x98\x4b\x83\xc5\xc1\x5c\xbe\xb9\x1b\xf6\x4b\x72\x75\xa4\x38\x7a\x00\x66\xbe\x4e\x1f\x80\x99\xf1\xe2\xd3\xb6\x1c\x74\x7d\xd8\xd8\x06\x1b\xd7\x83\x8d\x9b\xc2\x86\x36\xd8\xb0\x1e\x6c\xd8\x14\xe6\xda\x60\x6e\x3d\x98\xdb\x14\x76\x63\x83\xdd\xd4\x83\xdd\x34\x85\xf5\x6c\xb0\x5e\x3d\x58\xaf\x29\x6c\x6b\x83\x6d\xeb\xc1\xb6\x4d\x61\x6f\x36\xd8\x5b\x3d\xd8\x5b\x53\xd8\xe3\xa3\x8d\x66\xac\xd7\x71\x8f\x8f\x4d\x79\x4f\x4f\x36\x9e\xb1\x5e\xe7\x3d\x3d\x35\xe5\x4d\x6c\xb8\x49\x3d\xda\xa4\xf9\xd4\x6e\xa3\x19\xeb\x75\x5c\x36\xde\x37\x18\xe0\xc3\x70\xec\x6d\xc6\x9c\x21\x2c\xa9\x54\xe4\xbd\x3a\xb8\x93\xbf\x7c\x9e\x32\xd5\x77\xc8\xe0\xf5\xe4\x26\x5e\x5a\x1a\xdf\xbd\x9e\xde\xb3\x79\x18\xfe\xba\xb9\xb3\xdc\xb1\xf7\x8f\x35\x3d\x63\x12\x2f\xde\x80\xfe\xb2\xcb\x02\x0c\x40\x71\xd0\xc3\x68\xe6\x75\x7c\x97\x8d\xe6\x0c\x28\x0b\x70\x5d\x2e\x07\x9c\xb0\x61\x0c\x19\xbd\xa8\x45\x39\xe4\x6a\xed\xaf\xe2\x25\xfd\x07\x79\x78\x95\x6f\xb6\x9d\x09\xe0\xb0\xf8\x7a\x4d\x0d\x99\xca\x84\xfa\x94\xa7\x7a\x94\x5e\xa4\xe6\x01\xcb\x45\x80\xe2\x44\x0a\xa9\x44\xc2\x65\xff\xd6\xbd\x1d\x81\x23\x07\xc7\x62\xe4\x8b\x8e\x1c\xc1\xad\x7b\x3b\xa8\x3c\x40\xc5\x42\xbe\x68\x9f\x97\xf4\xd8\xb7\x45\x56\x13\x9f\x0b\x81\xbe\x82\x88\x6c\xa4\x22\xfe\x3b\xe8\xef\x28\x0c\x31\x88\xf1\x34\xce\x52\xea\x7a\x06\xd9\x9f\x98\x99\xdd\xc5\x20\xb2\x0f\xf1\xa3\xeb\x47\x44\x74\x77\x23\x70\x3e\xba\xc3\xee\x6e\xf0\x50\xac\xd3\x6b\x1b\xae\x78\x48\x3d\xa9\xc4\x2c\x57\xff\xa3\x2b\x55\x65\xd3\xee\x7c\x3f\x58\xd3\xcf\xff\x4b\x91\x09\xa6\x2b\x20\x53\xef\xbb\xb6\x9f\x2f\x41\x22\x70\x31\xcb\x8f\xf5\x9d\xd9\x08\xb2\x5f\x4b\x30\x3f\x13\x87\x89\x41\x09\xbd\xf1\xb2\x0a\xba\x10\xf9\xa1\xb2\x02\xb6\xc8\xce\x56\xbb\xb2\xb9\xdb\xfd\x5c\x35\x03\x8c\xe9\x92\xea\x8e\xad\x95\x07\xae\x93\x98\x07\x39\xde\x16\xe7\xae\xdd\xfa\x37\x00\x00\xff\xff\x01\xc2\xb1\x40\x53\x13\x00\x00")

func embeddedrulesRulesPhpBytes() ([]byte, error) {
	return bindataRead(
		_embeddedrulesRulesPhp,
		"embeddedrules/rules.php",
	)
}

func embeddedrulesRulesPhp() (*asset, error) {
	bytes, err := embeddedrulesRulesPhpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "embeddedrules/rules.php", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"embeddedrules/rules.php": embeddedrulesRulesPhp,
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
	"embeddedrules": {nil, map[string]*bintree{
		"rules.php": {embeddedrulesRulesPhp, map[string]*bintree{}},
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
