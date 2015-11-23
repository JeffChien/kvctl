package general

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"io"

	"github.com/JeffChien/kvctl/lib"
	"github.com/blang/semver"
	"github.com/docker/libkv/store"
)

type GeneralStorage struct {
	store.Store
	whoami string
}

func New(s store.Store, w string) *GeneralStorage {
	return &GeneralStorage{Store: s, whoami: w}
}

//{{{ implement Command interface
func (m *GeneralStorage) Cat(path string) (*store.KVPair, error) {
	return m.Get(path)
}

func (m *GeneralStorage) Ls(path string) ([]*store.KVPair, error) {
	return m.List(path)
}

func (m *GeneralStorage) Mkdir(path string, opt *lib.MkdirOption) error {
	return lib.ErrNotSupport
}

func (m *GeneralStorage) Rm(path string, recursive bool) error {
	var err error
	if recursive {
		err = m.DeleteTree(path)
	} else {
		err = m.Delete(path)
	}
	return err
}

func (m *GeneralStorage) Touch(path string, data []byte, opts *store.WriteOptions) error {
	return m.Put(path, data, opts)
}

func (m *GeneralStorage) Dump(path string) ([]byte, error) {
	pairs, err := m.List(path)
	if err != nil {
		return nil, err
	}
	kvBuf := new(bytes.Buffer)
	tw := tar.NewWriter(kvBuf)
	for _, v := range pairs {
		hdr := &tar.Header{
			Name: v.Key,
			Mode: 0755,
			Size: int64(len(v.Value)),
		}
		if hdr.Typeflag = tar.TypeReg; v.Dir {
			hdr.Typeflag = tar.TypeDir
		}
		if err = tw.WriteHeader(hdr); err != nil {
			return nil, err
		}
		if _, err := tw.Write([]byte(v.Value)); err != nil {
			return nil, err
		}
	}
	if err = tw.Close(); err != nil {
		return nil, err
	}
	// write with config file
	confData, err := json.Marshal(lib.ArchiveHeader{
		Version: lib.ArchiveVersion,
		Backend: m.whoami,
	})
	if err != nil {
		return nil, err
	}
	files := []struct {
		Name string
		Type byte
		Body []byte
	}{
		{"config.json", tar.TypeReg, confData},
		{"data", tar.TypeReg, kvBuf.Bytes()},
	}
	archiveBuf := new(bytes.Buffer)
	tw = tar.NewWriter(archiveBuf)
	for _, f := range files {
		hdr := &tar.Header{
			Name: f.Name,
			Mode: 0755,
			Size: int64(len(f.Body)),
		}
		hdr.Typeflag = f.Type
		if err = tw.WriteHeader(hdr); err != nil {
			return nil, err
		}
		if _, err = tw.Write(f.Body); err != nil {
			return nil, err
		}
	}
	if err = tw.Close(); err != nil {
		return nil, err
	}
	return archiveBuf.Bytes(), nil
}

func (m *GeneralStorage) Restore(archive []byte) error {
	var data []byte
	//read cofnig file first
	br := bytes.NewReader(archive)
	tr := tar.NewReader(br)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		bw := new(bytes.Buffer)
		if _, err := io.Copy(bw, tr); err != nil {
			return err
		}
		switch hdr.Name {
		case "config.json":
			config := new(lib.ArchiveHeader)
			if err := json.Unmarshal(bw.Bytes(), config); err != nil {
				return err
			}
			if config.Backend != m.whoami {
				return lib.ErrArchiveBackend
			}
			sv, err := semver.Parse(lib.ArchiveVersion)
			if err != nil {
				return err
			}
			tv, err := semver.Parse(config.Version)
			if err != nil {
				return err
			}
			if sv.Major != tv.Major {
				return lib.ErrArchiveVersion
			}
		case "data":
			data = bw.Bytes()
		}
	}
	//read data
	br = bytes.NewReader(data)
	tr = tar.NewReader(br)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		bw := new(bytes.Buffer)
		if _, err := io.Copy(bw, tr); err != nil {
			return err
		}
		if hdr.Typeflag == tar.TypeDir {
			err = m.Put(hdr.Name, nil, &store.WriteOptions{IsDir: true})
		} else {
			err = m.Put(hdr.Name, bw.Bytes(), nil)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

//}}}
