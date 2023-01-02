import { Marker, Popup } from 'react-leaflet';
import React, { useEffect, useState } from 'react';
import { getNaps } from '../helpers/getNaps';

export const Markers = () => {

    const [naps, setNaps] = useState([]);

    useEffect(() => {
        getNaps()
            .then(data => setNaps(data));
    }, []);

    return naps.map((nap) => (
        <Marker key={nap.id_nodo_arbol} position={[nap.latitud,nap.longitud]}>
            <Popup>
                {nap.descripcion}<br/>
                DISPONIBLES: {nap.disponibles}
            </Popup>           
        </Marker>
        )
    );
}