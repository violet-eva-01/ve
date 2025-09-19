// Package utils @author: Violet-Eva @date  : 2025/9/19 @notes :
package utils

import "os"

func WriteToTextFile(fileName string, path string, data []byte) error {

	create, err := os.Create(path + "/" + fileName)
	if err != nil {
		return err
	}
	defer create.Close()
	err = os.WriteFile(fileName, data, 0777)
	if err != nil {
		return err
	}

	return nil
}
