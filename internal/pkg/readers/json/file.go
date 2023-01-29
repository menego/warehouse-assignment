package file

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"warehouse-assignment/internal/pkg/common/structs"
)

type ReaderBase interface {
	Open(name string) (*os.File, error)
	ReadAll(r io.Reader) ([]byte, error)
}

type Reader struct{}

func (rd *Reader) Open(name string) (*os.File, error) {
	return os.Open(name)
}
func (rd *Reader) ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}

type InventoryList struct {
	Inventory []structs.Article `json:"inventory"`
}

type ProductList struct {
	Products []structs.Product `json:"products"`
}

func readFile(filePath string, reader ReaderBase) ([]byte, error) {
	jsonFile, err := reader.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error while opening file %s: %w", filePath, err)
	}

	defer jsonFile.Close()

	// this could be improved with stream reading to avoid saturating the memory in case of big files
	byteValue, err := reader.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("error while reading file %s: %w", filePath, err)
	}

	return byteValue, nil
}

func readJson(filePath string, reader ReaderBase, structure interface{}) error {
	fileContent, err := readFile(filePath, reader)

	err = json.Unmarshal(fileContent, structure)
	if err != nil {
		return fmt.Errorf("error while unmarshaling to json the file %s: %w", filePath, err)
	}

	return nil
}

func ReadInventory(filePath string, reader ReaderBase) ([]structs.Article, error) {
	i := InventoryList{}
	err := readJson(filePath, reader, &i)

	return i.Inventory, err
}

func ReadProducts(filePath string, reader ReaderBase) ([]structs.Product, error) {
	p := ProductList{}

	err := readJson(filePath, reader, &p)

	return p.Products, err
}
