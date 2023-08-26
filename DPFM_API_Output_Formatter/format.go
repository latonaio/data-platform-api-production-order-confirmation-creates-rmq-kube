package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-production-order-confirmation-creates-rmq-kube/DPFM_API_Input_Reader"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(sdc *dpfm_api_input_reader.SDC) (*Header, error) {
	data := sdc.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
