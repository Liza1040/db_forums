package DB_requests

import (
	. "DB_forums/models"
)

func ClearDB_all() (string, error) {
	_, err := DB.Exec(`TRUNCATE Forums, Posts, Threads, Users, Votes, Forums_to_users;`)
	if err != nil {
		return "Invalid database request", err
	}
	return "Database cleared", err
}

func GetStats() (Stats, error) {
	var stats Stats
	err := DB.QueryRow(`SELECT COUNT(*) FROM users;`).Scan(&stats.User)
	if err != nil {
		stats.User = 0
	}
	err = DB.QueryRow(`SELECT COUNT(*) FROM forums;`).Scan(&stats.Forum)
	if err != nil {
		stats.Forum = 0
	}
	err = DB.QueryRow(`SELECT COUNT(*) FROM threads;`).Scan(&stats.Thread)
	if err != nil {
		stats.Thread = 0
	}
	err = DB.QueryRow(`SELECT COUNT(*) FROM posts;`).Scan(&stats.Post)
	if err != nil {
		stats.Post = 0
	}
	return stats, err
}
