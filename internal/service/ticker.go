package service

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

func (s *Service) DeleteExpiredShortenerLinks(ctx context.Context, interval time.Duration) error {
	errCh := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				slog.Info("Stop deleting expired links")
				errCh <- ctx.Err()
			case <-ticker.C:
				if err := s.DeleteExpiredShortenerLink(ctx); err != nil {
					slog.Error(err.Error())
					errCh <- err
				} else {
					slog.Info("Successfully deleted expired shortener links")
				}
			}
		}
	}()
	wg.Wait()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
