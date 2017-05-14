package repository

import (
	"database/sql"
	"log"

	uuid "github.com/satori/go.uuid"
	// PostgresSQL
	_ "github.com/lib/pq"
	"github.com/weaver-ci/models"
)

// UserRepository user repository
type UserRepository interface {
	GetUsers() ([]models.User, error)
	GetUser(id uuid.UUID) (models.User, error)

	AddUser(models.User) error
	AddUsers([]models.User) error
}

// PgUserRepository PostgresSQL implementation
type PgUserRepository struct {
	db *sql.DB
}

// NewPgUserRepository creates repository
func NewPgUserRepository(db *sql.DB) *PgUserRepository {
	r := new(PgUserRepository)
	r.db = db
	return r
}

// GetUsers PostgresSQL implementation
func (userRepository PgUserRepository) GetUsers() ([]models.User, error) {
	rows, err := userRepository.db.Query("SELECT user_id, email_address, last_login, modified_on, created_on FROM users")

	if err != nil {
		log.Fatal(err)
	}

	var users []models.User

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.UserID, &user.EmailAddress, &user.LastLogin, &user.ModifiedOn, &user.CreatedOn); err != nil {
			log.Fatal(err)
			return users, err
		}

		users = append(users, user)
	}
	rows.Close()

	return users, nil
}

// GetUser PostgresSQL implementation
func (userRepository PgUserRepository) GetUser(id uuid.UUID) (models.User, error) {
	var user models.User

	err := userRepository.db.QueryRow("SELECT user_id, email_address, last_login, modified_on, created_on FROM users WHERE user_id=$1", id).Scan(&user.UserID, &user.EmailAddress, &user.LastLogin, &user.ModifiedOn, &user.CreatedOn)

	switch {
	case err == sql.ErrNoRows:
		return user, err
	case err != nil:
		log.Fatal(err)
		return user, err
	default:
		return user, nil
	}
}

// AddUser PostgresSQL implementation
func (userRepository PgUserRepository) AddUser(user models.User) error {
	return nil
}

// AddUsers PostgresSQL implementation
func (userRepository PgUserRepository) AddUsers(users []models.User) error {
	return nil
}
