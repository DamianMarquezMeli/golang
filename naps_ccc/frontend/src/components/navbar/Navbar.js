import React, { useEffect, useState } from 'react';
import { Link, NavLink, useHistory } from 'react-router-dom';



export const Navbar = () => {

    const history = useHistory();

    const handleLogout = () => {
        sessionStorage.clear();
        history.push('/login');
    }
    


    const [state, setState] = useState(false);

    useEffect(() => {
        (sessionStorage.getItem('role') === 'Admin') ? setState(true) : setState(false);
    }, []);

    return (
        <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
            <Link to="/" className="navbar-brand">ProductoresApp</Link>
            <div className="navbar-collapse collapse w-100 order-1 order-md-0 dual-collapse2">
                <div className="navbar-nav">

                    <NavLink
                    exact
                    activeClassName="active"
                    to="/"
                    className="nav-item nav-link"
                    >
                        Mapa
                    </NavLink>

                    {
                        (state &&
                            (
                                <div className="navbar-nav">
                                    <NavLink
                                    exact
                                    activeClassName="active"
                                    to="/register"
                                    className="nav-item nav-link"
                                    >
                                        Nuevo usuario
                                    </NavLink>
                                    <NavLink
                                    exact
                                    activeClassName="active"
                                    to="/adminpanel"
                                    className="nav-item nav-link"
                                    >
                                        Panel de administración
                                    </NavLink>
                                </div>
                            )
                        )
                    }
                </div>
            </div>
            <div className="navbar-collapse collapse w-100 order-3 dual-collapse2">
                <ul className="navbar-nav ml-auto">
                    <span className="nav-item nav-link text-info">
                        {sessionStorage.getItem('username')}
                    </span>
                    <button
                    className="nav-item nav-link btn"
                    onClick={ handleLogout }
                    >
                        Logout
                    </button>
                </ul>
            </div>
        </nav>
    )
}
