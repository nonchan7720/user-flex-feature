package feature

import (
	"context"
	"sync"
)

type updater struct{}

func NewUpdater() Updater {
	return &updater{}
}

func (u *updater) AppendOrUpdateRule(ctx context.Context, updateRetrievers []UpdateRetriever, key string, rule *Rule) error {
	var (
		wg    sync.WaitGroup
		errCh = make(chan error)
	)
	go func() {
		wg.Wait()
		close(errCh)
	}()
	for _, updater := range updateRetrievers {
		if updater.CanUpdate(ctx) {
			wg.Add(1)
			go func(ctx context.Context, updater UpdateRetriever, key string, rule *Rule) {
				defer wg.Done()
				if err := updater.AppendOrUpdateRule(ctx, key, rule); err != nil {
					errCh <- err
				}
			}(ctx, updater, key, rule)
		}
	}
	vErr := <-errCh
	return vErr
}
