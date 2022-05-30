package util

import "os"

func PathExists(paths ...string) (bool, error) {
	for _, path := range paths {
		_, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}
			return false, err
		}
	}
	return true, nil
}
