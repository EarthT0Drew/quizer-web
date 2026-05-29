# Standards
This is the place to standardize elements of this project, such as the database.

## Quizzes database
### Quizzes

| ID | Name | Author | Description | Created at | Last updated |
|----|------|--------|-------------|------------|--------------|
| id | name | author | description | created_at | updated_at   |
| int|string| int    | string      | TBD        | TBD          |

### Questions
| ID | Quiz ID | Position | Type | Author | Prompt | Answer | Created at | Last updated |
|----|---------|----------|------|--------|--------|--------|------------|--------------|
| id | quiz_id | position | type | author | prompt | answer | created_at | updated_at   |
| int| string  | int      |string| int    | string | string | TBD        | TBD          |

### MCQ Answers
| ID | Quiz ID | Question ID | Author | Position | Text | Is correct? | Created at | Last updated |
|----|---------|-------------|--------|----------|------|-------------|------------|--------------|
| id | quiz_id | question_id | author | position | text | is_correct  | created_at | updated_at   |
| int| int     | int         | int    | int      |string| bool        | TBD        | TBD          |

## Users database
### Users
| ID | Username | Display Name | Email | Password hash | Created at | Last updated |
|----|----------|--------------|-------|---------------|------------|--------------|
| id | username | display_name | email | password_hash | created_at | updated_at   |
| int| string   | string       | string| string        | TBD        | TBD          |