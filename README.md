# Solución

## Consideraciones

- La solución utiliza Docker. La guía mostrada a continuación muestra cómo iniciar los servicios con docker-compose
- Recomiendo usar Postman para realizar las pruebas

## Setup

1. `git clone https://github.com/rodsil01/go-microservices.git`
2. `cd go-microservices`
3. `docker-compose up -d`

## Servicios

La solución consiste en los siguientes servicios:

1. MySQL: Base de datos utilizada para la persistencia de datos
2. RabbitMQ: Servicio de mensajería para la comunicación entre los microservicios
3. Servicio de Reservas  (localhost:8080): Servicio encargado de realizar reservas de vuelos
4. Servicio de Equipajes (localhost:8081): Servicio encargado de administrar la información de equipajes de las reservas realizadas

## Enpoints

### Clientes

GET http://localhost:8080/clients


GET http://localhost:8080/clients/{client_id}

### Rutas

GET http://localhost:8080/routes


GET http://localhost:8080/routes/{route_id}

### Reservas

GET http://localhost:8080/reservations


GET http://localhost:8080/reservations/{reservation_id}




POST http://localhost:8080/reservations

### Equipaje

GET http://localhost:8081/baggage/reservations/{reservation_id}

---

**Se creó un script que initializa las bases de datos utilizadas y agrega valores para facilitar las pruebas**

**Los servicios pueden demorar unos segundos en inicializar completamente debido a las dependencias con RabbitMQ y MySQL**

## Ejemplo de petición de reserva

POST http://localhost:8080/reservations


body:

```json
{
    "routeId": "96de053d-fcfb-4409-8b96-0f7b136e62e0",
    "clientId": "4bc90533-ec62-45b3-9e43-6f1f79b239f0",
    "reservationDate": "2020-07-30T18:00:00.000Z",
    "seats": 2,
    "state": 23,
    "baggage": [
        {
            "description": "equipaje 1",
            "weight": "10.5"
        },
        {
            "description": "equipaje 2",
            "weight": "11.5"
        },
        {
            "description": "equipaje 3",
            "weight": "12.5"
        }
    ]
}
```

**Los valores de los ids para 'routeId' y 'clientId' deben ser valores existentes en la base de datos, los ids existentes son:**

**Clientes:**

```
- 12438f8b-c347-4782-a7a1-1b9e5ac0ae44

- be924f5a-18c1-4676-ab10-d4c0b0f13cd4

- 0f23b396-3a60-43f8-9254-8c2a0ee52502

- 4bc90533-ec62-45b3-9e43-6f1f79b239f0

- 42b94ad8-9850-4e4c-ac5b-81bef700b6ea
```

**Rutas:**

```
- 96de053d-fcfb-4409-8b96-0f7b136e62e0

- fff65364-1bed-4307-8bb4-813b2904fc02

- 25b44877-c1f0-4680-8cfc-c0718ecffa36

- 78bf5ed5-5bff-4eed-b520-6e1ba0b3cd4d

- 875bfecb-7c6b-4d06-adb6-0dab81bcab26
```
