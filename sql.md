# MYSQLの型まとめ

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

```
FLOAT  -3.402823466E+38 ~ -1.175494351E-38
DOUBLE -1.7976931348623157E+308 ~ -2.2250738585072014E-308

FLOAT系にも、UNSIGNEDとZEROFILLがある。
UNSIGNED...符号なし
ZEROFILL...桁数分0で埋める
```

```
CHAR    固定長文字列 0~255
VARCHAR 可変長文字列 0~65535 ※ 使用する文字コードによってはバイト数になる
```

```
BLOB バイナリデータを扱う (TINYBLOB, BLOB, MEDIUMTEXT, LONGBLOBがある)
TEXT 文字列データを扱う (TINYTEXT, TEXT, MEDIUMTEXT, LONGTEXTがある)
```

# 構文まとめ

ISUCON6のテーブルより例を示す
```
entry テーブル
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

user テーブル
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

### テーブル作成SQL

```
CREATE TABLE table_name (
  カラム名 型 (制約)

  # INDEXを追加
  KEY インデックスの名前(インデックスを貼るカラム)

  # PRIMARY KEY を設定
  PRIMARY KEY (カラム名)
)

※ 最終行に `,` があるとエラーが起きる


実際の作成例

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
