// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package merge

import (
	"bytes"
	"errors"

	"github.com/goccy/go-yaml"
)

func YAML(buf ...[]byte) (*bytes.Buffer, error) {
	if len(buf) < 2 {
		return nil, errors.New("2 file")
	}
	var dstMap map[string]interface{}
	err := yaml.Unmarshal(buf[0], &dstMap)
	if err != nil {
		return nil, err
	}
	if dstMap == nil {
		dstMap = make(map[string]interface{})
	}
	for _, src := range buf[1:] {
		var srcMap map[string]interface{}
		err = yaml.Unmarshal(src, &srcMap)
		if err != nil {
			return nil, err
		}
		if err = DeepMerge(dstMap, srcMap); err != nil {
			return nil, err
		}
	}
	bff := new(bytes.Buffer)
	if err := yaml.NewEncoder(bff).Encode(dstMap); err != nil {
		return nil, err
	}
	return bff, nil
}

var (
	ErrKeyWithPrimitiveValueDefinedMoreThanOnce = errors.New("error due to parameter with value of primitive type: only maps and slices/arrays can be merged, which means you cannot have define the same key twice for parameters that are not maps or slices/arrays")
)

func DeepMerge(dst, src map[string]interface{}) error {
	for srcKey, srcValue := range src {
		if srcValueAsMap, ok := srcValue.(map[string]interface{}); ok { // handle maps
			if dstValue, ok := dst[srcKey]; ok {
				if dstValueAsMap, ok := dstValue.(map[string]interface{}); ok {
					err := DeepMerge(dstValueAsMap, srcValueAsMap)
					if err != nil {
						return err
					}
					continue
				}
			} else {
				dst[srcKey] = make(map[string]interface{})
			}
			err := DeepMerge(dst[srcKey].(map[string]interface{}), srcValueAsMap)
			if err != nil {
				return err
			}
		} else if srcValueAsSlice, ok := srcValue.([]interface{}); ok { // handle slices
			if dstValue, ok := dst[srcKey]; ok {
				if dstValueAsSlice, ok := dstValue.([]interface{}); ok {
					// If both src and dst are slices, we'll copy the elements from that src slice over to the dst slice
					dst[srcKey] = append(dstValueAsSlice, srcValueAsSlice...)
					continue
				}
			}
			dst[srcKey] = srcValueAsSlice
		} else { // handle primitives
			if _, ok := dst[srcKey]; ok {
				return ErrKeyWithPrimitiveValueDefinedMoreThanOnce
			}
			dst[srcKey] = srcValue
		}
	}
	return nil
}
