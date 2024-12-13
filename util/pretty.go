package util

import (
	"encoding/json"
	"os"
)

func PrettyJSON(i interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(i)
}
