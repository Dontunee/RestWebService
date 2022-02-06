package controllers

import (
	"encoding/json"
	"github.com/dontunee/webservice/models"
	"net/http"
	"regexp"
	"strconv"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (controller userController) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/users" {
		switch request.Method {
		case http.MethodGet:
			controller.getAll(responseWriter, request)
		case http.MethodPost:
			controller.post(responseWriter, request)
		default:
			responseWriter.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := controller.userIDPattern.FindStringSubmatch(request.URL.Path)
		if len(matches) == 0 {
			responseWriter.WriteHeader(http.StatusNotFound)
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			responseWriter.WriteHeader(http.StatusNotFound)
		}

		switch request.Method {
		case http.MethodGet:
			controller.get(id, responseWriter)
		case http.MethodPut:
			controller.put(id, responseWriter, request)
		case http.MethodDelete:
			controller.delete(id, responseWriter)
		default:
			responseWriter.WriteHeader(http.StatusNotImplemented)
		}
	}

}

func (controller *userController) getAll(responseWriter http.ResponseWriter, request *http.Request) {
	encodeResponseAsJson(models.GetUsers(), responseWriter)
}

func (controller userController) get(id int, responseWriter http.ResponseWriter) {
	user, err := models.GetUserById(id)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJson(user, responseWriter)
}

func (controller *userController) post(responseWriter http.ResponseWriter, request *http.Request) {
	receivedUser, err := controller.parseRequest(request)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("could not parse user object"))
		return
	}
	addedUser, err := models.AddUser(receivedUser)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJson(addedUser, responseWriter)
}

func (controller *userController) put(id int, responseWriter http.ResponseWriter, request *http.Request) {
	receivedUser, err := controller.parseRequest(request)
	receivedUser.ID = id
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("could not parse user object"))
		return
	}
	updatedUser, err := models.UpdateUser(receivedUser)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJson(updatedUser, responseWriter)
}

func (controller *userController) delete(id int, responseWriter http.ResponseWriter) {
	err := models.RemoveUserById(id)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(err.Error()))
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}

func (controller *userController) parseRequest(request *http.Request) (models.User, error) {
	decoder := json.NewDecoder(request.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
