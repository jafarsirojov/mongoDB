package job

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log"
	"mongoDB/internal/job"
	"mongoDB/pkg/config"
	"sync/atomic"
	"time"
)

var Module = fx.Invoke(InitJobs)

type Params struct {
	fx.In
	Logger      *zap.Logger
	Config      *config.Config
	JobsService job.JobsService
}

func InitJobs(p Params) {
	go Worker(p.Config.ResetRecordsCacheDuration, p.JobsService.ResetRecordsCache)
}

func Worker(d time.Duration, f func()) {
	var reEntranceFlag int64
	for range time.Tick(d) {
		go func() {
			if atomic.CompareAndSwapInt64(&reEntranceFlag, 0, 1) {
				defer atomic.StoreInt64(&reEntranceFlag, 0)
			} else {
				log.Println("Previous worker in process now")
				return
			}
			f()
		}()
	}
}
