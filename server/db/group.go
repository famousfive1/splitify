package db

import "database/sql"

type GroupRepo interface {
	Select(int) (Group, error)
	Insert(string) (int, error)
}

type groupRepo struct {
	db *sql.DB
}

func newGroupRepo(db *sql.DB) (GroupRepo, error) {
	ctx, cancel := getStandardContext()
	defer cancel()

	_, err := db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS user_group(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);")
	if err != nil {
		return nil, err
	}
	return &groupRepo{
		db: db,
	}, nil
}

func (h *groupRepo) Select(id int) (Group, error) {
	ctx, cancel := getStandardContext()
	defer cancel()

	var ans Group
	err := h.db.QueryRowContext(ctx, "SELECT id, name FROM user_group WHERE id = $1", id).Scan(&ans.Id, &ans.Name)
	if err != nil {
		return Group{}, err
	}
	return ans, nil
}

func (h *groupRepo) Insert(name string) (int, error) {
	ctx, cancel := getStandardContext()
	defer cancel()

	var id int
	query := "INSERT INTO user_group(name) VALUES($1) RETURNING id"
	err := h.db.QueryRowContext(ctx, query, name).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}
