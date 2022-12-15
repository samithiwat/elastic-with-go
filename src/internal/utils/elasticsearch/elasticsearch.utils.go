package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func AppendDocToBuffer(indexName string, docId string, docData interface{}, buf *bytes.Buffer) error {
	meta := []byte(fmt.Sprintf(`{ "index" : { "_index": "%s", "_id" : "%s" } }%s`, indexName, docId, "\n"))
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
