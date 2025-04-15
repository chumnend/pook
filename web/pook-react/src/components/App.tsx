import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'

import LandingPage from './LandingPage';
import RegisterPage from './RegisterPage';
import LoginPage from './LoginPage';
import LibraryPage from './LibraryPage';
import BookCreationPage from './BookCreationPage';
import BookViewingPage from './BookViewingPage';
import UserProfilePage from './UserProfilePage';
import NotFoundPage from './NotFoundPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/library/:userId" element={<LibraryPage />} />
        <Route path="/book/new" element={<BookCreationPage />} />
        <Route path="/book/:bookId" element={<BookViewingPage />} />
        <Route path="/user/:userId" element={<UserProfilePage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </Router>
  )
}

export default App
