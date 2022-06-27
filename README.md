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
nerdctl --namespace k8s.io build -t sample-go:v1.0 .

##run image in k8s cluser
kubectl apply -f k8s/pod.yml
kubectl apply -f k8s/sbc-lb.yml

##GraphQL 
https://gqlgen.com/getting-started/
postman related commands in postman folder

To generate code from schema:
go generate ./...

mutation deleterecord {
DeleteRecord(input: { ID: "4" }) {
ID    
}
}

query record {
Record(input: { ID: "3" }) {
ID    
}
}

query {
Records  {
ID,
 title,
Content,
Views,
Timestamp
}   
}