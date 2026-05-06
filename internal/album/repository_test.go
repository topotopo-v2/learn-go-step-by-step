package album

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	cleanup := func() {
		db.Close()
	}

	return db, mock, cleanup
}

func TestRepository_GetAll_Success(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewRepository(db)

	rows := sqlmock.NewRows([]string{"id", "title", "artist", "price"}).
		AddRow(1, "Title1", "Artist1", 10.0).
		AddRow(2, "Title2", "Artist2", 20.0)

	mock.ExpectQuery(`SELECT \* FROM Album`).
		WillReturnRows(rows)

	albums, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Len(t, albums, 2)
	assert.Equal(t, "Title1", albums[0].Title)
	assert.Equal(t, "Artist2", albums[1].Artist)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_GetAll_QueryError(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewRepository(db)

	mock.ExpectQuery(`SELECT \* FROM Album`).
		WillReturnError(sql.ErrConnDone)

	albums, err := repo.GetAll()

	assert.Error(t, err)
	assert.Nil(t, albums)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_GetByID_Success(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewRepository(db)

	row := sqlmock.NewRows([]string{"id", "title", "artist", "price"}).
		AddRow(1, "Test", "Artist", 10.0)

	mock.ExpectQuery(`SELECT id, title, artist, price FROM Album WHERE id=\$1`).
		WithArgs(int64(1)).
		WillReturnRows(row)

	album, err := repo.GetByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, album)
	assert.Equal(t, int64(1), album.ID)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_GetByID_NotFound(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewRepository(db)

	mock.ExpectQuery(`SELECT id, title, artist, price FROM Album WHERE id=\$1`).
		WithArgs(int64(1)).
		WillReturnError(sql.ErrNoRows)

	album, err := repo.GetByID(1)

	assert.ErrorIs(t, err, ErrNotFound)
	assert.Nil(t, album)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Create_Success(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewRepository(db)

	a := &Album{
		Title:  "Test",
		Artist: "Artist",
		Price:  10.0,
	}

	mock.ExpectQuery(`INSERT INTO Album`).
		WithArgs(a.Title, a.Artist, a.Price).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(1),
		)

	err := repo.Create(a)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), a.ID)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Create_Error(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := NewRepository(db)

	a := &Album{
		Title:  "Test",
		Artist: "Artist",
		Price:  10.0,
	}

	mock.ExpectQuery(`INSERT INTO Album`).
		WithArgs(a.Title, a.Artist, a.Price).
		WillReturnError(sql.ErrConnDone)

	err := repo.Create(a)

	assert.Error(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
