-- Add down migration script here
DELETE FROM Items;
ALTER TABLE Items AUTO_INCREMENT = 1;
DELETE FROM Managers;
ALTER TABLE Managers AUTO_INCREMENT = 1;
DELETE FROM Departments;
ALTER TABLE Departments AUTO_INCREMENT = 1;

