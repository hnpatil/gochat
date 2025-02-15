package metadata

import (
	"encoding/json"
	jsonpatch "github.com/evanphx/json-patch/v5"
)

type Metadata map[string]interface{}

func ApplyPatch(source Metadata, patch Metadata) (Metadata, error) {
	if source == nil {
		return patch, nil
	}

	sourceBytes, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	patchBytes, err := json.Marshal(patch)
	if err != nil {
		return nil, err
	}

	responseBytes, err := jsonpatch.MergePatch(sourceBytes, patchBytes)
	if err != nil {
		return nil, err
	}

	var respose Metadata
	err = json.Unmarshal(responseBytes, &respose)
	if err != nil {
		return nil, err
	}

	return respose, nil
}

func (m Metadata) Marshall() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Metadata) UnMarshall(data []byte) error {
	return json.Unmarshal(data, &m)
}
