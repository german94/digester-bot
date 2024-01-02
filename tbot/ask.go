package tbot

import (
	"context"
	"digester-bot/user_usage"
	"errors"
)

var ErrRequestsLimitExceeded = errors.New("requests limit exceeded")

func (tb *TBot) Ask(ctx context.Context, s *State, question string) (string, error) {
	if s.User == nil {
		return "", errors.New("no user id")
	}
	userUsage, ok := tb.userRequests.Load(s.User.ID)
	if !ok {
		userUsage = user_usage.New(5)
		tb.userRequests.Store(s.User.ID, userUsage)
	}
	if userUsage.HasReachedLimit() {
		return "", ErrRequestsLimitExceeded
	}
	userUsage.Increment(string(s.Status))
	resp, err := s.Assistant.Ask(ctx, question)
	if err != nil {
		return "", err
	}
	return resp, err
}
