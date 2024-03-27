package data

import "time"

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// CreateAtDate format the data to display on the screen
func (t *Thread) CreatedAtDate() string {
	return t.CreatedAt.Format("Jan 2, 2006 at 3:00pm")
}

// CreateAtDate format the data to display on the screen
func (p *Post) CreatedAtDate() string {
	return p.CreatedAt.Format("Jan 2, 2006 at 3:00pm")
}

// NumReplies get the number of posts in a thread
func (t *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts WHERE threadId = $1", t.Id)
	if err != nil {
		return count
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return count
		}
	}
	rows.Close()
	return count

}

// Posts get posts to a thread
func (t *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, uuid, body, userId, threadId, createdAt FROM posts WHERE threadId = $1", t.Id)
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		p := Post{}
		if err = rows.Scan(&p.Id, &p.Uuid, &p.Body, &p.UserId, &p.ThreadId, &p.CreatedAt); err != nil {
			return posts, err
		}
		posts = append(posts, p)
	}
	rows.Close()
	return posts, nil
}

// CreateThread creates a new post to a thread
func (u *User) CreateThread(topic string) (conv Thread, err error) {
	query := `insert INTO threads (uuid, topic, userId, createdAt) VALUES ($1, $2, $3) RETURNING id, uuid, topic, userId, createdAt`
	stmt, err := Db.Prepare(query)
	if err != nil {
		return conv, err
	}
	defer stmt.Close()
	// use queryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), topic, u.Id, time.Now()).Scan(&conv.Id, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return conv, err
}
