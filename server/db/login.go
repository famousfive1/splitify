package db

import "database/sql"

type LoginRepo interface {
	Select(string) (Login, error)
	Insert(string, string) (bool, error)
	Update(string, string) (bool, error)
}

type loginRepo struct {
	db *sql.DB
}

func newLoginRepo(db *sql.DB) (LoginRepo, error) {
	ctx, cancel := getStandardContext()
	defer cancel()

	_, err := db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS login(username TEXT PRIMARY KEY, password TEXT)")
	if err != nil {
		return nil, err
	}
	return &loginRepo{
		db: db,
	}, nil
}

func (h *loginRepo) Select(username string) (Login, error) {
	ctx, cancel := getStandardContext()
	defer cancel()

	var ans Login
	query := "SELECT username, password FROM login WHERE username = $1"
	err := h.db.QueryRowContext(ctx, query, username).Scan(&ans.Username, &ans.Password)
	if err != nil {
		return Login{}, err
	}
	return ans, nil
}

func (h *loginRepo) Insert(username string, password string) (bool, error) {
	ctx, cancel := getStandardContext()
	defer cancel()

	query := "INSERT INTO login(username, password) VALUES($1, $2)"
	_, err := h.db.ExecContext(ctx, query, username, password)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (h *loginRepo) Update(username string, password string) (bool, error) {
	ctx, cancel := getStandardContext()
	defer cancel()

	query := "UPDATE login SET password = $2 WHERE username = $1"
	_, err := h.db.ExecContext(ctx, query, username, password)

	if err != nil {
		return false, err
	}

	return true, nil
}
