vesper-martini
=====================
**[go-martini/martini](https://github.com/go-martini/martini)** base web app

Install
-------
`go get github.com/h-miyamo/vesper-martini`

Routes
------
```
| uri           | detail            |
| ------------- | ----------------- |
| /hello        | print "hello"     |
| /user/:id     | find users.id     |
```

Used Package
------------
```
go get github.com/go-martini/martini
go get github.com/codegangsta/martini-contrib/render
go get github.com/go-sql-driver/mysql
go get github.com/jinzhu/gorm
```

SQL
---
```
CREATE DATABASE `practice` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

CREATE TABLE `users` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8_unicode_ci',
	`created_by` INT(11) NULL DEFAULT NULL,
	`updated_by` INT(11) NULL DEFAULT NULL,
	`lock_version` INT(11) NOT NULL DEFAULT '0',
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL,
	PRIMARY KEY (`id`)
)
COLLATE='utf8_unicode_ci'
ENGINE=InnoDB;

insert into users values (null, "hoge-ほげ", 1, 1, 0, now(), now());
```
