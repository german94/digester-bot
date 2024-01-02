package user_usage

import (
	"sync"
	"time"
)

type Request struct {
	At   time.Time
	Data string
}

type UserUsage struct {
	sync.RWMutex
	requests    []Request
	maxRequests int
}

func New(maxRequests int) *UserUsage {
	return &UserUsage{
		maxRequests: maxRequests,
	}
}

func (uu *UserUsage) Requests() []Request {
	uu.RLock()
	defer uu.RUnlock()
	return append([]Request{}, uu.requests...)
}

func (uu *UserUsage) HasReachedLimit() bool {
	uu.RLock()
	defer uu.RUnlock()
	if len(uu.requests) < uu.maxRequests {
		return false
	}
	for i := 0; i < uu.maxRequests; i++ {
		if !fromToday(uu.requests[len(uu.requests)-1-i].At) {
			return false
		}
	}
	return true
}

func (uu *UserUsage) Increment(data string) {
	uu.Lock()
	defer uu.Unlock()
	uu.requests = append(uu.requests, Request{
		At:   time.Now(),
		Data: data,
	})
}

func fromToday(t time.Time) bool {
	return t.Year() == time.Now().Year() &&
		t.Month() == time.Now().Month() &&
		t.Day() == time.Now().Day()
}
