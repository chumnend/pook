# Testing

Here are some curl commands to test the currently implemented routes in your project:

## User Routes

### Register a New User

```sh
curl -X POST http://localhost:8080/v1/register \
-H "Content-Type: application/json" \
-d '{
  "email": "tester@xchumz.com",
  "username": "tester",
  "password": "test123"
}'
```

### Login a User

```sh
curl -X POST http://localhost:8080/v1/login \
-H "Content-Type: application/json" \
-d '{
  "username": "tester",
  "password": "test123"
}'
```

### Get User Details

```sh
curl -X GET http://localhost:8080/v1/users/{user_id}
```

## Book Routes

### Create a New Book

```sh
curl -X POST http://localhost:8080/v1/books \
-H "Content-Type: application/json" \
-d '{
  "title": "My First Book",
  "userId": "user-uuid-here"
}'
```

### Get All Books

```sh
curl -X GET http://localhost:8080/v1/books
```

### Get Books By User ID

```sh
curl -X GET "http://localhost:8080/v1/books?user_id=user-uuid-here"
```

### Get a Specific Book

```sh
curl -X GET http://localhost:8080/v1/books/{book_id}
```

### Update a Book

```sh
curl -X PUT http://localhost:8080/v1/books/{book_id} \
-H "Content-Type: application/json" \
-d '{
  "title": "Updated Book Title"
}'
```

### Delete a Book

```sh
curl -X DELETE http://localhost:8080/v1/books/{book_id}
```

## Page Routes

### Create a Page

```sh
curl -X POST http://localhost:8080/books/{book_id}/pages \
-H "Content-Type: application/json" \
-d '{
  "imageUrl": "https://example.com/image.jpg",
  "caption": "Sample caption",
  "pageOrder": 1
}'
```

### Get All Pages of a Book

```sh
curl -X GET http://localhost:8080/books/{book_id}/pages
```

### Get a Specific Page

```sh
curl -X GET http://localhost:8080/books/{book_id}/pages/{page_id}
```

### Update a Page

```sh
curl -X PUT http://localhost:8080/books/{book_id}/pages/{page_id} \
-H "Content-Type: application/json" \
-d '{
  "imageUrl": "https://example.com/updated-image.jpg",
  "caption": "Updated caption",
  "pageOrder": 2
}'
```

### Delete a Page

```sh
curl -X DELETE http://localhost:8080/books/{book_id}/pages/{page_id}
```
