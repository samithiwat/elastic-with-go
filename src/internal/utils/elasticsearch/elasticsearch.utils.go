package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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
