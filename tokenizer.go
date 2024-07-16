package tokenizer

import (
//	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
//	"os"

	"github.com/sugarme/tokenizer"
	"github.com/sugarme/tokenizer/pretrained"
	"go.k6.io/k6/js/modules"
)

func init() {
    modules.Register("k6/x/tokenizer", new(TokenModule))
}

var tkz *tokenizer.Tokenizer

// LoadConfig loads tokenizer configuration from URL or local path
func LoadConfig(path string) error {
	var err error

	tkz, err = pretrained.FromFile(path)
	if err != nil {
		return fmt.Errorf("failed to load config")
	}

	return nil
}

func Tokenize(text string) ([]string, error) {
	if tkz == nil {
		return nil, fmt.Errorf("tokenizer is not loaded")
	}

	encoding, err := tkz.EncodeSingle(text, true)
	if err != nil {
		return nil, err
	}

	return encoding.Tokens, nil
}

func isValidURL(url string) bool {
	_, err := http.Get(url)
	return err == nil
}

func downloadFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

type TokenModule struct{}

func (m *TokenModule) Load(pathOrURL string) error {
	return LoadConfig(pathOrURL)
}

func (m *TokenModule) Tokenize(text string) ([]string, error) {
	return Tokenize(text)
}
