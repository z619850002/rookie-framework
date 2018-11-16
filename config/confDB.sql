create database DBCONFIG;
USE DBCONFIG;

CREATE TABLE `params` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `value` varchar(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


INSERT INTO `params` (`name`,`value`)VALUES
    ('RedisAddr','172.26.163.124:6379'),
    ('RedisPassword','ztjztj120'),
    ('RedisDB','0'),
    ('MysqlAddr','172.26.163.76'),
    ('MysqlPort','3306'),
    ('MysqlUser','rookie2'),
    ('MysqlPassword','12345678'),
    ('MysqlDB','sausage_db'),
    ('Port','10000'),
    ('WriteWait','10'),
    ('PongWait','60'),
    ('PingPeriod','54'),
    ('MaxMessageSize','512'),
    ('MessageBufferSize','256');

CREATE TABLE `dbinfo` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `datasourcename` varchar(45) DEFAULT NULL,
  `dbhost` varchar(45) DEFAULT NULL,
  `dbport` varchar(45) DEFAULT NULL,
  `dbuser` varchar(45) DEFAULT NULL,
  `dbpassword` varchar(45) DEFAULT NULL,
  `dbname` varchar(45) DEFAULT NULL,
  `dbtype` varchar(50) DEFAULT NULL ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `cacheinfo` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `addr` varchar(45) DEFAULT NULL,
  `password` varchar(45) DEFAULT NULL,
  `db` int DEFAULT NULL,
  `name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


INSERT INTO `dbinfo` (`datasourcename`,`dbhost`,`dbport`,`dbuser`,`dbpassword`,`dbname`) VALUES ('mysql','172.26.163.76','3306','rookie2','12345678','sausage_db');