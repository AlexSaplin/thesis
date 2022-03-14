package reporter


import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"rhino/pkg/clients/nalogi"
	"rhino/pkg/entities"
)

const (
	reportsChanSize = 1000
	reportsBufSize  = 1000
)

type Reporter interface {
	Submit(runReport entities.RunReport)
}

type reporter struct {
	client  nalogi.Client
	logger  log.Logger
	reports chan entities.RunReport
	ctx     context.Context
}

func NewReporter(client nalogi.Client, logger log.Logger) Reporter {
	reporter := &reporter{
		client:  client,
		logger:  logger,
		reports: make(chan entities.RunReport, reportsChanSize),
		ctx:     context.Background(),
	}
	go reporter.pushLoop()
	return reporter
}

func (r *reporter) Submit(runReport entities.RunReport) {
	select {
	case r.reports <- runReport:
	default:
		r.logger.Log("msg", "ignoring report, reports buffer full")
	}
}

func (r *reporter) pushLoop() {
	buffer := make([]entities.RunReport, 0, reportsBufSize)
	submit := func() {
		r.logger.Log("msg", "submitting reports", "len", len(buffer))
		err := r.client.SubmitReports(buffer)
		if err != nil {
			r.logger.Log("msg", "failed to submit reports", "err", err)
		}
		r.logger.Log("msg", "submitted reports", "len", len(buffer))

		buffer = buffer[:0]
	}

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.ctx.Done():
			r.logger.Log("msg", "closing push loop")
			if len(buffer) > 0 {
				submit()
			}
			r.logger.Log("msg", "push loop closed")
			return
		case report := <-r.reports:
			buffer = append(buffer, report)
			if len(buffer) == reportsBufSize {
				submit()
			}
		case <-ticker.C:
			if len(buffer) > 0 {
				submit()
			}
		}
	}
}
