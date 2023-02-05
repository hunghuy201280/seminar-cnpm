-- -------------------------------------------------------------
-- TablePlus 5.2.2(478)
--
-- https://tableplus.com/
--
-- Database: bds
-- Generation Time: 2023-02-05 14:42:37.7850
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


CREATE DATABASE IF NOT EXISTS th2;

USE th2;


DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` int NOT NULL AUTO_INCREMENT,
  `url` text NOT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `owner_id` int NOT NULL,
  `extension` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `owner_id` (`owner_id`),
  CONSTRAINT `files_ibfk_1` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `phone` varchar(15) NOT NULL,
  `fb_id` varchar(50) DEFAULT NULL,
  `gg_id` varchar(50) DEFAULT NULL,
  `password` varchar(100) NOT NULL,
  `salt` varchar(100) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `first_name` varchar(50) NOT NULL,
  `role` enum('student','teacher','admin') NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `avatar_url` varchar(255) DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `identity_number` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;

INSERT INTO `files` (`id`, `url`, `file_name`, `owner_id`, `extension`) VALUES
(11, 'https://d2asms0h14kd1r.cloudfront.net/files/623286505.sh', '623286505.sh', 5, '.sh'),
(12, 'https://d2asms0h14kd1r.cloudfront.net/files/514645213.json', '514645213.json', 5, '.json'),
(13, 'https://d2asms0h14kd1r.cloudfront.net/files/749289514.gif', '749289514.gif', 5, '.gif');

INSERT INTO `users` (`id`, `phone`, `fb_id`, `gg_id`, `password`, `salt`, `last_name`, `first_name`, `role`, `status`, `avatar_url`, `updated_at`, `created_at`, `identity_number`) VALUES
(1, '0933970824', NULL, NULL, '$2a$10$ksd/FvByuHUjXw160adYaehLYZQNh6909NtF3.MOsFXfGGap5zQtu', '$2a$10$cYHxrOk9TUD25/JhVY3tRe8L3Eg0lIC91yF2LqllImEySsHkWTF/q', 'truong', 'huhu', 'admin', 1, 'https://d2asms0h14kd1r.cloudfront.net/img/967902633.jpg', '2023-02-01 14:42:33', '2023-01-07 10:11:42', '093758275782'),
(5, '0902514621', NULL, NULL, '$2a$10$6AInIcicHuF1ZTVKnufqB.BSD0aok5wO/qANp6qSyDW0KASS3OUPa', '$2a$10$wvHJxXdOqWhcqzct2G0WweIiohtBmuvluiBV.llrbRcAZU01.FtNm', 'le', 'thinh', 'student', 1, 'https://d2asms0h14kd1r.cloudfront.net/img/967902633.jpg', '2023-02-05 06:18:07', '2023-02-05 06:16:19', NULL);



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;