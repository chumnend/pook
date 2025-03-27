CREATE TABLE Ratings (
    id UUID PRIMARY KEY,
    book_id UUID REFERENCES Books(id),
    user_id UUID REFERENCES Users(id),
    rating INT CHECK (rating >= 1 AND rating <= 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
