# Customer Information Web Api

Implemented PostgreSQL with Golang

```docker run --name postgresql-container -p 5432:5432 -e POSTGRES_PASSWORD=Password! -d postgres;```

```
docker build -t cmtapi .
docker run -d -p 8091:8091 cmtapi
```

<img width="1440" alt="Screen Shot 2021-11-04 at 11 37 26" src="https://user-images.githubusercontent.com/60069987/140282941-94900fd2-31a0-40d7-974e-91dca68bf5c2.png">


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
- Read Configuration file : go get github.com/spf13/viper
- Logging : go get github.com/sirupsen/logrus
- PostgreSQL : go get github.com/lib/pq



