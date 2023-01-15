SELECT * FROM ccc.web_sucursales_nuevas;
SELECT * FROM ccc.web_sucursales_viejas;

SELECT UPPER(alphanum(n.`agente`)) FROM ccc.web_sucursales_nuevas AS n;
SELECT UPPER(alphanum(v.`nombre sucursal`)) FROM ccc.web_sucursales_viejas AS v;

SELECT UPPER(limp_carac(n.`agente`)) FROM ccc.web_sucursales_nuevas AS n;
SELECT UPPER(limp_carac(v.`nombre sucursal`)) FROM ccc.web_sucursales_viejas AS v;

SELECT UPPER(remove_non_alphanum_char_f(n.`agente`)) FROM ccc.web_sucursales_nuevas AS n;
SELECT UPPER(remove_non_alphanum_char_f(v.`nombre sucursal`)) FROM ccc.web_sucursales_viejas AS v;


-- ----------------------------- --
-- ----------------------------- --
-- ----------------------------- --

#Funcion
#Elimina todos los caracteres NO alfanumericos de strings.
# hay combinar las 2 funciones, lo ideal sería hacer 1 que funcione bien
DROP FUNCTION IF EXISTS alphanum; 
DELIMITER | 
CREATE FUNCTION alphanum( str CHAR(255) ) RETURNS CHAR(255) DETERMINISTIC
BEGIN 
  DECLARE i, len SMALLINT DEFAULT 1; 
  DECLARE ret CHAR(255) DEFAULT ''; 
  DECLARE c CHAR(1); 
  SET len = CHAR_LENGTH( str ); 
  REPEAT 
    BEGIN 
      SET c = MID( str, i, 1 ); 
      IF c REGEXP '[[:alnum:]]' THEN 
        SET ret=CONCAT(ret,c); 
      END IF; 
      SET i = i + 1; 
    END; 
  UNTIL i > len END REPEAT; 
  RETURN ret; 
END 
| DELIMITER ;


###############################################################
# Parace q esta es la buena, controlar que de verdad funcione #
###############################################################
delimiter //
DROP FUNCTION IF EXISTS remove_non_alphanum_char_f //
CREATE FUNCTION remove_non_alphanum_char_f (prm_strInput varchar(255))
RETURNS VARCHAR(255)
DETERMINISTIC
BEGIN
  DECLARE i INT DEFAULT 1;
  DECLARE v_char VARCHAR(1);
  DECLARE v_parseStr VARCHAR(255) DEFAULT ' ';
 
WHILE (i <= LENGTH(prm_strInput) )  DO 
 
  SET v_char = SUBSTR(prm_strInput,i,1);
  IF v_char REGEXP  '^[A-Za-z0-9 ]+$' THEN  #alphanumeric
    
        SET v_parseStr = CONCAT(v_parseStr,v_char);  

  END IF;
  SET i = i + 1;
END WHILE; 
RETURN trim(v_parseStr);
END
// delimiter ;


DROP FUNCTION IF EXISTS `remove_accents`;

DELIMITER //
CREATE FUNCTION `remove_accents`(`str` TEXT)
    RETURNS text
    LANGUAGE SQL
    DETERMINISTIC
    NO SQL
    SQL SECURITY INVOKER
    COMMENT ''

BEGIN

    SET str = REPLACE(str,'Š','S');
    SET str = REPLACE(str,'š','s');
    SET str = REPLACE(str,'Ð','Dj');
    SET str = REPLACE(str,'Ž','Z');
    SET str = REPLACE(str,'ž','z');
    SET str = REPLACE(str,'À','A');
    SET str = REPLACE(str,'Á','A');
    SET str = REPLACE(str,'Â','A');
    SET str = REPLACE(str,'Ã','A');
    SET str = REPLACE(str,'Ä','A');
    SET str = REPLACE(str,'Å','A');
    SET str = REPLACE(str,'Æ','A');
    SET str = REPLACE(str,'Ç','C');
    SET str = REPLACE(str,'È','E');
    SET str = REPLACE(str,'É','E');
    SET str = REPLACE(str,'Ê','E');
    SET str = REPLACE(str,'Ë','E');
    SET str = REPLACE(str,'Ì','I');
    SET str = REPLACE(str,'Í','I');
    SET str = REPLACE(str,'Î','I');
    SET str = REPLACE(str,'Ï','I');
    SET str = REPLACE(str,'Ñ','N');
    SET str = REPLACE(str,'Ò','O');
    SET str = REPLACE(str,'Ó','O');
    SET str = REPLACE(str,'Ô','O');
    SET str = REPLACE(str,'Õ','O');
    SET str = REPLACE(str,'Ö','O');
    SET str = REPLACE(str,'Ø','O');
    SET str = REPLACE(str,'Ù','U');
    SET str = REPLACE(str,'Ú','U');
    SET str = REPLACE(str,'Û','U');
    SET str = REPLACE(str,'Ü','U');
    SET str = REPLACE(str,'Ý','Y');
    SET str = REPLACE(str,'Þ','B');
    SET str = REPLACE(str,'ß','Ss');
    SET str = REPLACE(str,'à','a');
    SET str = REPLACE(str,'á','a');
    SET str = REPLACE(str,'â','a');
    SET str = REPLACE(str,'ã','a');
    SET str = REPLACE(str,'ä','a');
    SET str = REPLACE(str,'å','a');
    SET str = REPLACE(str,'æ','a');
    SET str = REPLACE(str,'ç','c');
    SET str = REPLACE(str,'è','e');
    SET str = REPLACE(str,'é','e');
    SET str = REPLACE(str,'ê','e');
    SET str = REPLACE(str,'ë','e');
    SET str = REPLACE(str,'ì','i');
    SET str = REPLACE(str,'í','i');
    SET str = REPLACE(str,'î','i');
    SET str = REPLACE(str,'ï','i');
    SET str = REPLACE(str,'ð','o');
    SET str = REPLACE(str,'ñ','n');
    SET str = REPLACE(str,'ò','o');
    SET str = REPLACE(str,'ó','o');
    SET str = REPLACE(str,'ô','o');
    SET str = REPLACE(str,'õ','o');
    SET str = REPLACE(str,'ö','o');
    SET str = REPLACE(str,'ø','o');
    SET str = REPLACE(str,'ù','u');
    SET str = REPLACE(str,'ú','u');
    SET str = REPLACE(str,'û','u');
    SET str = REPLACE(str,'ý','y');
    SET str = REPLACE(str,'ý','y');
    SET str = REPLACE(str,'þ','b');
    SET str = REPLACE(str,'ÿ','y');
    SET str = REPLACE(str,'ƒ','f');


    RETURN str;
END
//
DELIMITER ;


-- ----------------------------- --
-- ----------------------------- --
-- ----------------------------- --

#Eliminar de la web
SELECT v.`nombre sucursal`, v.dirección, v.localidad
FROM ccc.web_sucursales_nuevas AS n
RIGHT JOIN ccc.web_sucursales_viejas AS v
ON UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(n.domicilio)))) = UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(v.dirección))))
AND UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(n.localidad)))) = UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(v.localidad))))
OR UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(n.sucursal)))) = UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(v.`nombre sucursal`))))
OR UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(n.agente)))) = UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(v.`nombre sucursal`))))
WHERE n.domicilio IS NULL
OR n.localidad IS NULL
OR n.sucursal IS NULL
OR n.agente IS NULL
ORDER BY v.localidad;

-- ----------------------------- --

#Agregar a la web PARECE q esta funciona de verdad, revisar!!!!
SELECT n.sucursal, n.agente, v.`nombre sucursal`, n.localidad, n.domicilio, n.horario_habiles, n.horario_sabados, n.horario_domingos
FROM ccc.web_sucursales_nuevas AS n
LEFT JOIN ccc.web_sucursales_viejas AS v
ON UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(n.domicilio)))) = UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(v.dirección))))
AND UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(n.localidad)))) = UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(v.localidad))))
OR UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(n.sucursal)))) = UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(v.`nombre sucursal`))))
OR UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(n.agente)))) = UPPER(alphanum(remove_non_alphanum_char_f(remove_accents(v.`nombre sucursal`))))
WHERE v.dirección IS NULL
OR v.`nombre sucursal` IS NULL
OR v.localidad IS NULL
ORDER BY n.localidad;