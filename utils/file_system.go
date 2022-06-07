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
