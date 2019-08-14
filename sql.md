# MYSQLの型まとめ

## INT
```
TINYINT         -128                  ~127
SMALLINT        -32768                ~32767
MEDIUMINT       -8388608              ~8388607
INT             -2147483648           ~2147483647
BIGINT          -9223372036854775808  ~9223372036854775807

INT系には、UNSIGNEDとZEROFILLがある。
UNSIGNED...符号なし
ZEROFILL...桁数分0で埋める
```

## FLOAT
```
FLOAT  -3.402823466E+38 ~ -1.175494351E-38
DOUBLE -1.7976931348623157E+308 ~ -2.2250738585072014E-308

FLOAT系にも、UNSIGNEDとZEROFILLがある。
UNSIGNED...符号なし
ZEROFILL...桁数分0で埋める
```

## 文字列系
```
CHAR    固定長文字列 0~255
VARCHAR 可変長文字列 0~65535 ※ 使用する文字コードによってはバイト数になる

BLOB バイナリデータを扱う (TINYBLOB, BLOB, MEDIUMTEXT, LONGBLOBがある)
TEXT 文字列データを扱う (TINYTEXT, TEXT, MEDIUMTEXT, LONGTEXTがある)
```

# 構文まとめ

ISUCON6のテーブルより例を示す
```sql
# entry テーブル
+-------------+---------------------+------+-----+---------+----------------+
| Field       | Type                | Null | Key | Default | Extra          |
+-------------+---------------------+------+-----+---------+----------------+
| id          | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
| author_id   | bigint(20) unsigned | NO   |     | NULL    |                |
| keyword     | varchar(191)        | YES  | UNI | NULL    |                |
| description | mediumtext          | YES  |     | NULL    |                |
| updated_at  | datetime            | NO   |     | NULL    |                |
| created_at  | datetime            | NO   |     | NULL    |                |
+-------------+---------------------+------+-----+---------+----------------+

# user テーブル
+------------+---------------------+------+-----+---------+----------------+
| Field      | Type                | Null | Key | Default | Extra          |
+------------+---------------------+------+-----+---------+----------------+
| id         | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
| name       | varchar(191)        | YES  | UNI | NULL    |                |
| salt       | varchar(20)         | YES  |     | NULL    |                |
| password   | varchar(40)         | YES  |     | NULL    |                |
| created_at | datetime            | NO   |     | NULL    |                |
+------------+---------------------+------+-----+---------+----------------+
```

## テーブル作成 CREATE TABLE

```
CREATE TABLE table_name (
  カラム名 型 (制約)

  # INDEXを追加
  KEY インデックスの名前(インデックスを貼るカラム)

  # PRIMARY KEY を設定
  PRIMARY KEY (カラム名)
)

※ 最終行に `,` があるとエラーが起きる
```

#### 実際の作成例

```sql
CREATE TABLE entry (
    id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    author_id BIGINT UNSIGNED NOT NULL,
    keyword VARCHAR(191) UNIQUE,
    description MEDIUMTEXT,
    updated_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE user (
    id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(191) UNIQUE,
    salt VARCHAR(20),
    password VARCHAR(40),
    created_at DATETIME NOT NULL
) Engine=InnoDB DEFAULT CHARSET=utf8mb4;
```

## テーブルの構造を変更 ALTER TABLE

```sql
# フィールド追加/削除
ALTER TABLE テーブル名 ADD/DROP フィールド名 型

# フィールドの変更
ALTER TABLE テーブル名 CHANGE フィールド名 新しいフィールド名 型

# 型のみを変更
ALTER TABLE テーブル名 MODIFY フィールド名 型

# フィールド名の変更
ALTER TABLE テーブル名 RENAME フィールド名 TO 新しいフィールド名

# インデックスの追加/削除
ALTER TABLE テーブル名 ADD/DROP インデックスの名前(インデックスを貼るカラム)
```

## 内部結合 INNER JOIN

```sql
SELECT * FROM テーブル1 INNER JOIN テーブル2 ON 結合条件 WHERE ...
```

### INDEX

```
IS NULLについては、INDEXが効いている
その他インデックスについての参考記事
https://qiita.com/NagaokaKenichi/items/44cabcafa3d02d9cd896
```
