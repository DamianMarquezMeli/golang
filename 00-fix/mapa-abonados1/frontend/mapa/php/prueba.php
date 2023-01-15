
<?php



$myObj[] =  ["lat" => -26.835621, "lng" => -65.208275, "numero_abonado" => 123456];

//echo count($myObj);

$myJSON = json_encode($myObj);

echo json_encode($info_mapa,JSON_PRETTY_PRINT);

echo $myJSON;
?>