package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Data        interface{} `json:"data"`
	contentType string
	respWriter  http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		respWriter:  rw,
		contentType: "application/json",
	}
}

func (resp *Response) NoFound() {
	resp.Data = "Resource Not Found"
}

func (resp *Response) Send() {
	resp.respWriter.Header().Set("Content-Type", resp.contentType)
	resp.respWriter.WriteHeader(http.StatusOK)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.respWriter, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

func SendNotFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NoFound()
	response.Send()
}