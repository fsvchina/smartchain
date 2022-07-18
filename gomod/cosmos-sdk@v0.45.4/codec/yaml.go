package codec

import (
	"encoding/json"

	"github.com/gogo/protobuf/proto"
	"gopkg.in/yaml.v2"
)




func MarshalYAML(cdc JSONCodec, toPrint proto.Message) ([]byte, error) {


	bz, err := cdc.MarshalJSON(toPrint)
	if err != nil {
		return nil, err
	}


	var j interface{}
	err = json.Unmarshal(bz, &j)
	if err != nil {
		return nil, err
	}

	return yaml.Marshal(j)
}
