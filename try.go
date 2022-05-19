package try

import (
	"errors"
	"fmt"
	"time"
)

func Do(f func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case string:
				err = errors.New(e)
			case error:
				err = e
			default:
				err = errors.New(fmt.Sprintf("unknown panic(%s)", e))
			}
		}
	}()

	err = f()

	return
}

func ReDoByCounterWithDelay(f func() error, maxCount int, delay time.Duration) (err error) {
	for i := 0; i < maxCount; i++ {
		err = Do(f)
		if err == nil {
			break
		}
		time.Sleep(delay)
	}

	return
}

func ReDoByCounter(f func() error, maxCount int) (err error) {
	return ReDoByCounterWithDelay(f, maxCount, 0)
}

func ReDoByDurationWithDelay(f func() error, duration, delay time.Duration) (err error) {
	startAt := time.Now()
	for time.Now().Sub(startAt) <= duration {
		err = Do(f)
		if err == nil {
			break
		}
		time.Sleep(delay)
	}

	return
}

func ReDoByDuration(f func() error, duration time.Duration) (err error) {
	return ReDoByDurationWithDelay(f, duration, 0)
}

func ReDoByConditionWithDelay(f func() error, condition *bool, delay time.Duration) (err error) {
	for *condition {
		err = Do(f)
		if err == nil {
			break
		}
		time.Sleep(delay)
	}

	return
}

func ReDoByCondition(f func() error, condition *bool) (err error) {
	return ReDoByConditionWithDelay(f, condition, 0)
}
