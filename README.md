# Golang-ninja-project-3

Stack:
- Golang
- Docker
- MongoDB
- AWS 
- GitHub Actions

Workflow:
- [x] Docker
docker run --rm -d --name uploader -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=123 -p 27017:27017 mongo:latest