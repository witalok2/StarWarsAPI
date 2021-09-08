package application_test

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/witalok2/starwarsplanet/application"
)

func TestPlanet_IsValid(t *testing.T) {
	planet := application.Planet{
		ID:      uuid.New(),
		Name:    "Tatooine",
		Climate: "Equatorial",
		Terrain: "Rochoso",
	}

	_, err := planet.IsValid()
	require.Nil(t, err)
}
