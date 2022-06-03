package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status        int         `json:"status"`
	Data          interface{} `json:"data"`
	Message       string      `json:"message"`
	contentType   string
	responseWrite http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:        http.StatusOK,
		responseWrite: rw,
		contentType:   "application/json",
	}
}

func (resp *Response) Send() {
	resp.responseWrite.Header().Set("Content-Type", resp.contentType)
	resp.responseWrite.WriteHeader(resp.Status)

	// Marshall devuelve 2 valores: Los valores transformados en tipo byte y un error
	output, _ := json.Marshal(&resp) // Para responder con json
	//output, _ := xml.Marshal(&resp) // Para responder con xml
	//output, _ := yaml.Marshal(&resp) // Para responder con yaml
	fmt.Fprintln(resp.responseWrite, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

func (resp *Response) NoFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource No Found"
}

func SendNoFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NoFound()
	response.Send()
}

func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity No Found"
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity()
	response.Send()
}
