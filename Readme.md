## Setup

#### Type in "sh start.sh" - boots up your Go and MYSQL environment

#### *localhost:8080/articles [GET]

Input (N/A):
```json
{}
```

Output:
```json
{
  "status": 200,
  "message": "Success",
  "data": [
    {
      "id": 2,
      "title": "Serpent's Stone",
      "content": "Tanz",
      "author": "J.K Rowling"
    }
  ]
}
```

#### *localhost:8080/articles [POST]

Input:
```json
{
  "title": "Serpent's Stone",
  "author": "J.K Rowling",
  "content": " Tanz"
}
```

Output:
```json
{
  "status": 201,
  "message": "Success",
  "data": {
    "id": 2
  }
}
```