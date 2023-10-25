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
	// I want to handle this by types
	// the idea comes from this two papers
	// https://blog.janestreet.com/effective-ml-revisited/
	// https://stianlagstad.no/2022/05/parse-dont-validate-python-edition/#:~:text=The%20point%20in%20%E2%80%9CParse%2C%20don,sure%20everything%20is%20in%20order.
	if t.State != AssignedTrip || t.AssignedAt.IsZero() || t.Assignee == "" {
		return errors.New("only assigned trip can be ended!")
	}
	// some thing is going here
	return nil
}
