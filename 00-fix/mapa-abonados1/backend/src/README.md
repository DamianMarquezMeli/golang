# Documentación General

IMPORTANTE: UNA VEZ REALIZADA LA INSTALACION, SE DEBE BORRAR LA CARPETA INSTALADOR DE SISTEMA EN SERVICIO.

## Configurar Laravel

## Editar laravel/.env

DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=ccc
DB_USERNAME=ccc_admin
DB_PASSWORD=sanlorenzo

## Instalar

Lo único que se tiene realmente configurar son:
MySQL:
    · nombre root.
    · contraseña root.
    · servidor mysql.

Datos:
    · path directorio de datos.

Si estos 3 datos son correctos, la instalacion y la carga de datos deberia correr sin problemas.

1. Se supone que para MySQL:
    usuario: root
    contraseña: rocky

   Si los datos son diferentes, actualizar la constante DB_PASSROOT en mapa_abonados/src/cfg/config.php

2. Path datos: especificar.

3. Correr instalador/instalador.php
    * Crea usuario mysql.
    * Base de datos.
    * Tablas.
    * Llena bases de datos.


Solucion de problemas:

1. 2000 - LOAD DATA LOCAL INFILE is forbidden, check mysqli.allow_local_infile
    /etc/php/7.2/cli/php.ini
    Descomentar esta linea: mysqli.allow_local_infile = On
