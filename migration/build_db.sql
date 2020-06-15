-- docker run --name olist-mysql -p 3306:8093 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:8
CREATE DATABASE olist;

USE olist;

CREATE TABLE `books` (
  `id`             INT         NOT NULL AUTO_INCREMENT,
  `name`           VARCHAR(50) NOT NULL               ,
  `edition`        INT                                ,
  `published_year` INT                                ,
  PRIMARY KEY (`id`)
);

CREATE TABLE `authors` (
  `id`             INT         NOT NULL AUTO_INCREMENT,
  `name`           VARCHAR(50) NOT NULL               ,
  PRIMARY KEY (`id`)
);

CREATE TABLE `books_authors` (
  `id`             INT         NOT NULL AUTO_INCREMENT,
  `book_id`        INT         NOT NULL,
  `author_id`      INT         NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`book_id`)   REFERENCES `books`(`id`)   ON DELETE CASCADE,
  FOREIGN KEY (`author_id`) REFERENCES `authors`(`id`)                  ,
  INDEX idx_book (`book_id`)
);
