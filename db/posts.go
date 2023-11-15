package db

import (
	"database/sql"
	"errors"

	"github.com/hosseinmirzapur/elk-example/models"
)

var (
	ErrNoRecord = errors.New("no matching record was found")
	insertOp    = "insert"
	deleteOp    = "delete"
	updateOp    = "update"
)

func (db *Database) SavePost(post *models.Post) error {
	var id int
	query := `INSERT INTO posts (title, content) VALUES ($1, $2) RETURNING id`

	err := db.Conn.QueryRow(query, post.Title, post.Body).Scan(&id)
	if err != nil {
		return err
	}

	logQuery := `INSERT INTO post_logs(post_id, operation) VALUES ($1, $2)`

	post.ID = id
	_, err = db.Conn.Exec(logQuery, post.ID, insertOp)
	if err != nil {
		db.Logger.Err(err).Msg("could not log operation for logstash")
	}
	return nil
}

func (db *Database) UpdatePost(postId int, post models.Post) error {
	query := "UPDATE posts SET title=$1, body=$2 WHERE id=$3"
	_, err := db.Conn.Exec(query, post.Title, post.Body, postId)
	if err != nil {
		return err
	}

	post.ID = postId
	logQuery := "INSERT INTO post_logs(post_id, operation) VALUES ($1, $2)"
	_, err = db.Conn.Exec(logQuery, post.ID, updateOp)
	if err != nil {
		db.Logger.Err(err).Msg("could not log operation for logstash")
	}
	return nil
}

func (db *Database) DeletePost(postId int) error {
	query := "DELETE FROM Posts WHERE id=$1"
	_, err := db.Conn.Exec(query, postId)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNoRecord
		}
		return err
	}

	logQuery := "INSERT INTO post_logs(post_id, operation) VALUES ($1, $2)"
	_, err = db.Conn.Exec(logQuery, postId, deleteOp)
	if err != nil {
		db.Logger.Err(err).Msg("could not log operation for logstash")
	}
	return nil
}
