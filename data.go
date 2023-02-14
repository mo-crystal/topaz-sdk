package topazsdk

import "encoding/json"

func CastToStruct(src map[string]interface{}, dstPtr interface{}) bool {
	j, err := json.Marshal(src)
	if err != nil {
		return false
	}

	err = json.Unmarshal(j, dstPtr)
	return err == nil
}
