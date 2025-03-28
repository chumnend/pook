# Design

## Usage

The user should visit pook.com and be presented with a landing page, showing a Hero slideshow of books. From this page, the
user can register / login and after successful login be presented to their "library" and a feed of other's librarires. From this page
they can view other's libraries or create a new "book". In the book creation process the user can upload an image and then add captions. The user should be able to rearrange the order of the "pages" in the book. In book viewing the user can flip through pages of a selected book and then leave a comment and rating for the book.

## Proposed Project Structure

```sh
pook
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── handlers.go  # HTTP request handlers
│   ├── models
│   │   └── models.go    # Data structures and models
│   └── routes
│       └── routes.go    # Application routes setup
├── pkg
│   └── utils
│       └── utils.go     # Utility functions
├── web
|    |_ react            # Frontend application
├── go.mod               # Module definition and dependencies
└── README.md            # Project documentation
```

## Proposed Database Schema

```sh
-- Users table to store user information
CREATE TABLE Users (
    id UUID PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Books table to store book information
CREATE TABLE Books (
    id UUID PRIMARY KEY,
    user_id INT REFERENCES Users(user_id),
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- Pages table to store pages of a book
CREATE TABLE Pages (
    id UUID PRIMARY KEY,
    book_id INT REFERENCES Books(book_id),
    image_url VARCHAR(255) NOT NULL,
    caption TEXT,
    page_order INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- Comments table to store comments on books
CREATE TABLE Comments (
    id UUID PRIMARY KEY,
    book_id INT REFERENCES Books(book_id),
    user_id INT REFERENCES Users(user_id),
    comment TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Ratings table to store ratings for books
CREATE TABLE Ratings (
    id UUID PRIMARY KEY,
    book_id INT REFERENCES Books(book_id),
    user_id INT REFERENCES Users(user_id),
    rating INT CHECK (rating >= 1 AND rating <= 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Suggested Client Routs

Based on the current design, here are the client routes that might need to be created:

### Landing Page - `/`

Displays a Hero slideshow of books.
Provides options to register or login.

### Registration Page - `/register`

Form for new users to register.

### Login Page - `/login`

Form for existing users to login.

### User Library Page - `/library/<userId>`

Displays the user's library of books.
Shows a feed of other users' libraries.

### Book Creation Page - `/book/new`

Form to create a new book.
Allows uploading images and adding captions.
Provides functionality to rearrange the order of pages.

### Book Viewing Page - `/book/<bookId>?page=<pageId>`

Allows users to flip through the pages of a selected book.
Provides options to leave comments and ratings.

### User Profile Page - `/user/<userId>`

Displays user details and their library.

These pages cover the main functionalities described in the design document.

## Suggested CRUD Routes

Based on the current design, here are the CRUD routes that might need to be created:

### Auth / User Routes

- POST /register - Register a new user
- POST /login - Login a user
- GET /users/{user_id} - Get user details

### Book Routes

- GET /books - Get a list of books
- POST /books - Create a new book
- GET /books/{book_id} - Get details of a specific book
- PUT /books/{book_id} - Update a book
- DELETE /books/{book_id} - Delete a book

### Page Routes

- POST /books/{book_id}/pages - Add a new page to a book
- GET /books/{book_id}/pages - Get pages of a book
- PUT /books/{book_id}/pages/{page_id} - Update a page
- DELETE /books/{book_id}/pages/{page_id} - Delete a page

### Comment Routes

- POST /books/{book_id}/comments - Add a comment to a book
- GET /books/{book_id}/comments - Get comments of a book

### Rating Routes

- POST /books/{book_id}/ratings - Add a rating to a book
- GET /books/{book_id}/ratings - Get ratings of a book

These routes cover the basic CRUD operations for users, books, pages, comments, and ratings as described in the design document.
