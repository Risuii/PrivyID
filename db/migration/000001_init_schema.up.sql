CREATE TABLE `privyID`.`cake` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NULL,
  `description` VARCHAR(255) NULL,
  `rating` FLOAT NULL,
  `image` VARCHAR(255) NULL,
  `created_at` DATETIME NULL DEFAULT (now()),
  `update_at` DATETIME NULL DEFAULT (now()),
  PRIMARY KEY (`ID`)
);