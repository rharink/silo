package store

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

//Codec provides a mechanism for storing/retriving objects as streams of data.
type Codec interface {
	NewEncoder(io.Writer) Encoder
	NewDecoder(io.Reader) Decoder
}

//Encoder is used to encode objects
type Encoder interface {
	Encode(interface{}) error
}

//Decoder is used to decode objects
type Decoder interface {
	Decode(interface{}) error
}

var (
	_ Codec = XMLCodec{}
	_ Codec = JSONCodec{}
)

//XMLCodec is used to encode/decode XML
type XMLCodec struct{}

//NewEncoder returns a new xml encoder which writes to w
func (c XMLCodec) NewEncoder(w io.Writer) Encoder {
	return xml.NewEncoder(w)
}

//NewDecoder returns a new xml decoder which reads from r
func (c XMLCodec) NewDecoder(r io.Reader) Decoder {
	return xml.NewDecoder(r)
}

//JSONCodec is used to encode/decode JSON
type JSONCodec struct{}

//NewEncoder retuns a new json encoder which writes to w
func (c JSONCodec) NewEncoder(w io.Writer) Encoder {
	return json.NewEncoder(w)
}

//NewDecoder returns a new json decoder which reads from r
func (c JSONCodec) NewDecoder(r io.Reader) Decoder {
	return json.NewDecoder(r)
}
