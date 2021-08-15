package pipefilter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("Input data should be string")

type SplitFilter struct {
	// 分隔符
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string) // 检查数据格式类型，是否可以处理
	if !ok {
		return nil, SplitFilterWrongFormatError
	}

	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
