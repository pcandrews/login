/**
 * 	Crear usuario con contraseña.
 */
CREATE USER 'pablo' IDENTIFIED BY 'rocky';


/**
 * 	Garantizar privilegios usuario.
 * 	Nota: aqui estoy dandole todos los privilegios, seria conveniente restringirlos en el futuro.
 */
REVOKE ALL PRIVILEGES ON *.* FROM 'ccc_admin'@'%'; 
GRANT ALL PRIVILEGES ON *.* 
TO 'ccc_admin'@'%' 
REQUIRE NONE WITH GRANT OPTION 
		MAX_QUERIES_PER_HOUR 0 
		MAX_CONNECTIONS_PER_HOUR 0 
		MAX_UPDATES_PER_HOUR 0 
		MAX_USER_CONNECTIONS 0;


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

/* 
	zerofill lo unico que hace es ocupar los espacios en blanco a la izquierda con ceros para visualizar mejor el numero.
	no tiene efecto alguno sobre los datos
*/
CREATE TABLE IF NOT EXISTS pruebas.personas (
	id_persona INT(11) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT UNIQUE,
	
	creado_en DATETIME DEFAULT CURRENT_TIMESTAMP,
	actualizado_en DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
	eliminado_en DATETIME DEFAULT NULL,

	dni_persona INT(11) UNSIGNED ZEROFILL NOT NULL UNIQUE,
	cuil_persona INT(11) UNSIGNED ZEROFILL NOT NULL UNIQUE,
	nombres_persona VARCHAR(255),
	apellidos_persona VARCHAR(255),
    sexo_persona VARCHAR(255),
	observaciones_persona TEXT,

	PRIMARY KEY(id_persona)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 DEFAULT COLLATE=utf8_spanish_ci;


CREATE TABLE IF NOT EXISTS pruebas.empleados (
	id_empleado INT(11) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT UNIQUE,
	id_persona INT(11) UNSIGNED ZEROFILL NOT NULL UNIQUE,

	creado_en DATETIME DEFAULT CURRENT_TIMESTAMP,
	actualizado_en DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
	eliminado_en DATETIME DEFAULT NULL,

	puesto_empleado VARCHAR(255),
	movil_empleado VARCHAR(255),
	numero_legajo_empleado  INT(11) UNSIGNED ZEROFILL NOT NULL UNIQUE,
	celular_empleado INT(11) UNSIGNED ZEROFILL,

	PRIMARY KEY(id_empleado),
	
    FOREIGN KEY (id_persona) REFERENCES pruebas.personas (id_persona) 
		ON DELETE CASCADE 
		ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 DEFAULT COLLATE=utf8_spanish_ci;

/*  
	Las contraseñas estan siendo guardadas en texto plano, hay que hashearlas. 
	ON DELETE CASCADE
	ON UPDATE CASCADE 

	hacen que al borrar o actualizar la table padre, afecte a la hija, o sea, si borro la persona, se borra el empleado, si borro el empleado no se borra la persona.
*/
CREATE TABLE IF NOT EXISTS pruebas.usuarios_empleados (
	id_usuario_empleado	INT(11) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT UNIQUE,
    id_empleado INT(11) UNSIGNED ZEROFILL NOT NULL UNIQUE,

	creado_en DATETIME DEFAULT CURRENT_TIMESTAMP,
	actualizado_en DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
	eliminado_en DATETIME DEFAULT NULL,

	nombre_usuario_empleado 	VARCHAR(255),
	contraseña_usuario_empleado VARCHAR(255),

	PRIMARY KEY(id_usuario_empleado),

    FOREIGN KEY (id_empleado) REFERENCES pruebas.empleados (id_empleado) 
		ON DELETE CASCADE 
		ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 DEFAULT COLLATE=utf8_spanish_ci;