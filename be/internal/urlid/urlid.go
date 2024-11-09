package urlid

import (
	"github.com/gosimple/slug"
	"strings"
)

func NaiveImpl(input string) string {
	return strings.ReplaceAll(input, " ", "-")
}

func Slug(input string) string {
	return slug.Make(input)
}

type Transformer struct {
	impl func(string) string
}

func NewTransformerWith(impl func(string) string) *Transformer {
	return &Transformer{
		impl: impl,
	}
}

func (transformer *Transformer) Process(input string) string {
	return transformer.impl(input)
}
