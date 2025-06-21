# DB Schema

### Events

|     Field     |     Type     |     Null     |     Key     |     Default     |     Extra     |     Description     |
|---------------|--------------|--------------|-------------|-----------------|---------------|---------------------|
| event_id      | CHAR(36)     | NO           | PRIMARY     |                 | UUID          |                     |
| title         | VARCHAR(100) | NO           |             |                 |               |                     |
| description   | TEXT         | NO           |             |                 |               |                     |
| date          | DATE         | NO           |             |                 |               |                     |
| is_open       | BOOLEAN      | NO           |             |                 |               |                     |
| is_deleted    | BOOLEAN      | NO           |             |                 |               |                     |
| created_at    | TIMESTAMP    | NO           |             |CURRENT_TIMESTAMP| ON CREATE     |                     |
| updated_at    | TIMESTAMP    | NO           |             |CURRENT_TIMESTAMP| ON UPDATE     |                     |

### Admins

|     Field     |     Type     |     Null     |     Key     |     Default     |     Extra     |     Description     |
|---------------|--------------|--------------|-------------|-----------------|---------------|---------------------|
| event_id      | CHAR(36)     | NO           | PRIMARY     |                 | UUID          |                     |
| traq_id       | VARCHAR(32)  | NO           | PRIMARY     |                 |               |                     |

> Composite Primary Key: (`event_id`, `traq_id`)

### attendees

|     Field     |     Type     |     Null     |     Key     |     Default     |     Extra     |     Description     |
|---------------|--------------|--------------|-------------|-----------------|---------------|---------------------|
| event_id      | CHAR(36)     | NO           | PRIMARY     |                 | UUID          |                     |
| traq_id       | VARCHAR(32)  | NO           | PRIMARY     |                 |               |                     |

> Composite Primary Key: (`event_id`, `traq_id`)

### lotteries

|     Field     |     Type     |     Null     |     Key     |     Default     |     Extra     |     Description     |
|---------------|--------------|--------------|-------------|-----------------|---------------|---------------------|
| lottery__id   | CHAR(36)     | NO           | PRIMARY     |                 | UUID          |                     |
| event_id      | CHAR(36)     | NO           |             |                 | UUID          |                     |
| title         | VARCHAR(100) | NO           |             |                 |               |                     |
| is_deleted    | BOOLEAN      | NO           |             |                 |               |                     |
| created_at    | TIMESTAMP    | NO           |             |CURRENT_TIMESTAMP| ON CREATE     |                     |
| updated_at    | TIMESTAMP    | NO           |             |CURRENT_TIMESTAMP| ON UPDATE     |                     |

### winners

|     Field     |     Type     |     Null     |     Key     |     Default     |     Extra     |     Description     |
|---------------|--------------|--------------|-------------|-----------------|---------------|---------------------|
| lottery_id    | CHAR(36)     | NO           | PRIMARY     |                 |               |                     |
| traq_id       | VARCHAR(32)  | NO           | PRIMARY     |                 |               |                     |
| event_id      | CHAR(36)     | NO           |             |                 | UUID          |                     |

> Composite Primary Key: (`lottery_id`, `traq_id`)