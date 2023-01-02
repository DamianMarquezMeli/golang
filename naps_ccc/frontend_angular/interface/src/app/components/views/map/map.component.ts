import { Component, OnInit, AfterViewInit } from '@angular/core';
import { BackendService } from "src/app/servicios/backend.service";
import { DataNap } from "src/app/models/data-nap";
import * as L from 'leaflet';

const iconRetinaUrl = 'assets/marker-icon-2x.png';
const iconUrl = 'assets/marker-icon.png';
const shadowUrl = 'assets/marker-shadow.png';
const iconDefault = L.icon({
  iconRetinaUrl,
  iconUrl,
  shadowUrl,
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [1, -34],
  tooltipAnchor: [16, -28],
  shadowSize: [41, 41]
});
L.Marker.prototype.options.icon = iconDefault;


@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css']
})
export class MapComponent implements OnInit, AfterViewInit {
  private map;
  _dataNap: DataNap[];
  constructor(private backservice: BackendService) { }
  private initMap(): void {
      this.map = L.map('map', {
        center: [ -26.80, -65.20 ],
        zoom: 14
      });
      //'http://172.30.0.141:8080/{z}/{x}/{y}.png'
      const tiles = L.tileLayer("http://172.30.0.141:8080/tile/{z}/{x}/{y}.png", {
      maxZoom: 18,
      minZoom: 3,
      attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    });
    tiles.addTo(this.map);
    this.map.invalidateSize()
    }
  ngOnInit(): void {
    let resu: any;
    this.backservice.getServicesDataNap().subscribe(
      res => {
        resu = res;
        this._dataNap= resu.data
        console.log('Valor devuelto: ', resu.data)
        this.makeNapPoints(this._dataNap)
      },
      err=>{
        console.log('Error:',err)
      }
    )

  }
  
  ngAfterViewInit(): void {
    this.initMap();
    
    
  }
  private makeNapPoints(naps: DataNap[]): void {
      for (const c of naps) {
        let info =  `` +
        `<div>Nombre: ${ c.descripcion }</div>` +
        `<div>Disponibles: ${ c.disponibles }</div>`
        const lon = c.longitud;
        const lat = c.latitud;
        const marker = L.marker([lat, lon]);
        marker.bindPopup(info)
        marker.addTo(this.map);
      }
  }
}
