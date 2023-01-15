import React from 'react';
import {
    Link
  } from "react-router-dom";
import './edit.css';
import Logo from '../../images/Logo.png';
import { useForm } from '../hooks/useForm';

export const Edit = () => {

    const {usuario, nombre, telefono} = JSON.parse(sessionStorage.getItem('user'));

    const [formValues, handleInputChange] = useForm({
        username: usuario,
        fullname: nombre,
        phone: telefono
    });

    const { username, fullname, phone } = formValues;

    const handleSubmit = (event) => {
        event.preventDefault();
        console.log('submitting!');
        // sessionStorage.setItem('useredit',JSON.stringify(formValues));
    }


    return (
        <>
            <div className="container">
                <div className="card" style={{ width: '18rem', border: 'none' }}>
                    <img src={Logo} className="card-img-top" alt="Logo" style={{background: '#c7c7c7'}}/>
                    <div className="card-body" style={{ color: 'black', textAlign: 'center' }}>
                        <h2 className="card-title">Editando a {nombre}</h2>
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
                            <Link to='/adminpanel'>
                                <button className="btn btn-light" style={{marginLeft: '20px'}}>Cancelar</button>
                            </Link>
                        </form>
                    </div>
                </div>
            </div>
        </>
    )
}
