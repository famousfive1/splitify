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
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS user_group(id INTEGER, name TEXT);")
	if err != nil {
		return nil, err
	}
	return &groupRepo{
		db: db,
	}, nil
}

func (h *groupRepo) Select(id int) (Group, error) {
	var ans Group
	// row := h.db.QueryRow("SELECT id, name FROM user_group WHERE id = $1", id)
	// if row.Err() != nil {
	// 	return Group{}, row.Err()
	// }
	// row.Scan(&ans.Id, &ans.Name)
	err := h.db.QueryRow("SELECT id, name FROM user_group WHERE id = $1", id).Scan(&ans.Id, &ans.Name)
	if err != nil {
		return Group{}, err
	}
	return ans, nil
}

func (h *groupRepo) Insert(name string) (int, error) {
	query := "INSERT INTO user_group(id, name) VALUES($1, $2)"
	_, err := h.db.Exec(query, 12, name)

	if err != nil {
		return -1, err
	}

	return 12, nil
}
