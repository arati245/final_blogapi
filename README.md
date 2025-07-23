# Final Blog API

A simple RESTful Blog API built in Go using Gin, GORM, and MySQL
## Features

- Create, Read, Update, Delete blog posts
- MySQL database connection
- Environment variable support with `.env` and `godotenv`
- Deployed-ready structure
## Setup

1. Clone the repo:
git clone https://github.com/arati245/final_blogapi.git
cd final_blogapi
2. Set up your `.env` file:
DB_USER=root
DB_PASS=@Arati2004
DB_HOST=127.0.0.1
DB_NAME=blogapplication

3. Run the app:
go run blogapp/main.go

## API Endpoints
- `POST /posts` — Create post
- `GET /posts` — Get all posts
- `GET /posts/:id` — Get single post
- `PUT /posts/:id` — Update post
- `DELETE /posts/:id` — Delete post

## License

MIT

Then run:

git add README.md
git commit -m "Add README"
git push

