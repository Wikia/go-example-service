package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Wikia/go-example-service/cmd/example_app/internal/metrics"
	"github.com/harnash/go-middlewares/logging"
)

type Message struct {
	Text string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	logger := logging.FromRequest(r)
	logger.Info("Greeting user")
	defer metrics.GreetCount.Inc()

	w.WriteHeader(http.StatusOK)
	m := Message{"Hello World"}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err) // no, not really
	}
	_, err = w.Write(b)
	if err != nil {
		logger.With("error", err).Error("could not write response")
	}
}
