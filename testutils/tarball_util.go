package testutils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
)

func CreateSampleTarGZ() (*bytes.Buffer, error) {
	contents := map[string]string{
		"file1.txt": "Hello, world!",
		"file2.txt": "Go is awesome!",
	}
	return CreateTarGz(contents)
}

func CreateEmptyTarGZ() (*bytes.Buffer, error) {
	contents := map[string]string{}
	return CreateTarGz(contents)
}

func CreateTarGz(contents map[string]string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	gzWriter := gzip.NewWriter(buf)
	tarWriter := tar.NewWriter(gzWriter)

	for name, content := range contents {
		header := &tar.Header{
			Name: name,
			Size: int64(len(content)),
			Mode: 0o644,
		}
		if err := tarWriter.WriteHeader(header); err != nil {
			return nil, err
		}
		if _, err := tarWriter.Write([]byte(content)); err != nil {
			return nil, err
		}
	}

	tarWriter.Close()
	gzWriter.Close()

	return buf, nil
}
