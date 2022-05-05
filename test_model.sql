-- test_study_gorm.users definition

CREATE TABLE `users`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` timestamp    NOT NULL,
    `updated_at` timestamp    NOT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `name`       varchar(100) NOT NULL,
    `gender`     tinyint      NOT NULL,
    `phone`      varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
