import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'

import LandingPage from './components/LandingPage';
import RegisterPage from './components/RegisterPage';
import LoginPage from './components/LoginPage';
import LibraryPage from './components/LibraryPage';
import BookCreationPage from './components/BookCreationPage';
import BookViewingPage from './components/BookViewingPage';
import UserProfilePage from './components/UserProfilePage';
import NotFoundPage from './components/NotFoundPage';

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
