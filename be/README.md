# Markdown Blog Server

This is a golang built backend.

### Entity Relationship Diagram

```text
                __ 1 _ Category
Article - n ---/
    |
    n
    |
    \______ 1 _ User
                 |
                 1                  
                 |   
Page - n --------/
```

### Database Schema Ideation

| Category        |
|-----------------|
| id UUID         |
| url_id TEXT     |
| name TEXT       |
| PRIMARY KEY(id) |

| Article                                           |
|---------------------------------------------------|
| id UUID                                           |
| url_id TEXT                                       |
| created TIMESTAMP UTC                             |
| edited TIMESTAMP UTC                              |
| title VARCHAR(100)                                |
| description VARCHAR(256)                          |
| content TEXT                                      |
| category_id UUID                                  |
| FOREIGN KEY (category_id) REFERENCES category(id) |
| PRIMARY KEY(id)                                   |

| Page               |
|--------------------|
| id UUID            |
| url_id TEXT        |
| title VARCHAR(100) |
| content TEXT       |