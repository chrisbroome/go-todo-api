package marshal

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type JSONMarshaller struct {
}

func (*JSONMarshaller) ContentType() string {
	return "application/json"
}

func (*JSONMarshaller) Unmarshal(reader io.Reader, result interface{}) error {
	return unmarshalJSON(reader, result)
}

func (*JSONMarshaller) Marshal(writer io.Writer, result interface{}) error {
	return marshalJSON(writer, result)
}

func unmarshalJSON(reader io.Reader, result interface{}) error {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, result)
}

func marshalJSON(writer io.Writer, result interface{}) error {
	bytes, err := json.Marshal(result)
	if err != nil {
		return err
	}

	_, err = writer.Write(bytes)
	return err
}
