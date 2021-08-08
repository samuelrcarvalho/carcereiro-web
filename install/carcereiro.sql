-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               5.6.34 - MySQL Community Server (GPL)
-- Server OS:                    Linux
-- HeidiSQL Version:             11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for carcereiro
CREATE DATABASE IF NOT EXISTS `carcereiro` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
USE `carcereiro`;

-- Dumping structure for table carcereiro.auths
CREATE TABLE IF NOT EXISTS `auths` (
  `user` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `authcode` varchar(6) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `expire_in` datetime DEFAULT NULL,
  PRIMARY KEY (`user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping structure for table carcereiro.configs
CREATE TABLE IF NOT EXISTS `configs` (
  `config` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`config`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table carcereiro.configs: ~14 rows (approximately)
/*!40000 ALTER TABLE `configs` DISABLE KEYS */;
INSERT INTO `configs` (`config`, `value`) VALUES
	('glpi_categoryid_autoapproved', ''),
	('glpi_categoryid_toapprove', ''),
	('glpi_integracao_app-token', ''),
	('glpi_integracao_authorization', ''),
	('glpi_requesttypeid', ''),
	('glpi_url', ''),
	('rocket_integracao_auth_token', ''),
	('rocket_integracao_user_id', ''),
	('rocket_url', ''),
	('target_database_host', ''),
	('target_database_port', ''),
	('target_database_pwd', ''),
	('target_database_schema', ''),
	('target_database_user', '');
/*!40000 ALTER TABLE `configs` ENABLE KEYS */;

-- Dumping structure for table carcereiro.importanttables
CREATE TABLE IF NOT EXISTS `importanttables` (
  `nome` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`nome`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping structure for table carcereiro.tables
CREATE TABLE IF NOT EXISTS `tables` (
  `nome` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`nome`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping structure for table carcereiro.users_excepts
CREATE TABLE IF NOT EXISTS `users_excepts` (
  `user` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
