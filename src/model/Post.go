package model

import (
	"config"
)

type Post struct {
	Id       int    `json:"id" form:"id"`
	Title    string `json:"title" form:"title"`
	Content  string `json:"content" form:"content"`
	Uid      int    `json:"uid" form:"uid"`
	Username string `json:"username" form:"username"`
}

var PostInfo []interface{}

func GetPostList() error {
	var posts []interface{}
	rows, err := config.Db.Query("select p.*,u.username from post as p left join user as u on p.uid=u.id where p.uid =? ", UserInfo.UId)
	if err == nil {
		for rows.Next() {
			post := Post{}
			rows.Scan(&post.Id, &post.Title, &post.Content, &post.Uid, &post.Username)
			posts = append(posts, post)
		}
	}
	PostInfo = posts
	rows.Close()
	return err
}
func GetDetail(id int) (post Post, err error) {
	err = config.Db.QueryRow("select p.*,u.username from post as p left join user as u on p.uid=u.id where p.uid =? and p.id = ?", UserInfo.UId, id).Scan(&post.Id, &post.Title, &post.Content, &post.Uid, &post.Username)
	return
}
func (p Post) Add() (err error) {
	_, err = config.Db.Exec("insert into post(title,content,uid) values (?,?,?)", p.Title, p.Content, UserInfo.UId)
	return
}
func (p Post) Update() (res int64, err error) {
	row, err := config.Db.Exec("update post set title=?,content=? where id=? and uid=?", p.Title, p.Content, p.Id, UserInfo.UId)
	res, err = row.RowsAffected()
	return
}

func Delete(id int) (res int64, err error) {
	row, err := config.Db.Exec("delete from post where id = ? and uid=?", id, UserInfo.UId)
	res, err = row.RowsAffected()
	return
}
