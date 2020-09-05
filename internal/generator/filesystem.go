package generator

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func validateTargetFolder(targetFolder string) error {
	if _, err := os.Stat(targetFolder); err != nil {
		return errors.New("targetFolder not found")
	}

	isEmpty, _ := isEmpty(targetFolder)

	if !isEmpty {
		return errors.New("targetFolder is not empty")
	}

	return nil
}

func isEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

func getTargetDirname(srcFolder, path string) string {
	s := strings.Replace(path, srcFolder, "", 1)
	return strings.TrimLeft(s, "/")
}

func mkdir(srcFolder, targetFolder, path string) error {
	targetDirname := getTargetDirname(srcFolder, path)
	if targetDirname == "layout" {
		log.Printf("Skiping layout directory %v", targetDirname)
		return nil
	}

	log.Printf("Creatings target dir %v", targetFolder+string(os.PathSeparator)+targetDirname)
	return os.MkdirAll(targetFolder+string(os.PathSeparator)+targetDirname, 0755)
}

func copyAsset(source, target string) error {

	if filepath.Ext(source) == ".yaml" {
		log.Printf("Skiping yaml file %v", source)
		return nil
	}

	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	newFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer newFile.Close()

	if _, err := io.Copy(newFile, sourceFile); err != nil {
		return err
	}

	return nil
}
func getSourceFilename(path string) string {
	filename := filepath.Base(path)
	extension := filepath.Ext(filename)
	sourceFilename := filename[0 : len(filename)-len(extension)]

	dirname := filepath.Dir(path)
	return dirname + string(os.PathSeparator) + sourceFilename + ".yaml"
}
