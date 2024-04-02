# Simple Todo App

## About
It's a minimalistic app demonstrates how to use [bun](https://github.com/uptrace/bun) + [templ](https://templ.guide) & [htmx](https://htmx.org). Project structure based on [slick](https://github.com/anthdm/slick). You can learn more about project structure from [this youtube video](https://www.youtube.com/watch?v=dJIUxvfSg6A).


## Implemented features
- [x] User authentication (login, logout, signup)
- [x] Add new todo
- [x] Edit todo
- [x] Delete todo
- [x] List user todos
- [x] Update user profile (username, password)
- [x] Flash alert messages
- [x] App landing page
- [ ] Use Alpine instead Debian in docker image
- [ ] Password recovery
- [ ] Show users and their todos in admin panel


## Run the app
You can run with different ways by `make run`, `slick run` or `go run cmd/main.go`.
You can run in watch mode using `air` or `make watch`.

## How to run in docker
```bash
docker-compose up -d
```

## How to stop the app
```bash
docker-compose down
```
