# blender-rendering

GitHubActions で Blender をレンダリングします。

## 機能

- 画像(PNG)をレンダリング
- 動画(mp4)を作成(Python により生成します。)

いずれも CYCLES のみでレンダリングすることができます。(EEVEE ではレンダリングできません。)

## 使い方

```sh
C:.
│  blend.sh
│  blender.mp4
│  make.sh
│  pngtomp4.py
│  README.md
│  render.go
│  renderConfig.txt
│
├─.github
│  └─workflows
│          blender.yml
│
├─blend
│      Miraikomachi.blend
│
├─img
│      anim1979.png
│      anim1980.png
│      anim1981.png
│      anim1982.png
│      anim1983.png
│      anim1984.png
│      anim1985.png
│      anim1986.png
│      anim1987.png
│      anim1988.png
│      anim1989.png
│      anim1990.png
│      anim1991.png
│      anim1992.png
│      anim1993.png
│      anim1994.png
│      anim1995.png
│      anim1996.png
│      anim1997.png
│      anim1998.png
│      anim1999.png
│      anim2000.png
│      anim2001.png
│      anim2002.png
│      anim2003.png
│      anim2004.png
│      anim2005.png
│      anim2006.png
│      anim2007.png
│      anim2008.png
│      anim2009.png
│      anim2010.png
│
└─textures
        bottoms.tga
        eye.tga
        eye_highlight.tga
        face.tga
        hair_back.tga
        hair_front.tga
        jipper.tga
        skin.tga
        tops.tga
        tops_shd.tga
```

`blend`ディレクトリにレンダリングしたい Blender のファイルを置きます。

`textures`ディレクトリには Blender のテクスチャデータを配置します。

`img`ディレクトリにレンダリングされた画像(PNG)が出力されます。

`blender.mp4`は`img`ディレクトリに出力された連番画像を動画にしたファイルです。

`renderConfig.txt`にレンダリングに関する設定を書きます。

```renderConfig.txt
FILENAME=Miraikomachi.blend
START=1978
END=1978
VIDEO=0
```

- `FILENAME`には、`blend`ディレクトリにあり、レンダリングしたい Blender のファイル名を指定します。
- `START`には、レンダリングしたい最初のフレーム番号を指定します。
- `END`には、レンダリングしたい最後のフレーム番号を指定します。
- `VIDEO`には、`blender.mp4`を出力するかを指定します。(出力する場合は 1 ,出力しない場合は 0 )

> 画像(1 フレームのみ)を出力したい場合は`START`と`END`に同じ番号を指定します。

### レンダリングするタイミング

push されたタイミングでレンダリングが開始されます。レンダリングは 1 フレームあたり 50 秒程かかります。

---

[このレポジトリの詳しい説明](https://zenn.dev/u_tan/articles/7a6cad307fa481)

[レンダリングコマンドについて](https://docs.blender.org/manual/en/latest/advanced/command_line/arguments.html)

## サンプルで使用した blend データ

https://www.miraikomachi.com/download/

https://github.com/Miraikomachi/MiraikomachiForBlender
