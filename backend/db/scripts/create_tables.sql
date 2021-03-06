--
-- Maridb database
--

-- database version 10.5.9

--
-- Definición de la tabla `users`
--
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
	`user_id` INT NOT NULL AUTO_INCREMENT COMMENT 'Clave primaria para los usuarios del blog',
	`email` VARCHAR(70) NOT NULL COMMENT 'Correo electronico de los usuarios',
	`password_hash` VARCHAR(100) NOT NULL COMMENT 'Hash de la cadena de la contraseña',
	`name` VARCHAR(50) NOT NULL COMMENT 'Nombre(s) del usuario',
	`surname` VARCHAR(80) NOT NULL COMMENT 'Apellidos del usuario',
	`gender` CHAR(1) NULL COMMENT 'Genero del usuario. F (Femenino), M (Masculino), NULL (No especificado)',
	`phone` VARCHAR(15) NULL COMMENT 'Numero telefonico del usuario',
	CONSTRAINT `users_pk` PRIMARY KEY (`user_id`),
	CONSTRAINT `users_email_un` UNIQUE (`email`)
) ENGINE=InnoDB DEFAULT CHAR SET=utf8mb4 COMMENT='Tabla donde se almacenan los usuarios del blog';

--
-- Definición de la tabla `articles`
--
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
	`article_id` INT NOT NULL AUTO_INCREMENT COMMENT 'Clave primaria para los posts creados en el blog',
	`title` VARCHAR(100) NOT NULL COMMENT 'Titulo del post',
	`slug` VARCHAR(70) NULL COMMENT 'Enlace personalizado para el post',
	`content` TEXT NULL COMMENT 'Contenido del post',
	`user_id` INT NOT NULL COMMENT 'Autor del articulo. Llave foranea que relaciona el post con el usuario',
	`published` TINYINT(1) UNSIGNED NOT NULL COMMENT 'Bandera que indica si el articulo se encuentra publicado o en borrador',
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
	`updated_at` TIMESTAMP NULL,
	CONSTRAINT `articles_pk` PRIMARY KEY (`article_id`),
	CONSTRAINT `articles_users_fk` FOREIGN KEY (`user_id`)
		REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHAR SET=utf8mb4 COMMENT='Tabla donde se almacenan los posts escritos por los usuarios';

--
-- Definición de la tabla `comments`
--
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
	`comment_id` INT NOT NULL AUTO_INCREMENT COMMENT 'Clave primaria para los comentarios escritos por los usuarios',
	`comment` VARCHAR(255) COMMENT 'Contenido del comentario',
	`user_id` INT NOT NULL COMMENT 'Autor del comentario',
	`article_id` INT NOT NULL COMMENT 'Indica el artículo donde fue escrito el comentario',
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT `comments_pk` PRIMARY KEY (`comment_id`),
	CONSTRAINT `comments_users_fk` FOREIGN KEY (`user_id`)
		REFERENCES `users` (`user_id`),
	CONSTRAINT `comments_articles_fk` FOREIGN KEY (`article_id`)
		REFERENCES `articles` (`article_id`)
) ENGINE=InnoDB DEFAULT CHAR SET=utf8mb4 COMMENT='Tabla donde se almacenan los comentarios hechos por los usuarios en los diferentes artículos';


--
-- Definición de la tabla `user_activity`
--
DROP TABLE IF EXISTS `user_activity`;
CREATE TABLE `user_activity` (
	`article_id` INT NOT NULL,
	`user_id` INT NOT NULL,
	`visited` TINYINT(1) UNSIGNED NULL COMMENT 'Bandera que indica si el artículo fue visitado por el usuario',
	`favorites` TINYINT(1) UNSIGNED NULL COMMENT 'Bandera que indica si el usuario agrego el artículo a su sección de favortios',
	`watch_later` TINYINT(1) UNSIGNED NULL COMMENT 'Bandera que indica si el usuario agrego el artículo a su seeción de ver más tarde',
	CONSTRAINT `user_activity_pk` PRIMARY KEY (`article_id`, `user_id`),
	CONSTRAINT `activity_articles_fk` FOREIGN KEY (`article_id`)
		REFERENCES `articles` (`article_id`),
	CONSTRAINT `activity_users_fk` FOREIGN KEY (`user_id`)
		REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHAR SET=utf8mb4 COMMENT='Tabla que almacena la actividad que tienen los usuarios con los articulos';

--
-- Definición de la tabla `categories`
--
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
	`category_id` INT NOT NULL AUTO_INCREMENT,
	`label` VARCHAR(50) NOT NULL COMMENT 'Nombre de la categoria',
	CONSTRAINT `categories_pk` PRIMARY KEY (`category_id`)
) ENGINE=InnoDB DEFAULT CHAR SET=utf8mb4 COMMENT='Catalogo que almacena las categorias de los artículos';


--
-- Definición de la tabla `article_categories`
--
DROP TABLE IF EXISTS `article_categories`;
CREATE TABLE `article_categories` (
	`category_id` INT NOT NULL,
	`article_id` INT NOT NULL,
	`level` SMALLINT NOT NULL,
	CONSTRAINT `article_categories_pk` PRIMARY KEY (`category_id`, `article_id`),
	CONSTRAINT `article_categories_cat_fk` FOREIGN KEY (`category_id`)
		REFERENCES `categories` (`category_id`),
	CONSTRAINT `article_categories_art_fk` FOREIGN KEY (`article_id`)
		REFERENCES `articles` (`article_id`)
) ENGINE=InnoDB DEFAULT CHAR SET=utf8mb4 COMMENT='Tabla que almacena las categorias';

--
-- Definición de la tabla `image_formats`
--
DROP TABLE IF EXISTS `image_formats`;
CREATE TABLE `image_formats`(
	`format_id` SMALLINT NOT NULL AUTO_INCREMENT,
	`format` VARCHAR(5) NOT NULL,
	CONSTRAINT image_formats_pk PRIMARY KEY (`format_id`), 
	CONSTRAINT image_formats_un UNIQUE(`format`)
) ENGINE=InnoDB DEFAULT CHAR SET=utf8mb4 COMMENT='Catalogo que almacena la informacion sobre los tipos de images';

--
-- Definición de la tabla `article_images`
--
DROP TABLE IF EXISTS `article_images`;
CREATE TABLE `article_images` (
	`image_id` INT NOT NULL AUTO_INCREMENT,
	`label` VARCHAR(20) NULL,
	`description` VARCHAR(70) NULL,
	`filename` VARCHAR(1024) NOT NULL,
	`size` INT NOT NULL,
	`filepath` VARCHAR(255) NOT NULL,
	`index` TINYINT UNSIGNED NULL,
	`article_id` INT NOT NULL,
	`format_id` SMALLINT NOT NULL,
	CONSTRAINT `article_images_pk` PRIMARY KEY (`image_id`),
	CONSTRAINT `article_images_art_fk` FOREIGN KEY (`article_id`)
		REFERENCES `articles` (`article_id`),
	CONSTRAINT `article_images_fmt_fk` FOREIGN KEY (`format_id`)
		REFERENCES `image_formats` (`format_id`)
) ENGINE=InnoDB DEFAULT CHAR SET=utf8mb4 COMMENT='Tabla que almacena la imagenes del artIculo';
