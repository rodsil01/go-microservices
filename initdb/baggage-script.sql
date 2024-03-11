CREATE SCHEMA `apuesta_total_baggage`;

USE `apuesta_total_baggage`;

CREATE TABLE `apuesta_total_baggage`.`baggage` (
  `id` CHAR(38) NOT NULL,
  `reservation_id`  CHAR(38) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `weight` DOUBLE NOT NULL,
  PRIMARY KEY (`id`));

INSERT INTO `apuesta_total_baggage`.`baggage` VALUES ('6ce5a38d-0688-419a-96a3-ec115db44fb3', '22abec06-783e-4f10-87ba-918b5e773444', 'Equipaje pequeño', 14.50);
