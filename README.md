# Intro

Do whatever this doc says to get started.

# Getting started

1. Clone this repo
2. Install postgresql and setup the db

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

3.
