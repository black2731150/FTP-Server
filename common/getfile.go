package common

import (
	"fmt"
	"os"
)

func OpenFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Faild to open file: ", err)
		return nil, err
	}
	// defer file.Close()

	return file, nil
}

func GetFileStat(filepath string) (os.FileInfo, error) {
	file, err := OpenFile(filepath)
	if err != nil {
		fmt.Println("OpenFile error :", err)
		return nil, err
	}
	defer file.Close()

	filestat, err := file.Stat()
	if err != nil {
		fmt.Println("Faild to Get fileInfo: ", err)
		return nil, err
	}

	return filestat, nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
