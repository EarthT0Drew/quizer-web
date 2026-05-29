# Standards
This is the place to standardize elements of this project, such as the database.

## Quizzes database
### quizzes

| ID | Name | Description | Created at | Last updated | Author |
|----|------|-------------|------------|--------------|--------|
| id | name | description | created_at | updated_at   | author |
| int|string| string      | TBD        | TBD          | int    |

### questions
| ID | Quiz ID | Position | Type | Prompt | Answer | Created at | Last updated | Author |
|----|---------|----------|------|--------|--------|------------|--------------|--------|
| id | quiz_id | position | type | prompt | answer | created_at | updated_at   | author |
| int| int     | int      |string| string | string | TBD        | TBD          | int    |

#### Question types
- Multiple choice ("mcq")
- Short answer ("answer")
- True/False ("true_false")

### mcq_answers
| ID | Quiz ID | Question ID | Position | Text | Is correct? | Created at | Last updated | Author |
|----|---------|-------------|----------|------|-------------|------------|--------------|--------|
| id | quiz_id | question_id | position | text | is_correct  | created_at | updated_at   | author |
| int| int     | int         | int      |string| bool        | TBD        | TBD          | int    |

## Users database
### users
| ID | Username | Display Name | Email | Password hash | Created at | Last updated |
|----|----------|--------------|-------|---------------|------------|--------------|
| id | username | display_name | email | password_hash | created_at | updated_at   |
| int| string   | string       | string| string        | TBD        | TBD          |