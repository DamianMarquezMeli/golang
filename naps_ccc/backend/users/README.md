sudo openvpn --config /root/pcristo.ovpn 
passw: 1n9.pcr1S70


Para verificar:
ping 172.30.0.141
ping 172.30.0.100

Para verificar conex a vianet
ping  172.30.0.100
http://172.30.0.100/spi40/login.php



[{"id_action":22015,"operador_ws":1,"descripcion":"Ayuda del webservice",}]

http://172.30.0.100/spi40/abn/webservice.php?request=W3siaWRfYWN0aW9uIjoyMjAxNSwib3BlcmFkb3Jfd3MiOjEsImRlc2NyaXBjaW9uIjoiQXl1ZGEgZGVsIHdlYnNlcnZpY2UiLH1d



MongoDB
mongodb://desarrollos3c:1n9d3s4rr0ll0s@172.30.0.141:27017

{
    "username":"Metallica",
    "password":"rock",
    "fullname":"Metallica Banda",
	"phone": "888-888-888",
	"role": "musica"
}


{
    "username":"pink.floyd",
    "password":"theWall",
    "email":"pink@floyd.com"
}


{"_
    id":{"$oid":"60a6bee0ea2e5a50fa571349"},
    "username":"pink.floyd","password":"theWall","email":"pink@floyd.com","created_at":{"$date":"2021-05-20T19:56:16.992Z"},"updated_at":{"$date":"2021-05-20T19:56:16.992Z"}}









En service se hace todo cq cuestion logica calculos o lo sea

repository interactua con la base de datos, o con los elementos de consulta, recibe todo procesado desde service

