import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'

import ProtectedRoute from './ProtectedRoute';
import LandingPage from './pages/LandingPage';
import RegisterPage from './pages/RegisterPage';
import LoginPage from './pages/LoginPage';
import LogoutPage from './pages/LogoutPage';
import LibraryPage from './pages/LibraryPage';
import BookCreationPage from './pages/BookCreationPage';
import BookViewingPage from './pages/BookViewingPage';
import UserProfilePage from './pages/UserProfilePage';
import NotFoundPage from './pages/NotFoundPage';
import useAuth from '../helpers/hooks/useAuth';

function App() {
  const { isLoggedIn, user } = useAuth()

  return (
    <Router>
      <Routes>
        {/* Authenticated Routes */}
        <Route 
          path="/book/new"
          element={
            <ProtectedRoute
              condition={isLoggedIn}
              redirect='/login'
            >
              <BookCreationPage />
            </ProtectedRoute>
          } 
        />
        <Route 
          path="/logout" 
          element={
            <ProtectedRoute
              condition={isLoggedIn}
              redirect='/login'
            >
              <LogoutPage />
            </ProtectedRoute>
          }
        />
        {/* Unauthenticated Routes */}
        <Route 
          path="/" 
          element={
            <ProtectedRoute
              condition={!isLoggedIn}
              redirect={`/library/${user?.id}`}
            >
              <LandingPage />
            </ProtectedRoute>
          } 
        />
        <Route 
          path="/register" 
          element={
            <ProtectedRoute
              condition={!isLoggedIn}
              redirect={`/library/${user?.id}`}
            >
              <RegisterPage />
            </ProtectedRoute>
          } 
        />
        <Route 
          path="/login" 
          element={
            <ProtectedRoute
              condition={!isLoggedIn}
              redirect={`/library/${user?.id}`}
            >
              <LoginPage />
            </ProtectedRoute>
          }
        />
        {/* Public Routes */}
        <Route path="/library/:userId" element={<LibraryPage />} />
        <Route path="/book/:bookId" element={<BookViewingPage />} />
        <Route path="/user/:userId" element={<UserProfilePage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </Router>
  )
}

export default App
