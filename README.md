# Naoto Nishihama's Portfolio & Blog

Next.js、Tailwind CSS、MDXを使用して構築されたポートフォリオ&ブログサイトです。

## 🚀 技術スタック

- **Next.js 14** - React フレームワーク
- **TypeScript** - 型安全な開発
- **Tailwind CSS** - ユーティリティファーストCSS
- **MDX** - Markdown + JSX for ブログ
- **Lucide React** - アイコン

## 📦 機能

- **レスポンシブデザイン** - モバイルファーストで最適化
- **ポートフォリオ** - プロジェクト一覧とスキル紹介
- **ブログ** - MDXベースの技術ブログ
- **お問い合わせ** - コンタクトフォーム
- **SEO最適化** - メタタグとサイトマップ

## 🛠️ 開発・デプロイ

### ローカル開発

```bash
# 依存関係のインストール
npm install

# 開発サーバーの起動
npm run dev
```

### ビルド・デプロイ

```bash
# 本番ビルド
npm run build

# GitHub Pagesへの自動デプロイ
git push origin main
```

## 📝 ブログ記事の追加

`content/blog/` ディレクトリに `.md` ファイルを追加してください。

```markdown
---
title: "記事タイトル"
date: "2024-07-26"
excerpt: "記事の概要"
tags: ["Tag1", "Tag2"]
---

# 記事の内容

Markdownで記事を書きます。
```

## 🌍 デプロイ

このサイトはGitHub Pagesで自動デプロイされます。`main` ブランチにプッシュすると、GitHub Actionsが自動的にビルド・デプロイを実行します。

## 📄 ライセンス

This project is licensed under the MIT License.