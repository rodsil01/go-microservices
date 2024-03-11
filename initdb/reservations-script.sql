CREATE SCHEMA `apuesta_total_reservations`;

USE `apuesta_total_reservations`;

CREATE TABLE `apuesta_total_reservations`.`clients` (
  `id` CHAR(38) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `apuesta_total_reservations`.`routes` (
  `id` CHAR(38) NOT NULL,
  `source` VARCHAR(255) NOT NULL,
  `destination` VARCHAR(255) NOT NULL,
  `departure_date` DATETIME NOT NULL,
  `arrival_date` DATETIME NOT NULL,
  `available_seats` INT NOT NULL,
  `price` DOUBLE NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `apuesta_total_reservations`.`reservations` (
  `id` CHAR(38) NOT NULL,
  `route_id` CHAR(38) NOT NULL,
  `client_id` CHAR(38) NOT NULL,
  `reservation_date` DATETIME NOT NULL,
  `seats` INT NOT NULL,
  `state` SMALLINT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`route_id`) REFERENCES `routes`(`id`),
  FOREIGN KEY (`client_id`) REFERENCES `clients`(`id`));

INSERT INTO `apuesta_total_reservations`.`clients` VALUES ('12438f8b-c347-4782-a7a1-1b9e5ac0ae44', 'Cliente 1', 'Apellido 1', 'cliente.apellido.1@gmail.com');
INSERT INTO `apuesta_total_reservations`.`clients` VALUES ('be924f5a-18c1-4676-ab10-d4c0b0f13cd4', 'Cliente 2', 'Apellido 2', 'cliente.apellido.2@gmail.com');
INSERT INTO `apuesta_total_reservations`.`clients` VALUES ('0f23b396-3a60-43f8-9254-8c2a0ee52502', 'Cliente 3', 'Apellido 3', 'cliente.apellido.3@gmail.com');
INSERT INTO `apuesta_total_reservations`.`clients` VALUES ('4bc90533-ec62-45b3-9e43-6f1f79b239f0', 'Cliente 4', 'Apellido 4', 'cliente.apellido.4@gmail.com');
INSERT INTO `apuesta_total_reservations`.`clients` VALUES ('42b94ad8-9850-4e4c-ac5b-81bef700b6ea', 'Cliente 5', 'Apellido 5', 'cliente.apellido.5@gmail.com');

INSERT INTO `apuesta_total_reservations`.`routes` VALUES ('96de053d-fcfb-4409-8b96-0f7b136e62e0', 'Origen 1', 'Destino 1', '2024-03-10 23:00:00', '2024-03-11 23:00:00', 100, 258.99);
INSERT INTO `apuesta_total_reservations`.`routes` VALUES ('fff65364-1bed-4307-8bb4-813b2904fc02', 'Origen 2', 'Destino 2', '2024-03-10 23:00:00', '2024-03-11 23:00:00', 200, 459.99);
INSERT INTO `apuesta_total_reservations`.`routes` VALUES ('25b44877-c1f0-4680-8cfc-c0718ecffa36', 'Origen 3', 'Destino 3', '2024-03-10 23:00:00', '2024-03-11 23:00:00', 300, 149.23);
INSERT INTO `apuesta_total_reservations`.`routes` VALUES ('78bf5ed5-5bff-4eed-b520-6e1ba0b3cd4d', 'Origen 4', 'Destino 4', '2024-03-10 23:00:00', '2024-03-11 23:00:00', 200, 522.12);
INSERT INTO `apuesta_total_reservations`.`routes` VALUES ('875bfecb-7c6b-4d06-adb6-0dab81bcab26', 'Origen 5', 'Destino 5', '2024-03-10 23:00:00', '2024-03-11 23:00:00', 100, 123.24);

INSERT INTO `apuesta_total_reservations`.`reservations` VALUES ('22abec06-783e-4f10-87ba-918b5e773444', '96de053d-fcfb-4409-8b96-0f7b136e62e0', '12438f8b-c347-4782-a7a1-1b9e5ac0ae44', '2024-03-04 23:00:00', 2, 0);
