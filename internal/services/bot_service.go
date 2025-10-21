package services

import (
	"context"
	"fmt"
	"sync"
	"talus_helper_windows/internal/workflowy"
	"time"
)

// BotService fetches content from web and updates Workflowy using API
type BotService struct {
	ctx             context.Context
	workflowyClient workflowy.WorkflowyClient
	interval        time.Duration
	cancelFunc      context.CancelFunc
	wg              sync.WaitGroup
	running         bool
	mu              sync.Mutex
}

// NewBotService initializes a new BotService.
// interval defines how often the cron job runs (e.g., 30m for every 30 minutes)
func NewBotService(ctx context.Context, wfClient workflowy.WorkflowyClient, interval time.Duration) *BotService {
	cctx, cancel := context.WithCancel(ctx)
	return &BotService{
		ctx:             cctx,
		workflowyClient: wfClient,
		interval:        interval,
		cancelFunc:      cancel,
	}
}

// Start begins the bot service, performing the sync immediately and then on a cron schedule.
func (b *BotService) Start() {
	b.mu.Lock()
	if b.running {
		b.mu.Unlock()
		return
	}
	b.running = true
	b.mu.Unlock()

	b.wg.Add(1)
	go func() {
		defer b.wg.Done()
		// Run immediately on start
		b.runOnce()

		ticker := time.NewTicker(b.interval)
		defer ticker.Stop()

		for {
			select {
			case <-b.ctx.Done():
				return
			case <-ticker.C:
				b.runOnce()
			}
		}
	}()
}

// runOnce fetches content from the web and updates Workflowy
func (b *BotService) runOnce() {
	// Example: Create a new node in Workflowy
	createReq := &workflowy.CreateNodeRequest{
		ParentID:   "parent123",
		Name:       "Bot Task",
		Note:       "Created by bot service",
		LayoutMode: "list",
	}

	_, err := b.workflowyClient.CreateNode(createReq)
	if err != nil {
		fmt.Printf("BotService: failed to create node: %v\n", err)
		return
	}

	fmt.Println("BotService: successfully created node in Workflowy")
}

// Stop gracefully stops the bot service and waits for goroutines to complete
func (b *BotService) Stop() {
	b.mu.Lock()
	if !b.running {
		b.mu.Unlock()
		return
	}
	b.running = false
	b.mu.Unlock()

	b.cancelFunc()
	b.wg.Wait()
}
