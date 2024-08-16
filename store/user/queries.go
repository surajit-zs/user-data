package user

import (
	"fmt"
	"strings"

	"github.com/user-data/models"
)

const (
	createUserQuery = `INSERT INTO "user" (id, name, user_name, password, created_at) VALUES ($1, $2, $3, $4, $5)`
	getUserQuery    = `SELECT id, name, user_name, password, created_at, updated_at FROM "user" where id = $1 and 
                    	deleted_at IS NULL`
	getAllUserQQuery = `SELECT id, name, user_name, password, created_at, updated_at FROM "user" `
	deleteUserQuery  = `UPDATE "user" SET deleted_at=$1 where id = $2`
)

func getAllBuildQuery(f *models.Filter) (query string, value []interface{}) {
	var sb strings.Builder

	args := make([]interface{}, 0)

	if f == nil {
		f = &models.Filter{Limit: 10, Offset: 0}
	}

	// Base query
	sb.WriteString(getAllUserQQuery)

	// Collect conditions
	conditions := make([]string, 0)
	placeholderIndex := 1

	if f.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name = $%d", placeholderIndex))
		args = append(args, f.Name)
		placeholderIndex++
	}

	if f.UserName != "" {
		conditions = append(conditions, fmt.Sprintf("user_name = $%d", placeholderIndex))
		args = append(args, f.UserName)
		placeholderIndex++
	}

	if !f.CreatedAt.IsZero() {
		conditions = append(conditions, fmt.Sprintf("created_at >= $%d", placeholderIndex))
		args = append(args, f.CreatedAt)
		placeholderIndex++
	}

	if !f.UpdatedAt.IsZero() {
		conditions = append(conditions, fmt.Sprintf("updated_at >= $%d", placeholderIndex))
		args = append(args, f.UpdatedAt)
		placeholderIndex++
	}

	conditions = append(conditions, "deleted_at is NULL")

	// Append WHERE clause if there are conditions
	if len(conditions) > 0 {
		sb.WriteString(" WHERE ")
		sb.WriteString(strings.Join(conditions, " AND "))
	}

	// Append ORDER BY clause (optional)
	sb.WriteString(" ORDER BY created_at DESC")

	// Append LIMIT and OFFSET if specified
	if f.Limit > 0 {
		sb.WriteString(fmt.Sprintf(" LIMIT $%d", placeholderIndex))

		args = append(args, f.Limit)

		placeholderIndex++
	}

	if f.Offset > 0 {
		sb.WriteString(fmt.Sprintf(" OFFSET $%d", placeholderIndex))

		args = append(args, f.Offset)
	}

	return sb.String(), args
}

// buildUpdateQuery builds the SQL update query based on the provided User struct.
func buildUpdateQuery(user *models.User) (query string, value []interface{}) {
	var sb strings.Builder

	args := make([]interface{}, 0)

	// Base query
	sb.WriteString(`UPDATE "user" SET `)

	// Collect SET clause components
	setClauses := make([]string, 0)
	placeholderIndex := 1

	if user.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", placeholderIndex))
		args = append(args, user.Name)
		placeholderIndex++
	}

	if user.UserName != "" {
		setClauses = append(setClauses, fmt.Sprintf("user_name = $%d", placeholderIndex))
		args = append(args, user.UserName)
		placeholderIndex++
	}

	if user.Password != "" {
		setClauses = append(setClauses, fmt.Sprintf("password = $%d", placeholderIndex))
		args = append(args, user.Password)
		placeholderIndex++
	}

	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", placeholderIndex))
	args = append(args, user.UpdatedAt)
	placeholderIndex++

	// Ensure there is at least one field to update
	if len(setClauses) == 0 {
		return "", nil
	}

	// Append SET clause components
	sb.WriteString(strings.Join(setClauses, ", "))

	// Add WHERE clause to specify which row to update
	sb.WriteString(fmt.Sprintf(" WHERE id = $%d", placeholderIndex))

	args = append(args, user.ID)

	return sb.String(), args
}
