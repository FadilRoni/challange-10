package repository

import (
	"challange-10/model"
	"challange-10/repository/query"
)

type BookRepo interface {
	AddBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookId(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (r Repo) AddBook(in model.Book) (res model.Book, err error) {
	err = r.db.QueryRow(
		query.AddBook,
		in.Title,
		in.Author,
		in.Description,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Author,
		&res.Description,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllBook() (res []model.Book, err error) {
	var book model.Book

	err = r.db.QueryRow(
		query.GetAllBook,
	).Scan(&book.ID, &book.Title, &book.Author, &book.Description)

	if err != nil {
		return res, err
	}

	res = append(res, book)

	return res, nil
}

func (r Repo) GetBookId(id int64) (res model.Book, err error) {
	err = r.db.QueryRow(
		query.GetBookId,
		id,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Author,
		&res.Description,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateBook(in model.Book) (res model.Book, err error) {

	// in = model.Book{}

	_, err = r.db.Exec(
		query.UpdateBook,
		in.ID,
		in.Title,
		in.Author,
		in.Description,
	)

	if err != nil {
		return res, nil
	}

	return res, nil
}

func (r Repo) DeleteBook(id int64) (err error) {
	in := model.Book{}
	err = r.db.QueryRow(
		query.DeleteBook,
		id,
	).Scan(
		&in.ID,
		&in.Title,
		&in.Author,
		&in.Description,
	)

	if err != nil {
		return err
	}

	return nil
}
