package api

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
)

const (
	defaultHttpServerPort    = 8880
	defaultHttpServerAddress = "0.0.0.0"
)

type Server struct {
	Debug bool
}

func (s *Server) Start() {
	logrus.Infof("Start http server, listening on %s:%d", defaultHttpServerAddress, defaultHttpServerPort)
	listen := fmt.Sprintf("%s:%d", defaultHttpServerAddress, defaultHttpServerPort)

	container := restful.NewContainer()
	container.Add(makeWebService(s))
	err := http.ListenAndServe(listen, container)
	logrus.Error("Server listen and server, err: %v", err)
}

func makeWebService(s *Server) *restful.WebService {
	ws := new(restful.WebService)

	ws.
		Path("/").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/health").To(probeHealth).
		Doc("probe health").
		Returns(http.StatusOK, "OK", ""))

	ws.Route(ws.POST("/predicates").To(predicates).
		Doc("scheduler predicate").
		Returns(http.StatusOK, "OK", ""))

	ws.Route(ws.GET("/prioritize").To(prioritizes).
		Doc("scheduler prioritize").
		Returns(http.StatusOK, "OK", ""))

	return ws
}

func probeHealth(r *restful.Request, w *restful.Response) {}
