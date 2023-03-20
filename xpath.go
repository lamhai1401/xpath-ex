package main

import (
	"io"
	"sync"

	"github.com/antchfx/jsonquery"
)

// JSONPath working with xpath json input
type JSONPath struct {
	doc   *jsonquery.Node
	mutex sync.RWMutex
}

// NewJSONPath return new JSONPath
func NewJSONPath() *JSONPath {
	return &JSONPath{}
}

func (j *JSONPath) getDoc() *jsonquery.Node {
	j.mutex.RLock()
	defer j.mutex.RUnlock()
	return j.doc
}

func (j *JSONPath) setDoc(doc *jsonquery.Node) {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	j.doc = doc
}

func (j *JSONPath) addNewDoc(r io.Reader) error {
	node, err := jsonquery.Parse(r)
	if err != nil {
		return err
	}
	j.setDoc(node)
	return nil
}

// QueryAll query to all matching value, return error if given xpath input query invalid
func (j *JSONPath) QueryAll(query string) ([]interface{}, error) {
	result := make([]interface{}, 0)
	doc := j.getDoc()

	if doc != nil {
		lst, err := jsonquery.QueryAll(doc, query)
		if err != nil {
			return nil, err
		}
		for _, data := range lst {
			result = append(result, data.Value())
		}
	}

	return result, nil
}
