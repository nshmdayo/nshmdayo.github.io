import Layout from '@/components/Layout'
import Link from 'next/link'
import { Github, Linkedin, Mail, ArrowRight } from 'lucide-react'

export default function Home() {
  return (
    <Layout>
      {/* Hero Section */}
      <section className="bg-gradient-to-r from-blue-600 to-purple-600 text-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center">
            <h1 className="text-4xl md:text-6xl font-bold mb-6">
              Naoto Nishihama
            </h1>
            <p className="text-xl md:text-2xl mb-8 text-blue-100">
              フロントエンド・エンジニア & UI/UXデザイナー
            </p>
            <p className="text-lg mb-8 max-w-3xl mx-auto text-blue-50">
              ユーザー体験を重視したWebアプリケーションの設計・開発を行っています。
              React、Next.js、TypeScriptを使用したモダンな開発手法を得意としています。
            </p>
            <div className="flex justify-center space-x-4">
              <Link href="/portfolio" className="btn btn-primary">
                ポートフォリオを見る
                <ArrowRight className="ml-2 h-5 w-5" />
              </Link>
              <Link href="/contact" className="btn btn-secondary bg-white text-blue-600 hover:bg-gray-100">
                お問い合わせ
              </Link>
            </div>
          </div>
        </div>
      </section>

      {/* About Section */}
      <section className="py-20 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              About Me
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              新技術への探求心を持ち、ユーザーファーストの開発を心がけています
            </p>
          </div>
          
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div className="text-center p-6">
              <div className="bg-blue-100 rounded-full w-16 h-16 flex items-center justify-center mx-auto mb-4">
                <span className="text-2xl">💻</span>
              </div>
              <h3 className="text-xl font-semibold mb-2">フロントエンド開発</h3>
              <p className="text-gray-600">
                React、Next.js、TypeScriptを使用したモダンなWebアプリケーション開発
              </p>
            </div>
            
            <div className="text-center p-6">
              <div className="bg-green-100 rounded-full w-16 h-16 flex items-center justify-center mx-auto mb-4">
                <span className="text-2xl">🎨</span>
              </div>
              <h3 className="text-xl font-semibold mb-2">UI/UXデザイン</h3>
              <p className="text-gray-600">
                ユーザビリティを重視したインターフェースデザインとプロトタイピング
              </p>
            </div>
            
            <div className="text-center p-6">
              <div className="bg-purple-100 rounded-full w-16 h-16 flex items-center justify-center mx-auto mb-4">
                <span className="text-2xl">⚡</span>
              </div>
              <h3 className="text-xl font-semibold mb-2">パフォーマンス最適化</h3>
              <p className="text-gray-600">
                Webサイトの高速化とSEO対策による最適なユーザー体験の提供
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Recent Blog Posts */}
      <section className="py-20 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              最新のブログ記事
            </h2>
            <p className="text-xl text-gray-600">
              技術的な学びや開発の気づきを共有しています
            </p>
          </div>
          
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {/* Blog post cards will be dynamically generated */}
            <div className="bg-white rounded-lg shadow-md overflow-hidden">
              <div className="p-6">
                <h3 className="text-xl font-semibold mb-2">
                  Next.js 14の新機能について
                </h3>
                <p className="text-gray-600 mb-4">
                  最新のNext.js 14で追加された新機能とその活用方法について解説します...
                </p>
                <div className="flex items-center text-sm text-gray-500">
                  <span>2024年7月15日</span>
                </div>
              </div>
            </div>
            
            <div className="bg-white rounded-lg shadow-md overflow-hidden">
              <div className="p-6">
                <h3 className="text-xl font-semibold mb-2">
                  TypeScriptのベストプラクティス
                </h3>
                <p className="text-gray-600 mb-4">
                  実際のプロジェクトで使えるTypeScriptのベストプラクティスを紹介...
                </p>
                <div className="flex items-center text-sm text-gray-500">
                  <span>2024年7月10日</span>
                </div>
              </div>
            </div>
            
            <div className="bg-white rounded-lg shadow-md overflow-hidden">
              <div className="p-6">
                <h3 className="text-xl font-semibold mb-2">
                  Tailwind CSSで効率的なスタイリング
                </h3>
                <p className="text-gray-600 mb-4">
                  Tailwind CSSを使った効率的なWebサイト制作のコツとテクニック...
                </p>
                <div className="flex items-center text-sm text-gray-500">
                  <span>2024年7月5日</span>
                </div>
              </div>
            </div>
          </div>
          
          <div className="text-center mt-12">
            <Link href="/blog" className="btn btn-primary">
              すべての記事を見る
              <ArrowRight className="ml-2 h-5 w-5" />
            </Link>
          </div>
        </div>
      </section>

      {/* Social Links */}
      <section className="py-20 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-8">
            つながりましょう
          </h2>
          <div className="flex justify-center space-x-6">
            <a
              href="https://github.com/nshmdayo"
              target="_blank"
              rel="noopener noreferrer"
              className="text-gray-600 hover:text-gray-900 transition-colors"
            >
              <Github size={32} />
            </a>
            <a
              href="https://linkedin.com/in/naoto-nishihama"
              target="_blank"
              rel="noopener noreferrer"
              className="text-gray-600 hover:text-blue-600 transition-colors"
            >
              <Linkedin size={32} />
            </a>
            <a
              href="mailto:naoto@example.com"
              className="text-gray-600 hover:text-red-600 transition-colors"
            >
              <Mail size={32} />
            </a>
          </div>
        </div>
      </section>
    </Layout>
  )
}
