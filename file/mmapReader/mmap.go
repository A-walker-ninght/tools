package mmapreader

import (
	"errors"
	"io"
	"os"
	"tools/file/mmap"

	"github.com/pkg/errors"
)

// mmap内存映射
type MMapFile struct {
	buf []byte
	fd  *os.File
}

func NewMMapFile(fileName string, fileSize int64) (*MMapFile, error) {
	if fileSize <= 0 {
		return nil, errors.Wrapf(err, "unable to open: %s", filename)
	}
	file, err := openFile(fileName, fileSize)
	if err != nil {
		return nil, err
	}
	buf, err := mmap.Mmap(file, true, fileSize)
	if err != nil {
		return nil, err
	}
	return &MMapFile{buf: buf, fd: file}, nil
}

func (m *MMapFile) Close() error {
	if m.fd == nil {
		return nil
	}
	if err := mmap.Msync(); err != nil {
		return nil
	}
	if err := mmap.Munmap(); err != nil {
		return nil
	}
	return m.fd.Close()
}

func (m *MMapFile) Delete() error {
	if m.fd == nil {
		return nil
	}
	if err := mmap.Munmap(m.buf); err != nil {
		return err
	}
	m.buf = nil
	if err := m.fd.Truncate(0); err != nil {
		return err
	}
	if err := m.fd.Close(); err != nil {
		return err
	}
	return os.Remove(m.fd.Name())
}

func (m *MMapFile) Sync() error {
	if m == nil {
		return nil
	}
	return mmap.Msync(m.buf)
}
