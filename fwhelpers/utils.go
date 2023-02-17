package fwhelpers

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"github.com/zclconf/go-cty/cty/json"
)

// Encode passed type into wire format.
// Any struct which contain interface internally need to be
// encoded before it can be sent over the wire.
func Encode(x interface{}) (string, error) {
	b := bytes.Buffer{}
	enc := gob.NewEncoder(&b)
	err := enc.Encode(x)
	if err != nil {
		return "", err
	} else {
		b64enc := base64.StdEncoding.EncodeToString(b.Bytes())
		return b64enc, nil
	}
}

// Decode input encoded string into passed type.
// We accept any interface, as the responsibility of passing
// the correct type is with the caller.
func Decode(s string, x interface{}) error {
	sBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	b := bytes.Buffer{}
	b.Write(sBytes)
	dec := gob.NewDecoder(&b)
	err = dec.Decode(x)
	if err != nil {
		return err
	}
	return nil
}

// Utility function to pack a data model over the wire.
// For all other types, use the Encode function.
// Looks for the model to be passed in either of the two
// input formats.
func PackModel(x *cty.Value, y interface{}) (string, error) {
	var val cty.Value
	var err error

	if x != nil {
		val = *x
	} else {
		ty, err := gocty.ImpliedType(y)
		if err != nil {
			return "", err
		}
		val, err = gocty.ToCtyValue(y, ty)
		if err != nil {
			return "", err
		}
	}
	b, err := json.Marshal(val, cty.DynamicPseudoType)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// Utility function to unpack a data model over the wire.
// For all other types, use the Decode function.
func UnpackModel(s string, x interface{}) error {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	val, err := json.Unmarshal(b, cty.DynamicPseudoType)
	if err != nil {
		return err
	}
	return gocty.FromCtyValue(val, x)
}
