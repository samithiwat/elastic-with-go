package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

func StringAdr(in string) *string {
	return &in
}

func BoolAdr(b bool) *bool {
	return &b
}

func UUIDAdr(in uuid.UUID) *uuid.UUID {
	return &in
}

func AppendDocToBuffer(docId string, docData interface{}, buf *bytes.Buffer) error {
	meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%s" } }%s`, docId, "\n"))
	data, err := json.Marshal(docData)
	if err != nil {
		return err
	}
	data = append(data, "\n"...)

	buf.Grow(len(meta) + len(data))
	buf.Write(meta)
	buf.Write(data)

	return nil
}
