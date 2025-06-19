import type { Book } from '../src/helpers/types';

export function createMockBooks(count: number): Book[] {
  return Array.from({ length: count }, (_, i) => ({
    title: `Book Title ${i + 1}`,
    author: `User${Math.floor(100000 + Math.random() * 900000)}`,
    imageUrl: "https://placehold.co/400x400",
  }));
}
  
const books = [
  { title: "Book Title 1", author: "User234313", imageUrl: "https://placehold.co/400x400" },
  { title: "Book Title 2", author: "User234313", imageUrl: "https://placehold.co/400x400" },
  { title: "Book Title 3", author: "User234313", imageUrl: "https://placehold.co/400x400" },
  { title: "Book Title 4", author: "User234313", imageUrl: "https://placehold.co/400x400" },
  { title: "Book Title 5", author: "User234313", imageUrl: "https://placehold.co/400x400" },
  { title: "Book Title 6", author: "User234313", imageUrl: "https://placehold.co/400x400" },
  { title: "Book Title 7", author: "User234313", imageUrl: "https://placehold.co/400x400" },
  { title: "Book Title 8", author: "User234313", imageUrl: "https://placehold.co/400x400" },
];

export default books;
  