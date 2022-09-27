-- MySQL dump 10.13  Distrib 5.7.34, for osx10.12 (x86_64)
--
-- Host: 127.0.0.1    Database: user
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
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `userIdentity` varchar(255) DEFAULT '' COMMENT '用户唯一标识',
  `userName` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `passWord` varchar(50) NOT NULL DEFAULT '' COMMENT '用户密码，MD5加密',
  `userNick` varchar(100) DEFAULT '' COMMENT '用户昵称',
  `userFace` varchar(255) DEFAULT '' COMMENT '用户头像地址',
  `UserSex` tinyint(1) DEFAULT '0' COMMENT '用户性别：0男，1女，2保密',
  `userEmail` varchar(255) DEFAULT '' COMMENT '用户邮箱',
  `userPhone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `loginAddress` varchar(255) DEFAULT '' COMMENT '用户登录IP地址',
  `create_time` timestamp default CURRENT_TIMESTAMP not null comment '创建时间',
  `update_time` timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `userName` (`userName`),
  UNIQUE KEY `userPhone` (`userPhone`),
  KEY `updateTime` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_receive_address`
--

DROP TABLE IF EXISTS `user_receive_address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_receive_address` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '收货人名称',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否为默认地址',
  `post_code` varchar(100) NOT NULL DEFAULT '' COMMENT '邮政编码',
  `province` varchar(100) NOT NULL DEFAULT '' COMMENT '省份/直辖市',
  `city` varchar(100) NOT NULL DEFAULT '' COMMENT '城市',
  `region` varchar(100) NOT NULL DEFAULT '' COMMENT '区',
  `detail_address` varchar(128) NOT NULL DEFAULT '' COMMENT '详细地址(街道)',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '数据创建时间[禁止在代码中赋值]',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '数据更新时间[禁止在代码中赋值]',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户收货地址表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_receive_address`
--

LOCK TABLES `user_receive_address` WRITE;
/*!40000 ALTER TABLE `user_receive_address` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_receive_address` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-08-24 16:52:40
