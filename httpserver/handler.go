package httpserver

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sonymuhamad/nexttalent-test/timeapi"
	"github.com/sonymuhamad/nexttalent-test/usecase"
)

type Handler struct {
	personUsecase *usecase.Person
	timeApi       *timeapi.TimeApi
}

func NewHandler(personUsecase *usecase.Person, timeApi *timeapi.TimeApi) *Handler {
	return &Handler{
		personUsecase: personUsecase,
		timeApi:       timeApi,
	}
}

type CountryResponse struct {
	CountryName string `json:"country_name"`
}

func (h *Handler) GetCountryByPersonName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	personName := chi.URLParam(r, "Name")

	countryName, err := h.personUsecase.GetCountryByPersonName(ctx, personName)
	if err != nil {
		h.writeError(w, err)

		return
	}

	res := CountryResponse{
		CountryName: countryName,
	}

	h.writeSuccess(w, res, http.StatusOK)
}

func (h *Handler) GetCurrentTimeWithTimezone(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	timeZone := r.URL.Query().Get("timeZone")
	if timeZone == "" {
		h.writeError(w, errors.New("timeZone needed"))

		return
	}

	res, err := h.timeApi.GetCurrentTimeByTimezone(ctx, timeZone)
	if err != nil {
		h.writeError(w, err)

		return
	}

	h.writeSuccess(w, res, http.StatusOK)
}
