package try_test

import (
	"errors"
	"github.com/aKorsi/try"
	"testing"
	"time"
)

func TestTryDo(t *testing.T) {
	SomeFunction := func() (string, error) {
		return "", nil
	}
	var value string
	err := try.Do(func() error {
		var err error
		value, err = SomeFunction()
		return err
	})
	_ = value
	if err != nil {
		t.Log("error:", err)
	}
}

func TestTryDoError(t *testing.T) {
	SomeFunction := func() (string, error) {
		return "", errors.New("error")
	}
	var value string
	err := try.Do(func() error {
		var err error
		value, err = SomeFunction()
		return err
	})
	_ = value
	if err != nil {
		t.Log("catch error:", err)
	}
}

func TestTryDoPanic(t *testing.T) {
	SomeFunction := func() (string, error) {
		panic("panic")
	}
	var value string
	err := try.Do(func() error {
		var err error
		value, err = SomeFunction()
		return err
	})
	_ = value
	if err != nil {
		t.Log("catch panic:", err)
	}
}

func TestTryReDoByCounterPanic(t *testing.T) {
	SomeFunction := func() (string, error) {
		panic("panic")
	}
	counter := 0
	var value string
	err := try.ReDoByCounter(func() error {
		counter++
		t.Logf("counter => %d", counter)
		var err error
		value, err = SomeFunction()
		return err
	}, 5)
	_ = value
	if err != nil {
		t.Log("catch panic:", err)
	}
}

func TestTryReDoByCounterWithDelayPanic(t *testing.T) {
	SomeFunction := func() (string, error) {
		panic("panic")
	}
	counter := 0
	var value string
	err := try.ReDoByCounterWithDelay(func() error {
		counter++
		t.Logf("counter => %d", counter)
		var err error
		value, err = SomeFunction()
		return err
	}, 5, time.Second)
	_ = value
	if err != nil {
		t.Log("catch panic:", err)
	}
}

func TestTryReDoByDurationPanic(t *testing.T) {
	SomeFunction := func() (string, error) {
		panic("panic")
	}
	counter := 0
	var value string
	err := try.ReDoByDuration(func() error {
		counter++
		t.Logf("counter => %d", counter)
		var err error
		value, err = SomeFunction()
		return err
	}, time.Microsecond*100)
	_ = value
	if err != nil {
		t.Log("catch panic:", err)
	}
}

func TestTryReDoByDurationWithDelayPanic(t *testing.T) {
	SomeFunction := func() (string, error) {
		panic("panic")
	}
	counter := 0
	var value string
	err := try.ReDoByDurationWithDelay(func() error {
		counter++
		t.Logf("counter => %d", counter)
		var err error
		value, err = SomeFunction()
		return err
	}, time.Second*5, time.Second)
	_ = value
	if err != nil {
		t.Log("catch panic:", err)
	}
}

func TestTryReDoByConditionPanic(t *testing.T) {
	SomeFunction := func() (string, error) {
		panic("panic")
	}
	condition := true
	go func() {
		time.Sleep(time.Microsecond * 100)
		condition = false
	}()
	counter := 0
	var value string
	err := try.ReDoByCondition(func() error {
		counter++
		t.Logf("counter => %d", counter)
		var err error
		value, err = SomeFunction()
		return err
	}, &condition)
	_ = value
	if err != nil {
		t.Log("catch panic:", err)
	}
}

func TestTryReDoByConditionWithDelayPanic(t *testing.T) {
	SomeFunction := func() (string, error) {
		panic("panic")
	}
	counter := 0
	condition := true
	go func() {
		time.Sleep(time.Second * 5)
		condition = false
	}()
	var value string
	err := try.ReDoByConditionWithDelay(func() error {
		counter++
		t.Logf("counter => %d", counter)
		var err error
		value, err = SomeFunction()
		return err
	}, &condition, time.Second)
	_ = value
	if err != nil {
		t.Log("catch panic:", err)
	}
}
