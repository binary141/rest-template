package db

import "time"

type Role struct {
	Id          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

func GetRoles() ([]Role, error) {
	var roles []Role
	err := DB.Select(&roles, `SELECT id, name, COALESCE(description, '') as description, created_at, updated_at FROM roles ORDER BY name`)
	return roles, err
}

func GetRole(id int64) (Role, error) {
	var role Role
	err := DB.QueryRow(
		`SELECT id, name, COALESCE(description, '') as description, created_at, updated_at FROM roles WHERE id = $1`,
		id,
	).Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	return role, err
}

func CreateRole(name, description string) (Role, error) {
	var role Role
	err := DB.QueryRow(
		`INSERT INTO roles (name, description) VALUES ($1, NULLIF($2, '')) RETURNING id, name, COALESCE(description, '') as description, created_at, updated_at`,
		name, description,
	).Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	return role, err
}

func UpdateRole(id int64, name, description string) (Role, error) {
	var role Role
	err := DB.QueryRow(
		`UPDATE roles SET name = $2, description = NULLIF($3, ''), updated_at = now() WHERE id = $1 RETURNING id, name, COALESCE(description, '') as description, created_at, updated_at`,
		id, name, description,
	).Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	return role, err
}

func DeleteRole(id int64) error {
	_, err := DB.Exec(`DELETE FROM roles WHERE id = $1`, id)
	return err
}

func GetUserRoles(userID int64) ([]Role, error) {
	var roles []Role
	err := DB.Select(&roles, `
		SELECT r.id, r.name, COALESCE(r.description, '') as description, r.created_at, r.updated_at
		FROM roles r
		JOIN user_roles ur ON ur.role_id = r.id
		WHERE ur.user_id = $1
		ORDER BY r.name
	`, userID)
	return roles, err
}

func AssignRole(userID, roleID int64) error {
	_, err := DB.Exec(
		`INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`,
		userID, roleID,
	)
	return err
}

func RemoveRole(userID, roleID int64) error {
	_, err := DB.Exec(`DELETE FROM user_roles WHERE user_id = $1 AND role_id = $2`, userID, roleID)
	return err
}
