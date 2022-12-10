package course

import "context"

type Handler interface {
	InsertData(context.Context)
}
