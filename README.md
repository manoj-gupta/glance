# glance

## Running Glance

### Using Docker

This mode is used for debugging. You should be running postgres with database created.

```
docker build . -t go-glance
docker build -f Dockerfile.dev -t go-glance . (development)
docker run -p 8080:8080 go-glance
```

## Using docker-compose (easiest)

- Start

```
docker-compose up -d --build
docker-compose -f docker-compose-dev.yml up -d --build (development)
```

- Stop

```
docker-compose down
docker-compose -f docker-compose-dev.yml down (development)
```

### database shell

```
docker-compose run database bash
```

# Installing docker and docker-compose on Raspberry Pi

1. Install Docker

```
curl -sSL https://get.docker.com | sh
```

2. Add permission to Pi User to run docker Commands without sudo

```
sudo usermod -aG docker pi
```

3. reboot raspbeery pi to run the docker commands without sudo

4. Test Docker installation

```
docker run hello-world
```

5. Install dependencies

```
sudo apt-get install -y libffi-dev libssl-dev

sudo apt-get install -y python3 python3-pip

sudo apt-get remove python-configparser
```

6. Install Docker Compose

```
sudo pip3 -v install docker-compose
```
