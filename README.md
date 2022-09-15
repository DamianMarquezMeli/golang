# conceptos-go

1. Manejo de errores.
2. Polimorfismo.
3. Inyeccion de dependencias.
4. Arquitectura hexagonal.
   1. Inversion de dependecias.
   2. Adaptadores y puertos.


Docker shortner
docker run --name redis-test-instance -p 6379:6379 -d redis

Correr HTTP: go run cmd/http/main.go

Archivos a revisar:
github.com/devpablocristo/conceptos-go/go-hexagonal-shortener/cmd/http/main.go
github.com/devpablocristo/conceptos-go/go-hexagonal-shortener/adapter/http/gin/gin_server.go
