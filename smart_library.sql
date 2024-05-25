-- MySQL dump 10.13  Distrib 8.0.35, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: smart_library
-- ------------------------------------------------------
-- Server version	8.0.35-0ubuntu0.22.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */
;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */
;
/*!50503 SET NAMES utf8mb4 */
;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */
;
/*!40103 SET TIME_ZONE='+00:00' */
;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */
;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */
;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */
;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */
;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;

CREATE TABLE `books` (
    `id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `author` varchar(255) NOT NULL,
    `publisher` varchar(255) NOT NULL,
    `published_date` date NOT NULL,
    `isbn` varchar(20) NOT NULL,
    `pages` int NOT NULL,
    `language` varchar(50) NOT NULL,
    `genre` varchar(100) NOT NULL,
    `description` text,
    `card_id` int DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `fk_card_id` (`card_id`),
    CONSTRAINT `fk_card_id` FOREIGN KEY (`card_id`) REFERENCES `card_rfid` (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 36 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */
;

--
-- Dumping data for table `books`
--

/*!40000 ALTER TABLE `books` DISABLE KEYS */
;

--
-- Table structure for table `borrows`
--

DROP TABLE IF EXISTS `borrows`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;

CREATE TABLE `borrows` (
    `id` int NOT NULL AUTO_INCREMENT,
    `book_id` int DEFAULT NULL,
    `student_id` int DEFAULT NULL,
    `transaction_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `borrow_date` datetime DEFAULT CURRENT_TIMESTAMP,
    `due_date` datetime DEFAULT NULL,
    `return_date` datetime DEFAULT NULL,
    `status` ENUM(
        'pending',
        'borrowed',
        'returned'
    ) DEFAULT 'pending',
    PRIMARY KEY (`id`),
    KEY `fk_book` (`book_id`),
    KEY `fk_student` (`student_id`),
    CONSTRAINT `fk_book` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`),
    CONSTRAINT `fk_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 19 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */
;

--
-- Dumping data for table `borrows`
--

/*!40000 ALTER TABLE `borrows` DISABLE KEYS */
;

--
-- Table structure for table `card_rfid`
--

DROP TABLE IF EXISTS `card_rfid`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;

CREATE TABLE `card_rfid` (
    `id` int NOT NULL AUTO_INCREMENT,
    `uid` varchar(100) DEFAULT NULL,
    `type` enum('book', 'student') DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 30 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */
;

--
-- Dumping data for table `card_rfid`
--

/*!40000 ALTER TABLE `card_rfid` DISABLE KEYS */
;

--
-- Table structure for table `students`
--

DROP TABLE IF EXISTS `students`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;

CREATE TABLE `students` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(30) DEFAULT NULL,
    `npm` varchar(8) DEFAULT NULL,
    `card_id` int DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `fk_student_card_id` (`card_id`),
    CONSTRAINT `fk_student_card_id` FOREIGN KEY (`card_id`) REFERENCES `card_rfid` (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 4 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */
;

--
-- Dumping data for table `students`
--

/*!40000 ALTER TABLE `students` DISABLE KEYS */
;

--
-- Dumping routines for database 'smart_library'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */
;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */
;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */
;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */
;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */
;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */
;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */
;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */
;

-- Dump completed on 2024-05-24 21:09:40