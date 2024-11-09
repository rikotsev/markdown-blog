package common

import (
	"fmt"
	"github.com/rikotsev/markdown-blog/be/gen/api"
)

func Problem(title string, description string, errorId string, code int) *api.Problem {
	return &api.Problem{
		Title:       api.NewOptString(title),
		Description: api.NewOptString(fmt.Sprintf("%s. Please provide this error id: %s", description, errorId)),
		Code:        api.NewOptInt(code),
	}
}
