-- MySQL dump 10.13  Distrib 5.7.34, for osx10.12 (x86_64)
--
-- Host: 127.0.0.1    Database: product
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
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `category` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `parent_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '父类别id当id=0时说明是根节点,一级类别',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '类别名称',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '类别状态1-正常,2-已废弃',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='商品类别表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,0,'男装',1,'2022-06-20 16:19:21','2022-06-20 16:19:21'),(2,1,'夹克',1,'2022-06-20 16:19:21','2022-06-20 16:19:21'),(3,1,'中山装',1,'2022-06-20 16:19:21','2022-06-20 16:19:21'),(4,0,'家电',1,'2022-06-20 16:19:21','2022-06-20 16:19:21'),(5,4,'洗衣机',1,'2022-06-20 16:19:21','2022-06-20 16:19:21'),(6,4,'冰箱',1,'2022-06-20 16:19:21','2022-06-20 16:19:21'),(7,4,'电风扇',1,'2022-06-20 16:19:21','2022-06-20 16:19:21'),(8,4,'电视',1,'2022-06-20 16:19:21','2022-06-20 16:19:21');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品id',
  `cateId` smallint(6) unsigned NOT NULL DEFAULT '0' COMMENT '类别Id',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '商品名称',
  `property` varchar(255) NOT NULL DEFAULT '' COMMENT '商品属性',
  `coverPic` varchar(255) NOT NULL DEFAULT '' COMMENT '商品封面',
  `images` mediumtext NOT NULL COMMENT '图片地址,逗号分隔',
  `detail` longtext NOT NULL COMMENT '商品详情',
  `mtPrice` decimal(20,2) NOT NULL DEFAULT '0.00' COMMENT '市场价格,单位-元保留两位小数',
  `disPrice` decimal(20,2) NOT NULL DEFAULT '0.00' COMMENT '折扣价格, 单位元-保留两位小数',
  `stock` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '库存数量',
  `state` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '商品状态.1-在售 2-下架 3-删除',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `salesVolume` int(11) NOT NULL DEFAULT '0' COMMENT '销量',
  PRIMARY KEY (`id`),
  KEY `cate_id` (`cateId`),
  KEY `name` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-10-05 22:24:26
