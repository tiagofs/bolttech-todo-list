# Bolttech Todo List

## Install
This Go + React + PostgreSQL project was developed in a devcontainer to allow for a consistent and reproducible development environment.

The easiest way to get it up and running is by using the VSCode extension Dev Containers from Microsoft.

Once installed you can either open the project in the container by clicking the popup on the right or in the menu through the bottom left corner button. 
![image](https://github.com/tiagofs/uphold-ticker-bot/assets/20630774/d01b6440-e1e6-4882-9f06-ca4171ce0e24)

The .devcontainer/docker_postgres_init.sql file will be used in the container build process and create the required database tables.

Once inside the container, to run the Go backend first change to the directory apps/backend and run:
```bash
$ go run main.go
```
To run the React frontend change to the directory apps/frontend and run:

```bash
$ npm i
$ npm run dev
```