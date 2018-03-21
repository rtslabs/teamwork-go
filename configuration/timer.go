package configuration

import (
	"time"
	"log"
	"github.com/rtslabs/teamwork-go/util"
)

type TimerConfig struct {
	Name     string
	Running  string
	Start    string
	Duration string
	Favorite string
}

func (timer *TimerConfig) CalculateDuration() time.Duration {
	dur, err := util.StringToDuration(timer.Duration)
	if err != nil {
		log.Fatal("Timer duration is in unexpected format", err)
	}

	if timer.IsRunning() {
	 	dur += time.Since(timer.GetStart())
	}

	return dur
}

func (timer *TimerConfig) IsRunning() bool {
	return timer.Running == "true"
}

func (timer *TimerConfig) SetRunning(val bool) {
	if val {
		timer.Running = "true"
	} else {
		timer.Running = "false"
	}
}

func (timer *TimerConfig) GetStart() time.Time {
	start, err := time.Parse(time.RFC3339, timer.Start)
	if err != nil {
		log.Fatal("Timer start is in unexpected format", err)
	}
	return start
}

func UpdateAllTimers() {

	year, month, day := time.Now().Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())

	for i := range Configs {
		conf := &Configs[i]
		updated := false
		filtered := conf.Timers[:0]
		for _, timer := range conf.Timers {
			if timer.GetStart().After(startOfDay) {
				filtered = append(filtered, timer)
			} else {
				updated = true
			}
		}

		if updated {
			conf.Timers = filtered
			WriteConfig(conf)
		}
	}
}

func GetAllTimers() (timers []*TimerConfig) {

	for _, conf := range Configs {
		for i := range conf.Timers {
			timers = append(timers, &conf.Timers[i])
		}
	}

}
