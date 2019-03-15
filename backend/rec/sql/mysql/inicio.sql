DROP DATABASE IF EXISTS pruebas;

/**
 *	Base de datos: 
 */
CREATE DATABASE IF NOT EXISTS pruebas
DEFAULT CHARACTER SET utf8
DEFAULT COLLATE utf8_spanish_ci;


/**
 *	Privilegios usuario.
 */
GRANT ALL PRIVILEGES ON pruebas.* TO pablo@localhost; 


/**
 * 	Refrescar todos los privilegios. siempre que haya un cambio de privilegios.
 */
FLUSH PRIVILEGES;

/**
 * 
 */
CREATE TABLE IF NOT EXISTS pruebas.personas (
	id_persona INT(11) UNSIGNED NOT NULL AUTO_INCREMENT UNIQUE,
	
	dni_persona INT(11) UNSIGNED NULL UNIQUE,
	cuil_persona INT(11) UNSIGNED NULL UNIQUE,
	nombres_persona VARCHAR(255),
	apellidos_persona VARCHAR(255),
    sexo_persona VARCHAR(255),

	PRIMARY KEY(id_persona)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 DEFAULT COLLATE=utf8_spanish_ci;

/**
 *  Las contraseñas estan siendo guardadas en texto plano, hay que hashearlas. 
**/
CREATE TABLE IF NOT EXISTS pruebas.usuarios_empleados (
	id_usuario_empleado	INT(11) UNSIGNED NOT NULL AUTO_INCREMENT UNIQUE,
    id_persona INT(11) UNSIGNED NOT NULL UNIQUE,

	nombre_usuario_empleado 	VARCHAR(255),
	contraseña_usuario_empleado VARCHAR(255),

	PRIMARY KEY(id_usuario_empleado),

    FOREIGN KEY (id_persona) REFERENCES pruebas.personas (id_persona) 
		ON DELETE CASCADE 
		ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 DEFAULT COLLATE=utf8_spanish_ci;