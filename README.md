# Shortly

This repo used for learning go-swagger, squirrel, go-migrate

## Development
To enable sqlite3 support for `migrate`, install using sqlite3 tag  
```
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Status
- [x] create short url
    - [ ] better short url generator, currently using non secure random generator
- [x] list short url
- [ ] redirect short url to original url
- [ ] implements authN & authZ
- [ ] record short url statistic