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

| Column     | Type                 | Max Length          |
| ---------- | :------------------- | ------------------- |
| ID         | SMALLINT(4) UNSINGED | 65,535              |
| Name       | TINYTEXT             | 255                 |
| Avatar     | MEDIUMBLOB           | 16,777,215          |
| Email      | TINYTEXT             | 255                 |
| JoinDate   | DATE                 | CCYY-MM-DD          |
| LastOnline | DATETIME             | CCYY-MM-DD hh:mm:ss |

### 問題 Questions：

| Column    | Type                  | Max Length |
| --------- | --------------------- | ---------- |
| ID        | MEDIUMINT(6) UNSINGED | 16,777,215 |
| User      | Users.ID              |            |
| Title     | TINYTEXT              | 255        |
| Content   | TEXT                  | 65535      |
| Created   | DATETIME              |            |
| Updated   | DATETIME              |            |
| Reference | TEXT                  | 65535      |
| Category  | Categories.ID         |            |

### 答案 Answers

### 回報 Issues

### 回覆 Replies

### 分類 Categories

### 標籤 Tags





* 活動紀錄系統：一段時間後刪除（或一定訊息量）
  * type Record struct :
    * user
    * time
    * description
* 題目管理系統：
  * type Question struct:
    * id
    * user - link
    * name
    * time
    * content
    * category
    * reference
    * 知識庫
    * 論壇
  * type Tag struct :
    * id
    * name
    * question - link
* 回報管理系統：
  * type Issue struct
    * id
    * user
    * time
    * question - link
    * content
* 作答管理系統
  * type Answer struct
    * ID
    * User
    * Content
    * Time
    * Question - link
* 使用者管理系統
  * type User struct
    * ID
    * Name
    * Account
    * []Answers - link
    * []Questions - link
    * []Issues - link

