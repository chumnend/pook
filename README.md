# Pook: An app for creating storybooks
Pook is a storybook creator application using React and Go. The idea if this app is to allow users to create fun novels in thier browser and share them with others.

## Demo
TBD

## Getting Started
### Prerequisites
- Node
- Go
- PostgreSQL

### Configuration
1) Clone this repo using `git clone https://github.com/chumnend/pook.git`

2) Create a postgresql database locally or online. The database will be connected through using a connection string. It should have the form of 
`postgresql://<username>:<password>@<address>/<dbname>`

3) Run `cp .env.example .env` and then open the `.env` and fill out the following fields,
```
PORT= # the port the app will run on
SECRET_KEY= # string used for hashing
DATABASE_URL= # database string used to connect to postgresql database
DATABASE_TEST_URL= # database string to database used for integration tests
```

4) This step is only needed if you plan to work on the React app. Go into the react folder ie. `cd react/` and copy run `cp .env.example .env` and then open the `.env` and fill out the following fields,

```
NODE_ENV= # production, dev or test
PORT= # the port the React app can run on, should be different than the previous step
BROWSER=none # by default stops the browser from opening when running thr React app

REACT_APP_API_PREFIX= # this should point to the address the Go app runs on
REACT_APP_SENTRY_DSN= # needed for configuration with sentry
```

5) Now the apps are ready to run. Go back to the root folder and run `make` this will build the React and Go assets and start the app on the given port. You can build assets on thier own using the `make build` command and just serve currently built assets using `make serve`.

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

## Contact
Nicholas Chumney - [nicholas.chumney@outlook.com](nicholas.chumney@outlook.com) 

## Acknowledgments
- [Building and Testing a REST API in GO](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)
