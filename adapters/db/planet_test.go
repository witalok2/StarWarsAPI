package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/witalok2/starwarsplanet/adapters/db"
	"github.com/witalok2/starwarsplanet/application"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("postgres", ":memory:")
	createTable(Db)
	createPlanet(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE planet (
    id uuid NOT NULL,
    name varchar(100) COLLATE pg_catalog."default" NOT NULL,
		climate varchar(30) NOT NULL,
		terrain varchar(30) NOT NULL,
		MovieAppearances integer,
    CONSTRAINT planet_pkey PRIMARY KEY (id)
	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createPlanet(db *sql.DB) {
	inserted := `insert into planet values("123", "Mercury", "Equatorial", "Rochoso")`

	stmt, err := db.Prepare(inserted)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func testePlanetDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	planetDb := db.NewPlanetDb(Db)
	planet, err := planetDb.Get(uuid.New())

	require.Nil(t, err)
	require.Equal(t, "Mercury", planet.GetName())
}

func testePlanetDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	planetDb := db.NewPlanetDb(Db)

	planet := application.NewPlanet()
	planet.Name = "Mercury"

	planetResult, err := planetDb.Save(planet)
	require.Nil(t, err)
	require.Equal(t, planet.Name, planetResult.GetName())
}
