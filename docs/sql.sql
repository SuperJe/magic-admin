CREATE TABLE `courses` (
    `id` bigint NOT NULL AUTO_INCREMENT,
   `name` varchar(255) NOT NULL,
    `course_type` tinyint(4) NOT NULL,
    `total_lesson_hours` int NOT NULL,
   `pre_required` varchar(255) DEFAULT NULL,
   `target` varchar(255) DEFAULT NULL,
   `recommend_competition` text DEFAULT NULL,
    `recommend_period` varchar(255) DEFAULT NULL,
   `created` datetime(3) DEFAULT NULL COMMENT '创建时间',
   `updated` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY `idx_course_type` (`course_type`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

insert into courses(name, course_type, total_lesson_hours, pre_required, target, recommend_competition, recommend_period) values ('Python基础课程', 1, 40, '零基础即可学习', '青少年编程能力等级考试7级', '蓝桥杯/STEAM测试', '二至六年级');

CREATE TABLE `lesson_record` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `course_type` tinyint(4) NOT NULL,
    `teacher` varchar(255) NOT NULL,
    `tags` varchar(255) DEFAULT NULL,
    `remark` text DEFAULT NULL,
    `created` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_user_course` (`user_id`, `course_type`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `lesson_record_002` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `course_type` tinyint NOT NULL,
    `teacher` varchar(255) NOT NULL,
    `tags` varchar(255) DEFAULT NULL,
    `remark` text,
    `created` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_user_course` (`user_id`,`course_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

insert into lesson_record_002 (user_id, course_type, teacher, tags, remark) values(2, 1, '颜总监', '拍马屁', '总监今天教我拍马屁呀真牛逼');