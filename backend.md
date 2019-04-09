# 後端

* 首頁 — Cookie 檢查登入，初次進入 Cookie 聲明
* 登入 — 登入成功 301 Dashboard
* Dashboard — 推薦題目、我的題目、最近活動（自己）、分類
  * 推薦題目：偏好及推薦系統。
  * 最近活動：Record 中 user == $user
* 推薦系統 函數：時間、作答人數、評比
* 答題 /\$username/\$question-name：
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

* 使用者資料：
  * type User struct:
    * ID string
    * Name
    * Email
    * NickName
    * JoinDate
    * LastOnlineDate
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

