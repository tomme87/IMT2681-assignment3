# Cloud Technologies - Assignement 2
This is a school assignment.

# Install
```
export PORT=8080
```

# Docker
We need a .env file in directory api, bot and monitor.

example:
```
PORT=8080
MGO_URL=mongodb://db
LATEST_URL=http://api:8080/exchange/latest
```

## Run
```
docker-compose up
```
In exchange_docker directory.