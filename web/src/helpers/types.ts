export type User = {
  id: string;
  username: string;
  email: string;
  token: string;
}

export type Book = {
  id: string;
  userId: string;
  imageUrl: string;
  title: string;
  createdAt: Date;
  updatedAt: Date;
}  
