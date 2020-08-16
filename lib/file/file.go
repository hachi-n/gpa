package file

import (
	"os"
)

func Read(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return nil, err
	}

	b := make([]byte, info.Size())
	if _, err := f.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func Create(name string, content []byte) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		return err
	}
	return nil
}

func Delete(filePath string) error {
	return os.Remove(filePath)
}
