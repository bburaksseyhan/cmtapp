[![Go Doc](https://godoc.org/github.com/gogf/gf?status.svg)](https://godoc.org/github.com/gogf/gf)
[![Go Report Card](https://goreportcard.com/badge/github.com/bburaksseyhan/cmtapp)](https://goreportcard.com/report/github.com/bburaksseyhan/cmtapp)

# Customer Information Web Api

Implemented PostgreSQL with Golang on Docker :ship:

```docker run --name postgresql-container -p 5432:5432 -e POSTGRES_PASSWORD=Password! -d postgres;```

```
docker build -t cmtapi .
docker run -d -p 8091:8091 cmtapi
```
<img width="1068" alt="Screen Shot 2021-11-04 at 13 29 03" src="https://user-images.githubusercontent.com/60069987/140298537-31b7d548-d902-439d-bbb9-48a2ce85aa51.png">

```
-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	firstname varchar NULL,
	lastname varchar NULL,
	id serial NOT NULL
);

```

```git clone https://github.com/bburaksseyhan/cmtapp.git```

```
go to > src/cmd/api
go run main.go
```

used packages :package:

- [ ] Web framework: go get github.com/labstack/echo/v4
- [ ] Read Configuration file : go get github.com/spf13/viper
- [ ] Logging : go get github.com/sirupsen/logrus
- [ ] PostgreSQL : go get github.com/lib/pq

documentation: [Related Post](https://dev.to/bseyhan/golang-with-database-operations-3jl0). :pencil2: :book:



