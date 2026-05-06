package album

import "database/sql"

var _ RepositoryI = (*Repository)(nil)

type RepositoryI interface {
	GetAll() ([]Album, error)
	GetByID(id int64) (*Album, error)
	Create(a *Album) error
}
type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]Album, error) {

	rows, err := r.db.Query(`SELECT * FROM Album`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums = make([]Album, 0)

	for rows.Next() {
		var a Album
		if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
			return nil, err
		}
		albums = append(albums, a)
	}

	return albums, nil
}

func (r *Repository) GetByID(id int64) (*Album, error) {
	var a Album

	err := r.db.QueryRow(
		"SELECT id, title, artist, price FROM Album WHERE id=$1", // $1 prevents SQL injection
		id,
	).Scan(&a.ID, &a.Title, &a.Artist, &a.Price)

	if err != nil {
		return nil, ErrNotFound
	}

	return &a, nil
}

func (r *Repository) Create(a *Album) error {
	return r.db.QueryRow(
		"INSERT INTO Album (title, artist, price) VALUES ($1, $2, $3) RETURNING id",
		a.Title, a.Artist, a.Price,
	).Scan(&a.ID)
}
