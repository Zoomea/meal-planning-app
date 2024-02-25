
curl 127.0.0.1:8080/api/recipe/ \
	-d '{
  "name": "test1",
  "ingredients": [
    {
      "id": 1,
      "name": "cumin",
      "quantity": 1,
      "unit": "gram"
    }
  ]
}'

curl 127.0.0.1:8080/api/recipe/

curl 127.0.0.1:8080/api/recipe/3

curl -XDELETE 127.0.0.1:8080/api/recipe/4
