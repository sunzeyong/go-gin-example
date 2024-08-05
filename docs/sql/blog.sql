CREATE TABLE `alarms` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `identification` varchar(255) NOT NULL,
      `site` varchar(255) DEFAULT NULL,
      `target` varchar(255) DEFAULT NULL,
      `name` varchar(255) DEFAULT NULL,
      `reason` text,
      `captain` varchar(255) DEFAULT NULL,
     `company` varchar(255) DEFAULT NULL,
     `status` int(11) DEFAULT NULL,
      `create_at` datetime DEFAULT NULL,
      `alarm_start_at` datetime DEFAULT NULL,
     `alarm_end_at` datetime DEFAULT NULL,
      `duration` float DEFAULT NULL,
     `site_type` varchar(255) DEFAULT NULL,
      PRIMARY KEY (`id`),
      KEY `idx_identification` (`identification`)
    ) ENGINE=InnoDB AUTO_INCREMENT=616 DEFAULT CHARSET=utf8;