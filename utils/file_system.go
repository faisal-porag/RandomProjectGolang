package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func IsExists(path string) bool {
   _, err := os.Stat(path)
	if err == nil {
	  return true, nil
	}
	
	if os.IsNotExist(err) {
		return false, nil
	}
	
	return false, err
}

func IsNotExists(path string) bool {
	return (!IsExists(path))
}


func File_Get_Contents(filename string) ([]byte, error) {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return nil, err
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	contents, _ := ioutil.ReadAll(reader)

	return contents, nil
}

func File_Put_Contents(filename string, content []byte) error {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	defer fp.Close()

	_, err = fp.Write(content)

	return err
}



