//** Modulos **//
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
//** Rutas **/
import { AppRoutingModule } from './app-routing.module';
//** Servicios e interceptores**//
import { BackendService } from "./servicios/backend.service";

//** Componentes **//
import { AppComponent } from './app.component';
import { ContentComponent } from './components/shared/content/content.component';
import { FooterComponent } from './components/shared/footer/footer.component';
import { HeaderComponent } from './components/shared/header/header.component';
import { MenuComponent } from './components/shared/menu/menu.component';
import { SettingComponent } from './components/shared/setting/setting.component';
import { MapComponent } from './components/views/map/map.component';

@NgModule({
  declarations: [
    AppComponent,
    ContentComponent,
    FooterComponent,
    HeaderComponent,
    MenuComponent,
    SettingComponent,
    MapComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [
    BackendService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
