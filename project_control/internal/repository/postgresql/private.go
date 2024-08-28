package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	pb "github.com/Mubinabd/project_control/internal/pkg/genproto"
	"strings"
	"time"

	"github.com/google/uuid"
)

type PrivateRepo struct {
	db *sql.DB
}

func NewPrivateManager(db *sql.DB) *PrivateRepo {
	return &PrivateRepo{db: db}
}

func (s *PrivateRepo) CreatePrivate(private *pb.CreatePrivateReq) (*pb.Void, error) {
	privateID := uuid.NewString()

	query := `INSERT INTO private (id, swagger_url, phone_number,telegram_username) VALUES ($1, $2, $3,$4)`
	_, err := s.db.Exec(query, privateID, private.SwaggerUrl, private.PhoneNumber, private.TelegramUsername)
	if err != nil {
		log.Println("error creating Private: ", err)
		return nil, err
	}

	docQuery := `INSERT INTO documentation (id, private_id, title, description, url) VALUES ($1, $2, $3, $4, $5)`
	for _, doc := range private.Documentation {
		docID := uuid.NewString()
		_, err := s.db.Exec(docQuery, docID, privateID, doc.Title, doc.Description, doc.Url)
		if err != nil {
			log.Println("error creating documentation: ", err)
			return nil, err
		}
	}

	return &pb.Void{}, nil
}

func (s *PrivateRepo) GetPrivate(req *pb.ById) (*pb.PrivateGet, error) {
	query := `SELECT 
			p.id, 
			p.swagger_url, 
			p.phone_number,
			p.telegram_username,
			p.created_at,
			d.title,
			d.description,
			d.url
		FROM 
			private p
		JOIN
			documentation d
		ON
			p.id = d.private_id
		WHERE 
			p.id = $1`

	row := s.db.QueryRow(query, req.Id)

	var Private pb.PrivateGet
	Private.Documentation = &pb.Documentation{}

	err := row.Scan(
		&Private.Id,
		&Private.SwaggerUrl,
		&Private.PhoneNumber,
		&Private.TelegramUsername,
		&Private.CreatedAt,
		&Private.Documentation.Title,
		&Private.Documentation.Description,
		&Private.Documentation.Url,
	)

	if err != nil {
		log.Println("Error while getting Private", err)
		return nil, err
	}

	return &Private, nil

}

func (s *PrivateRepo) ListPrivates(req *pb.PrivateListReq) (*pb.PrivateListRes, error) {
	query := `
	SELECT 
		p.id,
		p.swagger_url,
		p.phone_number,
		p.telegram_username,
		p.created_at,
		d.title,
		d.description,
		d.url
	FROM
		private p
	JOIN 
		documentation d
	ON
		p.id = d.private_id
		`

	var args []interface{}
	argCount := 1
	filters := []string{}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}
	if req.Pagination != nil {
		if req.Pagination.Limit > 0 {
			query += fmt.Sprintf(" LIMIT $%d", argCount)
			args = append(args, req.Pagination.Limit)
			argCount++
		}
		if req.Pagination.Offset > 0 {
			query += fmt.Sprintf(" OFFSET $%d", argCount)
			args = append(args, req.Pagination.Offset)
			argCount++
		}
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	Privates := []*pb.PrivateGet{}

	for rows.Next() {
		var Private pb.PrivateGet
		Private.Documentation = &pb.Documentation{}
		err := rows.Scan(
			&Private.Id,
			&Private.SwaggerUrl,
			&Private.PhoneNumber,
			&Private.TelegramUsername,
			&Private.Documentation.Title,
			&Private.Documentation.Description,
			&Private.Documentation.Url,
			&Private.CreatedAt,
		)
		if err != nil {
			log.Println("no rows result set")
			return nil, err
		}
		Privates = append(Privates, &Private)
	}
	return &pb.PrivateListRes{Private: Privates}, nil
}

func (s *PrivateRepo) UpdatePrivate(req *pb.UpdatePrivat) (*pb.Void, error) {
	var args []interface{}
	var conditions []string

	if req.Body.SwaggerUrl != "" && req.Body.SwaggerUrl != "string" {
		args = append(args, req.Body.SwaggerUrl)
		conditions = append(conditions, fmt.Sprintf("swagger_url = $%d", len(args)))
	}
	if req.Body.PhoneNumber != "" && req.Body.PhoneNumber != "string" {
		args = append(args, req.Body.PhoneNumber)
		conditions = append(conditions, fmt.Sprintf("phone_number = $%d", len(args)))
	}
	if req.Body.TelegramUsername != "" && req.Body.TelegramUsername != "string" {
		args = append(args, req.Body.TelegramUsername)
		conditions = append(conditions, fmt.Sprintf("telegram_username = $%d", len(args)))
	}

	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	query := `UPDATE private SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
	args = append(args, req.Id)

	_, err := s.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating Private", err)
		return nil, err
	}

	return &pb.Void{}, nil
}

func (s *PrivateRepo) DeletePrivate(req *pb.DeletePrivat) (*pb.Void, error) {
	query := `
	UPDATE
		private
	SET
		deleted_at = extract(epoch from now())
	WHERE
		id = $1
	`

	_, err := s.db.Exec(query, req.Id)
	if err != nil {
		log.Println("Error while deleting Privates", err)
		return nil, err
	}
	return &pb.Void{}, nil
}
