package utils

import (
	"github.com/pkg/errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func BoolAdr(b bool) *bool {
	return &b
}

func UUIDAdr(in uuid.UUID) *uuid.UUID {
	return &in
}

func GetCurrentTimePtr() *time.Time {
	tmp := time.Now()
	return &tmp
}

func GetCurrentYear2Digit() int {
	return (time.Now().Year() + 543) % 100
}

func CalYearFromID(sid string) (string, error) {
	if len(sid) != 10 {
		return "", errors.New("Invalid student id")
	}

	yearIn, err := strconv.Atoi(sid[:2])
	if err != nil {
		return "", errors.New("Invalid student id")
	}

	studYear := GetCurrentYear2Digit() - yearIn + 1
	if studYear <= 0 {
		return "", errors.New("Invalid student ID")
	}

	return strconv.Itoa(studYear), nil
}
