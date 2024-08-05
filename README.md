# Golang Demo

[🐋 Compose documentation 📝](https://docs.docker.com/compose/)

## Prompt

```
SYSTEM:
As an expert in Go and Docker, generate a Dockerfile and Docker Compose file for a typical Go project.

Regarding the Dockerfile:
1. Use an official Go base image (I use go version '1.22.1')
2. Use the /app working directory
3. Copy the project source files
4. Compile the Go application
5. Create a lightweight final image for execution 
    - Use the /app working directory
    - Copy the binary in the /app directory
    - Copy the static assets (and directory structure) from /public to /app/public
6. Start the Golang application

Regarding the Compose file:
The Golang application is using a Redis database
1. Add a redis service with "redis-server" as name to the compose file.
    - The redis service is listening on the default port
    - The data files are stored in the /data directory
    - Map a volume on this directory
2. Add a webapp service 
    - The webapp service uses an environment variable (REDIS_URL)  to connect to the "redis-server" service
    - To set the value of REDIS_URL, use only the name of the redis service and the default redis port
    - The webapp service uses these two environment variables: MESSAGE and TITLE
    - The webapp service is listening on the 8080 HTTP port
    - the webapp service depends on the redis-server service
        
Ensure the Dockerfile and the Compose file are well-commented and follows best practices. 
Briefly explain each step after the Dockerfile and the Compose file.

USER:
Here's the project structure:
.
├── go.mod
├── go.sum
├── main.go
├── data
│  ├── load.sh
│  └── read.sh
├── public
│  ├── components
│  │  ├── App.js
│  │  ├── Footer.js
│  │  ├── Header.js
│  │  └── Main.js
│  ├── css
│  │  └── pico.min.css
│  ├── index.html
│  ├── info.txt
│  └── js
│     └── preact-htm.js
└── README.md
```
