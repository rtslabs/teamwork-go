package configuration

import "time"

func (timer *TimerConfig) CalculateDuration() time.Duration {
	if timer.Running {
	 	return timer.Duration + time.Since(timer.Start)
	} else {
		return timer.Duration
	}
}

func GetAllTimers() (timers []*TimerConfig) {

	for _, conf := range Configs {
		for _, timer := range conf.Timers {
			timers = append(timers, &timer)
		}
	}

}
