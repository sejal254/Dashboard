CREATE DATABASE IF NOT EXISTS project3;
use project3;
CREATE TABLE IF NOT EXISTS user(
    user_id int PRIMARY KEY AUTO_INCREMENT,
    username varchar(50) NOT NULL,
    email varchar(50) NOT NULL,
    role ENUM("Admin","User") NOT NULL DEFAULT "User",
    status ENUM("Active","Inactive") NOT NULL DEFAULT "Active",
    password varchar(500) NOT NULL
);
CREATE TABLE IF NOT EXISTS project(
    project_id int PRIMARY KEY AUTO_INCREMENT,
    projectname varchar(200) NOT NULL,
    status ENUM("Active","Inactive") NOT NULL DEFAULT "Active"
   
);
CREATE TABLE IF NOT EXISTS userprojects(
    temp_id int PRIMARY KEY AUTO_INCREMENT,
    user_id int ,
    project_id int,
    project_name varchar(50) NOT NULL,
    status ENUM("Active","Inactive") NOT NULL DEFAULT "Active"
  
);
