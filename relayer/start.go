package relayer

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	updateExternalChainsLoopInterval = 1 * time.Minute
	signMessagesLoopInterval         = 1 * time.Second
	relayMessagesLoopInterval        = 1 * time.Second
	attestMessagesLoopInterval       = 1 * time.Second
	checkStakingLoopInterval         = 5 * time.Second
)

func (r *Relayer) checkStaking(ctx context.Context, locker sync.Locker) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	log.Info("checking if validator is staking")

	err := r.isStaking(ctx)
	if err == nil {
		log.Info("validator is staking")
		r.staking = true
	} else {
		log.Info("validator is not staking... waiting")
		r.staking = false
	}
	return nil
}

func (r *Relayer) startProcess(ctx context.Context, locker sync.Locker, tickerInterval time.Duration, requiresStaking bool, process func(context.Context, sync.Locker) error) {
	ticker := time.NewTicker(tickerInterval)
	defer ticker.Stop()

	logger := log.WithFields(log.Fields{})
	for {
		select {
		case <-ctx.Done():
			logger.Warn("exiting due to context being done")
			return
		case <-ticker.C:
			if !requiresStaking || r.staking {
				err := process(ctx, locker)
				if err != nil {
					logger.Error(err)
				}
			}
		}
	}
}

// Start starts the relayer. It's responsible for handling the communication
// with Paloma and other chains.
func (r *Relayer) Start(ctx context.Context) error {
	log.Info("starting pigeon")
	var locker sync.Mutex

	_ = r.checkStaking(ctx, &locker)

	// Start background goroutines to run separately from each other
	go r.startProcess(ctx, &locker, checkStakingLoopInterval, false, r.checkStaking)
	go r.startProcess(ctx, &locker, updateExternalChainsLoopInterval, true, r.UpdateExternalChainInfos)
	go r.startProcess(ctx, &locker, signMessagesLoopInterval, true, r.SignMessages)
	go r.startProcess(ctx, &locker, relayMessagesLoopInterval, true, r.RelayMessages)
	go r.startProcess(ctx, &locker, attestMessagesLoopInterval, true, r.AttestMessages)

	// Start the foreground process
	r.startProcess(ctx, &locker, r.relayerConfig.KeepAliveLoopTimeout, false, r.keepAlive)
	return nil
}
