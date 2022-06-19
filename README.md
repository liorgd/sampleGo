"# sampleGo"  

web service listening on port 3000

Has dockerfile support - Use port forward to communicate via container

POST record:

curl http://localhost:3000/records \
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"id": "4","title": "title1","content": "content2","views": 4,"timestamp": 5}'

Be sure the id is unique

GET all records: 
curl http://localhost:3000/records \
--header "Content-Type: application/json" \
--request "GET"