package endpoint

import (
	"TemplateApi/src/service"
	"net/http"
)

func HeartbeatHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var exampleModel models.ExampleModel

		//decoder := json.NewDecoder(r.Body)
		//if err := decoder.Decode(&exampleModel); err != nil {
		//	respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		//	return
		//}
		//defer r.Body.Close()

		heartbeatResponse, err := svc.Heartbeat()
		if err != nil {
			respondWithJSON(w, r, 500, r.Body)
			return
		}
		respondWithJSON(w, r, http.StatusCreated, heartbeatResponse)
	}
}
