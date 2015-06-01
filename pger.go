package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"os/exec"
	"flag"
	"fmt"
)

var (
	psql = flag.String("psql", "/usr/local/bin/psql", "The path to psql on your system.")
	host = flag.String("h", "localhost", "The -h (host) arg.")
	database = flag.String("d", "", "The -d (database) arg.")
	user = flag.String("U", "", "The -U (username) arg.")
	password = flag.String("W", "", "The password to react to -W. Leave blank if you use local trust.")
	dir = flag.String("dir", "", "The directory holding your migration files. Will default to the dir pger runs in.")
)

func usage() {
	fmt.Fprintf(os.Stderr, "\tpger migrate your sql scripts in order by 01-create-users.sql")
	fmt.Fprintf(os.Stderr, "\tpger -h <host> -d <database> -U <user> -W <password>")
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("pger: ")
	flag.Parse()

	if len(*database) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	if len(*dir) == 0 {
		wd, err := os.Getwd()

		*dir = wd

		if err != nil {
			log.Fatal(err.Error())
		}
	}

	infos, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal("Could not read from file system.")
	}

	sqls := loadMigrations(infos)

	for _, migration := range sqls {
		cmdArgs := commandString(*host, *database, *user, *dir, migration.filename)

		handlePassword := false

		if len(*password) > 0 {
			handlePassword = true
			cmdArgs = append(cmdArgs, "-W")
		}

		cmd := exec.Command(*psql, cmdArgs...)

		if handlePassword {
			cmd.Stdin = strings.NewReader(*password)
		}

		out, err := cmd.Output()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(out))
		fmt.Printf("Migrated %s.\n", migration.filename)
	}
}

func loadMigrations(files []os.FileInfo) []migration {
	sqls := make([]migration, 0)
	sqlPattern, err := regexp.Compile("[0-9].*\\-")

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		name := []byte(file.Name())
		if sqlPattern.Match(name) {
			bits := strings.Split(file.Name(), "-")
			sqls = append(sqls, migration{bits[0], file.Name()})
		}
	}

	return sqls
}

type migration struct {
	number, filename string
}

func commandString(host, database, user, workingDirectory, filename string) []string {
	var cmdStr []string
	cmdStr = []string{"-h", host, "-d", database, "-U", user, "-f", (fmt.Sprintf("%s/%s", workingDirectory, filename))}
	return cmdStr
}