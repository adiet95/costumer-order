package libs

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Code        int         `json:"-"`
	Status      string      `json:"status"`
	IsError     bool        `json:"isError"`
	Data        interface{} `json:"data,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

func (res *Response) Send(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")

	if res.IsError {
		c.Response().Writer.WriteHeader(res.Code)
	}
	err := json.NewEncoder(c.Response().Writer).Encode(res)
	if err != nil {
		c.Response().Writer.Write([]byte("Error When Encode respone"))
	}
	return err
}

func New(data interface{}, code int, isError bool) *Response {
	if isError {
		return &Response{
			Code:        code,
			Status:      getStatus(code),
			IsError:     isError,
			Description: data,
		}
	}
	return &Response{
		Code:    code,
		Status:  getStatus(code),
		IsError: isError,
		Data:    data,
	}
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 202:
		desc = "Accepted"
	case 204:
		desc = "Deleted"
	case 304:
		desc = "Not Modified"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 404:
		desc = "Not Found"
	case 500:
		desc = "Internal Server Error"
	default:
		desc = "Bad Gateway"
	}
	return desc
}
