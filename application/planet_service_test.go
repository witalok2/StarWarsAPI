package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/witalok2/starwarsplanet/application"
	mock_application "github.com/witalok2/starwarsplanet/application/mocks"
)

func TestPlanetService_Get(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	planet := mock_application.NewMockPlanetInterface(controller)
	persistence := mock_application.NewMockPlanetPersistenceInterface(controller)
	persistence.EXPECT().Get(gomock.Any()).Return(planet, nil).AnyTimes()
	service := application.PlanetService{Persistence: persistence}

	result, err := service.Get(uuid.New())
	require.Nil(t, err)
	require.Equal(t, planet, result)
}

func TestPlanetService_Create(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	planet := mock_application.NewMockPlanetInterface(controller)
	persistence := mock_application.NewMockPlanetPersistenceInterface(controller)
	persistence.EXPECT().Save(gomock.Any()).Return(planet, nil).AnyTimes()
	service := application.PlanetService{Persistence: persistence}

	result, err := service.Create("Tatooine", "Equatorial", "Rochoso", 3)
	require.Nil(t, err)
	require.Equal(t, planet, result)
}
