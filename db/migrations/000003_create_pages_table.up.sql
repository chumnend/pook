CREATE TABLE Pages (
    page_id UUID PRIMARY KEY,
    book_id UUID REFERENCES Books(book_id),
    image_url VARCHAR(255) NOT NULL,
    caption TEXT,
    page_order INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
