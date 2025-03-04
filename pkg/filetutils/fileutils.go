package fileutils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ExtractTarball(r io.Reader, dir string) ([]string, error) {
	gzReader, err := gzip.NewReader(r)
	if err != nil {
		return nil, fmt.Errorf("Failed to create gzReader: %w", err)
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)
	files := []string{}
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Error reading tar: %v", err)
		}

		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("Error creating code directory: %v", err)
		}

		filename := filepath.Join(dir, header.Name)
		dirPath := filepath.Dir(filename)

		if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return nil, fmt.Errorf("Error creating file directory: %v", err)
		}

		if header.Typeflag == tar.TypeDir {
			continue
		}

		outFile, err := os.Create(filename)
		if err != nil {
			return nil, fmt.Errorf("Error creating output file: %v", err)
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, tarReader)
		if err != nil {
			return nil, fmt.Errorf("Error copying content to output file: %v", err)
		} else {
			files = append(files, filename)
		}
	}

	return files, nil
}

func SaveContentToFile(content io.Reader, dir string) error {
	dirPath := filepath.Dir(dir)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Printf("Error creating directories: %v", err)
		return err
	}
	outFile, err := os.Create(dir)
	if err != nil {
		log.Printf("Error creating outfile: %v", err)
		return err
	}
	_, err = io.Copy(outFile, content)
	outFile.Close()
	return nil
}
