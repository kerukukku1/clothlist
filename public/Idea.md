# 実装項目アイデア
- Gallery機能(服が全部見れるとこ)
    - タグ付け機能
    - sort機能
    - detail書き込み機能
    - 削除機能
    - ジャンル機能(tagとは別)
    - **色別ソート機能!!**

- Gallery詳細ページ
    - 詳細な表示(買った場所や値段や思い出等)

- 洋服の組み合わせを保持する機能
    - ユーザが服の組み合わせを入力
    - 将来的には学習させて服をサジェスト
    - B2Cのリンクもサジェストしたい

- 季節による服の選別機能

## Galleryアイデア
mongooseにbase64 encodingされた画像を保存していく形式

- 画像DB設計
    - title : string =>   服のタイトル
    - image : string =>   encodingされた画像
    - date  : date   =>   追加日時等
    - id    : string =>   洋服のID
    - tag   : [string] => 洋服に対するタグ

- 画像詳細DB設計
    - parentObjectId    :画像DBに紐付けるOID
    - comment           :洋服への詳細(Markdown)
    - withCloth         :一緒に着る服

- VueComponent 設計
    - サイドバー
    - メインエリア
    - 衣服ボックス
    - 衣服ボックス内容