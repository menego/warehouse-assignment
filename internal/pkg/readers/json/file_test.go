package file

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"os"
	"testing"
)

type MockReader struct {
	mock.Mock
	ReaderBase
}

func (fwuc *MockReader) Open(name string) (*os.File, error) {
	args := fwuc.Called(name)

	return args.Get(0).(*os.File), args.Error(1)
}

func (fwuc *MockReader) ReadAll(r io.Reader) ([]byte, error) {
	args := fwuc.Called(r)

	return args.Get(0).([]byte), args.Error(1)
}


var readerMock = &MockReader{}
var testError = errors.New("test error")

func TestReadFile(t *testing.T) {

	assert := assert.New(t)

	t.Run("Should return error if file open fails", func(t *testing.T) {
		var nilOsFileHandle *os.File

		readerMock.On("Open", "test_file.txt").
			Return(nilOsFileHandle, testError).Times(1)

		content, err := readFile("test_file.txt", readerMock)

		assert.Nil(content)
		assert.EqualError(err, fmt.Sprintf("error while opening file test_file.txt: %s", testError.Error()))
		readerMock.AssertExpectations(t)
	})

	t.Run("Should return error if reading file content fails", func(t *testing.T) {
		fileHandle := &os.File{}

		readerMock.On("Open", "test_file.txt").
			Return(fileHandle, nil).Times(1)
		readerMock.On("ReadAll", fileHandle).
			Return([]byte{}, testError).Times(1)

		content, err := readFile("test_file.txt", readerMock)

		assert.Empty(content)
		assert.EqualError(err, fmt.Sprintf("error while reading file test_file.txt: %s", testError.Error()))
		readerMock.AssertExpectations(t)
	})

	t.Run("Should return the content and no error if everything goes fine", func(t *testing.T) {
		content, err := readFile("test_json_file.txt", &Reader{})

		assert.Equal(string(content), "{ \"message\": \"hello world!\" }")
		assert.NoError(err)
	})
}

func TestReadJson(t *testing.T) {

	assert := assert.New(t)
	fileHandle := &os.File{}
	type testMsg struct {
		Message string `json:"message"`
	}

	t.Run("Should return error if json unmarshaling fails", func(t *testing.T) {
		m:= testMsg{}

		readerMock.On("Open", "test_file.txt").
			Return(fileHandle, nil).Times(1)
		readerMock.On("ReadAll", fileHandle).
			Return([]byte("broken json }"), nil).Times(1)

		err := readJson("test_file.txt", readerMock, &m)

		assert.Contains(err.Error(), "error while unmarshaling to json the file")
	})

	t.Run("Should return the structure and no error if everything goes fine", func(t *testing.T) {
		m:= testMsg{}

		readerMock.On("Open", "test_file.txt").
			Return(fileHandle, nil).Times(1)
		readerMock.On("ReadAll", fileHandle).
			Return([]byte("{ \"message\": \"hello world!\" }"), nil).Times(1)

		err := readJson("test_file.txt", readerMock, &m)

		assert.Equal(m.Message, "hello world!")
		assert.NoError(err)
	})
}