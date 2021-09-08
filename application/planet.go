package application

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

const (
	Enabled = "enabled"
)

type PlanetInterface interface {
	IsValid() (bool, error)

	GetID() uuid.UUID
	GetName() string
	GetClimate() string
	GetTerrain() string
	GetMovieAppearances() int
}

type PlanetServiceInterface interface {
	Get(id uuid.UUID) (PlanetInterface, error)
	GetName(name string) (PlanetInterface, error)
	List() ([]PlanetInterface, error)
	Create(name, climate, terrain string, appearances int) (PlanetInterface, error)
	Delete(id uuid.UUID) error
}

type PlanetPersistenceInterface interface {
	PlanetReader
	PlanetWriter
}

type PlanetReader interface {
	Get(id uuid.UUID) (PlanetInterface, error)
	GetName(name string) (PlanetInterface, error)
	List() ([]PlanetInterface, error)
}

type PlanetWriter interface {
	Save(planet PlanetInterface) (PlanetInterface, error)
	Delete(id uuid.UUID) error
}

type Planet struct {
	ID               uuid.UUID `valid:"required"`
	Name             string    `valid:"required"`
	Climate          string    `valid:"required"`
	Terrain          string    `valid:"required"`
	MovieAppearances int       `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewPlanet() *Planet {
	planet := Planet{
		ID: uuid.New(),
	}

	return &planet
}

func (p *Planet) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Planet) GetID() uuid.UUID {
	return p.ID
}

func (p *Planet) GetName() string {
	return p.Name
}

func (p *Planet) GetClimate() string {
	return p.Climate
}

func (p *Planet) GetTerrain() string {
	return p.Terrain
}

func (p *Planet) GetMovieAppearances() int {
	return p.MovieAppearances
}
