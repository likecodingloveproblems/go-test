package main

import (
	"errors"
	"time"
)

type User string
type TripState uint

const (
	RequestedTrip TripState = iota
	AssignedTrip
	EndedTrip
	CanceledTrip
)

type ITrip interface {
	Request()
	Assign()
	End()
}

type Trip struct {
	State       TripState
	RequestedAt time.Time
	Assignee    User
	AssignedAt  time.Time
	EndedAt     time.Time
	CanceledAt  time.Time
}

func (t *Trip) End() error {
	if t.State != AssignedTrip || t.AssignedAt.IsZero() || t.Assignee == "" {
		return errors.New("only assigned trip can be ended!")
	}
	// some thing is going here
	return nil
}
