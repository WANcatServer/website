# 後端

* 首頁 — Cookie 檢查登入，初次進入 Cookie 聲明
* 登入 — 登入成功 301 Dashboard
* Dashboard — 推薦題目、我的題目、最近活動（自己）、分類
  * 推薦題目：偏好及推薦系統。
  * 最近活動：Record 中 user == $user
* 推薦系統 函數：時間、作答人數、評比
* 答題 /\$username/\$question-name：
  * /
  * /forum
  * /reference
  * /library
  * /issue



----



* 活動紀錄系統：一段時間後刪除（或一定訊息量）
  * type Record struct :
    * user
    * time
    * description
* 題目管理系統：
  * type Question struct:
    * id
    * user
    * name
    * time
    * content
    * category
    * tags
    * issue - link
    * reference
    * 知識庫
    * 論壇
* 回報管理系統：
  * type Issue struct
    * id
    * user
    * time
    * question - link
    * content

* 