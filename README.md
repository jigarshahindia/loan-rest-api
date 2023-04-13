# loan-rest-api
Rest API for Loan Application


Prerequites 
```
docker run --name loanTest -p 5432:5432 -e POSTGRES_USER=gorm -e POSTGRES_PASSWORD=gorm -e POSTGRES_DB=gorm -d postgres
```


API Sample

Post Loan
```
curl --location --request POST 'http://localhost:8080/loan' \
--header 'Authorization: Basic dGVzdDp0ZXN0' \
--header 'Content-Type: application/json' \
--data-raw '{
    "CustomerID": 1,
    "Amount": 1,
    "Status": "NEW"
}'
```


Get Loan
```
curl --location --request GET 'http://localhost:8080/loan/:id' \
--header 'Authorization: Basic dGVzdDp0ZXN0' \
--header 'Content-Type: application/json'
```