import React from 'react';
import {
    Link
  } from "react-router-dom";
import './register.css';
import Logo from '../../images/Logo.png';
import { useForm } from '../hooks/useForm';

export const Register = () => {

    const [formValues, handleInputChange] = useForm({
        username: '',
        password: '',
        fullname: '',
        phone: ''
    });

    const { username, password, fullname, phone } = formValues;

    const handleSubmit = (event) => {
        event.preventDefault();
        console.log('submitting!');
    }


    return (
        <>
            <div className="container">
                <div className="card" style={{ width: '18rem', border: 'none' }}>
                    <img src={Logo} className="card-img-top" alt="Logo" style={{background: '#c7c7c7'}}/>
                    <div className="card-body" style={{ color: 'black', textAlign: 'center' }}>
                        <h2 className="card-title">Registrar</h2>
                        <form>
                            <div className="form-group">
                                <label htmlFor="username">Nombre de usuario</label>
                                <input
                                    type="text"
                                    className="form-control"
                                    name="username"
                                    placeholder="Ingrese el nombre de usuario"
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
                            <div className="form-group">
                                <label htmlFor="fullname">Nombre completo</label>
                                <input
                                    type="text"
                                    className="form-control"
                                    name="fullname"
                                    placeholder="Ingrese el nombre completo"
                                    onChange={handleInputChange}
                                    value={fullname}
                                />
                            </div>
                            <div className="form-group">
                                <label htmlFor="phone">Telefono</label>
                                <input
                                    type="text"
                                    className="form-control"
                                    name="phone"
                                    placeholder="Ingrese el teléfono"
                                    onChange={handleInputChange}
                                    value={phone}
                                />
                            </div>
                            <button onSubmit={handleSubmit} className="btn btn-secondary">Enviar</button>
                            <Link to='/'>
                                <button className="btn btn-light" style={{marginLeft: '20px'}}>Cancelar</button>
                            </Link>
                        </form>
                    </div>
                </div>
            </div>
        </>
    )
}
