package utils

import (
	"io"
	"os"
)

func CopyFile(path string) error {
	src_file, err := os.Open(path + "data.ydb")
	if err != nil {
		return err
	}
	defer src_file.Close()

	dst_file, err := os.Create(path + "backup/backup.ydb")
	if err != nil {
		return err
	}

	defer dst_file.Close()
	_, err = io.Copy(dst_file, src_file)
	if err != nil {
		return err
	}
	return nil
}
