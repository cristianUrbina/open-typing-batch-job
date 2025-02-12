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

func ExtractTarball(r io.Reader, dir string) error {
	gzReader, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("Failed to create gzReader: %w", err)
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Error reading tar: %v", err)
		}

		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("Error creating code directory: %v", err)
		}

		filename := filepath.Join(dir, header.Name)
		dirPath := filepath.Dir(filename)

		if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return fmt.Errorf("Error creating file directory: %v", err)
		}

		if header.Typeflag == tar.TypeDir {
			continue
		}

		outFile, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("Error creating output file: %v", err)
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, tarReader)
		if err != nil {
			return fmt.Errorf("Error copying content to output file: %v", err)
		}
	}

	return nil
}

func saveContentToFile(content io.Reader, filename string) error {
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error saving content to file: %v", err)
		return err
	}
	filename = filepath.Join(cwd, "github", filename+".tar.gz")
	dirPath := filepath.Dir(filename)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Printf("Error creating directories: %v", err)
		return err
	}
	outFile, err := os.Create(filename)
	if err != nil {
		log.Printf("Error creating outfile: %v", err)
		return err
	}
	_, err = io.Copy(outFile, content)
	return nil
}
