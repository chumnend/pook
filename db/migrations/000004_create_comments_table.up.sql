CREATE TABLE Comments (
    id UUID PRIMARY KEY,
    book_id UUID REFERENCES Books(id),
    user_id UUID REFERENCES Users(id),
    comment TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
