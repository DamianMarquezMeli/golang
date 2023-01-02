<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <link rel="shortcut icon" type="image/ico" href="http://www.datatables.net/favicon.ico">
    <meta name="viewport" content="initial-scale=1.0, maximum-scale=2.0">
    <title>Información Instalaciones</title>
    <link rel="stylesheet" type="text/css" href="DataTables/datatables.css">
</head>
<body>
<table id="uhfapp" class="display" style="width:80%">
    <h1>Información Instalaciones</h1>

    <?php

        /**
         * 
        **/
        header('Content-Type: text/html; charset=UTF-8');
        ini_set("display_errors", "On");
        error_reporting(E_ALL | E_STRICT);
        header("Content-Type: text/html; charset=UTF-8");
        date_default_timezone_set('America/Argentina/Tucuman');
        setlocale(LC_ALL, 'es-AR');

        require_once('config.php');
        include('../backend/inicio.php');

        $bd = new BaseDeDatos("ccc");

        $moviles = $bd->fetch_all("SELECT DISTINCT * FROM ccc.moviles ORDER BY nombre_movil ASC");
                
        echo "<hr>";
        echo "<hr>";
        echo "<h2>Lista de moviles: </h2>";
        for($i=0; $i<count($moviles); $i++) {
            echo $moviles[$i]['nombre_movil'];
            echo "<br>";
        }
        echo "<hr>";
        echo "<hr>";

        $cantidad_instalaciones = $bd->num_rows("SELECT DISTINCT id_i_servicios FROM ccc.instalaciones_servicios");
        echo "<h2>Cantidad de total de instalaciones registradas: {$cantidad_instalaciones} </h2>";
        echo "<hr>";
        echo "<hr>";

        echo "<h2>Cantidad total de instalaciones por movil:</h2>";
        for($i=0; $i<count($moviles); $i++) {
            $inst_x_movil = $bd->num_rows("SELECT DISTINCT id_i_servicios FROM ccc.instalaciones_servicios WHERE id_movil = {$moviles[$i]['id_movil']}");        
            echo "Movil: {$moviles[$i]['nombre_movil']} - Instalaciones: {$inst_x_movil}";
            echo "<br>";
        }
        echo "<hr>";
        echo "<hr>";

        echo "<h2>Cantidad total de instalaciones por fecha y por movil:</h2>";
        echo "<hr>";
        $inst_x_fecha = $bd->fetch_all("SELECT DISTINCT fecha_i_servicios FROM ccc.instalaciones_servicios ORDER BY fecha_i_servicios DESC");
        for($i=0; $i<count($inst_x_fecha); $i++) { 

            echo "Fecha: ".gmdate("d/m/Y", $inst_x_fecha[$i]['fecha_i_servicios']);
            echo "<br>"; 

            for($j=0; $j<count($moviles); $j++) {  
                $inst_x_movil_x_fecha = $bd->num_rows(" SELECT DISTINCT id_i_servicios 
                                                        FROM ccc.instalaciones_servicios 
                                                        WHERE id_movil = {$moviles[$j]['id_movil']} 
                                                        AND fecha_i_servicios = {$inst_x_fecha[$i]['fecha_i_servicios']}"); 
                if($inst_x_movil_x_fecha > 0) {
                    echo "Movil: {$moviles[$j]['nombre_movil']} - Instalaciones: {$inst_x_movil_x_fecha}";  
                    echo "<br>";
                }
            }
            echo "<hr>";
        }
        echo "<br>";
    ?>

    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>        
    <script type="text/javascript" charset="utf8" src="DataTables/datatables.js"></script>
    <script type="text/javascript" language="javascript" class="init" src="js/conf.js"></script>
</body>
</html>