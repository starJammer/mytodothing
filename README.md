# Intro

Do whatever this doc says to get started.

# Getting started

1. Clone this repo
2. Install [golang](https://golang.org/dl/). Go lang 1.8 was used
at the time of writing this README.
3. Install postgresql. You can install either with `brew install postgresql@9.6`
or go [here](https://www.postgresql.org/download/).

4. Setup the db

```
# create a db for yourself
createdb `whoami`
# connect to postgresql on local
psql

-- drop the public schema in YOUR database
drop schema public cascade;
-- recreate the public schema in YOUR database
create schema public;

-- create the new user with the given password
create role "todouser" login password 'password00' superuser inherit createdb nocreaterole noreplication;

-- create the db
create database "todo" with encoding='UTF8' owner="todouser" connection limit=-1;

-- quit psql program
\q
```

4. Install a few tools for go that will help during development

```
# Cobra, necessary for adding commands
go get -v github.com/spf13/cobra/cobra

# used for managing dependencies for go and getting them into the vendor directory
go get -u -d github.com/mattes/migrate/cli
# If you don't have $GOPATH set then use /usr/local/bin/migrate instead
go build -tags 'postgres' -o $GOPATH/bin/migrate github.com/mattes/migrate/cli

go get -u -t github.com/vattle/sqlboiler
go get -u github.com/govend/govend
govend



```

5. Create config files for your machine by copying and editing the example
files.

```
cp mytodo-example.yaml mytodo.yaml
cp sqlboiler.example.yaml sqlboiler.yaml
```

6. Run database migrations

```
migrate -path migrations -database "postgres://todouser@localhost:5432/todo?sslmode=disable&password=password00" up
```

# Typical things you'll do

## Create a new migration file

```
go run main.go migrations new -n "name_of_migration"
```

## Run migrations

```
# update to the latest
migrate -path migrations -database "postgres://todouser@localhost:5432/todo?sslmode=disable&password=password00" up

# undo the last migration
migrate -path migrations -database "postgres://todouser@localhost:5432/todo?sslmode=disable&password=password00" down

# drop all migrations
migrate -path migrations -database "postgres://todouser@localhost:5432/todo?sslmode=disable&password=password00" drop
```

## Update vendor directory

```
govend
```

## Run SQLBoiler to generate models

```
# Never edit anything in the models folder
sqlboiler postgres
```

## Add a new cobra command

```
cobra add [command]
cobra add [command] -p [parent_command]
```
