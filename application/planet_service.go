package application

import (
	"github.com/google/uuid"
)

type PlanetService struct {
	Persistence PlanetPersistenceInterface
}

func NewPlanetService(persistence PlanetPersistenceInterface) *PlanetService {
	return &PlanetService{Persistence: persistence}
}

func (s *PlanetService) Get(id uuid.UUID) (PlanetInterface, error) {
	planet, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return planet, nil
}

func (s *PlanetService) GetName(name string) (PlanetInterface, error) {
	planet, err := s.Persistence.GetName(name)

	if err != nil {
		return planet, err
	}

	return planet, nil
}

func (s *PlanetService) List() ([]PlanetInterface, error) {
	planet, err := s.Persistence.List()
	if err != nil {
		return nil, err
	}

	return planet, nil
}

func (s *PlanetService) Create(name, climate, terrain string, appearances int) (PlanetInterface, error) {
	planet := NewPlanet()
	planet.Name = name
	planet.Climate = climate
	planet.Terrain = terrain
	planet.MovieAppearances = appearances

	_, err := planet.IsValid()
	if err != nil {
		return &Planet{}, err
	}

	result, err := s.Persistence.Save(planet)
	if err != nil {
		return &Planet{}, err
	}

	return result, nil
}

func (s *PlanetService) Delete(id uuid.UUID) error {
	err := s.Persistence.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
