-- phpMyAdmin SQL Dump
-- version 4.6.6deb5
-- https://www.phpmyadmin.net/
--
-- Servidor: localhost:3306
-- Tiempo de generación: 30-04-2019 a las 09:08:45
-- Versión del servidor: 5.7.26-0ubuntu0.18.04.1
-- Versión de PHP: 7.2.17-0ubuntu0.18.04.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `pruebas`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `empleados`
--

CREATE TABLE `empleados` (
  `id_empleado` int(11) UNSIGNED ZEROFILL NOT NULL,
  `id_persona` int(11) UNSIGNED ZEROFILL NOT NULL,
  `creado_en` datetime DEFAULT CURRENT_TIMESTAMP,
  `actualizado_en` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `eliminado_en` datetime DEFAULT NULL,
  `puesto_empleado` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `movil_empleado` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `numero_legajo_empleado` int(11) UNSIGNED ZEROFILL NOT NULL,
  `celular_empleado` int(11) UNSIGNED ZEROFILL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

--
-- Volcado de datos para la tabla `empleados`
--

INSERT INTO `empleados` (`id_empleado`, `id_persona`, `creado_en`, `actualizado_en`, `eliminado_en`, `puesto_empleado`, `movil_empleado`, `numero_legajo_empleado`, `celular_empleado`) VALUES
(00000000001, 00000000001, '2019-04-29 15:21:04', '2019-04-30 06:51:11', '2019-04-30 06:51:12', '123123', '5535', 00000000234, 00000052342),
(00000000002, 00000000002, '2019-04-30 05:54:08', '2019-04-30 06:56:24', '2019-04-30 06:56:25', 'gerente', '34', 00000000089, 00234512354),
(00000000003, 00000000003, '2019-04-30 08:46:34', NULL, NULL, 'vendedor', '1123', 00000000023, 00441231231);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `personas`
--

CREATE TABLE `personas` (
  `id_persona` int(11) UNSIGNED ZEROFILL NOT NULL,
  `creado_en` datetime DEFAULT CURRENT_TIMESTAMP,
  `actualizado_en` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `eliminado_en` datetime DEFAULT NULL,
  `dni_persona` int(11) UNSIGNED ZEROFILL NOT NULL,
  `cuil_persona` int(11) UNSIGNED ZEROFILL NOT NULL,
  `nombres_persona` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `apellidos_persona` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `sexo_persona` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `observaciones_persona` text COLLATE utf8_spanish_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

--
-- Volcado de datos para la tabla `personas`
--

INSERT INTO `personas` (`id_persona`, `creado_en`, `actualizado_en`, `eliminado_en`, `dni_persona`, `cuil_persona`, `nombres_persona`, `apellidos_persona`, `sexo_persona`, `observaciones_persona`) VALUES
(00000000001, '2019-04-29 15:21:04', NULL, NULL, 00005555555, 00000777777, 'wwwww', 'ttttttt', 'masculino', 'hhhhhhh'),
(00000000002, '2019-04-30 05:54:08', NULL, NULL, 00000000008, 00000000109, 'pitichi', 'cristo', 'femenino', 'gata gorda'),
(00000000003, '2019-04-30 08:46:34', NULL, NULL, 00081238123, 00009239234, 'jotajota', 'lopez', 'masculino', 'lalalal');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `usuarios_empleados`
--

CREATE TABLE `usuarios_empleados` (
  `id_usuario_empleado` int(11) UNSIGNED ZEROFILL NOT NULL,
  `id_empleado` int(11) UNSIGNED ZEROFILL NOT NULL,
  `creado_en` datetime DEFAULT CURRENT_TIMESTAMP,
  `actualizado_en` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `eliminado_en` datetime DEFAULT NULL,
  `nombre_usuario_empleado` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `contraseña_usuario_empleado` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `empleados`
--
ALTER TABLE `empleados`
  ADD PRIMARY KEY (`id_empleado`),
  ADD UNIQUE KEY `id_empleado` (`id_empleado`),
  ADD UNIQUE KEY `id_persona` (`id_persona`),
  ADD UNIQUE KEY `numero_legajo_empleado` (`numero_legajo_empleado`);

--
-- Indices de la tabla `personas`
--
ALTER TABLE `personas`
  ADD PRIMARY KEY (`id_persona`),
  ADD UNIQUE KEY `id_persona` (`id_persona`),
  ADD UNIQUE KEY `dni_persona` (`dni_persona`),
  ADD UNIQUE KEY `cuil_persona` (`cuil_persona`);

--
-- Indices de la tabla `usuarios_empleados`
--
ALTER TABLE `usuarios_empleados`
  ADD PRIMARY KEY (`id_usuario_empleado`),
  ADD UNIQUE KEY `id_usuario_empleado` (`id_usuario_empleado`),
  ADD UNIQUE KEY `id_empleado` (`id_empleado`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `empleados`
--
ALTER TABLE `empleados`
  MODIFY `id_empleado` int(11) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
--
-- AUTO_INCREMENT de la tabla `personas`
--
ALTER TABLE `personas`
  MODIFY `id_persona` int(11) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
--
-- AUTO_INCREMENT de la tabla `usuarios_empleados`
--
ALTER TABLE `usuarios_empleados`
  MODIFY `id_usuario_empleado` int(11) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT;
--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `empleados`
--
ALTER TABLE `empleados`
  ADD CONSTRAINT `empleados_ibfk_1` FOREIGN KEY (`id_persona`) REFERENCES `personas` (`id_persona`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Filtros para la tabla `usuarios_empleados`
--
ALTER TABLE `usuarios_empleados`
  ADD CONSTRAINT `usuarios_empleados_ibfk_1` FOREIGN KEY (`id_empleado`) REFERENCES `empleados` (`id_empleado`) ON DELETE CASCADE ON UPDATE CASCADE;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
