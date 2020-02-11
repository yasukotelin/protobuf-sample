package main

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/yasukotelin/protobuf-sample/go-server/protobuf"
)

func main() {
	e := echo.New()
	e.GET("/books", func(c echo.Context) error {
		books := makeBooks()

		data, err := proto.Marshal(books)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "application/protobuf", data)
	})

	e.Start(":80")
}

func makeBooks() *protobuf.Books {
	books := protobuf.Books{}
	books.Books = make([]*protobuf.Book, 3)

	books.Books[0] = &protobuf.Book{
		Id:       int64(1),
		Title:    "実践Vim 思考のスピードで編集しよう！",
		ImageUrl: "https://images-fe.ssl-images-amazon.com/images/I/51c5qMHV5VL.jpg",
		Label:    "ASCII",
	}
	books.Books[1] = &protobuf.Book{
		Id:       int64(2),
		Title:    "Vimテクニックバイブル～作業効率をカイゼンする150の技",
		ImageUrl: "https://images-fe.ssl-images-amazon.com/images/I/519T9-kMQxL.jpg",
		Label:    "技術評論者",
	}
	books.Books[2] = &protobuf.Book{
		Id:       int64(3),
		Title:    "Vim scriptテクニックバイブル～Vim使いの魔法の杖",
		ImageUrl: "https://images-fe.ssl-images-amazon.com/images/I/51inF679dJL.jpg",
		Label:    "技術評論者",
	}

	return &books
}
