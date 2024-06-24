package helper

import "go-saresep/dto"

type ResponseWithData struct {
	Code     int           `json:"code"`
	Success  bool          `json:"success"`
	Message  string        `json:"message"`
	Paginate *dto.Paginate `json:"paginate,omitempty"`
	Data     any           `json:"data"`
}

type ResponseWithNoData struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Response(params dto.ResponseParams) any {
	var response any
	var status bool

	if params.StatusCode >= 200 && params.StatusCode <= 299 {
		status = true
	} else {
		status = false
	}

	if params.Data != nil {
		response = &ResponseWithData{
			Code:     params.StatusCode,
			Success:  status,
			Message:  params.Message,
			Paginate: params.Paginate,
			Data:     params.Data,
		}
	} else {
		response = &ResponseWithNoData{
			Code:    params.StatusCode,
			Success: status,
			Message: params.Message,
		}
	}

	return response
}
