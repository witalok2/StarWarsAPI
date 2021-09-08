package db

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/witalok2/starwarsplanet/application"
)

type PlanetDb struct {
	db *sql.DB
}

func NewPlanetDb(db *sql.DB) *PlanetDb {
	return &PlanetDb{db: db}
}

func (p *PlanetDb) Get(id uuid.UUID) (application.PlanetInterface, error) {
	var planet application.Planet

	stmt, err := p.db.Prepare("select id, name, climate, terrain, movieappearances from planet where id = $1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&planet.ID, &planet.Name, &planet.Climate, &planet.Terrain, &planet.MovieAppearances)
	if err != nil {
		return nil, err
	}

	return &planet, nil
}

func (p *PlanetDb) GetName(name string) (application.PlanetInterface, error) {
	var planet application.Planet
	stmt, err := p.db.Prepare("select id, name, climate, terrain, movieappearances from planet where name = $1")
	if err != nil {
		return &planet, err
	}

	err = stmt.QueryRow(name).Scan(&planet.ID, &planet.Name, &planet.Climate, &planet.Terrain, &planet.MovieAppearances)
	if err != nil {
		return &planet, err
	}

	return &planet, nil
}

func (p *PlanetDb) List() (planets []application.PlanetInterface, err error) {
	query := "select id, name, climate, terrain, movieappearances from planet"

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var planet application.Planet
		if err := rows.Scan(&planet.ID, &planet.Name, &planet.Climate, &planet.Terrain, &planet.MovieAppearances); err != nil {
			return nil, err
		}
		planets = append(planets, &planet)
	}

	return planets, nil
}

func (p *PlanetDb) Save(planet application.PlanetInterface) (application.PlanetInterface, error) {
	var rows int
	p.db.QueryRow("select id from planet where id = $1", planet.GetID()).Scan(&rows)

	if rows == 0 {
		_, err := p.create(planet)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(planet)
		if err != nil {
			return nil, err
		}
	}

	return planet, nil
}

func (p *PlanetDb) create(planet application.PlanetInterface) (application.PlanetInterface, error) {
	stmt, err := p.db.Prepare(`insert into planet(id, name, climate, terrain, movieappearances) values($1, $2, $3, $4, $5)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(planet.GetID(), planet.GetName(), planet.GetClimate(), planet.GetTerrain(), planet.GetMovieAppearances())
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return planet, nil
}

func (p *PlanetDb) update(planet application.PlanetInterface) (application.PlanetInterface, error) {

	_, err := p.db.Exec("update planet set name = $1, climate = $2, terrain = $3, movieappearances = $4 where id = $5", planet.GetName(), planet.GetClimate(), planet.GetTerrain(), planet.GetMovieAppearances(), planet.GetID())
	if err != nil {
		return nil, err
	}

	return planet, nil
}

func (p *PlanetDb) Delete(id uuid.UUID) error {
	_, err := p.db.Exec("delete from planet where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
