import Layout from '@/components/Layout'
import Link from 'next/link'
import { Calendar, ArrowRight } from 'lucide-react'

export default function Blog() {
  // Sample blog posts data
  const blogPosts = [
    {
      id: 1,
      title: 'Next.js 14の新機能について',
      excerpt: 'Next.js 14で追加された新機能とその活用方法について詳しく解説します。App Routerの改善点や新しいパフォーマンス最適化機能に注目。',
      date: '2024-07-15',
      readTime: '5分',
      tags: ['Next.js', 'React', 'JavaScript'],
      slug: 'nextjs-14-new-features'
    },
    {
      id: 2,
      title: 'TypeScriptのベストプラクティス',
      excerpt: '実際のプロジェクトで使えるTypeScriptのベストプラクティスを紹介。型安全性を保ちながら開発効率を向上させる方法。',
      date: '2024-07-10',
      readTime: '7分',
      tags: ['TypeScript', 'JavaScript', 'Programming'],
      slug: 'typescript-best-practices'
    },
    {
      id: 3,
      title: 'Tailwind CSSで効率的なスタイリング',
      excerpt: 'Tailwind CSSを使った効率的なWebサイト制作のコツとテクニック。コンポーネント設計とパフォーマンスの最適化について。',
      date: '2024-07-05',
      readTime: '4分',
      tags: ['Tailwind CSS', 'CSS', 'Design'],
      slug: 'efficient-styling-with-tailwind'
    },
    {
      id: 4,
      title: 'React Hooksの実践的な使い方',
      excerpt: 'React Hooksの基本から応用まで、実際のプロジェクトで役立つパターンとアンチパターンを詳しく解説します。',
      date: '2024-06-30',
      readTime: '8分',
      tags: ['React', 'Hooks', 'JavaScript'],
      slug: 'practical-react-hooks'
    },
    {
      id: 5,
      title: 'WebアプリケーションのSEO最適化',
      excerpt: 'モダンなWebアプリケーションにおけるSEO最適化の手法。技術的なSEOからコンテンツ戦略まで包括的に解説。',
      date: '2024-06-25',
      readTime: '6分',
      tags: ['SEO', 'Web Development', 'Performance'],
      slug: 'web-app-seo-optimization'
    },
    {
      id: 6,
      title: 'GitHub Actionsで自動化するCI/CD',
      excerpt: 'GitHub Actionsを使った効率的なCI/CDパイプラインの構築方法。自動テスト、デプロイメント、コード品質チェックまで。',
      date: '2024-06-20',
      readTime: '9分',
      tags: ['GitHub Actions', 'CI/CD', 'DevOps'],
      slug: 'github-actions-cicd'
    }
  ]

  const categories = ['全て', 'JavaScript', 'React', 'Next.js', 'TypeScript', 'CSS', 'DevOps']

  return (
    <Layout>
      {/* Hero Section */}
      <section className="bg-gradient-to-r from-green-600 to-blue-600 text-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h1 className="text-4xl md:text-6xl font-bold mb-6">
            Tech Blog
          </h1>
          <p className="text-xl md:text-2xl text-green-100">
            技術的な学びと開発の気づきを共有しています
          </p>
        </div>
      </section>

      {/* Categories */}
      <section className="py-8 bg-white border-b">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex flex-wrap gap-2 justify-center">
            {categories.map((category) => (
              <button
                key={category}
                className="px-4 py-2 text-sm font-medium text-gray-600 bg-gray-100 rounded-full hover:bg-gray-200 transition-colors duration-200"
              >
                {category}
              </button>
            ))}
          </div>
        </div>
      </section>

      {/* Blog Posts */}
      <section className="py-20 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
            {/* Main Content */}
            <div className="lg:col-span-2">
              <div className="space-y-8">
                {blogPosts.map((post) => (
                  <article key={post.id} className="bg-white rounded-lg shadow-md overflow-hidden">
                    <div className="p-8">
                      <div className="flex items-center text-sm text-gray-500 mb-4">
                        <Calendar size={16} className="mr-2" />
                        <span>{new Date(post.date).toLocaleDateString('ja-JP')}</span>
                        <span className="mx-2">•</span>
                        <span>{post.readTime}</span>
                      </div>
                      
                      <h2 className="text-2xl font-bold text-gray-900 mb-4">
                        <Link 
                          href={`/blog/${post.slug}`}
                          className="hover:text-blue-600 transition-colors duration-200"
                        >
                          {post.title}
                        </Link>
                      </h2>
                      
                      <p className="text-gray-600 mb-6 leading-relaxed">
                        {post.excerpt}
                      </p>
                      
                      <div className="flex flex-wrap gap-2 mb-6">
                        {post.tags.map((tag) => (
                          <span
                            key={tag}
                            className="px-3 py-1 bg-blue-100 text-blue-800 text-sm rounded-full"
                          >
                            {tag}
                          </span>
                        ))}
                      </div>
                      
                      <Link
                        href={`/blog/${post.slug}`}
                        className="inline-flex items-center text-blue-600 hover:text-blue-800 font-medium transition-colors duration-200"
                      >
                        続きを読む
                        <ArrowRight size={16} className="ml-2" />
                      </Link>
                    </div>
                  </article>
                ))}
              </div>
              
              {/* Pagination */}
              <div className="mt-12 flex justify-center">
                <div className="flex space-x-2">
                  <button className="px-4 py-2 text-gray-500 bg-white border border-gray-300 rounded-lg hover:bg-gray-50">
                    前へ
                  </button>
                  <button className="px-4 py-2 text-white bg-blue-600 border border-blue-600 rounded-lg">
                    1
                  </button>
                  <button className="px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50">
                    2
                  </button>
                  <button className="px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50">
                    3
                  </button>
                  <button className="px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50">
                    次へ
                  </button>
                </div>
              </div>
            </div>

            {/* Sidebar */}
            <div className="space-y-8">
              {/* About */}
              <div className="bg-white p-6 rounded-lg shadow-md">
                <h3 className="text-xl font-bold text-gray-900 mb-4">このブログについて</h3>
                <p className="text-gray-600 text-sm leading-relaxed">
                  フロントエンド開発に関する技術的な学びや実践的なTipsを共有しています。
                  新しい技術の検証結果や開発プロセスでの気づきをお届けします。
                </p>
              </div>

              {/* Recent Posts */}
              <div className="bg-white p-6 rounded-lg shadow-md">
                <h3 className="text-xl font-bold text-gray-900 mb-4">最近の記事</h3>
                <div className="space-y-4">
                  {blogPosts.slice(0, 5).map((post) => (
                    <div key={post.id}>
                      <Link
                        href={`/blog/${post.slug}`}
                        className="block text-sm font-medium text-gray-900 hover:text-blue-600 transition-colors duration-200"
                      >
                        {post.title}
                      </Link>
                      <p className="text-xs text-gray-500 mt-1">
                        {new Date(post.date).toLocaleDateString('ja-JP')}
                      </p>
                    </div>
                  ))}
                </div>
              </div>

              {/* Tags */}
              <div className="bg-white p-6 rounded-lg shadow-md">
                <h3 className="text-xl font-bold text-gray-900 mb-4">タグ</h3>
                <div className="flex flex-wrap gap-2">
                  {Array.from(new Set(blogPosts.flatMap(post => post.tags))).map((tag) => (
                    <span
                      key={tag}
                      className="px-3 py-1 bg-gray-100 text-gray-700 text-sm rounded-full hover:bg-gray-200 cursor-pointer transition-colors duration-200"
                    >
                      {tag}
                    </span>
                  ))}
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </Layout>
  )
}
