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

4. Install `migrate` to manage the database migrations

```

```
