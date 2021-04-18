package handler

import (
	"github.com/nqmt/short-url/service"
)

type FiberHandler struct {
	sv service.Service
}

func NewFiberHandler(sv service.Service) *FiberHandler {
	return &FiberHandler{sv: sv}
}
