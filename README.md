# go-sample
Sample Dockerized Go Application 

## Dependencies

Git, Docker

## Running the application

Step 1: Download the repository and build image
```
git clone https://github.com/maroskukan/go-sample
cd go-sample/src
docker build -t .
```

Step 2: Verify the image
```
docker image ls
REPOSITORY                    TAG        IMAGE ID       CREATED          SIZE
go-sample                     latest     6d7e28a54cfd   21 seconds ago   845MB
```

Step 3: Run the container
```
docker run -d -p 5000:5000 --name go-web go-sample
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