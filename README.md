# sampleGo  

##references:
- https://go.dev/doc/database/querying
- https://go.dev/doc/tutorial/web-service-gin
- https://www.section.io/engineering-education/build-a-rest-api-application-using-golang-and-postgresql-database/
- https://docs.rancherdesktop.io/tutorials/working-with-images/

## container support
web service listening on port 3000
Has dockerfile support - Use port forward to communicate via container

## db support
db details:
const (
DB_USER     = "app1"
DB_PASSWORD = "p"
DB_NAME     = "app_db"
)

port forward is mostly to 5432


## commands
# POST record:
curl http://localhost:3000/records \
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"id": "4","title": "title1","content": "content2","views": 4,"timestamp": 5}'

Be sure the id is unique

#GET all records: 
curl http://localhost:3000/records \
--header "Content-Type: application/json" \
--request "GET"

#DELETE record by id
curl http://localhost:3000/records/1 \
--include \
--header "Content-Type: application/json" \
--request "DELETE" 

##Build from docker
nerdctl --namespace go build -t sample-go:v1.0 .
kubectl run --image sample-go:v1.0 sample-go