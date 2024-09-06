package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/Mubinabd/project_control/pkg/genproto"

	"github.com/google/uuid"
)

type GroupRepo struct {
	db *sql.DB
}

func NewGroupManager(db *sql.DB) *GroupRepo {
	return &GroupRepo{db: db}
}

func (s *GroupRepo) CreateGroup(group *pb.CreateGroupReq) (*pb.Void, error) {
	groupID := uuid.NewString()

	query := `INSERT INTO groups (id, swagger_url, name) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, groupID, group.SwaggerUrl, group.Name)
	if err != nil {
		log.Println("error creating group: ", err)
		return nil, err
	}

	devQuery := `INSERT INTO developers (id, group_id, name, phone_number, telegram_username) VALUES ($1, $2, $3, $4, $5)`
	for _, dev := range group.Developers {
		devID := uuid.NewString()
		_, err := s.db.Exec(devQuery, devID, groupID, dev.Name, dev.PhoneNumber, dev.TelegramUsername)
		if err != nil {
			log.Println("error creating developer: ", err)
			return nil, err
		}
	}

	docQuery := `INSERT INTO documentation (id, group_id, title, description, url) VALUES ($1, $2, $3, $4, $5)`
	for _, doc := range group.Documentation {
		docID := uuid.NewString()
		_, err := s.db.Exec(docQuery, docID, groupID, doc.Title, doc.Description, doc.Url)
		if err != nil {
			log.Println("error creating documentation: ", err)
			return nil, err
		}
	}

	return &pb.Void{}, nil
}

func (s *GroupRepo) GetGroup(req *pb.ById) (*pb.GroupGet, error) {
	query := `SELECT 
			g.id, 
			g.swagger_url, 
			g.name,
			g.created_at,
			d.name, 
			d.phone_number, 
			d.telegram_username,
			doc.title,
			doc.description,
			doc.url
		FROM 
			groups g
		JOIN
			developers d
		ON
			g.id = d.group_id
		JOIN
			documentation doc
		ON
			g.id = doc.group_id
		WHERE 
			g.id = $1
		AND
			g.deleted_at = 0`

	row := s.db.QueryRow(query, req.Id)

	var group pb.GroupGet
	group.Developers = &pb.Developer{}
	group.Documentation = &pb.Documentation{}

	err := row.Scan(
		&group.Id,
		&group.Name,
		&group.SwaggerUrl,
		&group.CreatedAt,
		&group.Developers.Name,
		&group.Developers.PhoneNumber,
		&group.Developers.TelegramUsername,
		&group.Documentation.Title,
		&group.Documentation.Description,
		&group.Documentation.Url,
	)

	if err != nil {
		log.Println("Error while getting group", err)
		return nil, err
	}

	return &group, nil

}

func (s *GroupRepo) ListGroups(req *pb.GroupListReq) (*pb.GroupListRes, error) {
	query := `
	SELECT 
		g.id,
		g.name,
		g.swagger_url,
		g.created_at,
		d.name,	
		d.phone_number,
		d.telegram_username,
		doc.title,
		doc.description,
		doc.url
	FROM
		groups g
	JOIN
		developers d
	ON
		g.id = d.group_id
	JOIN 
		documentation doc
	ON
		g.id = doc.group_id
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

	groups := []*pb.GroupGet{}

	for rows.Next() {
		var group pb.GroupGet
		group.Developers = &pb.Developer{}
		group.Documentation = &pb.Documentation{}
		err := rows.Scan(
			&group.Id,
			&group.Name,
			&group.SwaggerUrl,
			&group.Developers.Name,
			&group.Developers.PhoneNumber,
			&group.Developers.TelegramUsername,
			&group.Documentation.Title,
			&group.Documentation.Description,
			&group.Documentation.Url,
			&group.CreatedAt,
		)
		if err != nil {
			log.Println("no rows result set")
			return nil, err
		}
		groups = append(groups, &group)
	}
	return &pb.GroupListRes{Group: groups}, nil
}

func (s *GroupRepo) UpdateGroup(req *pb.UpdateGr) (*pb.Void, error) {
	var args []interface{}
	var conditions []string

	if req.Body.Name != "" && req.Body.Name != "string" {
		args = append(args, req.Body.Name)
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(args)))
	}

	if req.Body.SwaggerUrl != "" && req.Body.SwaggerUrl != "string" {
		args = append(args, req.Body.SwaggerUrl)
		conditions = append(conditions, fmt.Sprintf("swagger_url = $%d", len(args)))
	}

	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	if len(conditions) > 0 {
		query := `UPDATE groups SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
		args = append(args, req.Id)

		_, err := s.db.Exec(query, args...)
		if err != nil {
			log.Println("Error while updating group", err)
			return nil, err
		}
	}

	for _, dev := range req.Body.Developers {
		var existingDevID string
		checkQuery := `SELECT id FROM developers WHERE group_id = $1 AND phone_number = $2`
		err := s.db.QueryRow(checkQuery, req.Id, dev.PhoneNumber).Scan(&existingDevID)

		if err == nil {
			var devArgs []interface{}
			var devConditions []string

			if dev.Name != "" && dev.Name != "string" {
				devArgs = append(devArgs, dev.Name)
				devConditions = append(devConditions, fmt.Sprintf("name = $%d", len(devArgs)))
			}
			if dev.TelegramUsername != "" && dev.TelegramUsername != "string" {
				devArgs = append(devArgs, dev.TelegramUsername)
				devConditions = append(devConditions, fmt.Sprintf("telegram_username = $%d", len(devArgs)))
			}

			if len(devConditions) > 0 {
				devArgs = append(devArgs, existingDevID)
				updateQuery := `UPDATE developers SET ` + strings.Join(devConditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(devArgs))
				_, err := s.db.Exec(updateQuery, devArgs...)
				if err != nil {
					log.Println("Error while updating developer:", err)
					return nil, err
				}
			}
		} else if err == sql.ErrNoRows {
			insertQuery := `INSERT INTO developers (id, group_id, name, phone_number, telegram_username) VALUES ($1, $2, $3, $4, $5)`
			_, err := s.db.Exec(insertQuery, uuid.NewString(), req.Id, dev.Name, dev.PhoneNumber, dev.TelegramUsername)
			if err != nil {
				log.Println("Error while inserting developer:", err)
				return nil, err
			}
		} else {
			log.Println("Error while checking developer existence:", err)
			return nil, err
		}
	}

	// Update documentation
	for _, doc := range req.Body.Documentation {
		var docId string
		checkQuery := `SELECT id FROM documentation WHERE group_id = $1 AND url = $2`
		err := s.db.QueryRow(checkQuery, req.Id, doc.Url).Scan(&docId)

		if err == nil {
			var docArgs []interface{}
			var docConditions []string

			if doc.Description != "" && doc.Description != "string" {
				docArgs = append(docArgs, doc.Description)
				docConditions = append(docConditions, fmt.Sprintf("description = $%d", len(docArgs)))
			}
			if doc.Title != "" && doc.Title != "string" {
				docArgs = append(docArgs, doc.Title)
				docConditions = append(docConditions, fmt.Sprintf("title = $%d", len(docArgs)))
			}

			if len(docConditions) > 0 {
				docArgs = append(docArgs, docId)
				updateQuery := `UPDATE documentation SET ` + strings.Join(docConditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(docArgs))
				_, err := s.db.Exec(updateQuery, docArgs...)
				if err != nil {
					log.Println("Error while updating documentation:", err)
					return nil, err
				}
			}
		} else if err == sql.ErrNoRows {
			insertQuery := `INSERT INTO documentation (id, group_id, title, description, url) VALUES ($1, $2, $3, $4, $5)`
			_, err := s.db.Exec(insertQuery, uuid.NewString(), req.Id, doc.Title, doc.Description, doc.Url)
			if err != nil {
				log.Println("Error while inserting documentation:", err)
				return nil, err
			}
		} else {
			log.Println("Error while checking documentation existence:", err)
			return nil, err
		}
	}

	return &pb.Void{}, nil
}

func (s *GroupRepo) DeleteGroup(req *pb.DeleteGr) (*pb.Void, error) {
	query := `
	UPDATE
		groups
	SET
		deleted_at = extract(epoch from now())
	WHERE
		id = $1
	`
	

	_, err := s.db.Exec(query, req.Id)
	if err != nil {
		log.Println("Error while deleting groups", err)
		return nil, err
	}

	queryDocs := `
	UPDATE
		groups
	SET
		deleted_at = extract(epoch from now())
	WHERE
		group_id = $1
	`
	_, err = s.db.Exec(queryDocs, req.Id)
	if err != nil {
		log.Println("Error while deleting related documentation", err)
		return nil, err
	}

	queryDevs := `
	UPDATE
		groups
	SET
		deleted_at = extract(epoch from now())
	WHERE
		group_id = $1
	`
	_, err = s.db.Exec(queryDevs, req.Id)
	if err != nil {
		log.Println("Error while deleting related developers", err)
		return nil, err
	}
	return &pb.Void{}, nil
}
