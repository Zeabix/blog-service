# Blog Service API

### Environment variables
* `MONGO_CONNNECTION_URL`  :  Connection URL to mongoDB, e.g. mongodb://localhost:27017
* `MONGO_DATABASE_NAME`   :  Mongo Database name, e.g. blogs
* `MONGO_COLLECTION_NAME`  :  Mongo Collection name, e.g. blogs


During starting up the service will try to connect to database. If everything is ok, it will continue, but if not it will shutdown.
However, if database connection is lost `AFTER` the starting up process is done, the `/healthcheck` will return `HTTP 500` (See more detail regaring healthcheck below)


### Healthcheck
* `/health` : Return `HTTP 200` with body `{'status': 'ok' }` if service is healthy and can connect to database, otherwise return `HTTP 500`