package utils

import "github.com/google/uuid"

func StringAdr(in string) *string {
	return &in
}

func BoolAdr(b bool) *bool {
	return &b
}

func UUIDAdr(in uuid.UUID) *uuid.UUID {
	return &in
}
