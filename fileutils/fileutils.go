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
	fmt.Println(filename)
	outFile, err := os.Create(filename)
	if err != nil {
		log.Printf("Error creating outfile: %v", err)
		return err
	}
	_, err = io.Copy(outFile, content)
	fmt.Println("Tarball downloaded successfully")
	return nil
}

func ExtractTarball(r io.Reader) {
	gzReader, err := gzip.NewReader(r)
	if err != nil {
		log.Printf("Error creating gzReader: %v", err)
		return
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading tar: %v", err)
			return
		}

		codeDir := "./batch_job_pulled_repos"
		if err = os.MkdirAll(codeDir, os.ModePerm); err != nil {
			log.Fatalf("Error creating code directory: %v", err)
			return
		}

		filename := filepath.Join(codeDir, header.Name)
		log.Println("header name:", header.Name)
		log.Println("Extracting:", filename)
		dirPath := filepath.Dir(filename)
		log.Println("dir:", dirPath)

		fmt.Println("creating dir", dirPath)
		if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
			log.Fatalf("Error creating file directory: %v", err)
			return
		}
		fmt.Println("created dir", dirPath)

		if header.Typeflag == tar.TypeDir {
			continue
		}

		outFile, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Error creating output file: %v", err)
			return
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, tarReader)
		if err != nil {
			log.Fatalf("Error copying content to output file: %v", err)
			return
		}
	}
}
