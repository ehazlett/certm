package utils

import (
	"fmt"
	"os"
)

func CreateIfNotExists(path string, overwrite bool) error {
	f, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	if f != nil {
		if !overwrite {
			return fmt.Errorf("ca cert/key exists.  specify --overwrite to overwrite.")
		}

		if err := os.Remove(path); err != nil {
			return err
		}
	}

	nf, err := os.Create(path)
	if err != nil {
		return err
	}
	nf.Close()

	return nil
}
