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

func StringAdr(in string) *string {
	return &in
}

func IntAdr(in int) *int {
	return &in
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

	if time.Now().YearDay() < 213 {
		studYear = studYear - 1
	}

	return strconv.Itoa(studYear), nil
}
