module github.com/Komdosh/go-bookstore-items-api

go 1.15

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0

require (
	github.com/Komdosh/go-bookstore-oauth-client v0.0.0-20210116085116-b7b7c1f03e4a
	github.com/Komdosh/go-bookstore-utils v0.0.0-20210116090756-a7d3cfa03af1
	github.com/gorilla/mux v1.8.0
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/olivere/elastic v6.2.35+incompatible
)
