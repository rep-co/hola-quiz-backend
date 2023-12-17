package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Storage interface {
	CreatePack(*Pack) error
	DeletePack(int) error
	GetPackByID(int) (*Pack, error)
	GetPacks() ([]*Pack, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "postgres://admin:root@localhost:5432/holaquiz?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) Init() error {
	return s.createPackTable()
}

func (s *PostgresStorage) createPackTable() error {
	query := `create table if not exists pack (
        id serial primary key,
        name varchar(255),
        description varchar(255),
        category varchar(255),
        emojis varchar(3)
    );`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreatePack(pack *Pack) error {
	query := `insert into pack (
        name, 
        description,
        category,
        emojis
    ) values ($1, $2, $3, $4);`
	_, err := s.db.Query(
		query,
		pack.Name,
		pack.Description,
		pack.Category,
		pack.Emojis)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) DeletePack(id int) error {
	return nil
}

func (s *PostgresStorage) GetPackByID(id int) (*Pack, error) {
	return nil, nil
}

func (s *PostgresStorage) GetPacks() ([]*Pack, error) {
	query := `select * from pack;`
	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	packs := []*Pack{}
	for rows.Next() {
		pack := new(Pack)
		err := rows.Scan(
			&pack.ID,
			&pack.Name,
			&pack.Description,
			&pack.Category,
			&pack.Emojis)

		if err != nil {
			return nil, err
		}

		packs = append(packs, pack)
	}

	return packs, nil
}
