<?php
	
	/** 
	*	Dependecias: 
	*		config.h
	*	
	*	Descripcion: 
	*		Clase para manejo de Mysql.
	*/

	header('Content-Type: text/html; charset=UTF-8'); 
	ini_set("display_errors", "On");
	error_reporting(E_ALL | E_STRICT);
 	header("Content-Type: text/html; charset=UTF-8");
 	date_default_timezone_set('America/Argentina/Tucuman');
 	setlocale(LC_ALL, 'es-AR');

	require_once(LIB_PATH.DS."config/config.php");

	class MysqlDatabase {

	 	private $connection;
		public $last_query;
		private $magic_quotes_active;
		private $real_escape_string_exists;

		function __construct($nombreBD="") {
			//echo $nombreBD;
			//echo "a";
			//$this->create_db();			
			$this->open_connection($nombreBD);
			$this->magic_quotes_active = get_magic_quotes_gpc();
			$this->real_escape_string_exists = function_exists("mysql_real_escape_string");
		}

		/**
		*
		*/
		public function open_connection ($nombreBD) {
			$this->connection = new mysqli(DB_SERVER, DB_USER, DB_PASS, $nombreBD);
			if ($this->connection->connect_errno) {
	    		die( "Fallo al conectar a MySQL: (" . $mysqli->connect_errno . ") " . $mysqli->connect_error);
			}
			else {
				$this->query ("SET NAMES 'UTF8'"); //Este punto es crucial para tener la informacion correctamente codificada en español
				//echo "Conexión exitosa <br />";
				//echo $this->connection->host_info . "<br />";
			}
		}

		/*********************************************************************************
		//No se usa
		//funciona

		function create_db ($db_name=DB_NAME) {
			echo "entra";
			$this->connection = new mysqli(DB_SERVER, DB_USERROOT, DB_PASSROOT);
			if ($this->connection->connect_errno) {
				die( "Fallo al conectar a MySQL: (" . $mysqli->connect_errno . ") " . $mysqli->connect_error);
			}
			else {
				$sql = "CREATE DATABASE IF NOT EXISTS " . $db_name;
				$sql .= " DEFAULT CHARACTER SET " . DB_CHARSET;
				$sql .= " DEFAULT COLLATE " . DB_COLLATION; 
				$this->query($sql);	  	
				//$this->grant_privileges();
			}
		}

		*********************************************************************************/

		/** 
		*	PHP cierra todos los archivos y conexiones a bases de datos al final del script. Es una buena practica cerrarlas manualmente, pero no es indispensable. 
		*/
		public function close_connection () {
			$desconexion = $this->connection->close();
			if (!$desconexion) {
				error_log(date("d/m/Y-G:i") . " - ERROR: Fallo al cerrar la conexión con MySQL: (" . $mysqli->connect_errno . ") " . $mysqli->connect_error."\n", 3, LOG_ERRORS);
				echo log::leer_ultimas_n_lineas(1,LOG_ERRORS);
			}
			else {
				//echo "Desconexión exitosa <br />";
			}
		}

		/**
		*
		*/
		public function query ($sql) {
			/*
			echo $sql;
			echo "<br />";
			echo "<br />";
			echo "<br />";
			echo "<br />";
			*/
			$this->last_query = $sql;
			//echo $this->last_query;
			
			$result = $this->connection->query($sql);

			$this->confirm_query($result);
			return $result;
		}

		/** 
		*	Revisa version de PHP.
		*	Arregla problema con caracteres especiales para mysql.
		*/
		public function escape_value ($value) {
			if( $this->real_escape_string_exists ) { // PHP v4.3.0 or higher
				// undo any magic quote effects so mysql_real_escape_string can do the work
				if( $this->magic_quotes_active ) {
				 $value = stripslashes( $value ); 
				}

				$value = $this->connection->real_escape_string($value);
			}
			else { // before PHP v4.3.0
				// if magic quotes aren't already on then add slashes manually
				if( $this->magic_quotes_active ) { 
					$value = addslashes( $value ); 
				}
				// if magic quotes are active, then the slashes already exist
			}
			return $value;
		}

		/**
		*	Devuelve la consulta en un array.
		*/
		public function fetch_array ($flag=MYSQLI_BOTH) {
			$resultado = $this->connection->fetch_array(MYSQLI_BOTH);
			return $resultado;
		}

		/**
		*	Inserta	informacion en una columna de una tabla dada.		
		*
		*	Entradas: 
		*		$table(string)
		*		$column(array)
		*		$value(array)
		*	Salida:
		*		Resultado del query
		*
		*	Nota: 
		*		Crea una nueva columna al insertar la info.
		*		No creo que sea muy util.
		*/
		public function insert_into_column($table, $columns, $values) {
			$scolumns = Arreglo::arreglo_a_string ($columns);
			$svalues = Arreglo::arreglo_a_string ($values);
			echo "columnas: " . $scolumns;
			echo "<br />";
			echo "valores:" . $svalues;
			echo "<br />";
			$query = "INSERT INTO {$table} ({$scolumns}) VALUES ({$svalues})";
			echo "Query:" . $query;
			echo "<br />";
			//$resultado = $this->query($query);

			return $resultado;
		}

		/**
		*	Actualiza la informacion de n columnas en una determinada fila una columnadad.
		*
		*	Entradas: 
		*		$table(string)
		*		$column(array)
		*		$value(array)
		*	Salida:
		*		Resultado del query
		*/
		public function update_column ($table, $columns, $values, $id) {
			//print_r($columns);
			$sentencia="";
			if(count($columns)>1) {
				for($i=0; $i<count($columns); $i++) {		
					//echo $columns[$i];		
					//echo "<br />";
					$sentencia .= $columns[$i] . '=' . $values[$i] . ',';
					//echo $sentencia;
				}
				$sentencia = substr($sentencia, 0, -1);
			}
			else {
				$sentencia = $columns[$i] . '=' . $values[$i];
				echo "entra";
			}
			//column1=value, column2=value2
			$query = "UPDATE {$table} SET {$sentencia} WHERE id={$id}";
			//echo "Query:" . $query;
			//echo "<br />";
			$resultado = $this->query($query);
			return $resultado;
		}



		/* Video 6x03 */
		/**
		*	Devuelve la cantidad de filas que en el result_set.
		*/
		public function num_rows ($result_set) {
		}

		/**
		*	Devuelve el ultima id conectado.
		*/
		public function insert_id () {
		}

		/**
		*	Cuantas filas fueron afectadas por el ultimo query.
		*/
		public function affected_rows () {
		}
		/***/


		/**
		*
		*/
		private function confirm_query ($result_set) {
			//echo "entra";
			if(!$result_set) {
				$output = "<br /><br />";
				$output .= "Fallo al ejecutar SQL query:" . $this->connection->error;
				$output .= "<br /><br />";
				$output .= "Ultimo SQL query: " . $this->last_query; 
				die($output);
			}
			else {
				//echo "SQL query confirmado.";
			}
		}

		/**
		*	$nombreArchivo: string - Nombre del archivo.csv.
		*	$nombreTabla: string - Nombre de la tabla donde se importaran los datos.
		*	$registros: array - Array con los nombres de todos los registros a utilizar. Hay que seguir el orden del csv.
		*	$evitarPrimeraLinea: bool - Selecciona si se evita la primera linea o no. Por defecto TRUE, o sea, se evita.
		*	Importa un cvs a la base de datos, en una tabla dada.	
		* 	Hay que tener muy en cuenta el problema de comillas simples y dobles para poder hacer una correcta importacion.
		* 	Con la configuracion actual NO se debe utilizar.
		*/		
		public function csv_a_mysql ($nombreArchivoCSV, $nombreTabla, $registros, $evitarPrimeraLinea=TRUE) {  
			//implota el array en un string separado por comas.
			$registrosAString = implode(', ', $registros); 

			if (file_exists($nombreArchivoCSV)) {

	 			$fh = fopen($nombreArchivoCSV,"r");// || die("<br /><br />No se pudo abrir el archivo .csv!<br /><br />");

	 			if($fh != false) {
	  				while ( ($data = fgetcsv($fh) ) !== FALSE ) {

	 					//$data = fgetcsv($fh);

	 					//Es necesario que los VALUES vaya entre "", ejemplo: 'data1'.
	  					$dataAString = "'".implode("', '", $data)."'";
						
						if(!$evitarPrimeraLinea) {
							
							echo $dataAString;
							echo "<br />";
							echo "<br />";
							echo "<br />";
							echo "<br />";

							$sql = "INSERT INTO {$nombreTabla} ( {$registrosAString } ) VALUES ( {$dataAString} )";
 
							echo $sql;
							echo "<br />";
							echo "<br />";
							echo "<br />";
							echo "<br />";

							

							$result = $this->query($sql);
	        				//return $result;
						}

						$evitarPrimeraLinea = FALSE;
					}
				}
				fclose($fh);
			}
			else {
				error_log(date("d/m/Y-G:i") . " - ERROR: No existe {$nombreArchivoCSV}'.\n", 3, LOG_ERRORS);
			}


			/*$sql ="INSERT INTO abonados 
					(usuario, 
					exUsuario, 
					apellido, 
					nombre, 
					calle, 
					nroCalle, 
					direccion, 
					municipio,
					zona, 
					telefono, 
					barrio) 
					VALUES ('7', '0', 'CATSAP', 'ANA', 'LORENZO', '620', 'B', 'CIU', '8', '4306463', '52')";

	        $result = $this->query($sql);
	        return $result;*/
    	}

    	/**
    	*
    	*/
		public function xml_a_mysql ($nombreArchivoXml, $nombreTabla="", $registros="")  {			
			if (file_exists($nombreArchivoXml)) {
				
				$xml = simplexml_load_file($nombreArchivoXml);

				//print_r($xml);

				//public void addAttribute ( string $name [, string $value [, string $namespace ]] )	
				//public SimpleXMLElement addChild ( string $name [, string $value [, string $namespace ]] )
				//public mixed asXML ([ string $filename ] )
				//public SimpleXMLElement attributes ([ string $ns = NULL [, bool $is_prefix = false ]] )
				//public SimpleXMLElement children ([ string $ns [, bool $is_prefix = false ]] )
				
				/*

				echo $xml->count();	
				echo "<br />";
				echo "<br />";
				echo "<br />";
				echo "<br />";
				
				$namespaces = $xml->getDocNamespaces(true);
				var_dump($namespaces);

				echo "<br />";
				echo "<br />";
				echo "<br />";
				echo "<br />";		
				
				echo $xml->getName();
				echo "<br />";
				echo "<br />";
				echo "<br />";
				echo "<br />";	

				$namespaces = $xml->getNamespaces(true);
				var_dump($namespaces);
				echo "<br />";
				echo "<br />";
				echo "<br />";
				echo "<br />";				

				echo $xml->EventSchedule;
				echo "<br />";
				echo "<br />";
				echo "<br />";
				echo "<br />";		


				*/		

				//public bool registerXPathNamespace ( string $prefix , string $ns )
				//public string __toString ( void )
				//public array xpath ( string $path )




				
				//echo "Llenando ikon.EPG";
				//echo "<br />";

				/*
				
				


				foreach ($xml->EventSchedule as $canales) {			
			   		foreach ($canales->Event as $info) {
						echo $canales["sService"]; 
						echo "<br />";
						echo "<br />";	
						echo $info["sTitle"];
						echo "<br />";
						echo "<br />";	
						echo $info->ShortDescriptor;
						echo "<br />";
						echo "<br />";	
			   		}
			   	}

			   	*/

			   	//var_dump($xml->EventSchedule);

			   //	var_dump($xml->EventSchedule["sService"]);

			   		var_dump($xml->EventSchedule->Event["sTitle"]);




//object(SimpleXMLElement)#6 (2) { ["@attributes"]=> array(3) { ["sService"]=> string(3) "106" ["tStart"]=> string(19) "2014.03.25-00:00:00" ["tEnd"]=> string(19) "2014.04.01-23:59:59" } ["Event"]=> array(3) { [0]=> object(SimpleXMLElement)#8 (3) { ["@attributes"]=> array(4) { ["uId"]=> string(1) "1" ["tBoxStart"]=> string(19) "2014.03.25-00:00:00" ["dBoxDur"]=> string(8) "01:00:00" ["sTitle"]=> string(20) "Vivo a la medianoche" } ["ShortDescriptor"]=> string(8) "Noticias" ["ExtendedDescriptor"]=> string(34) "Las noticias para estar enterados." } [1]=> object(SimpleXMLElement)#9 (3) { ["@attributes"]=> array(4) { ["uId"]=> string(1) "2" ["tBoxStart"]=> string(19) "2014.03.25-01:00:00" ["dBoxDur"]=> string(8) "01:30:00" ["sTitle"]=> string(20) "Noticias / Musicales" } ["ShortDescriptor"]=> object(SimpleXMLElement)#11 (1) { ["@attributes"]=> array(2) { ["sLang"]=> string(3) "eng" ["sName"]=> string(20) "Noticias / Musicales" } } ["ExtendedDescriptor"]=> string(29) "CCC la TV Digital de Tucumán" } [2]=> object(SimpleXMLElement)#10 (3) { ["@attributes"]=> array(4) { ["uId"]=> string(3) "136" ["tBoxStart"]=> string(19) "2014.04.01-23:30:00" ["dBoxDur"]=> string(8) "00:30:00" ["sTitle"]=> string(12) "Líderes hoy" } ["ShortDescriptor"]=> string(11) "Entrevistas" ["ExtendedDescriptor"]=> string(82) "Entrevistas a los que construyen opinión y tendencias. Con: Cecilia Luchía Puig." } } }




			   	/*

				foreach ($xml->EventSchedule as $canales) {	
						echo $canales["sService"]; 
						echo "<br />";
						echo "<br />";
						var_dump( $canales->Event);	
						echo "<br />";
						echo "<br />";		
			   		
			   	}
				*/



			   	/*	
			   	<EventSchedule sService="106" tStart="2014.03.25-00:00:00" tEnd="2014.04.01-23:59:59">
				   	<Event uId="1" tBoxStart="2014.03.25-00:00:00" dBoxDur="01:00:00" sTitle="Vivo a la medianoche">
						<ShortDescriptor sLang="eng" sName="Vivo a la medianoche">Noticias</ShortDescriptor>
						<ExtendedDescriptor sLang="eng">Las noticias para estar enterados.</ExtendedDescriptor>
					</Event>
				*/

				/*foreach ($xml->EventSchedule as $canales) {			
			   		foreach ($canales->Event as $info) {
				   		$horaActual = date("G");
				    	$horaE = horaEmision($info);
					  	$minE = minutoEmision($info);					
					   	$horaFinal = horaFinalPrograma($info);	
					   	$sin = $info->ExtendedDescriptor;

						$numeroCanal = $canales["sService"]; 
						$tituloPrograma = $info["sTitle"];
						$fechaEmision = $info["tBoxStart"];
						$horaInicioPrograma = $horaE.":".$minE;
						$horaFinPrograma = $horaFinal;
						$duracionPrograma = $info["dBoxDur"];
						$generoPrograma = $info->ShortDescriptor;
						$sinopsisPrograma = $info->ExtendedDescriptor;
						$uid = $info["uId"];
						
						$sinopsisPrograma = $this->escape_value($sinopsisPrograma);

						//echo $sinopsisPrograma;
						//echo "<br />";

						$sql = "INSERT INTO {$nombreTabla} ( {$registrosAString } ) 
							VALUES (
							'$numeroCanal',
							'$tituloPrograma',
							'$fechaEmision',
							'$horaInicioPrograma',
							'$horaFinPrograma',
							'$duracionPrograma',
							'$generoPrograma',
							'$sinopsisPrograma',
							'$uid'
						)";
							
						$base_de_datos->query($insertar);
				 	}
			   	}			
			   	//echo "Finalizacion exitosa";
			   	//echo "<br />";
			}
			else {
				error_log(date("d/m/Y-G:i") . "No existe epg.xml", 3, "../logs/errores.log");
			*/
			}
		}
	}

	/*$base_de_datos = new MysqlDatabase();
	$bd =& $base_de_datos;*/

?>