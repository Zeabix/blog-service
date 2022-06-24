# Blog Service API

### Environment variables
* `MONGO_CONNNECTION_URL`  :  Connection URL to mongoDB, e.g. mongodb://localhost:27017
* `MONGO_DATABASE_NAME`   :  Mongo Database name, e.g. blogs
* `MONGO_COLLECTION_NAME`  :  Mongo Collection name, e.g. blogs


During starting up the service will try to connect to database. If everything is ok, it will continue, but if not it will shutdown.
However, if database connection is lost `AFTER` the starting up process is done, the `/healthcheck` will return `HTTP 500` (See more detail regaring healthcheck below)


### Healthcheck
* `/health` : Return `HTTP 200` with body `{'status': 'ok' }` if service is healthy and can connect to database, otherwise return `HTTP 500`


### Metrics
* `/metrics` : is the endpoint that you should configure Prometheus to scrape the metrics information. This microservice export `api_blog_service_request_count` and `api_blog_service_request_latency_microseconds_sum` . In order to make it available you might need to call the apis a couple times before they are available

Example

```
# HELP api_blog_service_request_count number of requests recieved
# TYPE api_blog_service_request_count counter
api_blog_service_request_count{method="create_blog"} 1
api_blog_service_request_count{method="list_blogs"} 6
api_blog_service_request_count{method="publis_blog"} 1
# HELP api_blog_service_request_latency_microseconds Total duration of requests in microseconds.
# TYPE api_blog_service_request_latency_microseconds summary
api_blog_service_request_latency_microseconds_sum{method="create_blog"} 0.012395208
api_blog_service_request_latency_microseconds_count{method="create_blog"} 1
api_blog_service_request_latency_microseconds_sum{method="list_blogs"} 0.014551166
api_blog_service_request_latency_microseconds_count{method="list_blogs"} 6
api_blog_service_request_latency_microseconds_sum{method="publish_blog"} 0.015850583
api_blog_service_request_latency_microseconds_count{method="publish_blog"} 1
```


### API
#### Create Blog API

```
curl -XPOST -H 'Content-type: application/json' http://localhost:8080/blogs/v1/blogs -d '{"topic": "Test4", "content": "This is a test blog post 4", "author": "Prawit Chaivong"}' -vvv
```

#### Get Blog by ID
`curl http://localhost:8080/blogs/v1/blogs/{id}`

#### List All blog posts
`curl http://localhost:8080/blogs/v1/blogs`

#### Publish blog posts
`curl -XPUT http//localhost:8080/blogs/v1/blogs/{id}`