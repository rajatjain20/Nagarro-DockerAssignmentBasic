/*
Enter custom T-SQL here that would run after SQL Server has started up. 
*/

ALTER LOGIN SA WITH PASSWORD='admin@1234';
GO
CREATE DATABASE USERDATA;
GO
CREATE TABLE USERDATA.dbo.USERINFO(ID INT NOT NULL PRIMARY KEY, NAME VARCHAR(50) NOT NULL);
GO