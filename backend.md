# 後端

* 首頁 — Cookie 檢查登入，初次進入 Cookie 聲明
* 登入 — 登入成功 301 Dashboard
* Dashboard — 推薦題目、我的題目、最近活動（自己）、分類
  * 推薦題目：偏好及推薦系統。
  * 最近活動：Record 中 user == $user
* 推薦系統 函數：時間、作答人數、評比
* 題目網址 /\$username/\$question-name：
  * /
  * /answering   //答題頁面
  * /show-answers   //答題後頁面
  * /forum
  * /reference
  * /library
  * /issue

----

* 通知系統：
  * 出題者：
    * 批改通知
    * 錯誤回報通知
  * 作答者：
    * 已批改通知
    * 受理錯誤回報通知
  * 使用者：
    * 訂閱通知
    * 成就達成通知

  ## 資料表格設計

  儲存技術：SQL （SQLite3, MariaDB, MySQL, PostgreSQL ...）

  滿足資料庫第三正規化

  **Max Length 用 Bytes 表示**

  

  ### 使用者 Users：

| Column     | Type       | Max Length          |
| ---------- | :--------- | ------------------- |
| ID         | INT        | 2,147,483,647       |
| Name       | VARCHAR    | 65535               |
| Avatar     | MEDIUMBLOB | 16,777,215          |
| Email      | VARCHAR    | 65535               |
| JoinDate   | DATETIME   | CCYY-MM-DD hh:mm:ss |
| LastOnline | DATETIME   | CCYY-MM-DD hh:mm:ss |

### 問題 Questions：

| Column     | Type          | Max Length    |
| ---------- | ------------- | ------------- |
| ID         | INT           | 2,147,483,647 |
| User       | Users.ID      |               |
| Title      | VARCHAR       | 255           |
| Content    | MEDIUMTEXT    | 16,772,215    |
| BestAnswer | Answers.ID    |               |
| Created    | DATETIME      |               |
| Updated    | DATETIME      |               |
| Reference  | TEXT          | 65535         |
| Category   | Categories.ID |               |

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



### 回覆 Replies

| Column   | Type                          | Max Length    |
| -------- | ----------------------------- | ------------- |
| ID       | INT                           | 2,147,483,647 |
| User     | Users.ID                      |               |
| Type     | ENUM('Answer', 'Issue')       |               |
| ObjectID | INT //Answers.ID OR Issues.ID |               |
| Time     | DATETIME                      |               |
| Message  | TEXT                          | 65,535        |



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

