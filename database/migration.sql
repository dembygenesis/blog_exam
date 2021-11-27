CREATE TABLE IF NOT EXISTS blog.article (
    id INT(11) NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY title (title,author) COMMENT 'Title, and author must be unique TOGETHER'
);