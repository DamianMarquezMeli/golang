<?php

	header('Content-Type: text/html; charset=UTF-8');
	ini_set("display_errors", "On");
	error_reporting(E_ALL | E_STRICT);
	header("Content-Type: text/html; charset=UTF-8");
	date_default_timezone_set('America/Argentina/Tucuman');
	setlocale(LC_ALL, 'es-AR');

	//require_once("config/config_mapa_abonados.php");
	require_once("config/config.php");

	$bd = new BaseDeDatos("ccc");
	$info_mapa = [];
	$fotos  = "";
	$eleminar_del_path_foto = "/home/pablo/Proyectos/Web";

	$datos = $bd->fetch_all("SELECT n_abonado_c_uhfapp, lat_c_uhfapp, lng_c_uhfapp FROM ccc.crudo_uhfapp");

	
	foreach ($datos as $clave => $dato) {

		$lat = floatval($dato['lat_c_uhfapp']);
		$lng = floatval($dato['lng_c_uhfapp']);
		$numero_abonado = $dato["n_abonado_c_uhfapp"];

		$info_mapa[] =  ["lat" => $lat, "lng" => $lng, "numero_abonado" => $numero_abonado];
		//$info_mapa =  ["lat" => $lat, "lng" => $lng, "numero_abonado" => $numero_abonado];
	}
	


	/*$info_mapa->lat = -26.835621;
	$info_mapa->lng = -65.208275;
	$info_mapa->numero_abonado = 123456;*/


	//print_r($info_mapa);
	//
	

	//echo count($info_mapa);
	
	echo json_encode($info_mapa,JSON_PRETTY_PRINT);

	$bd->cerrar_conexion();

?>