package category

import (
	"github.com/rikotsev/markdown-blog/be/gen"
)

type Mapper struct {
}

func (m *Mapper) ToPersistenceLayer(input gen.Category) Entity {
	return Entity{
		Id:    input.Id,
		Name:  input.Name,
		UrlId: input.UrlId,
	}
}

func (m *Mapper) ToHttpLayer(input Entity) gen.Category {
	return gen.Category{
		EntityType: gen.CategoryTypeCategory,
		Id:         input.Id,
		Name:       input.Name,
		UrlId:      input.UrlId,
	}
}

func NewMapper() *Mapper {
	return &Mapper{}
}
