# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: localhost (MySQL 5.7.17)
# Database: testcs455
# Generation Time: 2017-04-28 17:41:34 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table catalog
# ------------------------------------------------------------

DROP TABLE IF EXISTS `catalog`;

CREATE TABLE `catalog` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `year` varchar(50) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `catalog` WRITE;
/*!40000 ALTER TABLE `catalog` DISABLE KEYS */;

INSERT INTO `catalog` (`id`, `year`)
VALUES
	(1,'2016-2017');

/*!40000 ALTER TABLE `catalog` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table cis_major
# ------------------------------------------------------------

DROP TABLE IF EXISTS `cis_major`;

CREATE TABLE `cis_major` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dept` varchar(50) NOT NULL DEFAULT '0',
  `course` varchar(50) NOT NULL DEFAULT '0',
  `credits` int(11) NOT NULL DEFAULT '0',
  `year` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `cis_major` WRITE;
/*!40000 ALTER TABLE `cis_major` DISABLE KEYS */;

INSERT INTO `cis_major` (`id`, `dept`, `course`, `credits`, `year`)
VALUES
	(1,'cs','135',3,1),
	(2,'cis','225',3,1),
	(3,'cis','315',3,1),
	(4,'cis','330',3,1),
	(5,'cis','366',3,1),
	(6,'cis','376',3,1),
	(7,'cis','344',3,1),
	(8,'cis','476',3,1),
	(9,'cis','486',3,1),
	(10,'cis','289',3,1),
	(11,'cis','249',3,1),
	(12,'cis','444',3,1),
	(13,'cis','445',3,1),
	(14,'cis','446',3,1),
	(15,'cis','480',3,1),
	(16,'fi','393',3,1),
	(17,'mg','330',3,1),
	(18,'mg','382',3,1),
	(19,'mg','395',3,1),
	(20,'mg','491',3,1),
	(21,'mg','498',3,1),
	(22,'mk','360',3,1),
	(23,'ac','291',3,1),
	(24,'ac','292',3,1),
	(25,'bl','240',3,1),
	(26,'cis','125',3,1),
	(27,'qm','291',3,1),
	(28,'qm','292',3,1);

/*!40000 ALTER TABLE `cis_major` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table cis_minor
# ------------------------------------------------------------

DROP TABLE IF EXISTS `cis_minor`;

CREATE TABLE `cis_minor` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dept` varchar(50) NOT NULL DEFAULT '0',
  `course` varchar(50) NOT NULL DEFAULT '0',
  `credits` int(11) NOT NULL DEFAULT '0',
  `year` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `cis_minor` WRITE;
/*!40000 ALTER TABLE `cis_minor` DISABLE KEYS */;

INSERT INTO `cis_minor` (`id`, `dept`, `course`, `credits`, `year`)
VALUES
	(1,'cis','225',3,1),
	(2,'cis','236',3,1),
	(3,'cis','330',3,1),
	(4,'cis','366',3,1),
	(5,'cis','249',3,1),
	(6,'cis','289',3,1),
	(7,'cis','444',3,1),
	(8,'cis','445',3,1),
	(9,'cis','480',3,1);

/*!40000 ALTER TABLE `cis_minor` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table cs_major
# ------------------------------------------------------------

DROP TABLE IF EXISTS `cs_major`;

CREATE TABLE `cs_major` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dept` varchar(50) NOT NULL DEFAULT '0',
  `course` varchar(50) NOT NULL DEFAULT '0',
  `credits` int(11) NOT NULL DEFAULT '0',
  `year` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `cs_major` WRITE;
/*!40000 ALTER TABLE `cs_major` DISABLE KEYS */;

INSERT INTO `cs_major` (`id`, `dept`, `course`, `credits`, `year`)
VALUES
	(1,'ma','125',4,1),
	(2,'ma','126',4,1),
	(3,'ma','345',3,1),
	(4,'ma','237',3,1),
	(5,'ma','431',3,1),
	(6,'ma','227',3,1),
	(7,'cs','155',3,1),
	(8,'cs','245',3,1),
	(9,'cs','255',3,1),
	(10,'cs','310',3,1),
	(11,'cs','311',3,1),
	(12,'cs','355',3,1),
	(13,'cs','410',3,1),
	(14,'cs','420',3,1),
	(15,'cs','455',3,1),
	(16,'cs','315',3,1),
	(17,'cs','325',3,1),
	(18,'cs','335',3,1),
	(19,'cs','390',3,1),
	(20,'cis','315',3,1),
	(21,'cs','360',3,1),
	(22,'cis','344',3,1),
	(23,'cs','447',3,1),
	(24,'cis','445',3,1),
	(25,'cs','421',3,1),
	(26,'cs','430',3,1),
	(27,'cs','470',3,1),
	(28,'cs','249',3,1),
	(29,'cs','480',3,1),
	(30,'cs','490',3,1),
	(31,'cis','486',3,1),
	(32,'cis','489',3,1);

/*!40000 ALTER TABLE `cs_major` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table cs_minor
# ------------------------------------------------------------

DROP TABLE IF EXISTS `cs_minor`;

CREATE TABLE `cs_minor` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dept` varchar(50) NOT NULL DEFAULT '0',
  `course` varchar(50) NOT NULL DEFAULT '0',
  `credit` int(11) NOT NULL DEFAULT '0',
  `year` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `cs_minor` WRITE;
/*!40000 ALTER TABLE `cs_minor` DISABLE KEYS */;

INSERT INTO `cs_minor` (`id`, `dept`, `course`, `credit`, `year`)
VALUES
	(1,'cs','155',3,1),
	(2,'cs','255',3,1),
	(3,'cs','355',3,1),
	(4,'cs','315',3,1),
	(5,'cs','325',3,1),
	(6,'cs','335',3,1),
	(7,'cs','390',3,1),
	(8,'cis','315',3,1),
	(9,'cs','249',3,1),
	(10,'cs','480',3,1),
	(11,'cs','490',3,1),
	(12,'cis','486',3,1),
	(13,'cis','489',3,1);

/*!40000 ALTER TABLE `cs_minor` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table general
# ------------------------------------------------------------

DROP TABLE IF EXISTS `general`;

CREATE TABLE `general` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dept` varchar(50) NOT NULL DEFAULT '0',
  `course` varchar(50) NOT NULL DEFAULT '0',
  `credits` int(11) NOT NULL DEFAULT '0',
  `year` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `general` WRITE;
/*!40000 ALTER TABLE `general` DISABLE KEYS */;

INSERT INTO `general` (`id`, `dept`, `course`, `credits`, `year`)
VALUES
	(1,'en','111',3,1),
	(2,'en','112',3,1),
	(3,'en','121',3,1),
	(4,'en','122',3,1),
	(5,'com','201',3,1),
	(6,'en','211',3,1),
	(7,'en','212',3,1),
	(8,'en','221',3,1),
	(9,'en','231',3,1),
	(10,'en','232',3,1),
	(11,'en','233',3,1),
	(12,'en','234',3,1),
	(13,'ar','170',3,1),
	(14,'ar','281',3,1),
	(15,'ar','282',3,1),
	(16,'fr','101',4,1),
	(17,'fr','102',4,1),
	(18,'fr','201',3,1),
	(19,'fr','202',3,1),
	(20,'gr','101',4,1),
	(21,'gr','102',4,1),
	(22,'gr','201',3,1),
	(23,'gr','202',3,1),
	(24,'mu','222',3,1),
	(25,'mu','244',3,1),
	(26,'phl','201',3,1),
	(27,'phl','205',3,1),
	(28,'re','221',3,1),
	(29,'re','231',3,1),
	(30,'sp','101',4,1),
	(31,'sp','102',4,1),
	(32,'sp','201',3,1),
	(33,'sp','202',3,1),
	(34,'th','210',3,1),
	(35,'ma','112',3,1),
	(36,'ma','113',3,1),
	(37,'ma','115',4,1),
	(38,'ma','125',4,1),
	(39,'ma','126',4,1),
	(40,'ma','227',4,1),
	(41,'ma','237',3,1),
	(42,'ma','238',3,1),
	(43,'bi','111',4,1),
	(44,'bi','112',4,1),
	(45,'ch','111',4,1),
	(46,'ch','112',4,1),
	(47,'es','131',4,1),
	(48,'es','132',4,1),
	(49,'es','133',4,1),
	(50,'ge','111',4,1),
	(51,'ge','112',4,1),
	(52,'ph','251',5,1),
	(53,'ph','252',5,1),
	(54,'hi','101',3,1),
	(55,'hi','102',3,1),
	(56,'hi','201',3,1),
	(57,'hi','202',3,1),
	(58,'ec','251',3,1),
	(59,'ec','252',3,1),
	(60,'ed','299',3,1),
	(61,'ge','102',3,1),
	(62,'ge','260',3,1),
	(63,'ps','241',3,1),
	(64,'py','201',3,1),
	(65,'so','221',3,1),
	(66,'so','222',3,1);

/*!40000 ALTER TABLE `general` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table general_area1
# ------------------------------------------------------------

DROP TABLE IF EXISTS `general_area1`;

CREATE TABLE `general_area1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dept` varchar(50) NOT NULL DEFAULT '0',
  `course` varchar(50) NOT NULL DEFAULT '0',
  `credits` int(11) NOT NULL DEFAULT '0',
  `year` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=COMPACT;

LOCK TABLES `general_area1` WRITE;
/*!40000 ALTER TABLE `general_area1` DISABLE KEYS */;

INSERT INTO `general_area1` (`id`, `dept`, `course`, `credits`, `year`)
VALUES
	(1,'en','111',3,1),
	(2,'en','112',3,1),
	(3,'en','121',3,1),
	(4,'en','122',3,1);

/*!40000 ALTER TABLE `general_area1` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table general_area2
# ------------------------------------------------------------

DROP TABLE IF EXISTS `general_area2`;

CREATE TABLE `general_area2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dept` varchar(50) NOT NULL DEFAULT '0',
  `course` varchar(50) NOT NULL DEFAULT '0',
  `credits` int(11) NOT NULL DEFAULT '0',
  `year` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=COMPACT;

LOCK TABLES `general_area2` WRITE;
/*!40000 ALTER TABLE `general_area2` DISABLE KEYS */;

INSERT INTO `general_area2` (`id`, `dept`, `course`, `credits`, `year`)
VALUES
	(1,'com','201',3,1),
	(2,'en','211',3,1),
	(3,'en','212',3,1),
	(4,'en','221',3,1),
	(5,'en','231',3,1),
	(6,'en','232',3,1),
	(7,'en','233',3,1),
	(8,'en','234',3,1),
	(9,'ar','170',3,1),
	(10,'ar','281',3,1),
	(11,'ar','282',3,1),
	(12,'fr','101',4,1),
	(13,'fr','102',4,1),
	(14,'fr','201',3,1),
	(15,'fr','202',3,1),
	(16,'gr','101',4,1),
	(17,'gr','102',4,1),
	(18,'gr','201',3,1),
	(19,'gr','202',3,1),
	(20,'mu','222',3,1),
	(21,'mu','244',3,1),
	(22,'phl','201',3,1),
	(23,'phl','205',3,1),
	(24,'re','221',3,1),
	(25,'re','231',3,1),
	(26,'sp','101',4,1),
	(27,'sp','102',4,1),
	(28,'sp','201',3,1),
	(29,'sp','202',3,1),
	(30,'th','210',3,1);

/*!40000 ALTER TABLE `general_area2` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table general_area3
# ------------------------------------------------------------

DROP TABLE IF EXISTS `general_area3`;

CREATE TABLE `general_area3` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dept` varchar(50) NOT NULL DEFAULT '0',
  `course` varchar(50) NOT NULL DEFAULT '0',
  `credits` int(11) NOT NULL DEFAULT '0',
  `year` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=COMPACT;

LOCK TABLES `general_area3` WRITE;
/*!40000 ALTER TABLE `general_area3` DISABLE KEYS */;

INSERT INTO `general_area3` (`id`, `dept`, `course`, `credits`, `year`)
VALUES
	(1,'ma','112',3,1),
	(2,'ma','113',3,1),
	(3,'ma','115',4,1),
	(4,'ma','125',4,1),
	(5,'ma','126',4,1),
	(6,'ma','227',4,1),
	(7,'ma','237',3,1),
	(8,'ma','238',3,1),
	(9,'bi','111',4,1),
	(10,'bi','112',4,1),
	(11,'ch','111',4,1),
	(12,'ch','112',4,1),
	(13,'es','131',4,1),
	(14,'es','132',4,1),
	(15,'es','133',4,1),
	(16,'ge','111',4,1),
	(17,'ge','112',4,1),
	(18,'ph','251',5,1),
	(19,'ph','252',5,1);

/*!40000 ALTER TABLE `general_area3` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table general_area4
# ------------------------------------------------------------

DROP TABLE IF EXISTS `general_area4`;

CREATE TABLE `general_area4` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dept` varchar(50) NOT NULL DEFAULT '0',
  `course` varchar(50) NOT NULL DEFAULT '0',
  `credits` int(11) NOT NULL DEFAULT '0',
  `year` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=COMPACT;

LOCK TABLES `general_area4` WRITE;
/*!40000 ALTER TABLE `general_area4` DISABLE KEYS */;

INSERT INTO `general_area4` (`id`, `dept`, `course`, `credits`, `year`)
VALUES
	(1,'hi','101',3,1),
	(2,'hi','102',3,1),
	(3,'hi','201',3,1),
	(4,'hi','202',3,1),
	(5,'ec','251',3,1),
	(6,'ec','252',3,1),
	(7,'ed','299',3,1),
	(8,'ge','102',3,1),
	(9,'ge','260',3,1),
	(10,'ps','241',3,1),
	(11,'py','201',3,1),
	(12,'so','221',3,1),
	(13,'so','222',3,1);

/*!40000 ALTER TABLE `general_area4` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` text NOT NULL,
  `password` text NOT NULL,
  `department` text NOT NULL,
  `firstname` text NOT NULL,
  `lastname` text NOT NULL,
  `email` text NOT NULL,
  `supervisor` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;

INSERT INTO `user` (`id`, `username`, `password`, `department`, `firstname`, `lastname`, `email`, `supervisor`)
VALUES
	(1,'john','12345','cs','john','smith','jsmith@una.edu',0);

/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
