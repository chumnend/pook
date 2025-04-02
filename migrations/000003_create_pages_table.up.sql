CREATE TABLE pages (
    id UUID NOT NULL UNIQUE,
    book_id UUID NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    caption TEXT,
    page_order INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    UNIQUE (book_id, page_order)
);
