package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUserByID(t *testing.T) {
	// Membuat mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock: %v", err)
	}
	defer db.Close()

	// Menetapkan ekspektasi untuk query SELECT
	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "John Doe")
	mock.ExpectQuery("SELECT id, name FROM users WHERE id = ?").
		WithArgs(1).
		WillReturnRows(rows)

	// initialization user repo
	repo := NewUserRepository(db)

	// Memanggil fungsi yang akan diuji
	user, err := repo.GetUserByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verifikasi hasil
	if user.ID != 1 || user.Name != "John Doe" {
		t.Errorf("unexpected user: got %v", user)
	}

	// Verifikasi bahwa semua ekspektasi terpenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	// Membuat mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock: %v", err)
	}
	defer db.Close()

	// Menetapkan ekspektasi untuk query SELECT yang tidak mengembalikan hasil
	mock.ExpectQuery("SELECT id, name FROM users WHERE id = ?").
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	// initialization user repo
	repo := NewUserRepository(db)

	// Memanggil fungsi yang akan diuji
	_, err = repo.GetUserByID(1)
	if err == nil || err.Error() != "user not found" {
		t.Errorf("expected error 'user not found', got %v", err)
	}

	// Verifikasi bahwa semua ekspektasi terpenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}
