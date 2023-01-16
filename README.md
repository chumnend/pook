# Pook: An app for creating storybooks

## About this project
TBD

### Demo
TBD

### Built With
- Go (Golang) 1.18
- Gin Web Framework

## Getting Started
TBD

1) Clone this repository, `git clone https://github.com/chumnend/pook.git`

2) Create a copy of the .env.example file and rename it to .env, `cp .env.example .env`

3) Open the .env file and fill in the required fields

4) Start the server by typing the command `make run`

5) Tests can be ran by typing the command `make test`

## API Documentation

### Authentication Routes

POST ```/v1/register``` (create a new user)

- Data Payload:
  ```
    {
        "email": <string>,
        "password": <string>,
        "firstName": <string>,
        "lastName": <string>,
    }
    ```
- Success Response 
    ```
    {
        "token": <string>,
    }
    ```
- Error Response
    ```
    {
        "error": <string>,
    }
    ```


POST ```/v1/login``` (acquire jwt for authentication)
- Data Payload:
  ```
    {
        "email": <string>,
        "password": <string>,
    }
    ```
- Success Response 
    ```
    {
        "token": <string>,
    }
    ```
- Error Response
    ```
    {
        "error": <string>,
    }
    ```


### Book Routes
GET ```/v1/books?userId=<integer>``` (get all of a user's books)
- Query Params

| Parameter | Type | Notes | Required |
| --- | --- | --- | --- |
| userId | integer | Filters books by this user value | No |

- Success Response 
    ```
    {
        "books": [
            {
                "id": <integer>,
                "title": <string>,
                "createdAt": <time>,
                "updatedAt": <time>,
                "userID": <integer>,
                "pages": [
                    {
                    "id": <integer>,
                    "content": <string>,
                    "createdAt": <time>,
                    "updatedAt": <time>,
                    "bookID": <integer>,
                    },
                    (+ more)
                ],
            },
            (+ more)
        ]
    }
    ```
- Error Response
    ```
    {
        "error": <string>,
    }
    ```


POST```/v1/books``` (create a book)
- Data Payload:
  ```
    {
        "title": <string>,
        "userId": <integer>,
    }
    ```
- Success Response 
    ```
    {
        "book": {
            {
                "id": <integer>,
                "title": <string>,
                "createdAt": <time>,
                "updatedAt": <time>,
                "userID": <integer>,
                "pages": [
                    {
                        "id": <integer>,
                        "content": <string>,
                        "createdAt": <time>,
                        "updatedAt": <time>,
                        "bookID": <integer>,
                    },
                    (+ more)
                ],
            }
        }
    }
    ```
- Error Response
    ```
    {
        "error": <string>,
    }
    ```

GET ```/v1/books/<BOOK_ID>``` (get a book)
- Success Response 
    ```
    {
        "book": {
            {
                "id": <integer>,
                "title": <string>,
                "createdAt": <time>,
                "updatedAt": <time>,
                "userID": <integer>,
                "pages": [
                    {
                        "id": <integer>,
                        "content": <string>,
                        "createdAt": <time>,
                        "updatedAt": <time>,
                        "bookID": <integer>,
                    },
                    (+ more)
                ],
            }
        }
    }
    ```
- Error Response
    ```
    {
        "error": <string>,
    }
    ```
PUT ```/v1/books/<BOOK_ID>``` (update a book)
- Data Payload:
  ```
    {
        "title": <string>,
    }
    ```
- Success Response 
    ```
    {
        "book": {
            {
                "id": <integer>,
                "title": <string>,
                "createdAt": <time>,
                "updatedAt": <time>,
                "userID": <integer>,
                "pages": [
                    {
                        "id": <integer>,
                        "content": <string>,
                        "createdAt": <time>,
                        "updatedAt": <time>,
                        "bookID": <integer>,
                    },
                    (+ more)
                ],
            }
        }
    }
    ```
- Error Response
    ```
    {
        "error": <string>,
    }
    ```
DELETE ```/v1/books/<BOOK_ID>``` (delete a book)
- Success Response 
    ```
    {
        "result": <string>,
    }
    ```
- Error Response
    ```
    {
        "error": <string>,
    }
    ```

### Page Routes
GET ```/v1/pages``` (get all of a book's pages)
- Query Params

| Parameter | Type | Notes | Required |
| --- | --- | --- | --- |
| bookId | integer | the book to get pages of | Yes |

- Success Response 
    ```
    {
        "pages": [
            {
                "id": <integer>,
                "content": <string>,
                "createdAt": <time>,
                "updatedAt": <time>,
                "bookID": <integer>,
            },
            (+ more)
        ]
    }
    ```
- Error Response
    ```
    {
        "error": <string>,
    }
    ```

POST```/v1/pages``` (create a page)
- Data Payload:
  ```
    {
        "content": <string>,
        "bookId": <integer>,
    }
    ```
- Success Response 
    ```
    {
        "result": {
            "id": <integer>,
            "content": <string>,
            "createdAt": <time>,
            "updatedAt": <time>,
            "bookID": <integer>,
        },
    }
    ```
- Error Response
    ```
    {
        "error": <string>,
    }

GET ```/v1/pages/<PAGE_ID>``` (get a book's page)
- Success Response 
    ```
    {
        "result": {
            "id": <integer>,
            "content": <string>,
            "createdAt": <time>,
            "updatedAt": <time>,
            "bookID": <integer>,
        },
    }
    ```
- Error Response
    ```
    {
        "error": <string>,
    }
    ```

PUT ```/v1/pages/<PAGE_ID>``` (update a book's page)
- Data Payload:
  ```
    {
        "content": <string>,
    }
    ```
- Success Response 
    ```
    {
        "result": {
            "id": <integer>,
            "content": <string>,
            "createdAt": <time>,
            "updatedAt": <time>,
            "bookID": <integer>,
        },
    }
    ```
- Error Response
    ```
    {
        "error": <string>,
    }
    ```

DELETE ```/v1/pages/<PAGE_ID>``` (delete a book's page)
- Success Response 
    ```
    {
        "result": <string>,
    }
    ```
- Error Response
    ```
    {
        "error": <string>,
    }
    ```

## Deployment 
TBD

## Contact
Nicholas Chumney - [nicholas.chumney@outlook.com](nicholas.chumney@outlook.com) 

## Acknowledgments
- [Building and Testing a REST API in GO](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)
- [Go-clean-template: Clean Architectur](https://evrone.com/go-clean-template)
- [Create your first Go REST API with JWT Authentication in Gin Framework](https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817)