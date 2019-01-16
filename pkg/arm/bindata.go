// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// data/master-startup.sh
// data/node-startup.sh
package arm

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

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x7c\x6f\x73\xdb\x36\xf2\xf0\xeb\xe3\xa7\xd8\x93\x73\x4f\xda\x4c\x28\xd9\x8e\x93\x26\x6c\x7c\x33\xb2\x2c\x27\x9a\x38\xb6\x4e\x56\x72\xd7\xa7\xd7\xf1\x40\xe4\x4a\x42\x4d\x01\x2c\x00\x2a\x56\xd2\x7c\xf7\xdf\xe0\x0f\x29\x50\xa2\x14\xd9\x49\xaf\x75\x67\x52\x11\xd8\x5d\x2c\x76\x17\xbb\x8b\x05\xc8\xbd\xbf\xb7\x46\x94\xb5\x46\x44\x4e\x21\xc4\xdb\x20\xd8\x83\xe1\xe5\xe9\x65\x04\x2d\x54\x71\x2b\x61\x72\x46\xe4\x6f\xcd\xa4\xc5\x05\x9d\x50\x16\xe6\x99\x54\x02\xc9\x2c\x4c\x98\x6c\xc6\x9c\x8d\x81\x4a\x88\x73\x21\x90\xa9\x74\x01\x53\x22\x92\x98\x27\x98\xfc\x08\x54\x05\x7b\x90\x09\x3e\x22\xa3\x74\x01\x72\xca\xf3\x34\x61\x0f\x15\x8c\x30\x08\xae\xba\x83\xf7\xbd\x4e\xf7\x7a\xf8\x53\xbf\x7b\x6c\x29\x07\x74\x0c\x3f\x43\x38\x86\x86\x19\x58\x2e\xa4\xa6\x4e\x27\x2d\xa2\xf8\x8c\xc6\x21\xcf\x90\xc9\x29\x1d\xab\x90\xf1\x04\x1b\xf0\xcb\x8f\xa0\xa6\xc8\x02\x00\x80\x0a\xb9\x55\xf8\x60\x4c\xf5\xa4\x04\xce\xf8\x1c\x41\xe0\x84\x4a\x25\x16\x10\xa3\x50\x74\x4c\x63\xa2\x10\x24\x1f\xab\x94\xb2\x1b\x18\x0b\x3e\x83\x84\xc7\x37\x28\x82\x9c\x99\x26\x2b\x06\xd3\xd4\xd2\x38\xb2\x99\xb4\x0a\x22\x4d\x12\xc7\x28\x65\x53\x60\x32\x25\xaa\x19\xf3\x59\xcb\xfe\x0c\x63\xd2\x8c\x85\x0a\xf4\xa4\xfe\x0e\x13\x81\x19\xb4\xe6\x44\xb4\x52\x3a\x72\xb4\x2c\xdd\xb1\x54\x64\x54\x4e\x44\x2e\xa4\xc2\x59\xac\x52\x90\x8a\x67\x8e\x8f\xa6\x44\x31\xa7\x31\x06\x00\xb3\x9b\xb1\x6c\xde\x8e\xa5\x96\x52\x2b\xc1\x79\x2b\xa1\xf2\xa6\x45\x3e\xe6\x02\x5b\x02\x25\xcf\x45\x8c\x61\x46\x84\x3a\x08\x00\x30\x9e\x72\x78\xb8\x1d\x0c\xd6\xb8\x02\x4d\x1e\x26\x22\xfb\x2d\xe7\x8a\x00\xec\xc3\xfe\x43\xf8\xe7\x3f\x97\xcc\x6a\x36\x78\xce\xd4\x2a\x66\x00\x20\x50\x2a\x2e\x30\xe6\x0c\xc2\x41\x4d\x7f\x4c\x14\xfc\xd3\x17\x67\x42\x70\xc6\x59\xf3\x57\xc9\x19\xbc\x7c\xf9\xb0\x7b\x79\xf6\x30\xf8\x14\x00\x34\x52\x3e\x09\x13\x41\xe7\x28\x1a\x11\x34\x7e\xe5\xb9\x60\x24\x4d\x1a\xc1\xe7\xa0\x7b\x79\xb6\x22\x28\x22\xd4\xaa\xa4\xac\xc2\xad\x15\xc7\x9c\x49\x9a\xa0\x80\x31\x89\x15\xa8\x29\x51\x6b\xa2\x93\xb1\xa4\x07\xad\x34\x67\xfb\x9b\x6c\x59\x1b\x32\x11\x8a\x2a\xca\xd9\x16\xf4\x1f\xbd\xd1\x72\x95\x0b\x04\xa9\x04\x51\x38\x59\xc0\x98\x0b\x2d\x1f\xfa\x11\x25\xd0\x71\xb0\x07\x0c\x31\xc1\xa4\xce\x3e\x50\xc5\x49\x9d\x75\x94\xda\xdf\xc2\xff\xef\xbf\x83\x12\x39\x6e\x54\xbf\x07\xba\x32\xa0\x55\x7c\x82\x63\x92\xa7\x4a\xee\xa4\x78\x8d\xb7\x59\xed\xa6\x57\x6b\xc2\x70\xd2\x38\xb9\xbc\x1c\x5e\x0d\x07\xed\xfe\x75\xe7\xf2\xe2\xac\xf7\xea\xfa\xa2\xfd\xb6\x7b\xac\x97\x72\x68\xd7\x79\x38\x23\x52\xa1\x68\x14\x83\x2e\x1d\xc0\x83\x4f\xfe\xfa\xfe\x6c\xd6\x7f\x10\x7c\xfa\x14\x02\x1d\xc3\x83\x66\xf7\x56\x09\xd2\x1c\xa2\x54\x1d\x03\xdf\x1c\xe4\x8c\x51\x36\x79\xc7\x12\x14\xba\x19\x3e\x7f\x0e\x24\x26\x10\x52\x08\x11\x1a\x72\xef\xb4\x7b\xf2\xee\xd5\xf5\xf9\xe5\xab\xf3\xee\xfb\xee\xf9\xf1\xe1\x6a\xc3\xd1\x5e\x03\x76\xe1\x41\xb3\x80\x2c\xd1\xf4\x03\xad\xe0\x44\x2a\xa0\x0c\x54\x9c\x3d\x3e\x7c\xf2\x7c\xdf\xfc\x38\x3a\x3a\xfa\x11\x12\x1e\xec\xad\x01\xfc\xf0\xa2\x0a\xf9\xfc\xe8\xe8\x49\xf1\xe3\xc8\xfe\xd8\x7f\xfa\x04\xf2\xc4\xfd\xd0\x2d\x2f\xf6\x5f\xec\x1b\x72\x7f\xcb\x04\x57\xfc\xf8\xc1\xa7\x44\xaa\x7f\xfc\xe3\xf1\xa3\xcf\xc1\xdf\x32\x2e\x94\x6d\xd8\xdb\x7b\xf4\xf8\x73\xf0\x37\x9a\x29\x32\x4a\x51\x42\xd8\x86\xcb\xab\xeb\xb3\xde\xa0\xfb\xef\xf6\xf9\xf9\x75\xfb\xfc\xfc\xf2\xdf\x10\x66\xf0\xc0\x10\x81\x70\xa6\x97\x91\x42\x08\x43\xfb\xff\x8b\xee\xbf\x75\x63\xd1\x1d\x26\x9a\x34\x3c\x30\xff\x86\xbf\x42\xbb\xd3\xe9\xf6\x87\x41\xc2\x19\x06\x41\x31\x48\x28\xc9\x1c\x61\x55\x75\x45\x6f\x10\x88\x19\x84\x62\x6c\xc5\xaa\x4d\xa3\xf5\xc8\xfe\xb6\xce\xbf\x65\x95\xdf\x7a\x14\x04\x23\x22\xf1\xd9\x11\x84\x09\xbc\x7c\xf9\x12\x3e\x7d\x82\x0e\x0a\xd5\x96\x27\x0b\x85\x12\x9a\x4e\xc5\x9d\xa5\xf3\x96\xcd\xae\x8a\x93\x0e\x31\x6d\xf0\x3b\x9c\x18\xfc\x2e\xd3\x2b\x17\x3e\x7f\x76\x2c\x99\x21\x9d\x53\xbe\xdf\x08\x57\x28\xe6\x28\x76\x18\x45\x5a\xc0\xda\x91\xfa\x82\xce\x89\xc2\x37\xb8\xd8\x75\xbc\x37\xb8\xd8\x69\xb8\x1b\x5c\xdc\x73\x62\x7d\xdc\x69\x5a\x19\x7e\x83\x49\x99\xb1\xbe\x38\x25\x33\x94\x9e\xd0\xfa\x58\x3f\x91\x59\xfa\x96\x08\x39\x25\x69\x39\x4a\x3b\x99\x51\xf6\x26\x1f\xa1\x35\xba\x8d\xb4\x9d\xa9\xe9\xa5\x6b\xfe\x69\xde\x94\x38\xf7\x10\xdd\x97\x2c\xce\x1f\xad\xc8\x06\x66\x37\x09\x15\x7a\xe5\xd5\x98\x3e\x23\x33\x4c\xfe\x20\xeb\xaf\x8e\x64\xff\xd7\xd4\xa2\x0e\xef\xbd\x24\xee\x36\xe4\xa6\x61\x76\x34\x9d\x0e\xd9\x6a\x34\x6b\x63\xdd\x6f\x31\x5c\x66\xc8\xae\x74\xc6\xd8\xe1\x4c\xf2\x14\xef\x32\x41\xa3\xbd\x56\xec\x10\xbf\x66\xae\x6b\x5c\xec\x3e\xf3\x2a\x13\xf7\x13\xc2\x99\xe0\x4c\xf5\x05\xbf\x5d\xdc\x4d\xc3\x63\x8d\x17\x66\x1a\xf1\xfe\x46\x75\x65\x13\xb8\x2b\x3a\xd1\x01\xfc\x6e\x0c\xb8\xe4\x2f\x94\x74\xc2\xbe\xd2\x53\xad\xb1\xb1\xbb\x0a\x56\xb8\xb8\xbf\x57\xee\xa4\x14\x99\xba\xf7\xb2\xb6\xd8\x5f\xeb\xae\x1d\x13\xbb\x4f\xbf\x86\x87\xfb\x89\xa0\x3d\x99\x08\x9c\x10\xc5\xc5\xd2\x20\xef\x22\x0c\x52\xe2\x87\x9e\x65\x7e\x95\x40\x6a\x59\xda\x5d\x34\x1b\x38\xba\x9f\x78\xde\x1a\x9a\x3a\xe6\xa5\xa8\xee\x6d\x2a\x37\x16\xff\x5b\x58\x4b\x1d\x43\x77\x36\x9b\x15\x7e\xbe\x46\x34\xd6\x83\xdd\x57\x30\xce\x8d\x7d\x2b\xb1\xf8\xcc\xdc\x59\x28\x15\x5e\xbe\x46\x24\xbb\x24\xb0\xb5\x1c\x7c\x83\x84\xb6\xc2\xc1\x9d\x45\xb0\x2d\xc5\xed\xe7\xa3\x94\xc6\x35\xe3\x3b\x27\xde\x8e\x63\xbd\x6b\x7d\x83\x8b\x66\x09\x7a\x37\x5f\x4e\x2c\x05\xd9\xcc\x0c\xfe\x06\x36\x36\xca\x61\x8d\x8f\xfb\x8e\x6e\x47\xa8\x1f\xfe\x1b\xa4\xc5\x85\x9b\xd2\x58\x5b\x33\xe3\xba\xb1\x96\xcb\xff\x4e\x83\x2d\x6b\x7a\x9e\x07\x70\xe3\xde\x71\x83\xbf\xc6\x66\xc1\xda\x6b\xd5\x27\x52\x7e\x48\x76\x64\x69\xaa\x32\x03\x5e\xd9\xdc\x2f\x2b\x57\x66\x7f\xa2\xff\xb1\xb5\xcf\x97\x2f\xbb\x97\x67\x41\x77\xd8\x39\xbd\x6e\x9f\xbe\xef\x0e\x86\xbd\xab\xee\x75\xe7\xbc\xd7\xbd\x18\x5e\xbf\x1b\x9c\x5f\x1d\x4f\x95\xca\x64\xd4\x6a\x3d\xf8\x6e\xca\xa5\xd2\x09\xda\xf7\x91\xde\xfa\x5b\x9c\x4e\x77\x30\xbc\x3e\xeb\x9d\x77\x8f\x6b\xb7\x8f\x16\xc6\x52\x33\xa0\xed\x77\xc3\xd7\xc7\xa6\xc0\x63\xba\x4e\xdb\xc3\xf6\xf5\x69\x6f\x70\x5c\x2d\xba\x98\xbe\xee\x79\xb7\x33\xec\x5d\x5e\x5c\x0f\x7b\x6f\xbb\x97\xef\x86\xc7\x87\x4f\xf7\xf7\x6d\xd7\xeb\x6e\x7b\x30\x3c\xe9\xb6\x87\xd7\xbd\x8b\x61\x77\xf0\xbe\x7d\x7e\x5c\xf6\xf5\x2e\x7a\xc3\x5e\xfb\xdc\x9b\x4d\xbf\xdb\x1d\x6c\x9b\xcb\xf3\x15\xcc\xce\xf9\xbb\xab\x61\x77\x70\x6c\xe5\x19\xee\x9b\xbf\x12\xb7\xd2\x6a\xb0\x1f\xfb\x4d\x07\xb5\x80\x07\xeb\x80\x87\xb5\x80\x87\x1e\x3f\x6f\xba\x3f\x6d\x10\xad\x5e\x41\x06\xe4\xbc\x77\x35\xec\x5e\xd4\xea\x6b\xbf\x69\xfe\xf3\x74\xe5\x80\xd7\xc5\xb1\x04\x2d\x86\x36\xa5\x2d\x4f\x4a\xb6\xd5\x60\xd6\x69\xbc\xdc\x59\x7b\x60\x9b\x95\x6e\xfa\x6b\x26\x57\x6e\x9a\x97\x50\xc3\x81\x56\xc5\xe9\x75\xa7\xbd\x0a\xec\x32\x74\x03\xfa\xaf\x77\x97\xc3\xf6\xf5\x49\xbb\xf3\xa6\x7b\x71\x7a\x7d\xf2\xd3\xb0\x7b\x75\x7c\x74\xf8\xe2\xe8\xc5\xb3\x1f\x0e\x5f\x3c\xb3\x30\x5f\xa6\x74\x79\x16\x04\x71\x75\x63\xeb\x6d\x7d\x6b\xda\x4d\x40\x2b\xb6\x0a\x5f\xc2\xcc\x6e\x68\x2b\x26\xa1\x12\xb9\x54\x2d\x5b\xc1\x6e\x11\x16\x4f\xb9\x90\x9e\x13\x71\xc4\xf2\x2c\x21\x0a\xc3\x02\xde\x5f\xbe\x75\xe1\xc5\xd5\x1f\x9b\x0b\x32\x4b\xdd\x82\xd6\x4e\x50\x4a\xca\x99\xf5\x21\x51\x00\x90\xa5\xf9\x84\x7a\xcf\x00\x27\x39\x4d\x93\x53\x57\x30\xb5\x4d\x00\x96\x56\x2e\x88\xa2\x9c\x15\x8d\x00\x24\xa3\xef\x51\x68\x8a\x11\xcc\x0f\xca\x66\x64\xf3\x08\x7e\xfe\xa5\x7c\xbe\xa1\x2c\x89\xaa\x84\xed\x88\x25\x44\x51\xbf\x97\x4b\xda\x00\x29\x9d\x51\x25\x23\xf8\xf4\xd9\x6b\x14\xf8\x5b\x8e\xd2\x6b\x36\x64\x2f\xe7\x28\x04\x4d\xf0\x9e\x0c\x7b\x0c\x96\x94\x3c\x0e\x4b\x55\x34\x29\x6f\xf5\x66\x64\x82\x7d\x9e\xd2\x78\x71\x4f\xe9\xdc\x62\x9c\x6b\xc8\x41\x9e\xfa\x13\x0e\x61\x46\x54\x3c\x35\xf4\xdb\x8c\x71\x65\xc8\x55\x24\x12\xc2\x0d\x2e\x22\xa0\x1a\x44\x36\x2b\x6c\x25\xc8\x16\x61\x49\xda\xc3\x01\x98\x93\x34\xc7\x08\x1e\xea\xc5\xf6\xd0\xeb\xd1\x8b\x38\x5a\xb2\x13\x26\xc8\x28\x26\x1e\x00\x67\x83\x3a\xbd\x84\xa5\xba\x22\xc8\x78\x22\x37\x74\x8d\xb4\x38\x65\x45\x75\xbf\x62\xac\xa2\xa2\x94\x5f\xfc\xc9\x1b\x9a\x5d\x9a\x91\x52\xc3\xc7\x19\xa1\x69\x2e\x70\x05\xce\xaa\xc8\x13\xbe\xd3\xcf\x72\xff\xe1\x19\xf5\x32\x31\xed\xb1\x31\xb7\xbc\xc7\x28\xd4\x19\x4d\x31\x82\x2d\x9b\x28\x33\x12\x2e\xb6\xc2\x69\x67\x44\x32\x7a\x8e\x73\x4c\x65\x14\x84\x5a\xb7\x2b\xaa\x26\xb9\x9a\x2e\xd9\x71\x26\xfb\x1a\x49\x82\xc2\x31\x63\x98\xeb\xb4\x23\xa8\x29\x30\x78\x00\x7c\x36\xe3\xec\x82\xcc\x0a\x05\x84\x1b\x98\x0a\xac\x61\x29\x41\xec\x28\x7d\x81\x63\x7a\xbb\xc4\xfa\x4f\x38\xc0\x19\x57\x18\x9a\xa4\x23\x34\xad\x13\xc1\xf3\xcc\x82\xaf\xc3\xbd\xd2\x9d\xa6\x31\x97\x28\xb4\xa5\x6c\x82\x7c\x27\x51\x04\x31\x67\x4a\xf0\x34\x45\x4f\x0b\x98\x62\xbc\x5c\x10\x29\x8f\x6f\x2e\x8c\xc1\xad\x66\x47\xe1\x12\x59\x5b\x8b\x4b\x10\x4d\x7e\xc9\x26\x3a\xdd\xb6\x04\x6c\xe9\xa1\x5c\x72\xa5\x36\x6b\x0a\x24\xce\x62\x0a\x3d\xd6\x14\x2f\xbc\x21\x23\x68\x3c\x6a\x04\x31\x17\xb2\x9d\xa6\xfc\x03\x26\x97\xc6\x99\xca\x28\x48\x98\x5c\xce\x66\x44\x59\xd2\x4e\x12\x81\x52\x46\x50\xc4\xc6\xe7\xfb\x4f\x9f\xb8\xbe\x0b\x54\x1f\xb8\xb8\x89\x40\xc5\xd9\x51\x80\x65\x95\xa1\x30\xc0\x98\x44\x50\x53\xa1\xf4\x67\xb2\xa1\xd2\xe1\xcd\x64\x43\x1d\x02\x20\x17\xa9\xd1\x4c\x08\x1b\x33\x33\x8d\x74\xa5\xb8\x20\x13\x5c\xce\x4a\x27\xa6\x82\xa1\x42\xe9\xba\xac\xe1\x44\x5e\x47\x93\xf2\x3a\xc0\xaa\x67\xe3\x45\x75\x6f\x85\x8c\xef\xa2\x6a\xc0\x7c\x22\xc6\xa9\x2d\x39\x1b\x73\x31\x23\x2a\xf2\x53\xde\x9e\x75\x7b\x67\xa6\x07\x7e\x07\x94\x31\xc9\x74\xc6\x6b\x71\x7d\xb7\xa0\x29\x50\xa6\xb4\xe5\xa6\x03\x77\xaa\xfd\xda\xc9\x23\x72\xe7\xaa\x61\x79\xdc\xed\x8e\x07\x9b\x72\x1e\x47\x4f\xf7\xf7\xf7\x03\xeb\x6c\x6c\xf6\xef\xfc\xcc\x8d\x5f\x03\xf0\x95\xba\x59\x91\x35\x75\x88\x75\x5d\xd6\x14\x07\x00\x32\x2e\x54\x04\x07\xfb\x87\x4f\xf7\x83\xa5\xe4\x7d\x7e\xf4\xe8\x24\xa3\x76\xeb\xd9\x16\x93\x7c\x86\xac\x08\xd7\x71\xca\xf3\xc4\x85\xff\x62\xb9\xfa\x69\x82\xe9\xcf\x04\x9f\xd3\x44\xef\x8c\x3e\xe6\x02\x4d\xda\xef\x21\x17\xbd\xa5\xd7\xd1\x40\x76\x19\x5a\xcd\x85\x23\x12\xdf\x20\x4b\x0a\x00\x6d\x5d\x4f\x2a\x00\x33\x4c\x28\x09\xd5\x22\xc3\x92\x48\x96\xa5\x7a\xe3\x4c\x39\x6b\xcd\x59\xd2\xf4\x6c\xcc\x1c\xc9\x8d\x72\xcd\xc2\x72\x69\xfe\x2f\xa7\x15\xa7\xb9\xf1\x45\xd2\x16\x47\x43\xad\xcc\x70\xac\x15\x55\x33\x52\xb5\xfa\x5f\x87\x7e\x83\x8b\x1d\xb0\xad\xb2\x67\x4e\xad\x39\x53\x11\x3c\x29\x1b\x7a\xfd\x08\x0e\x0e\x7f\x30\x9e\xe6\xe0\xcb\x41\x6d\x53\x89\xa7\xe2\x09\x37\xd5\x5e\x00\x64\x3c\xc5\x24\x2f\xfd\xb7\x05\xaf\xdb\xbe\x17\x70\xe6\x7a\xc3\xd2\x5d\xcb\x0b\x9e\x60\x9f\x0b\x35\x20\x6c\xa2\x93\x8d\x87\x5e\xdf\x55\x3e\x62\xa8\x0d\xfa\x87\xc3\xe6\x13\xe3\x39\x5b\x07\xcf\x74\xbf\x4e\x71\x62\x8d\x69\x43\x9c\xce\x19\x9d\x34\x0c\x6f\x46\xf1\x78\x6b\x97\xf1\x9b\xd2\x5a\x3a\x2e\x36\x32\x66\x03\xcc\x4a\xf2\x47\xe2\x18\x33\xdd\xad\x90\xa9\xe1\x22\xd3\x84\x77\x30\xbd\xc7\x3e\x8c\x9b\x1c\xc0\x28\x17\x52\x45\x70\xb4\xbf\x1f\xb8\x4c\xaf\xa0\xba\x13\x51\x83\xf4\x5b\x26\x23\x38\x34\x14\xd6\x27\xa3\x7f\xb9\x25\x6d\x85\x56\xba\xcc\x73\xce\x33\xbd\xca\xfe\x84\xe9\x3e\xfb\xea\xe9\x3e\x31\x14\xd6\xe6\xe2\xcf\x76\x5b\x95\xc4\xb6\xd8\xca\xd6\xbb\xc1\xb9\x09\x03\x99\xa0\x4c\x41\xa3\x08\x70\x0d\xf8\xae\x79\x8a\x82\xce\x31\x71\x25\xb0\xc2\xc1\x9b\x80\xa1\x08\x65\xb6\x34\x47\x63\xfc\x1e\x7e\x87\xdf\x72\xae\x4c\xb0\x60\x36\x52\x2f\xfd\xa8\x5b\xbc\x2e\x82\xbb\x38\x1a\xd3\x44\x68\x17\xdc\x3c\x38\x7c\x6e\x2d\xf6\xc8\xcc\x4e\x47\x55\x6b\xcf\xe7\xc8\x26\x6a\x1a\xc1\x0b\x4f\xaf\xbd\xbe\xa3\xd2\xe9\x9d\x0e\x1c\x25\x97\x2c\xb4\xb4\x40\xdc\xd8\x7d\xb3\xeb\xb2\xe9\x90\xbd\x59\xe5\x6d\xf7\xf8\x5c\x86\x0e\x2e\x33\x21\x6d\xb9\x92\x3c\xe2\xab\xab\x89\x57\x73\x4e\x22\x25\xaa\x6f\x28\xbf\x46\x71\xfe\xd6\x6a\xf8\xb2\xd4\x59\x24\x61\xca\xdf\x3f\xce\x50\x4d\x79\x12\x01\xc9\x95\x0e\xf9\x34\x41\xa6\xa8\x5a\xf4\x9d\xeb\x75\x42\x49\xf9\x84\x32\x2f\xc5\x9f\x91\x2c\xa3\x6c\xf2\xd6\x21\xc7\x29\xa1\xb3\x60\xb9\x49\xf9\xf4\x09\xbe\xa3\x2c\xc1\xdb\x75\xde\x9a\x7d\xc1\x33\x14\x8a\xa2\x6c\xb6\x73\x35\xed\x0b\xae\x3d\x6f\xb3\xb7\x3a\x30\xec\x7f\xdf\xd4\x22\xaf\xf2\x6f\x1c\xab\x17\x14\x36\x6c\xdb\x0c\x43\xde\x2e\x08\x67\x84\xa6\xfe\xde\xcd\x34\x94\xcf\x34\xf1\xfb\x64\x3e\x0a\x2a\x9b\x2e\xaf\x4f\x3f\x97\x8f\x99\xc0\x31\x0a\x81\xc9\x3b\x97\x74\xfb\x90\x39\xa3\xbf\xe5\x78\xed\x21\x58\x0f\xde\x3b\xfd\x76\xf2\x29\x1e\x9a\x2e\xd0\x9c\xae\x0a\xab\x18\xf4\x0a\x63\x81\xea\x0f\x18\xd8\x12\x5e\x1f\xd6\xa6\x64\x97\x19\xb2\xde\xe9\x2a\x05\x07\x52\x24\xc1\x4e\x8d\xb9\x9a\x72\x41\x3f\x62\x9d\xf1\x1b\xf3\x6b\xce\x68\x2c\xb8\xe4\x63\xc5\x59\x4a\x19\x9a\x8b\x8e\x8d\x6f\x3c\x9f\x21\x32\x62\x04\xd9\x68\x99\x25\x7a\xd8\x2a\x19\x6b\xac\xcf\x12\x40\xf1\x1b\x64\x7f\x35\x96\x0d\x53\x15\x76\xef\x56\xc5\x36\x0e\x75\x4a\xd2\x14\x4d\x72\x50\xae\xfa\x3b\x7a\x81\x73\x1e\x93\x14\x4c\x21\x9b\x8b\x64\xf7\xc5\x3b\xde\x94\xcd\x94\x55\x71\xdf\xc6\x5e\x0f\x6d\x65\xbd\xef\xc6\x59\xb3\x36\xaf\x84\x5e\x66\x6f\x6d\x6f\x23\xf0\x47\x06\xaf\x82\xfa\x46\xba\x5b\x6d\xe0\xec\x5f\xa7\x17\x55\x6a\x12\x57\x0a\x82\x65\xd3\x5b\x72\xdb\x9e\xe0\x95\x8e\xc8\x89\x0e\xe8\x45\x4e\xe0\xba\x6d\xf0\x92\x92\xf9\x8d\x76\xf1\xca\xcd\xd9\xa3\x05\x0b\xa5\x85\x33\xb5\xc9\xc0\xd9\xbc\xcf\x82\xbd\x7f\x3c\xd4\xcd\x2b\x6c\x3c\x7f\x56\xa4\x62\xe5\x32\xaa\x03\x7b\xba\xbf\x1f\x64\x82\xff\x8a\xb1\x17\x9b\xdc\x36\x4f\xe7\x9b\x57\xa6\x38\xc1\x45\x04\xe6\xaa\xa6\x30\x17\x4b\xfc\x2d\x6f\x2b\xe6\xb3\x2c\x57\x78\xec\x6c\x53\x62\x9c\x0b\xaa\x16\xed\x34\xe5\x31\xd1\x98\xd6\x60\x63\x59\xb6\xb8\xcc\x57\xee\x47\xad\xc3\xa2\xf3\x9c\x8c\x30\x95\x7d\x73\x56\x6a\x8b\x5f\x4f\x6d\x51\x85\x26\xab\x78\x07\xfb\xc5\x5f\x78\xf0\xa2\xf8\x6b\x99\xd6\x40\xf0\x5c\x51\x36\x59\x4e\x45\xe6\xa3\x84\xcf\x08\x65\x3b\x7b\xe0\x01\xcf\xed\x91\xad\x5e\x0c\x76\xbd\x1b\xbb\xbb\x2a\x28\xf9\x76\x21\x2b\x87\x7b\xcb\x61\x67\x84\x91\x09\x26\x65\x41\x2a\x2c\x64\x6a\x7e\x9b\x82\x9f\xf1\xc6\xba\x3d\x4b\xf9\xc2\x3c\xd4\xac\x90\xac\x3c\x56\xac\x54\x68\x6a\x0f\x05\x01\xb2\xe2\x80\x53\x03\xbb\x71\xb7\x1c\x62\x4a\x5b\x39\x2a\x36\x49\xb5\x85\x9b\xa3\xa3\xfa\xba\x4d\xcd\x8e\xca\x3b\xb8\xf2\xeb\x76\xe5\x5c\x56\x37\x57\xde\x69\x8c\x9e\xfb\xed\xc0\x15\xac\x7b\xec\x2c\xa5\x93\xa9\xb2\xc6\x59\x56\x05\x87\x74\x86\x3c\x57\xab\xeb\xcc\xdc\x79\xf2\xcf\x9a\x8d\xbd\x85\x1e\x7b\x3b\xdd\xd6\xaa\x16\xc2\x76\xba\x5b\x55\x66\x29\x65\x2c\x0d\x4d\x11\x66\x57\x77\xe5\xdb\xd1\x9c\xa7\xf9\xcc\x2b\xea\x24\x0b\x46\x66\x34\x36\x9e\x54\x3b\x02\xca\x26\x5d\x46\x46\x29\x26\x2e\x08\xd8\xf3\x96\x0d\xe7\x1a\x75\xde\xa3\xbc\x60\xbf\xe2\xf8\xad\x2b\xbf\xaa\xb8\xa5\xc0\x21\x46\x41\x68\xbc\x87\x76\xeb\x76\x3f\x53\x29\x33\x55\x90\x74\xd4\x5c\x3d\x62\xad\x38\x50\x64\xb1\x58\x64\xdb\x89\x74\x59\xbc\x85\xc6\xf6\x39\x57\xf6\xdb\xd5\xd7\x09\x96\x53\x6e\x44\xd0\x98\x1f\x34\x1e\xeb\x56\x3d\x73\xfd\x6c\xcb\x61\xb6\x2d\x13\x98\x58\x33\x6a\x44\xf0\xb3\x51\xeb\x27\xa7\xdc\x86\xd6\xa2\x86\xbf\xe0\xef\x8d\xb6\xfe\x3f\x67\x46\x63\x29\x8d\x55\xc3\x00\x7d\x7e\x5c\x8f\xf1\x96\xdc\x76\x4f\xae\xde\x3b\x1d\xe7\xec\xcb\xe0\xaf\x3a\xdd\xfe\xe9\x5d\x10\xda\x1f\x73\x81\xa7\x54\xde\xdc\x01\x49\xc5\xd3\x1e\xd3\xbe\x8e\x27\xed\xf1\x98\x32\xaa\x16\xdb\x51\x2e\xb8\x1e\x61\xb7\x39\xbf\x42\x86\x82\xa4\xfd\xa5\x40\xb7\x82\xf7\x79\x32\xe4\x29\x0a\x0d\xa9\xe3\xce\x90\x50\xa6\xbe\x80\xd3\x99\x62\x7c\xa3\x81\xdf\xe2\x8c\x8b\x45\x5f\x7b\xae\x5c\xe0\x8e\x48\x7a\x2a\x77\x40\xb1\x72\x3d\xa1\x2c\xa1\x6c\x52\x0f\x4f\x5c\x19\xae\x11\x95\x6d\x00\x8d\x22\x44\x14\x12\xf6\x3b\x01\x1a\xa9\x09\x7c\xa5\xb5\x95\xed\x02\x27\xda\x5e\xbd\xc6\xe5\x31\x61\x91\x04\x3b\x06\x3c\x56\x07\x1e\x96\x86\xfa\xc5\x59\x35\xe5\x82\xea\xe8\xb6\xd9\xaa\x8b\x40\x7f\x95\x09\x24\x49\xdf\x62\xd8\x55\x61\xe0\x3e\xa0\xf6\xc9\x8d\x08\x0e\xb6\xca\x6a\xd5\xa0\xee\x4d\xe8\x1c\x89\x54\x2e\x24\xe0\xfd\xf9\x39\x21\x29\x61\x31\x26\xc5\xf9\x9c\x4b\x26\xb4\x90\xee\x4a\xca\x14\xee\xcc\xbe\xb3\x3d\xe7\x34\xe9\xf3\x44\x6e\x63\xcb\xe4\x23\x5f\xa2\xf7\xd5\x62\x32\xeb\xc4\xad\x1c\xca\xd9\x9d\xe9\x6c\xb7\x59\xa6\xe8\x56\xbb\xd5\x1c\x7c\xe4\x0c\x1b\x3b\x98\xa6\x76\x97\x35\x6c\x1d\x2e\x6d\xd5\xbd\xde\x65\x5f\x18\x7a\x68\x22\xab\x49\x11\xe0\xe0\xd9\xf3\xe6\xb3\x27\xcd\x83\xc3\x17\xcd\x83\x67\x0f\x6b\xee\xd7\x0b\x94\x3c\x9d\xdb\x22\x76\xed\x1d\xfb\x4a\xc9\xbb\x26\x88\x6c\x2a\x89\x2f\xe3\x88\x17\xd8\x3b\x1a\xb8\xd8\xe4\x68\x6f\x58\x1b\xdb\xaf\x94\xa0\x6c\x52\xc6\xad\xe2\x7d\x34\xca\xe6\x28\x15\x9d\x10\x85\xa0\xa6\x08\x61\x38\x23\x8c\x8e\x51\xaa\x30\x17\x29\xb8\x1b\x94\x90\x11\x41\x66\xa8\x50\x00\x61\x09\x48\x44\xbd\x89\xa4\x0a\x66\x5a\x6c\xc1\x1e\x4c\x31\xcd\x20\x97\x40\x14\x90\x34\x5d\x9f\x90\x11\x4b\xc6\x13\x69\x2f\x2b\xf9\x77\x1b\xea\xe2\x7f\x9f\x27\xc1\x0c\x15\x49\x88\x22\xa6\x2a\xb6\x7a\xac\xbe\x0c\xaf\x24\xcd\xa6\x64\x75\x0f\xa0\xbd\x4b\x4c\xd2\x30\xd3\xbb\x51\x53\x9d\xb5\x7e\xcd\x22\x57\xce\xdf\xf5\x7e\x81\x33\x64\x2a\x02\xf7\x02\xd9\x1a\x80\x39\xd7\x08\xb3\x94\x30\xff\x34\xde\x6e\x6d\xdd\x51\xa8\x43\x35\x46\x92\x91\x18\xed\x41\x5c\x68\x5f\x0c\x0c\x64\x86\x71\xe4\x8e\x48\x8c\x5a\x5c\x32\x4c\xc4\xa4\x3c\x97\xfd\xdd\x99\xa2\x44\x05\x21\x71\x0f\x4d\xa8\xb9\xe5\xe5\xfa\xf0\x16\xe3\x25\xcb\x31\x9f\xcd\xc8\xf2\x58\xc7\xbc\x43\x2b\xa7\xee\x29\x8c\xcd\x0f\x73\xda\x56\x77\x32\x67\x0f\xa8\xba\x2a\x4e\x56\x2b\x1b\xf6\x80\x2e\x4f\x53\x77\x71\x02\xda\xe9\x07\xb2\xb0\x97\x03\x52\x3a\x47\x86\x52\xf6\x05\x1f\x95\x05\x2f\xcd\xd5\xb2\x96\x53\xe1\x0a\xca\x03\xa7\x58\xa5\x5e\x4b\x18\xc6\xc4\x1c\xba\x78\x6d\x6b\x17\x7a\x2a\xe0\xc5\x19\x4f\x2d\x42\x79\x7d\xc9\x47\x29\xce\x75\x36\x63\x94\x29\xb4\xc3\x40\x96\x64\x5c\x07\x7d\xaf\x75\xe3\x19\xed\x12\xa4\x38\x50\x9a\x22\x49\xd5\xd4\x75\x68\x8f\x45\x49\x7a\x8a\x29\x59\x94\x1b\x86\xa3\xa7\x5e\x85\xa4\xd4\x63\xb1\x5f\x35\xe7\x02\xb7\xaa\x90\x9c\xde\x58\xd1\x14\x27\x65\xc2\xad\x1b\x6d\x96\xfe\xd6\x6c\xa4\x0a\xc5\x9b\x77\x22\xfb\x44\x4d\xa3\xe5\x04\x1d\x8d\xe5\x48\xee\x30\xce\xb5\xeb\x00\x7b\xc9\xd2\x85\x47\xb9\x4a\x67\xe5\x05\xcb\x15\x5a\x7a\x8d\x9a\x56\xbd\x1d\xa3\x6c\x72\x4a\xc5\x3a\x8e\x96\xd7\x72\xc7\x66\x87\x71\x69\xc0\xa2\x93\x12\x29\x5d\x59\xc2\xac\x97\xd0\xbe\x86\xe9\x16\x71\x50\x4c\xb4\x38\x28\xe7\xd2\x32\x56\x88\xa6\x66\xb2\x75\x53\xdd\x88\xb9\x36\xbd\xd5\xc9\xd5\x27\xfa\x4b\x97\x46\x32\xea\x36\x8f\x5f\xde\xd9\xfc\x39\x9e\x8d\x64\xf4\x9e\x8e\xcd\x62\x7e\x85\x5f\x33\x6f\x44\x17\x46\x65\x68\xba\x87\x82\x27\xb3\xa2\x8d\x92\x8e\x77\xba\x13\x77\xe7\x12\xa6\x1d\x22\xe5\x93\x14\xe7\x98\x1e\x1f\xd9\x3a\x60\x2a\xb1\xb6\xfb\xb0\x5a\x26\x5c\xf3\xac\xcb\xf7\xf8\xb7\xba\xd4\x8e\x15\x6e\x5f\xcb\xf6\xab\x9d\xaa\xf6\x3b\xaf\x50\x2d\xdd\xa8\xb5\x5c\xeb\x63\x3e\x2e\x5b\xcd\xed\x03\x5b\x22\xb1\x7f\xda\x98\xb4\x32\x5f\x0f\x87\xfd\xab\x1d\x9c\x11\x80\x5a\x29\x6b\x1c\xec\x7b\x4b\xa2\x50\x99\xf6\x18\xf4\x8e\x5c\xb6\x34\xd2\xe2\x5b\xf0\xea\x58\xda\xc8\xeb\xb7\x76\xa1\x15\x83\xac\xf8\xbf\x8a\x71\xee\xea\x4d\x37\xa6\x81\x75\x94\x2b\xb7\x1f\xee\xea\xaf\xdd\x28\x75\x84\x4b\xa7\xbd\xce\x5b\x76\x43\xeb\x30\x6c\xf3\xff\xc6\x8f\xaf\x4b\xbc\x5e\xde\xbb\x10\x59\x97\xf1\x76\x09\x7f\x31\x4c\xb8\xcf\x83\x6c\x12\xe9\x56\x9e\x0a\xd9\xae\x49\xf6\x4b\x21\xc6\xbb\xf7\x76\x8f\x20\x63\x87\xab\x5e\xd7\xfb\x23\x5c\x7a\x75\x84\x7b\xb8\xf6\xd2\x1b\x53\xa9\x90\xad\x5d\x69\x3f\x3a\x3a\xfa\x4b\x79\x7f\x64\xf3\xa2\xc3\x8a\xf8\xcd\xbb\x93\xee\xf5\xdb\xf6\x7f\xae\xfb\xa7\xd7\xef\x2f\xcf\x0b\x37\xe6\x2e\xf2\xfa\x9b\xb6\xb7\xe4\xf6\x94\x28\x72\x4a\xe5\x8d\xec\xa3\x78\xff\x76\x7b\x41\xf6\x2f\x1b\x6b\x8e\xb6\xf8\xef\x3a\xb3\xfb\x83\xdd\x73\xad\xb3\xf3\xfc\xf3\x6e\xae\xb8\x9e\xca\xba\x2f\x5e\xa7\xb6\x5c\x36\xc5\x47\x79\x28\xca\x66\x52\x25\x28\xe9\x84\x11\x95\x0b\x0c\xe9\x4c\xcb\xb1\x86\x56\x2e\x8d\xb3\xd1\xbb\xa9\xd6\x32\x03\x6c\x15\x97\x0e\xed\xed\xff\x62\x6b\x62\x11\x05\xcf\xc8\xc4\x55\xc5\x5f\x73\xa9\x86\xbc\x34\xa8\xca\xe8\xeb\x34\xfe\x04\xa7\xde\xfa\xab\x78\xf5\x2d\x4a\x0b\xb6\x6a\x6c\x23\xc5\x3b\xa8\xae\x5e\x21\x26\x16\xd0\x31\xfc\xfc\x33\x34\xbc\x9d\x66\x03\x8e\x8f\xa1\x51\x79\x61\xa9\x01\xbf\x2c\x3f\x3b\xb5\x25\x7a\xc8\x05\x8b\xef\x1d\x36\x34\xf2\xdd\xe3\xc5\x46\x67\x75\xb5\x30\x87\x2a\x3b\x3b\xa9\x0a\x13\xdf\xd6\x7b\x5c\xeb\x89\xb6\xae\x79\xae\xee\x9f\x7b\xd5\xad\x9d\x3b\xae\x88\xbb\xd9\xb2\xb6\x8e\x31\x0d\xca\x2f\x08\x09\xf3\x09\xa1\xef\xe0\x91\xdd\x82\x46\xf0\x7d\xf3\xd1\xde\x7f\x0f\x6a\x4a\x15\xc5\x57\x84\x76\x7f\xe7\xf1\xab\xc7\xa8\xfd\xb6\xc9\x1e\xbc\x6e\x77\xde\xe8\xc0\x90\x2d\x60\xa5\x13\x14\x87\x11\xe7\x4a\x2a\x41\x32\xbf\x5d\x72\xfb\xb1\xac\x92\x5d\xf7\xa9\x2d\x8d\x0f\xb9\x66\x94\x32\x53\xc0\x34\x96\xb2\x07\xf6\x5b\x5e\x12\x15\x7c\xa0\x69\x0a\x8c\x2b\x18\x13\x9a\xba\xf2\xa5\x32\xa0\x76\xc2\x96\x84\xdd\x45\x43\xcc\x85\xc0\x58\xa5\x8b\x66\xed\x3b\x60\xab\xdc\xae\x01\xd4\xf1\xbe\xf6\xea\x68\x4f\x0e\x30\xe6\x73\x14\x0b\xbd\x00\x52\x3e\x99\xa0\x80\x50\x15\x0a\x37\x13\xcb\xb3\xa6\x9c\x42\xc3\xfc\xa6\x6c\x02\xa2\xc0\xe0\x0c\x7c\xa7\x10\xec\x81\x54\x98\xc1\x01\x4c\xdc\xac\x46\x24\xbe\xc9\x33\xf3\x81\xa3\xc1\xb8\x5a\xdd\x68\x3d\x0a\x14\xce\xb2\x13\x12\xdf\x9c\x52\x51\x7d\x4d\xb3\xe5\xf0\x6c\xbd\xfa\x81\x07\x17\x7c\x98\xd2\x14\xe1\xef\xee\xfa\x3d\x64\x79\x9a\xd6\x2c\xec\xae\x8a\x93\x13\x43\x03\x3e\x7f\xfe\x11\x12\x0e\x32\x45\xcd\x99\xf9\x42\x14\xc3\xc0\xa1\x8b\x9c\x41\x18\x6a\xf6\x8a\xdb\x92\x66\x8d\xc0\x7f\xf5\x72\x99\xd7\x98\x68\xb4\xde\x54\x07\x5c\x8d\xe0\x91\xb7\xba\x0b\x68\x7f\x52\x51\x8b\xe7\x2a\xfa\x68\xba\xbe\x30\x17\x8b\x3e\x4a\xf9\x48\x8b\xfc\x58\x43\x5b\x3d\x5a\x88\x93\x94\x8f\xcc\xf5\xc4\x02\x32\x41\xa9\x28\x33\x91\xf8\x58\x8f\xe2\x24\xdb\x4c\x46\xd0\x48\xf8\x07\x96\x72\x92\x34\xb6\xab\xdd\x62\x40\x01\x8d\x89\x5e\x17\x0d\x78\xf0\x5d\x2a\x2b\xb3\xf8\x3e\x28\xf4\x7f\x58\x7c\xf7\x6c\x3b\x61\x0b\xa4\x0d\x4a\x32\x92\xc9\x29\x57\x8d\x1d\xd5\xeb\xd5\x88\xbf\x5a\xbd\xda\xde\xa2\xf2\x57\xd9\xe5\x9b\x63\x54\x7d\xb2\x8a\x0a\x11\xba\xc3\xce\x69\x67\x78\x7e\xdd\xee\xf7\x8e\x1b\x4f\x1a\x1b\xf4\x57\x61\xd6\xc0\xb8\xb2\x73\x39\xed\x42\x5c\x15\x71\x7a\x9a\x32\xc3\x99\x3d\x5d\xa8\xd7\x43\x75\xa9\x30\xfc\xe0\x00\xcc\x55\x0a\x6f\x3d\xba\x66\x57\xb7\x08\x5d\x35\x78\x25\x6c\xff\x4f\xdf\x3e\x6e\xd4\xb3\x14\x9a\xbb\x5b\xb6\xc8\x39\xe6\x22\x34\x67\x4c\x2b\xa0\x24\x99\xa3\x50\x54\x62\x98\x21\x8a\x30\x17\xa9\xdc\x50\x02\x7f\xbe\x1f\x04\xb3\xf9\xba\x94\x5a\x8f\x56\xda\xca\xef\xae\xad\x42\x16\x1d\x15\xc7\x53\xf9\x94\xdf\x0a\xf9\x5d\xec\x1c\x8d\x65\x36\x2a\xef\xce\xef\xe9\x58\x80\x11\xd4\x7d\x41\x0f\x62\x41\xe4\x14\x52\xce\x33\x09\x39\x53\x34\x2d\x62\x04\x95\x90\x67\xc1\xf2\x3b\x8f\x68\x6e\xbf\xd4\x12\x29\x3f\xfb\xb8\xfa\x55\xc8\x6d\xc0\xf0\xff\x34\x67\x09\x95\x64\x94\x1a\x77\x2f\x17\x32\xe5\x13\x90\x94\xc5\x08\x1f\xd0\xdd\xe0\x02\xd4\x31\x40\x4d\x35\x88\x9a\x0a\x9e\x4f\xa6\x50\x7c\x8e\xd2\x1b\xcf\xd2\xc1\x82\x4a\x2d\x47\x3c\x5b\xeb\xf6\xce\x2a\x05\xe7\xaa\x65\x62\xd8\x6a\x1c\xdc\xf0\x5d\x06\x1f\xa5\xe5\xe2\xde\xff\x05\x00\x00\xff\xff\x2c\xd6\x52\xab\xcb\x55\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\x6d\x4f\xe4\x36\x17\xfd\xee\x5f\x71\x77\x66\xf5\xcc\xd3\x95\x12\xc3\x76\x85\x5a\x16\x90\x60\x18\xd0\x0a\x96\x41\xc0\xae\xb4\xaa\x2a\xe4\xd8\x37\x89\x77\x12\x3b\x5c\x3b\x53\xd2\xe9\xfc\xf7\xca\xc9\xbc\xf0\x32\x65\xa9\xfa\x2d\xb2\x8f\xcf\x3d\x3e\xb9\xf7\xb8\xff\x86\x27\xda\xf0\x44\xb8\x1c\x22\xbc\x67\xac\x0f\x37\xe3\xe3\xf1\x2e\x70\xf4\x92\x2b\xe3\x4a\xe1\xee\x62\xc5\x2d\xe9\x4c\x9b\xa8\xae\x9c\x27\x14\x65\xa4\x8c\x8b\xa5\x35\x29\x68\x07\xb2\x26\x42\xe3\x8b\x06\x72\x41\x4a\x5a\x85\xea\x23\x68\xcf\xfa\x50\x91\x4d\x44\x52\x34\xe0\x72\x5b\x17\xca\x0c\x3c\x24\xc8\xd8\xf5\xe8\xea\xeb\xa7\xe1\xe8\xf6\xe6\xdb\xe5\x68\xbf\x63\x66\x3a\x85\xdf\x20\x4a\xa1\xd7\x16\x76\x8d\x0b\xec\x3a\xe3\xc2\xdb\x52\xcb\xc8\x56\x68\x5c\xae\x53\x1f\x19\xab\xb0\x07\xbf\x7f\x04\x9f\xa3\x61\x00\x00\x8f\xe8\x9e\xe2\x59\xaa\xc3\xa5\x08\x4b\x3b\x45\x20\xcc\xb4\xf3\xd4\x80\x44\xf2\x3a\xd5\x52\x78\x04\x67\x53\x5f\x68\x33\x81\x94\x6c\x09\xca\xca\x09\x12\xab\x4d\xbb\xd4\xd9\xd0\x2e\xf1\x70\xc6\xc5\x8a\x2f\x49\x62\x21\x25\x3a\x17\x13\xaa\x5c\xf8\x58\xda\x92\x77\x9f\x91\x14\xb1\x24\x0f\x2c\xdc\xea\x0d\x64\x84\x15\xf0\xa9\x20\x5e\xe8\x64\x41\xd6\x11\xa7\xce\x8b\x64\x75\x13\xd7\x38\x8f\xa5\xf4\x05\x38\x6f\xab\x85\x90\xd8\x21\x4d\xb5\x44\x06\x50\x4e\x52\x17\xdf\xa7\x2e\xd8\xc4\x15\x4e\xb9\xd2\x6e\xc2\xc5\x9f\x35\x21\x27\x74\xb6\x26\x89\x51\x25\xc8\x6f\x33\x00\x94\xb9\x85\xc1\xcb\x30\x78\xa6\x0a\x02\x3d\x64\x54\xdd\xd5\xd6\x0b\x80\x2d\xd8\x1a\xc0\xc1\xc1\x5a\x6c\x90\x61\x6b\xe3\x9f\x9e\x64\x00\x84\xce\x5b\x42\x69\x0d\x44\x57\xcf\xf6\x67\xb3\x08\x74\x0a\x78\x07\xf1\xe8\xde\x93\x88\xaf\x6c\x81\xd0\xd3\x26\x25\xd1\x83\xf9\x9c\x01\x48\xe1\xe1\xe0\xa1\xe1\x4a\x60\x69\x4d\xfc\xdd\x59\x03\x7b\x7b\x83\xd1\xf8\x64\xc0\x66\x0c\xa0\x57\xd8\x2c\x52\xa4\xa7\x48\xbd\x5d\xe8\x7d\xb7\x35\x19\x51\xa8\x1e\x9b\xb3\xd1\xf8\xa4\x2d\x85\x46\x75\xa4\x0f\x4d\x15\xe4\x9f\xba\x1a\xba\xa3\xb3\xea\x68\x3c\xbe\xb9\xbe\xb9\x3a\xbc\xbc\x1d\x8e\x2f\x4e\x3e\x9d\xde\x5e\x1c\x7e\x1e\xed\x87\x6e\x8b\xba\x56\x8c\x66\xb3\x47\xda\xe7\xf3\x95\x35\xeb\x76\x7d\x3b\x7b\xd8\x8d\xf3\xb6\x5b\xd9\xf2\xf2\x6f\x17\xc7\x6f\xd0\xf9\x61\x8b\x8f\xaf\x6a\x63\xb4\xc9\xbe\x18\x85\x14\x96\x83\x68\x87\x0a\x22\x0d\x11\x42\xcf\xf5\x8f\x47\x47\x5f\x4e\x6f\xcf\xc7\xa7\xe7\xa3\xaf\xa3\xf3\xfd\xf7\x4f\x17\x3e\xf4\x7b\xf0\x1a\x0d\x0f\x4c\x61\x54\x42\x44\x69\x77\x0c\xbd\x54\xfc\x5d\xf7\xdd\x8d\x22\x2f\x85\xf3\x48\xfc\x1d\x63\x89\x70\xb8\xf3\x01\x22\x05\x7b\x7b\x7b\x30\x9b\xc1\x51\xbb\x30\x32\x61\xc8\xe1\xff\xdf\x44\x59\x7c\x16\xe4\x72\x51\x40\xbc\xb8\xd0\x85\x55\x78\x64\xad\x77\x9e\x44\x75\x56\x27\xd8\x89\xfa\x09\xe6\xf3\xc5\xcf\x5d\x54\x09\xaa\x78\xb2\x44\xc6\x93\x15\xf4\x47\x55\x87\x48\xfe\xd0\x1d\x35\x1e\xdd\xaa\xea\x70\x3d\xce\xee\xb1\x84\x76\xeb\x1f\xaa\xb7\x3f\x77\x25\xa1\x42\x0a\x53\xfb\xa3\xf2\x97\xa4\xa7\xc2\xe3\x19\x36\xff\x42\xc4\x19\x36\xaf\xd6\x30\xc1\x86\xc9\xbc\xb4\x0a\xb6\x76\xb6\xb6\xe0\x75\x27\x9e\xc3\x36\x5a\xfb\x9f\xbd\x1d\x8a\x97\x0c\xed\x72\x8f\xc9\xea\xb9\x9c\x45\x24\xb6\xeb\xd5\x44\x73\x29\x22\x4f\xb5\xf3\xbc\x0b\x24\x2e\x8c\xcc\x2d\x39\xbe\x0e\xfa\x05\x59\x5d\x29\xe1\x31\x5a\xe2\x97\xe3\x6a\x44\x89\x61\x88\x91\x60\x7b\xe7\x97\x78\xe7\xe7\x78\xfb\xfd\xaf\xf1\xf6\xce\x60\x83\xac\x90\x7a\xc5\xb4\x7d\xaf\x58\x39\x51\x9a\x20\x7a\xac\x50\x16\xb6\x56\x15\xd9\xa9\x56\x48\x8c\xad\x83\x68\xd3\x7e\x17\xa5\xdd\xeb\xb7\x4a\xa5\x19\xc4\xc7\x18\x12\x49\xc5\xc3\x00\xbe\x5c\x80\x83\x87\xad\x93\x5e\x68\x83\x74\xdd\xc5\x0e\xfc\x05\xd7\x9e\xb4\xc9\xc2\x38\x86\xc8\x62\x7d\x30\xd6\xe3\x2e\x6c\x9a\x5d\x90\x14\x9e\xe7\xc2\xda\xca\x41\x6d\xbc\x2e\xa0\x1b\xd1\xf0\xf8\xd6\x15\x5b\x47\x1c\x1a\x91\x14\xb8\x91\x64\x95\x78\x4f\x03\xf1\x25\x30\xfc\x2f\x28\x53\xda\x89\xa4\x08\x6a\xc9\x35\xae\xb0\x19\x38\x6d\x24\xc2\x1f\x08\xa5\x30\x22\x43\xc0\x29\x52\xe3\xf3\x00\xf1\x39\xd9\x3a\xcb\x61\x19\xca\x0f\xea\x75\x3c\xb8\x64\xd9\xa8\xc8\x56\xcf\xb6\xff\x0e\x00\x00\xff\xff\x7f\x5b\x3f\x7f\xa3\x08\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
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
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
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
