# 資料表格設計

儲存技術：SQL （SQLite3, MariaDB, MySQL, PostgreSQL ...）

滿足資料庫第三正規化

**Max Length 用 Bytes 表示**



### 使用者 Users：

| Column     | Type     | Max Length              |
| ---------- | :------- | ----------------------- |
| ID         | INT      | 2,147,483,647 (4 bytes) |
| Name       | VARCHAR  | 255                   |
| Avatar     | VARCHAR  | 65,535                  |
| Email      | VARCHAR  | 65535                   |
| JoinDate   | DATETIME | CCYY-MM-DD hh:mm:ss     |
| LastOnline | DATETIME | CCYY-MM-DD hh:mm:ss     |
| Password   | BLOB     | 65,535                  |
| Nickname | VARCHAR | |
| Description | TEXT | |

### 問題 Questions：

| Column     | Type       | Max Length    |
| ---------- | ---------- | ------------- |
| ID         | INT        | 2,147,483,647 |
| User       | Users.ID   |               |
| Title      | VARCHAR    | 255           |
| Content    | MEDIUMTEXT | 16,772,215    |
| BestAnswer | Answers.ID |               |
| Created    | DATETIME   |               |
| Updated    | DATETIME   |               |
| Reference  | TEXT       | 65535         |

### 答案 Answers

| Column    | Type         | Max Length    |
| --------- | ------------ | ------------- |
| ID        | INT          | 2,147,483,647 |
| User      | Users.ID     |               |
| Question  | Questions.ID |               |
| Content   | TEXT         | 65535         |
| Time      | DATETIME     |               |
| Displayed | BOOL         |               |
| Solved    | BOOL         |               |



### 回報 Issues

| Column   | Type         | Max Length    |
| -------- | ------------ | ------------- |
| ID       | INT          | 2,147,483,647 |
| User     | Users.ID     |               |
| Question | Questions.ID |               |
| Title    | VARCHAR      | 255           |
| Content  | TEXT         | 65,535        |
| Time     | DATETIME     |               |



### 回報回覆 IssueReplies

| Column  | Type      | Max Length    |
| ------- | --------- | ------------- |
| ID      | INT       | 2,147,483,647 |
| User    | Users.ID  |               |
| Issue   | Issues.ID |               |
| Time    | DATETIME  |               |
| Content | TEXT      | 65,535        |



### 答案回覆 AnswerReplies

| Column  | Type       | Max Length    |
| ------- | ---------- | ------------- |
| ID      | INT        | 2,147,483,647 |
| User    | Users.ID   |               |
| Answer  | Answers.ID |               |
| Time    | DATETIME   |               |
| Content | TEXT       | 65,535        |



### 題目評價 QuestionVotes

| Column     | Type         | Max Length |
| ---------- | ------------ | ---------- |
| *User      | Users.ID     |            |
| *Question  | Questions.ID |            |
| Evaluation | BOOL         |            |
| Time       | DATETIME     |            |

### 作答評價 AnswerVotes

| Column     | Type       | Max Length |
| ---------- | ---------- | ---------- |
| *User      | Users.ID   |            |
| *Answer    | Answers.ID |            |
| Evaluation | BOOL       |            |
| Time       | DATETIME   |            |

### 分類 Categories

| Column    | Type            | Max Length |
| --------- | --------------- | ---------- |
| *Name     | ENUM( 分類 ...) |            |
| *Question | Questions.ID    |            |



### 標籤 Tags

| Column    | Type         | Max Length |
| --------- | ------------ | ---------- |
| *Name     | VARCHAR      | 255        |
| *Question | Questions.ID |            |



### 紀錄 Records

| Column      | Type     | Max Length |
| ----------- | -------- | ---------- |
| *ID         | INT      |            |
| User        | Users.ID |            |
| Description | VARCHAR  | 255        |
| Time        | DATETIME |            |

### 登入狀態 Session

| Column   | Type      | Max Length |
| -------- | --------- | ---------- |
| ID       | BLOB      | 65,535     |
| User     | Users.ID  |            |
| LastUsed | TIMESTAMP |            |

