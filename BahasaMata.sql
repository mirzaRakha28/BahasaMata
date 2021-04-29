-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 29, 2021 at 05:31 PM
-- Server version: 10.4.18-MariaDB
-- PHP Version: 8.0.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bahasa_mata`
--

-- --------------------------------------------------------

--
-- Table structure for table `alarm`
--
CREATE DATABASE bahasa_mata;
USE  bahasa_mata;
CREATE TABLE `alarm` (
  `id` int(11) NOT NULL,
  `id_difabel` int(11) NOT NULL,
  `id_perawat` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `jam` time NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `alarm`
--

INSERT INTO `alarm` (`id`, `id_difabel`, `id_perawat`, `title`, `jam`) VALUES
(3, 13, 12, '11', '00:00:01'),
(4, 1, 1, '1', '00:00:00'),
(5, 1, 1, '1', '00:00:24');

-- --------------------------------------------------------

--
-- Table structure for table `difabel`
--

CREATE TABLE `difabel` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `difabel`
--

INSERT INTO `difabel` (`id`, `username`, `password`, `name`) VALUES
(13, 'mirzarakha12332s33rre', 'b95b2313abc196917f977003177438df', 'mirza rakha12'),
(14, 'mirzarakha12332s33rre1', 'b95b2313abc196917f977003177438df', 'mirza rakha12');

-- --------------------------------------------------------

--
-- Table structure for table `emergency_call`
--

CREATE TABLE `emergency_call` (
  `id` int(11) NOT NULL,
  `id_difabel` int(11) NOT NULL,
  `id_perawat` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `no_tlp` time NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `notifikasi`
--

CREATE TABLE `notifikasi` (
  `id_perawat` int(11) NOT NULL,
  `id_difabel` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `pasien`
--

CREATE TABLE `pasien` (
  `id_perawat` int(11) NOT NULL,
  `id_difabel` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `perawat`
--

CREATE TABLE `perawat` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `perawat`
--

INSERT INTO `perawat` (`id`, `username`, `password`, `name`) VALUES
(1, 'mirzarakha12332s33rre', 'b95b2313abc196917f977003177438df', 'mirza rakha12');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `alarm`
--
ALTER TABLE `alarm`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `difabel`
--
ALTER TABLE `difabel`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `emergency_call`
--
ALTER TABLE `emergency_call`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `perawat`
--
ALTER TABLE `perawat`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `alarm`
--
ALTER TABLE `alarm`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `difabel`
--
ALTER TABLE `difabel`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `emergency_call`
--
ALTER TABLE `emergency_call`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `perawat`
--
ALTER TABLE `perawat`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
