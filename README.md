# go-sample
Sample Dockerized Go Application 

- [go-sample](#go-sample)
  - [Dependencies](#dependencies)
  - [Building the application](#building-the-application)
  - [Running the application](#running-the-application)
  - [Verifying the application](#verifying-the-application)
  - [Pushing image to registry](#pushing-image-to-registry)
  - [Cleanup](#cleanup)

## Dependencies

Git, Docker

## Building the application

Step 1: Download the repository and build image

```bash
git clone https://github.com/maroskukan/go-sample
cd go-sample/src
docker build -t go-sample .
```

Step 2: Verify the image

```bash
docker image ls
REPOSITORY                    TAG        IMAGE ID       CREATED          SIZE
go-sample                     latest     6d7e28a54cfd   21 seconds ago   845MB
```

## Running the application

Step 1: Run the container

```bash
docker run -d -p 5000:5000 --name go-web go-sample
```

Step 2: Verify the container

```bash
docker ps -f name=go-web
```

## Verifying the application

Step 1: Verify the application using curl

```bash
curl http://localhost:5000

Hello from 4c8581f38fc3
```

## Pushing image to registry

Step 1: Login to Docker Hub

```bash
docker login
```

Step 2: Tag local image

```bash
docker tag go-sample maroskukan/go-sample
```

Step 3: Push to Docker Hub

```bash
docker push maroskukan/go-sample
```

Step 4: Verify the remote image (may take a couple of minutes to get indexed)

```bash
docker search maroskukan/go-sample
```

## Cleanup

Step 1: Stop and remove the container

```bash
docker rm -f go-web
```

Step 2: Optionally, remove the created image

```bash
docker rmi go-sample
```