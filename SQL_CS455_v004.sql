-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.1.21-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win32
-- HeidiSQL Version:             9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for test
CREATE DATABASE IF NOT EXISTS `test` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `test`;

-- Dumping structure for table test.catalog
CREATE TABLE IF NOT EXISTS `catalog` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `year` varchar(11) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

-- Dumping data for table test.catalog: ~6 rows (approximately)
/*!40000 ALTER TABLE `catalog` DISABLE KEYS */;
INSERT INTO `catalog` (`id`, `year`) VALUES
	(1, '2016-2017'),
	(2, '2015-2016'),
	(3, '2014-2015'),
	(4, '2013-2014'),
	(5, '2012-2013'),
	(6, '2011-2012');
/*!40000 ALTER TABLE `catalog` ENABLE KEYS */;

-- Dumping structure for table test.general
CREATE TABLE IF NOT EXISTS `general` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `courses` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- Dumping data for table test.general: ~0 rows (approximately)
/*!40000 ALTER TABLE `general` DISABLE KEYS */;
INSERT INTO `general` (`id`, `courses`) VALUES
	(1, 'en111:3,en112:3,en121:3,en122:3,com201:3,en211:3,en212:3,en221:3,en222:3,en231:2,en232:3,en233:3,en234:3, ar170:3,ar281:3,ar282:3,com133:3,fl100:3,fr101:3,fr102:3,fr201:3,fr202:3,gr101:3,gr102:3,gr201:3,gr202:3,mu222:3,mu244:3,phl201:3,phl205:3,phl250:3,re111:3,re221:3,re231:3,sp101:3,sp102:3,sp201:3,sp202:3,th210:3,ma110:3,ma111:3,ma112:3,ma113:3,ma115:4,ma125:4,ma126:3,ma147:3,ma227:4,ma237:3,ma238:3,bl101:4,bl102:4,bl111:4,bl112:4,ch101:3,ch101l:1,ch102:3,ch102l:1,ch111:3,ch111l:1,ch112:3,ch112l:1,es131:4,es132:4,es133:4,ge111:4,ge112:4,ph101:4,ph121:4,ph125:4,ph241:4,ph242:4,ph251:5,ph252:5,hi101:3,hi102:3,hi201:3,hi202:3,com205:3,ec251:3,ec252:3,ed299:3,fl101:3,fl101h:3,fl201:3,fl204:3,ge102:3,ge260:3,hpe175:3,hpe213:3,ps241:3,ps251:3,py201:3,so221:3,so222:3,srm200:3');
/*!40000 ALTER TABLE `general` ENABLE KEYS */;

-- Dumping structure for table test.major
CREATE TABLE IF NOT EXISTS `major` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `majorName` varchar(30) DEFAULT '',
  `courses` text NOT NULL,
  `year` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

-- Dumping data for table test.major: ~7 rows (approximately)
/*!40000 ALTER TABLE `major` DISABLE KEYS */;
INSERT INTO `major` (`id`, `majorName`, `courses`, `year`) VALUES
	(1, 'Computer Science', 'cs155:3,cs245:3,cs255:3,cs310:3,cs311:3,cs355:3,cs410:4,cs420:3,cs455:3,cs315:3,cs325:3,cs335:3,cs390:3,cis315:3,cs360:3,cis344:3,cs421:3,cs430:3,cs447:3,cis445:3,cs470:3,cs480:3,cis486:3,cis489:3,cs490:3', 1),
	(2, 'Math', 'ma125:4,ma126:4,ma227:4,ma238:3,ma355:3,ma325:3,ma345:3,ma447:3,ma431:3,ma437:3,ma451:3,ma471W:3,cs155:3,ma(344-491):3', 1),
	(3, 'Computer Information Systems', 'cis135:3,cis225:3,cis315:3,cis330:3,cis366:3,cis376:3,cis344:3,cis476:3,cis486:3,cis289:3,cis249:3,cis444:3,cis445:3,cis446:3,cis480:3', 1),
	(4, 'Computer Science', '', 2),
	(5, 'Math', '', 2),
	(6, 'Computer Information Systems', '', 2),
	(7, 'Computer Science', '', 3);
/*!40000 ALTER TABLE `major` ENABLE KEYS */;

-- Dumping structure for table test.minor
CREATE TABLE IF NOT EXISTS `minor` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `minorName` varchar(30) DEFAULT NULL,
  `courses` text NOT NULL,
  `year` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- Dumping data for table test.minor: ~3 rows (approximately)
/*!40000 ALTER TABLE `minor` DISABLE KEYS */;
INSERT INTO `minor` (`id`, `minorName`, `courses`, `year`) VALUES
	(1, 'Computer Science', 'cs155:3,cs255:3,cs355:3', 1),
	(2, 'Math', 'ma125:3,ma126:3,ma227:3', 1),
	(3, 'Computer Information Systems', 'cs225:3,cis236:3,cis330:3,cis366:3,', 1);
/*!40000 ALTER TABLE `minor` ENABLE KEYS */;

-- Dumping structure for table test.user
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL,
  `username` text NOT NULL,
  `password` text NOT NULL,
  `department` text NOT NULL,
  `firstname` text NOT NULL,
  `lastname` text NOT NULL,
  `email` text NOT NULL,
  `supervisor` tinyint(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Dumping data for table test.user: ~0 rows (approximately)
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` (`id`, `username`, `password`, `department`, `firstname`, `lastname`, `email`, `supervisor`) VALUES
	(1, 'john', '12345', 'cs', 'john', 'smith', 'jsmith@una.edu', 0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
