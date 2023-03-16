# golag-study

- this is practice for golang
aaaaa

## golang-migration
[golang-migration](https://github.com/golang-migrate/migrate)

- migrationができるよ

into docker container of app(golang container)
```
$ docker exec -it app sh
```

create migration file
```
$ migrate create -ext sql -dir db/sql/  -seq  create_xxxxxx_xxxxxx_table

/app/db/sql/000001_create_xxxxxxx_xxxxxx_table.up.sql
/app/db/sql/000001_create_xxxxxxx_xxxxxx_table.down.sql
```

you can do migration 
```
$ migrate -database ${DATABASE_CONECTION_URL} -path ./db/sql up
```

you can do rollback 
```
$ migrate -database ${DATABASE_CONECTION_URL} -path ./db/sql down
```
