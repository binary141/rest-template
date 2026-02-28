package db

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type Session struct {
	ID        int64  `db:"id"`
	SessionID string `db:"session_id"`
	ExpiresAt int64  `db:"expires_at"`
	IsValid   bool   `db:"is_valid"`
	UserID    int64  `db:"user_id"`
}

const sessionTimeout = 15 * time.Minute

func newToken() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func CreateSession(userID int64) (Session, error) {
	s := Session{
		SessionID: newToken(),
		ExpiresAt: time.Now().Add(sessionTimeout).Unix(),
		IsValid:   true,
		UserID:    userID,
	}
	_, err := DB.Exec(
		`INSERT INTO sessions (session_id, expires_at, is_valid, user_id) VALUES ($1, $2, $3, $4)`,
		s.SessionID, s.ExpiresAt, s.IsValid, s.UserID,
	)
	return s, err
}

func getSessionByToken(token string) (Session, error) {
	var s Session
	err := DB.QueryRow(
		`SELECT id, session_id, expires_at, is_valid, user_id FROM sessions WHERE session_id = $1`,
		token,
	).Scan(&s.ID, &s.SessionID, &s.ExpiresAt, &s.IsValid, &s.UserID)
	return s, err
}

func IsValidSession(token string) (Session, bool) {
	s, err := getSessionByToken(token)
	if err != nil {
		return Session{}, false
	}
	return s, s.IsValid && s.ExpiresAt >= time.Now().Unix()
}

func ExtendSession(id int64) error {
	_, err := DB.Exec(
		`UPDATE sessions SET expires_at = $2 WHERE id = $1`,
		id, time.Now().Add(sessionTimeout).Unix(),
	)
	return err
}

func DeleteSession(token string) error {
	_, err := DB.Exec(`DELETE FROM sessions WHERE session_id = $1`, token)
	return err
}
