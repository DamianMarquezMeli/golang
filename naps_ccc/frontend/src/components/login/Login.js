import React from 'react';
import {
    useHistory
  } from "react-router-dom";
import './login.css';
import Logo from '../../images/Logo.png';
import { useForm } from '../hooks/useForm';

export const Login = () => {

    let history = useHistory();
    

    const [formValues, handleInputChange] = useForm({
        username: '',
        password: ''
    });

    const { username, password } = formValues;

    const handleSubmit = (event) => {
        event.preventDefault();
        console.log('submitting!');
    }
    
    const prueba = () => {
        sessionStorage.setItem('role','Admin');
        sessionStorage.setItem('username','Emir');
        history.push('/');
    }
    

    return (
        <>
            <div className="container">
                <div className="card" style={{ width: '18rem', border: 'none' }}>
                    <img src={Logo} className="card-img-top" alt="Logo" style={{background: '#c7c7c7'}}/>
                    <div className="card-body" style={{ color: 'black', textAlign: 'center' }}>
                        <h2 className="card-title">Iniciar sesión</h2>
                        <form>
                            <div className="form-group">
                                <label htmlFor="username">Nombre de usuario</label>
                                <input
                                    type="text"
                                    className="form-control"
                                    name="username"
                                    placeholder="Ingrese su nombre de usuario"
                                    onChange={handleInputChange}
                                    value={username}
                                />
                            </div>
                            <div className="form-group">
                                <label htmlFor="password">Contraseña</label>
                                <input
                                    type="password"
                                    className="form-control"
                                    name="password"
                                    placeholder="Contraseña"
                                    onChange={handleInputChange}
                                    value={password}
                                />
                            </div>
                            <button type="submit" onSubmit={handleSubmit} className="btn btn-secondary">Ingresar</button>
                            
                                <button className="btn btn-light" style={{marginLeft: '20px'}} onClick={prueba}>Prueba</button>
                            
                        </form>
                    </div>
                </div>
            </div>
        </>
    )
}
