package server

import (
	"github.com/ptrsd/webdavMock/pkg/repository"
	"net/http"
	"strings"
)

type Server struct {
	Config       Config
	fallbackRepo repository.ReadRepository
}

func Create(config Config, fallbackRepo repository.ReadRepository) Server {
	return Server{
		Config:       config,
		fallbackRepo: fallbackRepo,
	}
}

func (server *Server) Start() error {
	http.HandleFunc("/", server.handleRoot)
	return http.ListenAndServe(":"+server.Config.Port, nil)
}

func (server Server) handleRoot(writer http.ResponseWriter, request *http.Request) {
	var (
		body []byte
		err  error
	)

	switch request.Method {
	case http.MethodGet:
		body, err = server.handleGet(request)
	}

	if err != nil {
		switch err.(type) {
		case repository.NotFoundErr:
			writer.WriteHeader(404)
		default:
			writer.WriteHeader(500)
		}

		return
	}

	writer.Write(body)
}

func (server Server) handleGet(request *http.Request) ([]byte, error) {
	split := strings.Split(request.RequestURI, ".")
	ext := split[len(split)-1]

	return server.fallbackRepo.Find(ext)
}
