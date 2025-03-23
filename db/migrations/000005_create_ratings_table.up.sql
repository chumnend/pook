CREATE TABLE Ratings (
    rating_id UUID PRIMARY KEY,
    book_id UUID REFERENCES Books(book_id),
    user_id UUID REFERENCES Users(user_id),
    rating INT CHECK (rating >= 1 AND rating <= 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
