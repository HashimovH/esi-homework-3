package http

import (
	// "encoding/json"
	// "github.com/HashimovH/esi-homework-3/pkg/domain"
	"github.com/gorilla/mux"
	// log "github.com/sirupsen/logrus"
	// "net/http"
)

type InvoiceService interface {
	// GetStatus(i string, s string, e string) (int, error)
}

type InvoiceHandler struct {
	InvoiceService InvoiceService
}

func NewInvoiceHandler(pS InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{
		InvoiceService: pS,
	}
}

func (h *InvoiceHandler) RegisterRoutes(router *mux.Router) {
	// router.HandleFunc("/remittance/:id", h.SubmitRemittance).Methods(http.MethodPost)
}