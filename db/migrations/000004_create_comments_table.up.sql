CREATE TABLE Comments (
    comment_id UUID PRIMARY KEY,
    book_id UUID REFERENCES Books(book_id),
    user_id UUID REFERENCES Users(user_id),
    comment TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
