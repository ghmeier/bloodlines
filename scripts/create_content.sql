CREATE TABLE content(
	id VARCHAR(36) NOT NULL PRIMARY KEY,
	contentType VARCHAR(20) NOT NULL,
	text varchar(4096) NOT NULL,
	parameters VARCHAR(1024),
	status VARCHAR(20) NOT NULL,
	subject varchar(1024)
);