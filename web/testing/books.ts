import type { Book } from '../src/helpers/types';

export function createMockBooks(count: number): Book[] {
  return Array.from({ length: count }, (_, i) => ({
    id: `book-${i + 1}-${Math.random().toString(36).substr(2, 9)}`,
    userId: `user-${Math.floor(100000 + Math.random() * 900000)}`,
    title: `Book Title ${i + 1}`,
    imageUrl: "https://placehold.co/400x400",
    createdAt: new Date(Date.now() - Math.random() * 30 * 24 * 60 * 60 * 1000), // Random date within last 30 days
    updatedAt: new Date(),
  }));
}
  
const books: Book[] = [
  { 
    id: "book-1",
    userId: "user-123456",
    title: "Book Title 1", 
    imageUrl: "https://placehold.co/400x400",
    createdAt: new Date('2024-12-01'),
    updatedAt: new Date('2024-12-01')
  },
  { 
    id: "book-2",
    userId: "user-234567",
    title: "Book Title 2", 
    imageUrl: "https://placehold.co/400x400",
    createdAt: new Date('2024-12-02'),
    updatedAt: new Date('2024-12-02')
  },
  { 
    id: "book-3",
    userId: "user-345678",
    title: "Book Title 3", 
    imageUrl: "https://placehold.co/400x400",
    createdAt: new Date('2024-12-03'),
    updatedAt: new Date('2024-12-03')
  },
  { 
    id: "book-4",
    userId: "user-456789",
    title: "Book Title 4", 
    imageUrl: "https://placehold.co/400x400",
    createdAt: new Date('2024-12-04'),
    updatedAt: new Date('2024-12-04')
  },
  { 
    id: "book-5",
    userId: "user-567890",
    title: "Book Title 5", 
    imageUrl: "https://placehold.co/400x400",
    createdAt: new Date('2024-12-05'),
    updatedAt: new Date('2024-12-05')
  },
  { 
    id: "book-6",
    userId: "user-678901",
    title: "Book Title 6", 
    imageUrl: "https://placehold.co/400x400",
    createdAt: new Date('2024-12-06'),
    updatedAt: new Date('2024-12-06')
  },
  { 
    id: "book-7",
    userId: "user-789012",
    title: "Book Title 7", 
    imageUrl: "https://placehold.co/400x400",
    createdAt: new Date('2024-12-07'),
    updatedAt: new Date('2024-12-07')
  },
  { 
    id: "book-8",
    userId: "user-890123",
    title: "Book Title 8", 
    imageUrl: "https://placehold.co/400x400",
    createdAt: new Date('2024-12-08'),
    updatedAt: new Date('2024-12-08')
  },
];

export default books;
  