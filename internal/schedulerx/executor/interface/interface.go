package executor_interface

import "context"

type ExecutorService interface {
	Execute(ctx context.Context) error
}
