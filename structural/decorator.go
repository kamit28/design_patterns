package structural

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Interface defining basic operations on a data source
type DataSource interface {
	WriteData(data string)
	ReadData() string
}

// Our Simple File Data source implementation
type FileDataSource struct {
	fileName string
}

func NewFileDataSource(name string) *FileDataSource {
	return &FileDataSource{name}
}

// A simplistic implementation to write data in the file
func (f *FileDataSource) WriteData(d string) {
	file, err := os.OpenFile(f.fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(d); err != nil {
		panic(err)
	}
}

// A simplistic implementation to read data from the file
func (f *FileDataSource) ReadData() string {
	data, err := ioutil.ReadFile(f.fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Decorator to encrypt the data before writing to file
// and decrypt data after reading from file.
type EncryptionDecorator struct {
	DataSource
	encryptionKey string
}

func NewEncryptionDecorator(ds DataSource) *EncryptionDecorator {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(b)
	return &EncryptionDecorator{
		ds,
		key,
	}
}

func (f *EncryptionDecorator) WriteData(d string) {
	enc := encrypt(d, f.encryptionKey)
	f.DataSource.WriteData(enc)
}

func (f *EncryptionDecorator) ReadData() string {
	enc := f.DataSource.ReadData()
	return decrypt(enc, f.encryptionKey)
}

// Decorator impl to compress the data before writing to the file
// and to flate the data after reading from file.
type CompressionDecorator struct {
	DataSource
}

func NewCompressionDecorator(ds DataSource) *CompressionDecorator {
	return &CompressionDecorator{ds}
}

func (f *CompressionDecorator) WriteData(d string) {
	c := compress(d)
	f.DataSource.WriteData(c)
}

func (f *CompressionDecorator) ReadData() string {
	c := f.DataSource.ReadData()
	return uncompress(c)
}

// Internal use methods - Not visible to client

// Compression mechanism
// Changes in compression algorithm can break the client.
func compress(source string) string {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(source)); err != nil {
		log.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		log.Fatal(err)
	}
	return b.String()
}

// UnCompress mechanism.
func uncompress(source string) string {
	reader := bytes.NewReader([]byte(source))
	r, _ := gzip.NewReader(reader)
	s, _ := ioutil.ReadAll(r)
	return string(s)
}

// Encryption mechanism.
// Can change - but will break client, if encryption algorithm changes.
func encrypt(source string, key string) (encryptedString string) {
	k, _ := hex.DecodeString(key)
	plaintext := []byte(source)

	block, err := aes.NewCipher(k)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

// Decryption mechanism
func decrypt(source string, key string) (decryptedString string) {
	k, _ := hex.DecodeString(key)
	enc, _ := hex.DecodeString(source)

	block, err := aes.NewCipher(k)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)
}
