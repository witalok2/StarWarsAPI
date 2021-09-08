package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	swapi "github.com/witalok2/starwarsplanet/adapters/api/swapi"
	"github.com/witalok2/starwarsplanet/adapters/dto"
	"github.com/witalok2/starwarsplanet/application"
)

func MakePlanetHandlers(r *mux.Router, n *negroni.Negroni, service application.PlanetServiceInterface) {
	r.Handle("/planet/{id}", n.With(negroni.Wrap(getPlanetID(service)))).Methods("GET", "OPTIONS")
	r.Handle("/planet/search/{name}", n.With(negroni.Wrap(getPlanetName(service)))).Methods("GET", "OPTIONS")
	r.Handle("/planet", n.With(negroni.Wrap(listPlanet(service)))).Methods("GET", "OPTIONS")
	r.Handle("/planet", n.With(negroni.Wrap(createPlanet(service)))).Methods("POST", "OPTIONS")
	r.Handle("/planet/{id}", n.With(negroni.Wrap(deletePlanet(service)))).Methods("DELETE", "OPTIONS")
}

func getPlanetName(service application.PlanetServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		name := vars["name"]

		planet, err := service.GetName(name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(planet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func getPlanetID(service application.PlanetServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		planet, err := service.Get(uuid.MustParse(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(planet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createPlanet(service application.PlanetServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var planetDTO dto.Planet

		err := json.NewDecoder(r.Body).Decode(&planetDTO)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		planetIsExist, _ := service.GetName(planetDTO.Name)

		if planetIsExist.GetName() == planetDTO.Name {
			w.Write(jsonError("Planeta já cadastrado"))
			return
		}

		clientSWAPI := swapi.Client{
			Host: "https://swapi.dev/api/",
		}

		var errSWAPI error

		response, errSWAPI := clientSWAPI.GetPlanetInSwapiByName(planetDTO.Name)
		if errSWAPI != nil {
			log.Println("Não foi possivel encontrar a quantidade de aparição do planeta nos filmes de star wars.")
		}

		var MovieAppearances int

		if response.Count > 0 && len(response.Results[0].Films) > 0 {
			MovieAppearances = len(response.Results[0].Films)
		} else {
			MovieAppearances = 0
		}

		planet, err := service.Create(planetDTO.Name, planetDTO.Climate, planetDTO.Terrain, MovieAppearances)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(planet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func listPlanet(service application.PlanetServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		planet, err := service.List()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(planet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func deletePlanet(service application.PlanetServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		planet, err := service.Get(uuid.MustParse(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = service.Delete(planet.GetID())
		if err != nil {
			w.Write(jsonError("Erro ao remover planeta do universo"))
			return
		}

		w.Write(jsonError("Planeta removido do universo."))
	})
}
