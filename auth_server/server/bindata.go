// Code generated by go-bindata.
// sources:
// data/google_auth.tmpl
// DO NOT EDIT!

package server

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _dataGoogle_authTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x56\x6d\x6f\xeb\x34\x14\xfe\xdc\xfe\x0a\x2b\x5c\x94\x54\xec\x3a\xe2\x7e\x40\x57\xbd\x6d\xd1\x00\x09\x0d\x21\x36\xb1\xf1\x09\xa1\xca\x73\x4e\x52\x6f\xae\x1d\xec\x93\xb1\x6a\xea\x7f\xe7\xd8\x4e\x5f\xb6\xb5\xd5\x90\x10\x08\x2a\x6d\x4d\xce\xf1\x79\x7b\xfc\x3c\xb5\x27\x0b\x5c\x6a\xa6\x10\x96\x5e\xda\x16\xe2\x13\xae\x5a\x98\x66\x0b\xc4\x76\x5c\x96\x5e\x2e\x60\x29\xb8\x75\x4d\x79\xee\x50\x49\x0d\xd9\x6c\x38\x59\x80\xa8\x66\x43\xc6\x26\x5e\x3a\xd5\x22\xf3\x4e\x4e\xb3\xb2\x14\x77\xe2\x91\x37\xd6\x36\x1a\x44\xab\x3c\x97\x76\x19\x6d\xa5\x56\xb7\xbe\xbc\xfb\xbd\x03\xb7\x2a\xbf\xe4\x1f\xf9\x87\xfe\x85\x2f\x95\xe1\x77\x3e\x9b\x4d\xca\x94\xe9\x55\xd2\xd0\x86\xa7\x3e\x62\xbe\x94\x3a\xa6\xbd\xf3\xa5\xd4\x0a\x0c\x8e\x5b\x2d\xb0\xb6\x6e\x49\x79\xbe\xb6\x46\x5b\x51\x4d\x3d\x0a\x87\x19\x13\x7e\x65\x24\xab\xa0\x06\x77\xa8\xc0\x6c\x38\x18\xd4\x9d\x91\xa8\xac\x61\x34\xa6\xbc\xff\xd1\x36\xca\x14\x23\xf6\x44\x9e\xc1\x83\x70\x4c\x74\xb8\xf8\xc0\xa6\xac\xa1\xf2\x3c\xbe\xf0\x06\xf0\x9c\x1e\x2e\x0c\x15\x31\x12\x8a\xd1\xa7\xb0\x58\xd5\xac\x48\x7e\xe5\xaf\x55\x63\xa0\xba\x30\x61\x69\x31\xea\xb3\x0d\xde\x15\xf9\x67\x0e\x7c\xa7\x31\x1f\x71\x84\x47\x2c\xf2\x07\xa1\x55\x25\x50\x99\x86\xc1\xa3\xf2\xf1\x01\xed\x3d\x18\xce\x79\x9e\xf2\xc6\x2e\x54\x35\x8f\x66\x6a\x24\xd5\x90\x9d\x73\x34\xfa\x2f\x1e\x5c\x2a\xb2\xe9\xea\x67\xf0\xad\x35\x9e\xba\xe2\x9b\xa0\x94\xe6\x1d\x0f\xfb\x50\xa4\x56\x06\x61\x87\xc7\x2c\xbf\xba\xbc\xbe\xc9\xcf\x92\xa9\x73\x9a\x2c\x65\x42\x78\x1e\xca\x6c\x3c\xd2\x1a\xa4\x62\x37\x29\x46\xb4\xad\x56\x52\x04\xcc\x68\x0f\xac\xf9\x44\xc8\x09\xe7\x01\xa7\x1d\xd6\xef\x3f\x6e\x82\x5a\x67\x25\x78\xff\x9d\x40\x31\x66\xb5\xd0\x1e\x7a\x47\x15\x2d\x3f\x5c\x5f\xfe\xc4\x3d\x3a\x9a\x58\xd5\xab\xe2\x29\x17\x71\x17\x72\x2a\x10\x37\x22\x3f\x63\x79\xec\x9e\x2c\x9b\x41\xd6\xa3\x3e\x85\xef\x64\xc8\x4d\x79\xfb\xcd\x2b\x12\xae\x1b\xa4\x0f\x61\xdd\xaf\x48\x60\x0c\xd6\x7d\x26\x70\xce\xba\xbd\x3c\x8f\x0b\x77\x2a\x49\xde\xaf\xcf\xd9\x17\x8c\x96\x72\xd7\xa3\x7d\x43\xce\x17\xa9\xd7\xe9\x7d\xcd\x80\x46\xef\x53\x12\x90\xde\x12\x7d\xb5\x6d\x8a\xdc\x58\x64\xf4\xd0\x40\xc5\x94\xe9\x37\x7b\x3d\x0c\x7f\x8c\x3e\x5b\x5a\x46\x26\x47\x46\xb2\xf8\x89\x44\x0c\x1c\x2f\xf2\x48\x05\x02\x6a\xdb\xfd\xa6\xf5\x3d\xb2\x2a\xa3\xb0\x78\x4a\x42\x99\xab\x8a\x5a\xff\xdc\xe7\x6b\x1a\x67\x01\xa6\xd8\x51\xbe\x2f\x4f\x5f\xa1\x46\x68\x61\x27\x97\x49\x99\xc4\x3e\xb9\xb5\xd5\x8a\x24\x33\xb9\xed\x10\xa9\x33\x55\x4d\x33\x4f\x4c\x57\xe6\x9b\x68\xc8\x66\x81\xf7\x34\x0c\xfb\x43\xe1\x82\x7d\x1f\x89\x34\x29\xd3\xea\x10\xb7\x13\x5d\x40\x76\x3f\x94\xf0\xa5\x16\xe5\x7d\xf1\x72\x94\xb2\x64\x61\xdd\x85\xf9\x56\x68\x7d\x2b\xe4\x7d\x10\xb3\x32\x11\x33\x82\x06\x5a\xf6\x15\xff\xcb\x52\xed\x7d\x4e\x18\xbc\xac\x6b\x4d\xe9\xce\x23\x9f\x88\x85\x0e\x2a\xe5\x40\xe2\xbc\x73\x2a\x70\xb1\xb5\x1e\x97\xe4\x12\x0d\x6c\x51\xdb\xf6\x28\x92\xde\xf6\x88\xb7\xbf\xc1\x7b\xde\xff\x9a\x02\x03\xe4\x73\x22\x25\x69\x50\xda\x0a\xc8\xb4\x1b\xe6\xd7\x64\xfa\xed\x6f\x14\xe3\x33\xd8\xb2\xe4\x1b\x67\x67\xec\x1f\x93\xec\x73\x61\xa6\x00\x1a\xfe\x8d\x0a\x8f\x5f\xf1\xff\x4e\x33\xaf\x44\x72\xd9\xe1\x33\x95\xd8\x0e\x4f\x4b\x63\x1b\x70\x54\x1b\xff\xab\xe3\x89\x84\x7e\x05\x2e\x9c\xe2\x8c\x82\x1e\xc0\xbd\xf7\xaa\x82\x28\xfe\x80\x15\xff\x57\x14\xf4\x36\xa1\x50\x7b\x27\x4f\xab\xa3\x3a\x7c\xad\x1c\xda\x8b\x63\x34\xdf\x73\xd1\xed\xa7\xd5\x80\x70\x4c\x04\xcf\xd8\xbc\x41\xb0\x57\xd3\x29\x5e\x0f\x4e\xb1\x88\x57\x8a\x2e\x88\xc6\xd0\x8f\x63\xb1\x0d\x38\xc0\x19\x1f\xc9\x15\x2a\xe6\x6f\x38\x11\x0f\x24\x38\x7a\x2e\xbe\x54\x58\xa5\x1e\xa2\xbc\x52\x7c\xb8\x41\x92\x25\x1c\x57\xe9\x9c\xa2\x63\x8b\x6e\xb6\xb3\xe1\x9f\x01\x00\x00\xff\xff\xf7\x6e\x04\x9b\xe1\x0a\x00\x00")

func dataGoogle_authTmplBytes() ([]byte, error) {
	return bindataRead(
		_dataGoogle_authTmpl,
		"data/google_auth.tmpl",
	)
}

func dataGoogle_authTmpl() (*asset, error) {
	bytes, err := dataGoogle_authTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/google_auth.tmpl", size: 2785, mode: os.FileMode(436), modTime: time.Unix(1433433554, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"data/google_auth.tmpl": dataGoogle_authTmpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"google_auth.tmpl": &bintree{dataGoogle_authTmpl, map[string]*bintree{
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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

