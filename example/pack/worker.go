package factory

import "context"

const DefaultURL = "http://localhost:8080"

type Worker struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewWorker(ctx context.Context) *Worker {
	ctx, cancel := context.WithCancel(ctx)
	return &Worker{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (w *Worker) Start() error {
	for {
		select {
		case <-w.ctx.Done():
			return nil
		default:
			err := w.fetchAndProcess()
			if err != nil {
				return err
			}

		}
	}
}

func (w *Worker) fetchAndProcess() error {
	return nil
}

func (w *Worker) Stop() {
	w.cancel()
}
