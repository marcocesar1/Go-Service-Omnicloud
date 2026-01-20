# API People con Go

Este proyecto es una pequeña aplicación de ejemplo para el desarrollo de APIs RESTful con Go.

---

## Requisitos previos

Herramientas necesarias para el desarrollo de este proyecto:

- [Go](https://go.dev/doc/install)
- [MongoDB Client](https://www.mongodb.com/docs/drivers/go/current/connect/mongoclient/#std-label-golang-mongoclient)
- [Chi Router](https://github.com/go-chi/chi)
- [City API](https://random-city-api.vercel.app/api/random-city)
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

---

## Inicio

### Verificación de requisitos

Para verificar si tus requisitos previos están cumplidos, ejecuta los siguientes comandos:

```bash
docker version
```

```bash
docker compose version
```

### 🐳 Docker



Una vez instalado docker, en la raiz de proyecto copia el archivo `.env.example` a `.env`.

El archivo `.env.example` contiene las variables de entorno ya configuradas para el proyecto así que no es necesario cambiar nada (en un proyecto de producción esto no es lo recomendado pero es una app de practica), puedes cambiar los valores según tus necesidades si así lo deseas.


```bash
MONGO_URL # url de la base de datos
CITY_API_URL # url de la API de ciudades
APP_PORT # puerto de la aplicación
```

Una vez hecho esto, ejecuta los siguientes comandos para levantar el contenedor de la aplicación:

```bash
docker compose up
```

Si todo ha ido bien, deberías ver algo como esto:

```bash
app_mongodb  | {"t":{"$date":"2026-01-20T19:51:38.305+00:00"},"s":"I",  "c":"NETWORK",  "id":6788700, "ctx":"conn2","msg":"Received first command on ingress connection since session start or auth handshake","attr":{"elapsedMillis":1}}
app_api      | Successfully connected to MongoDB!
app_api      | Server running on: http://localhost:3000
```

Por defecto, la aplicación estará disponible en `http://localhost:3000`.
Al ingresar a la ruta `/health` deberías ver algo como esto:

```bash
{"status":"ok","message":"Server is running"}
```

### Ejecución de tests

Para ejecutar los tests, ejecuta el siguiente comando para entrar en el contenedor de la aplicación:

```bash
docker exec -it app_api sh 
```

Luego ejecuta el siguiente comando para ejecutar los tests:

```bash
go test ./... -v
```

### CI/CD con GitHub Actions a Digital Ocean

Actualmente este proyecto está configurado para que se ejecute en [Digital Ocean](https://www.digitalocean.com/) mediante GitHub Actions para la integración continua y despliegue automático.

Las configuraciones de CI/CD se encuentran en el archivo `.github/workflows/ci_cd.yml`.

Cuando se hace un cambio en el código, se ejecutará el archivo `ci_cd.yml` que ejecutará los tests y desplegará la aplicación en Digital Ocean a través de una conexión SSH.

La URL de la aplicación desplegada en Digital Ocean es `http://167.172.211.101`.

### Postman

Para probar la API con Postman, puedes usar el archivo `People Api.postman_collection.json` que contiene todas las rutas de la API ubicada en la carpeta `postman` de la raíz del proyecto.