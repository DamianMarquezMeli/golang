import React, { useState } from 'react';
import BootstrapTable from 'react-bootstrap-table-next';
import { useHistory } from 'react-router-dom';
import './adminpanel.css';
import {data} from './data.json';

export const Adminpanel = () => {

    const history = useHistory();

    const [userInfo, setUserInfo] = useState({});
    const [selected, setSelected] = useState(false);
    
    const selectRow = {
        mode: 'radio',
        clickToSelect: true,
        hideSelectAll: true,
        onSelect: (row) => {
            setUserInfo(row);
            console.log(userInfo);
            setSelected(true);
        }
    };


    const handleClick = () => {
            sessionStorage.setItem('user',JSON.stringify(userInfo));
            history.push('/edit');
    }
    

    const columns = [
    {
        dataField: 'nombre',
        text: 'Nombre',
        sort: true
    },
    {
        dataField: 'usuario',
        text: 'Usuario'
    },
    {
        dataField: 'contraseña',
        text: 'Contraseña'
    },
    {
        dataField: 'telefono',
        text: 'Teléfono'
    }
]

    return (
        
        <div className="container">

            <BootstrapTable
            keyField='nombre'
            data={data}
            columns={columns}
            selectRow={ selectRow }
            headerClasses='table-text'
            bodyClasses='table-text'
            />
            
            {
            selected
            ?
            <div>
                <button type="button" className="btn btn-warning" onClick={handleClick}>Editar seleccionado</button>
                <button type="button" className="btn btn-danger" style={{marginLeft: '10px'}}>Eliminar seleccionado</button>
            </div>
            :
            <div>
                <button type="button" className="btn btn-warning" disabled>Editar seleccionado</button>
                <button type="button" className="btn btn-danger" disabled style={{marginLeft: '10px'}}>Eliminar seleccionado</button>
            </div>
            }
            
        </div>
    
    )
}
