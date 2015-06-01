## pger

`pger` lets you migrate, create, drop or reset your database via command lines.

It will load .sql files based on their natural order and run them in order. You can't migrate down or anything like that.

It expects filenames to fit the pattern `NUM-filename.sql`. The NUM section is used for ordering.

## Install

`go install github.com/nicholasf/pger`
