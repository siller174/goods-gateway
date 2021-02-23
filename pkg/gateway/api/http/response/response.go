package response

import (
	"github.com/siller174/goodsGateway/pkg/utils/converter"
	"net/http"

	"github.com/siller174/goodsGateway/pkg/logger"
)

func WriteJSON(w http.ResponseWriter, code int, val interface{}) error {

	js, err := converter.StructToJsonByte(val)
	if err != nil {
		logger.Error("Could not write middleware %v to response", val)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	writeBytes, err := w.Write(js)

	if err != nil {
		logger.Error("Could not write json message in response")
		return err
	}

	if writeBytes == 0 {
		logger.Error("Wrote empty message in response")
	}

	return nil
}
