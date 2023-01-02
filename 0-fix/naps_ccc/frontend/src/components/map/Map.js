import React from 'react';
import { MapContainer, TileLayer } from 'react-leaflet';
import { Markers } from './Markers';


export const Map = () => {
    return (

        <MapContainer center={[-26.83567, -65.20838]} zoom={18} scrollWheelZoom={true} style={{ height: '100vh' }}>
            <TileLayer
                attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
                url="http://172.30.0.141:8080/tile/{z}/{x}/{y}.png"
            />
            <Markers />
        </MapContainer>

    )
}
