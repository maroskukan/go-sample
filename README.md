# go-sample
Sample Dockerized Go Application 

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
```
git clone https://github.com/maroskukan/go-sample
cd go-sample/src
docker build -t go-sample .
```

Step 2: Verify the image
```
docker image ls
REPOSITORY                    TAG        IMAGE ID       CREATED          SIZE
go-sample                     latest     6d7e28a54cfd   21 seconds ago   845MB
```

## Running the application

Step 1: Run the container
```
docker run -d -p 5000:5000 --name go-web go-sample
```

Step 2: Verify the container
```
docker ps -f name=go-web
```

## Verifying the application

Step 1: Verify the application using curl
```
curl http://localhost:5000
StatusCode        : 200
StatusDescription : OK
Content           : Hello World
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 11
                    Content-Type: text/plain; charset=utf-8
                    Date: Tue, 29 Dec 2020 09:11:31 GMT

                    Hello World
Forms             : {}
Headers           : {[Content-Length, 11], [Content-Type, text/plain; charset=utf-8], [Date, Tue, 29 Dec 2020 09:11:31
                    GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 11
```

## Pushing image to registry

Step 1: Login to Docker Hub
```
docker login
```

Step 2: Tag local image
```
docker tag go-sample maroskukan/go-sample
```

Step 3: Push to Docker Hub
```
docker push maroskukan/go-sample
```

Step 4: Verify the remote image (may take a couple of minutes to get indexed)
```
docker search maroskukan/go-sample
```

## Cleanup

Step 1: Stop and remove the container
```
docker rm -f go-web
```

Step 2: Optionally, remove the created image
```
docker rmi go-sample
```