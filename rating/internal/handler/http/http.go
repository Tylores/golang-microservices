package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"golang-microservices/rating/internal/controller/rating"
	"golang-microservices/rating/pkg/model"
)

type Handler struct {
	ctrl *rating.Controller
}

func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) Handle(w http.ResponseWriter, req *http.Request) {
	recordID := model.RecordID(req.FormValue("id"))
	if recordID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recordType := model.RecordType(req.FormValue("type"))
	if recordID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch req.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedRating(req.Context(), recordID, recordType)
		if err != nil && errors.Is(err, rating.ErrNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}
	case http.MethodPut:
		userID := model.UserID(req.FormValue("userId"))
		v, err := strconv.ParseFloat(req.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		model_ref := model.Rating{UserID: userID, Value: model.RatingValue(v)}
		if err := h.ctrl.PutRating(req.Context(), recordID, recordType, &model_ref); err != nil {
			log.Printf("Repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
