import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

import useAuth from '../../../helpers/hooks/useAuth';


const LogoutPage = () => {
    const { logout } = useAuth();
    const navigate = useNavigate();

    useEffect(() => {
        logout();
        navigate('/login');
    }, [logout, navigate]);

    return (
        <div>
            <p>Logging out...</p>
        </div>
    )

}

export default LogoutPage;
