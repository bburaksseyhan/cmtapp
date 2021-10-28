# Customer Information Web Api

Implemented PostgreSQL with Golang

```docker run --name postgresql-container -p 5432:5432 -e POSTGRES_PASSWORD=Password! -d postgres;```

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

used packages

- Web framework: go get github.com/labstack/echo/v4
- Logging : go get github.com/sirupsen/logrus
- PostgreSQL : go get github.com/lib/pq


