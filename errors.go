package gonavitia

import (
	"net/http"
)

type ErrorCode int

const (
	ErrorOk ErrorCode = iota
	ErrorServiceUnavailable
	ErrorInternalError
	ErrorDateOutOfBounds
	ErrorNoOrigin
	ErrorNoDestination
	ErrorNOriginNorDestination
	ErrorUnknownObject
	ErrorUnableToParse
	ErrorBadFilter
	ErrorUnkownApi
	ErrorBadFormat
	ErrorNoSolution
)

func (e ErrorCode) HTTPCode() int {
	switch e {
	case ErrorOk:
		return http.StatusOK
	case ErrorServiceUnavailable:
		return http.StatusServiceUnavailable
	case ErrorInternalError:
		return http.StatusInternalServerError
	case ErrorDateOutOfBounds:
		return http.StatusNotFound
	case ErrorNoOrigin:
		return http.StatusNotFound
	case ErrorNoDestination:
		return http.StatusNotFound
	case ErrorNOriginNorDestination:
		return http.StatusNotFound
	case ErrorUnknownObject:
		return http.StatusNotFound
	case ErrorUnableToParse:
		return http.StatusBadRequest
	case ErrorBadFilter:
		return http.StatusBadRequest
	case ErrorUnkownApi:
		return http.StatusBadRequest
	case ErrorBadFormat:
		return http.StatusBadRequest
	case ErrorNoSolution:
		return http.StatusOK
	default:
		return http.StatusOK
	}
}

type Error struct {
	Id      *string   `json:"id"`
	Message *string   `json:"message"`
	Code    ErrorCode `json:"-"`
}
