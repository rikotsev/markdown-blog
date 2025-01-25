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

### CI/CD pipeline
Currently only the CI part is implemented. There are two gateways the code needs to pass:
- on a PR - the code should pass all tests and conform to all linting rules.
- on a release - the code should pass all tests and conform to all linting rules.

### Generating the API stubs
The development of this module is done API first - meaning the API specification should be
considered the main contract. From the API specification the code is generated for handling the HTTP
layer.