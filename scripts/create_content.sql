CREATE TABLE content(
	id VARCHAR(36) NOT NULL PRIMARY KEY,
	contentType ENUM('EMAIL') NOT NULL,
	text varchar(4096) NOT NULL,
	parameters VARCHAR(1024)
);