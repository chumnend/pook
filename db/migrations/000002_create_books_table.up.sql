CREATE TABLE Books (
    book_id UUID PRIMARY KEY,
    user_id UUID REFERENCES Users(user_id),
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
