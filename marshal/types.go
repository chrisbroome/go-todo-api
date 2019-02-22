package marshal

import "io"

type (
	Marshaller interface {
		Marshal(writer io.Writer, value interface{}) error
	}

	Unmarshaller interface {
		Unmarshal(reader io.Reader, result interface{}) error
	}

	SymmetricMarshal interface {
		Marshaller
		Unmarshaller
	}
)
