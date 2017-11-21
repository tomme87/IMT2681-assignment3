# Cloud Technologies - Assignement 3
This is a school assignment.

# Install (without docker)
You need to have the code running for assignment 2.

Run this (and optionally add it to you .bashrc)
```
export PORT=8080
LATEST_URL=http://assignment2/exhange/latest
```

Fetch and install
```
go get github.com/tomme87/IMT2681-assignment3/cmd/webapp
go install github.com/tomme87/IMT2681-assignment3/cmd/webapp
```

Run
```
webapp
```

# Docker
We need a .env file in directory api, bot and monitor.

example for api:
```
PORT=8080
MGO_URL=mongodb://db
```

ecample for bot:
```
PORT=8081
LATEST_URL=http://api:8080/exchange/latest
```

example for monitor:
```
MGO_URL=mongodb://db
```

### Run
```
docker-compose up -d
```
In exchange_docker directory. (`docker-compose down` to stop)
