import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import { environment } from '../../../src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class BackendService {
  API_URI = environment.BACK_URI;
  constructor(private http: HttpClient) { }
  
  getServicesDataNap(){
    return this.http.get(`${this.API_URI}/vianet`);
  }
}
