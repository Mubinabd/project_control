package postgres

import (
	"database/sql"
	"fmt"

	pb "github.com/Mubinabd/project_control/internal/genproto/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) Register(req *pb.RegisterReq) (*pb.Void, error) {
	res := &pb.Void{}

	tr, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	var id string
	query := `INSERT INTO users (username, email, password, full_name, date_of_birth) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = tr.QueryRow(query, req.Username, req.Email, req.Password, req.FullName, req.DateOfBirth).Scan(&id)
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	query = `INSERT INTO settings (user_id) VALUES ($1)`
	_, err = tr.Exec(query, id)
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	err = tr.Commit()
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	return res, nil
}

func (r *AuthRepo) Login(req *pb.LoginReq) (*pb.User, error) {
	res := &pb.User{}

	var passwordHash string
	query := `SELECT id, username, email, role, password FROM users WHERE username = $1`
	err := r.db.QueryRow(query, req.Username).Scan(
		&res.Id,
		&res.Username,
		&res.Email,
		&res.Role,
		&passwordHash,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid password for username: %s", req.Username)
	}

	return res, nil
}
func (r *AuthRepo) ForgotPassword(req *pb.GetByEmail) (*pb.Void, error) {
	res := &pb.Void{}

	query := `SELECT email FROM users WHERE email = $1`

	var email string
	err := r.db.QueryRow(query, req.Email).Scan(&email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s Email not found", req.Email)
		}
		return nil, err
	}

	return res, nil
}

func (r *AuthRepo) ResetPassword(req *pb.ResetPassReq) (*pb.Void, error) {
	res := &pb.Void{}

	query := `UPDATE users SET password = $1, updated_at=now() WHERE email = $2`

	_, err := r.db.Exec(query, req.NewPassword, req.Email)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *AuthRepo) SaveRefreshToken(req *pb.RefToken) (*pb.Void, error) {
	res := &pb.Void{}

	query := `INSERT INTO tokens (user_id, token) VALUES ($1, $2)`

	_, err := r.db.Exec(query, req.UserId, req.Token)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *AuthRepo) GetAllUsers(req *pb.ListUserReq) (*pb.ListUserRes, error) {
	res := &pb.ListUserRes{}

	query := `SELECT 
				id, 
				username, 
				full_name,
				email, 
				date_of_birth,
				role 
			FROM 
				users 
			WHERE 
				deleted_at=0 AND role = 'user'`

	var args []interface{}

	if req.Username != "" && req.Username != "string" {
		args = append(args, "%"+req.Username+"%")
		query += fmt.Sprintf(" AND username ILIKE $%d", len(args))
	}

	if req.FullName != "" && req.FullName != "string" {
		args = append(args, "%"+req.FullName+"%")
		query += fmt.Sprintf(" AND full_name ILIKE $%d", len(args))
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, req.Filter.Limit, req.Filter.Offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user pb.UserRes
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.FullName,
			&user.Email,
			&user.DateOfBirth,
			&user.Role,
		)
		if err != nil {
			return nil, err
		}
		res.Users = append(res.Users, &user)
	}

	return res, nil
}
