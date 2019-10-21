![web-ai-speaker](https://github.com/OkanoShogo0903/web-ai-speaker)

## Status code

- 200 StatusOK  
  検索候補が見つかった  
  検索候補一番目の説明テキストを返す

- 204 StatusNoContent  
  入力されたテキストにウェイクアップワードや検索するべき文字列が入っていない  
  クライアントから渡されたテキストをそのまま返す

- 210
  検索候補が見つからなかった時

- 500 StatusInternalServerError
  外部API周辺でエラー発生してる
