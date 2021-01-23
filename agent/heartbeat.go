package agent

import "time"

func (a *Agent) heartbeat() {
	for {
		err := a.doHeartbeat(10 * time.Second)
		if err != nil {
			time.Sleep(30 * time.Second)
		}
	}
}

// send heartbeat with sleep interval duration
// return when fail to send heartbeat
func (a *Agent) doHeartbeat(si time.Duration) error {
	// fake
	time.Sleep(si)
	return nil
}
