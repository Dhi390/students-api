package students

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Dhi390/students-api/internal/storage"
	"github.com/Dhi390/students-api/internal/types"
	"github.com/Dhi390/students-api/internal/utils/responce"
	"github.com/go-playground/validator/v10"
)

// creating a new student--->POST /api/students

func New(storage storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating a student")

		// decode request body
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {

			// body empty
			responce.WriteJSON(w, http.StatusBadRequest, responce.GeneralError(fmt.Errorf("empty body")))
			return

		}

		if err != nil {
			responce.WriteJSON(w, http.StatusBadRequest, responce.GeneralError(err))
			return
		}

		//REQUEST VALIDATION*--->using go-playground/validator
		if err := validator.New().Struct(student); err != nil {

			validateErrs := err.(validator.ValidationErrors)
			responce.WriteJSON(w, http.StatusBadRequest, responce.ValidationError(validateErrs))
			return
		}

		// save into DB
		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user created successfully", slog.String("userId", fmt.Sprint(lastId)))

		if err != nil {
			responce.WriteJSON(w, http.StatusInternalServerError, err)
			return
		}

		responce.WriteJSON(w, http.StatusCreated, map[string]int64{"id": lastId})

	}
}

//Get student by Id -->GET /api/students/{id}

func GetById(storage storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		slog.Info("getting a student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responce.WriteJSON(w, http.StatusBadRequest, responce.GeneralError(err))
			return
		}

		student, err := storage.GetStudentById(intId)

		if err != nil {
			slog.Error("error getting user", slog.String("id", "id"))
			responce.WriteJSON(w, http.StatusInternalServerError, responce.GeneralError(err))
			return

		}

		responce.WriteJSON(w, http.StatusOK, student)

	}
}

//Get Student List -->GET /api/students

func GetList(storage storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("getting all students")

		students, err := storage.GetStudents()
		if err != nil {
			responce.WriteJSON(w, http.StatusInternalServerError, err)

			return
		}

		responce.WriteJSON(w, http.StatusOK, students)

	}
}

//updating a student by Id---->PUT /api/students/{id}

func Update(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responce.WriteJSON(w, http.StatusBadRequest, responce.GeneralError(err))
			return
		}

		var student types.Student
		err = json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			responce.WriteJSON(w, http.StatusBadRequest, responce.GeneralError(fmt.Errorf("empty body")))
			return
		}
		if err != nil {
			responce.WriteJSON(w, http.StatusBadRequest, responce.GeneralError(err))
			return
		}

		// Validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			responce.WriteJSON(w, http.StatusBadRequest, responce.ValidationError(validateErrs))
			return
		}

		err = storage.UpdateStudent(intId, student.Name, student.Email, student.Age)
		if err != nil {
			responce.WriteJSON(w, http.StatusInternalServerError, responce.GeneralError(err))
			return
		}

		responce.WriteJSON(w, http.StatusOK, map[string]string{"message": "student updated successfully"})
	}
}

// deleting a student by ID---> DELETE /api/students/{id}

func Delete(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responce.WriteJSON(w, http.StatusBadRequest, responce.GeneralError(err))
			return
		}

		err = storage.DeleteStudent(intId)
		if err != nil {
			responce.WriteJSON(w, http.StatusInternalServerError, responce.GeneralError(err))
			return
		}

		responce.WriteJSON(w, http.StatusOK, map[string]string{"message": "student deleted successfully"})
	}
}
