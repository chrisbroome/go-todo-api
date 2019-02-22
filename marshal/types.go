package marshal

import "io"

type (
	Marshaller interface {
		Marshal(writer io.Writer, value interface{}) error
	}

	Unmarshaller interface {
		Unmarshal(reader io.Reader, result interface{}) error
	}

	HttpMarshaller interface {
		Marshaller
		ContentType() string
	}

	SymmetricMarshal interface {
		Marshaller
		Unmarshaller
	}
)

type withContentType struct {
	marshaller  Marshaller
	contentType string
}

func (this *withContentType) Marshal(writer io.Writer, value interface{}) error {
	return this.marshaller.Marshal(writer, value)
}

func (this *withContentType) ContentType() string {
	return this.contentType
}

func WithContentType(marshaller Marshaller, contentType string) HttpMarshaller {
	return &withContentType{
		marshaller: marshaller,
		contentType: contentType,
	}
}
