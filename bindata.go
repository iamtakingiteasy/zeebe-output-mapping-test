package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _dia_bpmn = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x5d\x73\xe2\x36\x14\x7d\xcf\xaf\x50\xd5\x99\xbe\x2d\xfe\x02\x36\xd0\x98\x9d\x6e\x21\xed\xce\x84\xed\x4e\x93\xee\x74\xfa\xc2\x28\xf6\x05\x94\xd8\x92\x63\xc9\xe0\xf4\xd7\x77\x6c\xcb\x9f\xd8\x60\x36\x6c\x9f\x36\xb2\xef\x95\xce\xbd\xe7\xe8\xfa\xb0\x37\x1f\x62\xdf\x43\x3b\x08\x05\xe5\xcc\xc6\xc6\x40\xc7\x08\x98\xc3\x5d\xca\x36\x36\xfe\xeb\xe1\xf6\xdd\x35\xfe\x30\xbb\xba\x79\x0c\x7c\x36\x75\x61\x4d\x19\x95\x94\x33\x81\x62\xdf\x63\x62\x9a\x3c\xb6\xf1\x56\xca\x60\xaa\x69\xfb\xfd\x7e\xc0\xfd\xcd\x80\x87\x1b\x4d\x04\xe0\x68\x1f\xbf\x2c\x3f\x6b\xa6\x6e\xe8\xfa\xc8\x1c\x6a\xcb\x3f\xe6\x8b\x3b\x5c\x49\x74\x69\xdf\xd4\xf9\xa7\x3c\xcf\x75\xba\x73\xe6\xf3\x4a\xc6\xaf\x79\xc6\xbf\x00\x8f\x50\x24\x39\xc4\x8f\x98\x4b\xb2\x24\x67\x0b\x3e\xd1\xd2\x00\x2d\xad\x3c\xcb\x88\x45\x1d\xd8\xde\x4a\xc3\x4d\x5d\x37\xb4\xbf\x97\x77\xf7\x69\xda\x3b\xca\x84\x24\xcc\x81\x02\xd9\x91\x6a\x6a\xc8\x3e\x61\x44\x5d\x1b\xcf\xcb\x66\xae\x8c\x97\x35\x8f\x36\x4f\x18\x49\x12\x6e\x40\x7e\x26\x3e\x88\x80\x38\x25\xec\xa4\x5d\x03\xca\x73\xc8\xc9\x12\x23\x88\x03\x1e\x4a\x08\x6d\xfc\x4f\x52\x02\x5a\x72\x17\x3c\x08\xcb\x17\x5f\x73\x5a\xf5\xc1\x78\x60\xe2\xd9\x15\x42\x19\x91\x41\xc8\x1d\x10\x22\xc5\x21\x41\xc8\x77\xea\x01\x46\x54\x2c\x62\x70\x22\x49\x1e\x3d\xb0\xb1\x0c\x23\x48\xd3\xf2\x44\x21\x49\x28\x17\x3b\x60\x32\xcd\xbd\x2f\x96\x2b\x43\xc5\xe5\x91\x3c\x92\x1b\x4e\xd9\x66\x76\x0f\x2f\x11\x30\x07\x6e\x3d\xbe\x5f\xe9\xdb\xd1\x13\xf1\xdf\xdf\x68\xf5\x98\xec\x04\xad\x71\x44\xed\x60\x08\x77\xd4\x81\x07\x22\x9e\xb3\x93\xcb\xf5\x4a\xdf\x8f\xd6\xe6\x7e\x8d\x11\x23\x3e\xd8\xf8\x01\x84\x44\x2a\x1e\x39\xc4\xf3\x1a\xc8\x20\x96\xc0\x92\xbe\x2c\x3c\xf0\x81\x49\x91\xbf\x46\xe8\x26\xd5\xc2\x54\x12\xf1\x5c\xd2\x83\xe4\x6b\x00\xaa\x4f\x6a\x5b\x8c\xb4\x83\x24\xca\x97\x24\x08\x8a\x6a\xea\xef\x58\x10\x49\x24\x78\x14\x26\x9c\x52\x37\x27\x3a\xfb\x5b\x6b\xc9\xe0\x91\xac\xa6\x10\xc7\xe1\x11\x93\x95\x3c\xb6\xa3\x12\xc2\x5f\xf2\xe7\x55\x3c\x5a\x27\xa0\x4a\x7d\xbf\x03\x71\x21\x14\x2d\x47\x6f\xd3\x37\xe8\x19\x5e\x6d\xcc\x03\x8c\x76\xc4\x8b\x54\xfd\xad\xe7\xb4\xec\xa6\xa8\xec\xec\x74\x46\x04\x65\x0e\xf7\x4f\x48\xa4\x88\xe9\x21\x2e\x83\x12\x27\x70\x9f\x8e\x8a\xab\x94\x4d\x5d\x5d\xe5\x36\x4a\x5e\x87\x88\xb0\x62\xe3\x4f\x58\x37\x94\xaf\x58\xc9\x5e\xb4\x09\x53\xab\x1e\x06\xb1\xe3\x45\x82\xee\xe0\x37\x22\x61\x4f\x5e\xd3\x03\x17\x8d\x87\x2b\x43\xe8\x93\x98\x3f\xe5\xa2\x5e\xbc\x44\xc4\x13\x48\x72\xf4\xd3\x8f\xd6\xf0\x67\xc3\xb4\xd2\x7f\x3f\x60\xe4\xc2\x9a\x44\x9e\x6c\x62\xbe\x8e\xc0\x58\x33\xdc\xa7\xe3\xf5\xbe\x9d\xd3\x71\x3d\x5e\x73\x97\xe9\xad\x1d\x3f\xc1\x15\x03\xc7\xdd\xf1\x6f\xc8\x54\x95\x1d\x63\xb9\xd9\xe1\xfe\x54\xab\x56\xd4\xa9\x6e\x63\xb4\xc2\x78\x37\x75\x75\xda\x99\x5b\x0e\xce\x85\x5a\xac\x8c\x8d\x03\x54\x9f\xf4\xe2\xa9\xde\xed\x3a\x4f\x79\xe5\x6a\xdf\x33\xc4\x9d\x6d\x9a\xeb\x2c\x9d\xf8\xd5\xea\xbb\xab\xab\xb6\xe0\x78\x39\x0e\x67\x6e\x3a\x4a\x17\x71\x10\x82\x48\x26\x02\x8a\x05\x9d\x66\x83\x35\x0d\x91\xb7\x3c\xf4\x89\x57\x06\xe0\x59\x7d\xc4\x0d\xa8\x8b\x6c\x1b\x61\xc3\xb4\xb0\x2a\xb6\x65\xdb\xc6\x65\x2f\xeb\x3c\xcd\x84\xce\xa5\x0c\xc7\x8f\xfd\x6e\x4c\x4d\xbd\x17\x62\x42\x6d\x9a\x33\xb1\x26\x9e\x78\x0b\x15\xed\xf5\x5c\x8a\x8a\x1f\xbe\x27\x15\x86\xc5\xb7\xcf\x0e\xe9\x77\x29\x6a\xe3\xe0\x52\x97\x42\x4d\x4f\x45\x05\x03\x2a\xb7\x89\xb1\xfa\xf6\x7b\xa1\x2a\xca\x46\x82\xc2\xa5\x4c\x57\x61\xcb\x5c\x3a\x4d\x6c\xef\x9c\x92\x4d\x48\xfc\x14\x55\x65\x5d\x38\xac\x6a\xec\x17\x8f\x30\x28\x22\xd3\x55\xf2\x3d\x4a\x22\xd4\x37\xb7\x61\xf0\x6a\x2d\x55\x9b\xdc\x6f\x49\x90\x6d\xb2\x2a\x96\xab\xca\x07\xce\x6c\x6c\xd8\xea\xfa\x10\xba\x71\x9d\xe9\x47\x1e\x31\x57\xa0\xd8\xc6\xc6\xfb\x09\x46\xaf\x36\x36\x0d\x0b\xa3\x3d\x75\xe5\xd6\xc6\xd6\x18\xa3\x2d\xd0\xcd\x56\x66\x7f\x6b\x75\xbf\x50\x85\x73\x1c\x67\xcb\x48\x5e\xb9\xb4\x09\xb3\x65\x6e\x77\xa1\x35\xc7\xa3\x14\xad\x31\x31\x0a\xb4\x86\xae\x97\x70\xaf\xf5\xb3\xe1\x2e\xdc\x0d\x74\x9a\x89\x36\xb8\x2d\x96\xa3\x8a\x97\x4e\xf7\xe4\x35\xe0\x94\xc9\x14\xb1\x91\x21\x36\x2d\xa3\xee\xca\x9a\x71\xe3\xb6\xb8\x5a\x05\x09\xd0\xe3\xfd\xee\x12\xfb\x61\x15\xdd\xd7\x82\x8a\x25\x09\x9f\x21\xfc\x4a\x05\x6d\xfe\xb8\x38\x24\x64\x98\x97\xa7\x8f\x0b\x42\x46\x15\x3e\x46\x7a\xbd\xea\x0a\xec\x3b\xf2\x08\x5e\xcd\xd7\xd6\x76\xb6\x86\x8a\xea\x71\x29\xcc\x6b\xb3\xdc\xd9\x18\x36\x5c\x6e\xd7\xd6\x6f\x51\x81\xf2\x19\x27\x54\x90\xbb\x91\x6e\x76\xad\x56\x76\x0f\xe3\x86\xad\x6a\x39\x57\x05\x8d\xaf\x7c\x0b\xfb\x5d\x3e\xa0\x49\xc2\x48\xe1\x31\xf4\x8b\x4d\x87\x8e\xeb\x96\xd9\x9b\x53\xd7\x4d\x99\xa0\x23\x0d\x1c\xea\xa5\x1e\x8f\x35\x5a\xc5\x19\xe6\x71\x42\x8a\x06\x34\xe3\xfa\x0b\x79\x38\x1c\x2b\x21\x57\x66\xd6\xe4\xad\x42\x3e\x43\x06\xca\x61\x1c\x91\x41\xd3\x83\x74\xc9\xc0\x32\xbf\xb3\x0c\x94\xb7\x3a\x75\xdf\x94\x03\xeb\x21\x83\x51\x3f\x19\x58\xc3\x7e\x32\x38\x88\x3b\x47\x06\x56\x06\x69\x52\x4e\x4a\x73\xf8\x3f\xca\x40\x79\x9b\x63\xd3\xa0\xe1\xe7\xba\x64\x70\x41\xaf\xd0\x31\x0d\x32\x5f\x77\x6a\x1a\x34\x7e\x3b\xb7\xd0\xdb\x73\xec\x8e\xba\x3e\xd2\x67\xd0\xfb\xde\x3a\xec\xcd\x05\xe9\xad\x3d\x4e\x0d\x64\xe9\x50\xeb\x86\x74\x76\xa5\x7c\x6b\xe5\xbf\x81\x67\x57\xff\x05\x00\x00\xff\xff\xf8\x79\xa9\x64\x41\x16\x00\x00")

func dia_bpmn() ([]byte, error) {
	return bindata_read(
		_dia_bpmn,
		"dia.bpmn",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"dia.bpmn": dia_bpmn,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"dia.bpmn": &_bintree_t{dia_bpmn, map[string]*_bintree_t{}},
}}
