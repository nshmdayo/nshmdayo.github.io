---
title: "Hello World - ブログを始めました"
date: "2025-10-18"
description: "Go言語で静的サイトジェネレーターを作成し、GitHub Pagesでブログを公開しました。"
tags: ["Go", "GitHub Pages", "Blog"]
---

## はじめに

Go言語で静的サイトジェネレーターを作成し、GitHub Pagesでポートフォリオ兼ブログサイトを公開しました。

## 技術スタック

- **Go言語** - 静的サイト生成
- **goldmark** - Markdownパーサー
- **GitHub Pages** - ホスティング

## なぜGo言語？

Go言語を選んだ理由：

1. **高速なビルド** - コンパイル言語のため、サイト生成が高速
2. **シンプルな構文** - 学習コストが低く、保守性が高い
3. **標準ライブラリが充実** - テンプレートエンジンやHTMLパーサーが標準で提供
4. **単一バイナリ** - デプロイが簡単

## サイト生成の仕組み

```go
// Markdownファイルを読み込み
data, _ := os.ReadFile("content/blog/post.md")

// goldmarkでHTMLに変換
var buf bytes.Buffer
md := goldmark.New()
md.Convert(data, &buf)

// テンプレートに埋め込み
tmpl.Execute(writer, data)
```

## 今後の予定

- ダークモード対応
- タグによる記事フィルタリング
- RSS/Atom フィード
- 検索機能

引き続き、技術的な学びをシェアしていきます！
