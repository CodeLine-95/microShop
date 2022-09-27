-- MySQL dump 10.13  Distrib 5.7.34, for osx10.12 (x86_64)
--
-- Host: 127.0.0.1    Database: home
-- ------------------------------------------------------
-- Server version	5.7.34

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `banner`
--

DROP TABLE IF EXISTS `banner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `banner` (
                          `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                          `title` varchar(255) NOT NULL DEFAULT '',
                          `href` varchar(255) NOT NULL DEFAULT '',
                          `target` varchar(20) NOT NULL DEFAULT '',
                          `pic` varchar(255) NOT NULL DEFAULT '',
                          `status` tinyint(1) DEFAULT '0',
                          `createTimeTamp` int(11) NOT NULL DEFAULT '0',
                          `updateTimeTamp` int(11) NOT NULL DEFAULT '0',
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `banner_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='首页banner';
/*!40101 SET character_set_client = @saved_cs_client */;

