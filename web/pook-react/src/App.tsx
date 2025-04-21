import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'

import LandingPage from './pages/LandingPage';
import RegisterPage from './pages/RegisterPage';
import LoginPage from './pages/LoginPage';
import LibraryPage from './pages/LibraryPage';
import BookCreationPage from './pages/BookCreationPage';
import BookViewingPage from './pages/BookViewingPage';
import UserProfilePage from './pages/UserProfilePage';
import NotFoundPage from './pages/NotFoundPage';
import AuthProvider from './providers/AuthProvider';

function App() {
  return (
    <AuthProvider>
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
    </AuthProvider>
  )
}

export default App
