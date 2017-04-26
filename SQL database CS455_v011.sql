-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.1.19-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win32
-- HeidiSQL Version:             9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for testcs455
CREATE DATABASE IF NOT EXISTS `testcs455` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `testcs455`;

-- Dumping structure for table testcs455.catalog
CREATE TABLE IF NOT EXISTS `catalog` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `year` varchar(50) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- Dumping data for table testcs455.catalog: ~0 rows (approximately)
/*!40000 ALTER TABLE `catalog` DISABLE KEYS */;
INSERT INTO `catalog` (`id`, `year`) VALUES
	(1, '2016-2017');
/*!40000 ALTER TABLE `catalog` ENABLE KEYS */;

-- Dumping structure for table testcs455.cs_major
CREATE TABLE IF NOT EXISTS `cs_major` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dept` varchar(50) NOT NULL DEFAULT '0',
  `course` varchar(50) NOT NULL DEFAULT '0',
  `credits` int(11) NOT NULL DEFAULT '0',
  `year` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=latin1;

-- Dumping data for table testcs455.cs_major: ~32 rows (approximately)
/*!40000 ALTER TABLE `cs_major` DISABLE KEYS */;
INSERT INTO `cs_major` (`id`, `dept`, `course`, `credits`, `year`) VALUES
	(1, 'ma', '125', 4, 1),
	(2, 'ma', '126', 4, 1),
	(3, 'ma', '345', 3, 1),
	(4, 'ma', '237', 3, 1),
	(5, 'ma', '431', 3, 1),
	(6, 'ma', '227', 3, 1),
	(7, 'cs', '155', 3, 1),
	(8, 'cs', '245', 3, 1),
	(9, 'cs', '255', 3, 1),
	(10, 'cs', '310', 3, 1),
	(11, 'cs', '311', 3, 1),
	(12, 'cs', '355', 3, 1),
	(13, 'cs', '410', 3, 1),
	(14, 'cs', '420', 3, 1),
	(15, 'cs', '455', 3, 1),
	(16, 'cs', '315', 3, 1),
	(17, 'cs', '325', 3, 1),
	(18, 'cs', '335', 3, 1),
	(19, 'cs', '390', 3, 1),
	(20, 'cis', '315', 3, 1),
	(21, 'cs', '360', 3, 1),
	(22, 'cis', '344', 3, 1),
	(23, 'cs', '447', 3, 1),
	(24, 'cis', '445', 3, 1),
	(25, 'cs', '421', 3, 1),
	(26, 'cs', '430', 3, 1),
	(27, 'cs', '470', 3, 1),
	(28, 'cs', '249', 3, 1),
	(29, 'cs', '480', 3, 1),
	(30, 'cs', '490', 3, 1),
	(31, 'cis', '486', 3, 1),
	(32, 'cis', '489', 3, 1);
/*!40000 ALTER TABLE `cs_major` ENABLE KEYS */;

-- Dumping structure for table testcs455.cs_minor
CREATE TABLE IF NOT EXISTS `cs_minor` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dept` varchar(50) NOT NULL DEFAULT '0',
  `course` varchar(50) NOT NULL DEFAULT '0',
  `credit` int(11) NOT NULL DEFAULT '0',
  `year` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;

-- Dumping data for table testcs455.cs_minor: ~13 rows (approximately)
/*!40000 ALTER TABLE `cs_minor` DISABLE KEYS */;
INSERT INTO `cs_minor` (`id`, `dept`, `course`, `credit`, `year`) VALUES
	(1, 'cs', '155', 3, 1),
	(2, 'cs', '255', 3, 1),
	(3, 'cs', '355', 3, 1),
	(4, 'cs', '315', 3, 1),
	(5, 'cs', '325', 3, 1),
	(6, 'cs', '335', 3, 1),
	(7, 'cs', '390', 3, 1),
	(8, 'cis', '315', 3, 1),
	(9, 'cs', '249', 3, 1),
	(10, 'cs', '480', 3, 1),
	(11, 'cs', '490', 3, 1),
	(12, 'cis', '486', 3, 1),
	(13, 'cis', '489', 3, 1);
/*!40000 ALTER TABLE `cs_minor` ENABLE KEYS */;

-- Dumping structure for table testcs455.general
CREATE TABLE IF NOT EXISTS `general` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dept` varchar(50) NOT NULL DEFAULT '0',
  `course` varchar(50) NOT NULL DEFAULT '0',
  `credits` int(11) NOT NULL DEFAULT '0',
  `year` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=67 DEFAULT CHARSET=latin1;

-- Dumping data for table testcs455.general: ~66 rows (approximately)
/*!40000 ALTER TABLE `general` DISABLE KEYS */;
INSERT INTO `general` (`id`, `dept`, `course`, `credits`, `year`) VALUES
	(1, 'en', '111', 3, 1),
	(2, 'en', '112', 3, 1),
	(3, 'en', '121', 3, 1),
	(4, 'en', '122', 3, 1),
	(5, 'com', '201', 3, 1),
	(6, 'en', '211', 3, 1),
	(7, 'en', '212', 3, 1),
	(8, 'en', '221', 3, 1),
	(9, 'en', '231', 3, 1),
	(10, 'en', '232', 3, 1),
	(11, 'en', '233', 3, 1),
	(12, 'en', '234', 3, 1),
	(13, 'ar', '170', 3, 1),
	(14, 'ar', '281', 3, 1),
	(15, 'ar', '282', 3, 1),
	(16, 'fr', '101', 4, 1),
	(17, 'fr', '102', 4, 1),
	(18, 'fr', '201', 3, 1),
	(19, 'fr', '202', 3, 1),
	(20, 'gr', '101', 4, 1),
	(21, 'gr', '102', 4, 1),
	(22, 'gr', '201', 3, 1),
	(23, 'gr', '202', 3, 1),
	(24, 'mu', '222', 3, 1),
	(25, 'mu', '244', 3, 1),
	(26, 'phl', '201', 3, 1),
	(27, 'phl', '205', 3, 1),
	(28, 're', '221', 3, 1),
	(29, 're', '231', 3, 1),
	(30, 'sp', '101', 4, 1),
	(31, 'sp', '102', 4, 1),
	(32, 'sp', '201', 3, 1),
	(33, 'sp', '202', 3, 1),
	(34, 'th', '210', 3, 1),
	(35, 'ma', '112', 3, 1),
	(36, 'ma', '113', 3, 1),
	(37, 'ma', '115', 4, 1),
	(38, 'ma', '125', 4, 1),
	(39, 'ma', '126', 4, 1),
	(40, 'ma', '227', 4, 1),
	(41, 'ma', '237', 3, 1),
	(42, 'ma', '238', 3, 1),
	(43, 'bi', '111', 4, 1),
	(44, 'bi', '112', 4, 1),
	(45, 'ch', '111', 4, 1),
	(46, 'ch', '112', 4, 1),
	(47, 'es', '131', 4, 1),
	(48, 'es', '132', 4, 1),
	(49, 'es', '133', 4, 1),
	(50, 'ge', '111', 4, 1),
	(51, 'ge', '112', 4, 1),
	(52, 'ph', '251', 5, 1),
	(53, 'ph', '252', 5, 1),
	(54, 'hi', '101', 3, 1),
	(55, 'hi', '102', 3, 1),
	(56, 'hi', '201', 3, 1),
	(57, 'hi', '202', 3, 1),
	(58, 'ec', '251', 3, 1),
	(59, 'ec', '252', 3, 1),
	(60, 'ed', '299', 3, 1),
	(61, 'ge', '102', 3, 1),
	(62, 'ge', '260', 3, 1),
	(63, 'ps', '241', 3, 1),
	(64, 'py', '201', 3, 1),
	(65, 'so', '221', 3, 1),
	(66, 'so', '222', 3, 1);
/*!40000 ALTER TABLE `general` ENABLE KEYS */;

-- Dumping structure for table testcs455.user
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` text NOT NULL,
  `password` text NOT NULL,
  `department` text NOT NULL,
  `firstname` text NOT NULL,
  `lastname` text NOT NULL,
  `email` text NOT NULL,
  `supervisor` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- Dumping data for table testcs455.user: ~0 rows (approximately)
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` (`id`, `username`, `password`, `department`, `firstname`, `lastname`, `email`, `supervisor`) VALUES
	(1, 'john', '12345', 'cs', 'john', 'smith', 'jsmith@una.edu', 0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
