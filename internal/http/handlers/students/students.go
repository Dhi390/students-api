package students

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Dhi390/students-api/internal/types"
	"github.com/Dhi390/students-api/internal/utils/responce"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {

			responce.WriteJSON(w, http.StatusBadRequest, responce.GeneralError(fmt.Errorf("empty body")))
			return

		}

		if err != nil {
			responce.WriteJSON(w, http.StatusBadRequest, responce.GeneralError(err))
			return
		}

		//request validation*--->using go-playground/validator

		if err := validator.New().Struct(student); err != nil {

			validateErrs := err.(validator.ValidationErrors)
			responce.WriteJSON(w, http.StatusBadRequest, responce.ValidationError(validateErrs))
			return
		}

		responce.WriteJSON(w, http.StatusCreated, map[string]string{"success": "ok"})

	}
}
