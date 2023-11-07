package cache

import "time"

type IWork interface {
	work()
}

type janitor struct {
	interval time.Duration
	stop     chan bool
}

func (j *janitor) Run(c IWork) {
	ticker := time.NewTicker(j.interval)
	for {
		select {
		case <-ticker.C:
			c.work()
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}
