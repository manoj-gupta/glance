# glance

## Running Glance

### Using Docker

This mode is used for debugging. You should be running postgres with database created.

- `docker build . -t go-glance`
- `docker build -f Dockerfile.dev -t go-glance .` (development)
- `docker run -p 8080:8080 go-glance`

## Using docker-compose (easiest)

- `docker-compose up`

### database shell

- `docker-compose run database bash`
