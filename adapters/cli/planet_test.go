package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"github.com/stretchr/testify/require"
	"github.com/witalok2/starwarsplanet/adapters/cli"
	mock_application "github.com/witalok2/starwarsplanet/application/mocks"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	planetID := uuid.New()
	planetName := "Mercury"
	planetClimate := "Equatorial"
	planetTerrain := "Rochoso"

	planetMock := mock_application.NewMockPlanetInterface(ctrl)
	planetMock.EXPECT().GetID().Return(planetID).AnyTimes()
	// planetMock.EXPECT().GetPlanet().Return(planetName).AnyTimes()
	planetMock.EXPECT().GetName().Return(planetName).AnyTimes()

	service := mock_application.NewMockPlanetServiceInterface(ctrl)
	service.EXPECT().Create(planetName, planetClimate, planetTerrain, 2).Return(planetMock, nil).AnyTimes()
	service.EXPECT().Get(planetID).Return(planetMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Planet ID %s with the name %s has been created", planetID, planetName)
	result, err := cli.Run(service, "create", uuid.Nil, planetName, planetClimate, planetTerrain, 2)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
