## pger

`pger` lets you migrate your database via command lines.

It will load .sql files based on their natural order and run them in order. You can't migrate down or anything like that.

It expects filenames to fit the pattern `NUM-filename.sql`. The NUM section is used for ordering.

## Install

`go install github.com/nicholasf/pger`

## Usage

```
pger -d=someapp

CREATE TABLE

Migrated 01-create-users.sql.

CREATE TABLE

Migrated 02-create-groups.sql.
```

`pger` will take the following flags: 

* `-d dbname`. Database name is the only mandatory arg.
* `-psql /usr/bin/psql`. Allows you to specify the path to `psql`. Defaults to being '/usr/local/bin/psql'.
* `-h 192.1680.1`. Host otherwise defaults to being 'localhost'.
* `-U username`. If you are running in local trust mode just leave this out.
* `-dir /path/to/migrations`. If you are storing your migration files in a directory different from where you are executing `pger`
* `-W password`. This is still unimplemented. 