package common

import (
	"crypto/md5"
	"fmt"
	"io"
)

func HashMD5(filepath string) string {
	fh, err := OpenFile(filepath)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}

	hash := md5.New()
	if _, err := io.Copy(hash, fh); err != nil {
		fmt.Println("error copy error: ", err)
	}

	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum)
}
