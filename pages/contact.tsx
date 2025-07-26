import Layout from '@/components/Layout'
import { Mail, Github, Linkedin, MapPin, Phone } from 'lucide-react'

export default function Contact() {
  return (
    <Layout>
      {/* Hero Section */}
      <section className="bg-gradient-to-r from-indigo-600 to-purple-600 text-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h1 className="text-4xl md:text-6xl font-bold mb-6">
            お問い合わせ
          </h1>
          <p className="text-xl md:text-2xl text-indigo-100">
            お気軽にご連絡ください
          </p>
        </div>
      </section>

      {/* Contact Section */}
      <section className="py-20 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-12">
            {/* Contact Form */}
            <div className="bg-white rounded-lg shadow-lg p-8">
              <h2 className="text-2xl font-bold text-gray-900 mb-6">
                メッセージを送る
              </h2>
              
              <form className="space-y-6">
                <div>
                  <label htmlFor="name" className="block text-sm font-medium text-gray-700 mb-2">
                    お名前 *
                  </label>
                  <input
                    type="text"
                    id="name"
                    name="name"
                    required
                    className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    placeholder="山田太郎"
                  />
                </div>
                
                <div>
                  <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-2">
                    メールアドレス *
                  </label>
                  <input
                    type="email"
                    id="email"
                    name="email"
                    required
                    className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    placeholder="example@email.com"
                  />
                </div>
                
                <div>
                  <label htmlFor="subject" className="block text-sm font-medium text-gray-700 mb-2">
                    件名 *
                  </label>
                  <input
                    type="text"
                    id="subject"
                    name="subject"
                    required
                    className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    placeholder="お問い合わせの件名"
                  />
                </div>
                
                <div>
                  <label htmlFor="message" className="block text-sm font-medium text-gray-700 mb-2">
                    メッセージ *
                  </label>
                  <textarea
                    id="message"
                    name="message"
                    rows={6}
                    required
                    className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    placeholder="お問い合わせの内容をご記入ください"
                  ></textarea>
                </div>
                
                <button
                  type="submit"
                  className="w-full bg-blue-600 text-white py-3 px-6 rounded-lg font-medium hover:bg-blue-700 transition-colors duration-200"
                >
                  送信する
                </button>
              </form>
            </div>

            {/* Contact Info */}
            <div className="space-y-8">
              <div className="bg-white rounded-lg shadow-lg p-8">
                <h2 className="text-2xl font-bold text-gray-900 mb-6">
                  連絡先情報
                </h2>
                
                <div className="space-y-6">
                  <div className="flex items-start">
                    <Mail className="text-blue-600 mr-4 mt-1" size={20} />
                    <div>
                      <h3 className="font-medium text-gray-900">メール</h3>
                      <p className="text-gray-600">naoto@example.com</p>
                    </div>
                  </div>
                  
                  <div className="flex items-start">
                    <MapPin className="text-blue-600 mr-4 mt-1" size={20} />
                    <div>
                      <h3 className="font-medium text-gray-900">所在地</h3>
                      <p className="text-gray-600">東京都, 日本</p>
                    </div>
                  </div>
                  
                  <div className="flex items-start">
                    <Phone className="text-blue-600 mr-4 mt-1" size={20} />
                    <div>
                      <h3 className="font-medium text-gray-900">電話</h3>
                      <p className="text-gray-600">ご相談後に開示いたします</p>
                    </div>
                  </div>
                </div>
              </div>

              <div className="bg-white rounded-lg shadow-lg p-8">
                <h2 className="text-2xl font-bold text-gray-900 mb-6">
                  SNS・プロフィール
                </h2>
                
                <div className="space-y-4">
                  <a
                    href="https://github.com/nshmdayo"
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex items-center text-gray-600 hover:text-gray-900 transition-colors duration-200"
                  >
                    <Github size={20} className="mr-3" />
                    <span>GitHub - nshmdayo</span>
                  </a>
                  
                  <a
                    href="https://linkedin.com/in/naoto-nishihama"
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex items-center text-gray-600 hover:text-blue-600 transition-colors duration-200"
                  >
                    <Linkedin size={20} className="mr-3" />
                    <span>LinkedIn - Naoto Nishihama</span>
                  </a>
                </div>
              </div>

              <div className="bg-blue-50 rounded-lg p-6">
                <h3 className="text-lg font-semibold text-blue-900 mb-3">
                  お仕事について
                </h3>
                <p className="text-blue-800 text-sm leading-relaxed">
                  フロントエンド開発、UI/UXデザイン、
                  Webアプリケーション制作のご相談を承っております。
                  小規模なサイト制作から大規模なシステム開発まで対応可能です。
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* FAQ Section */}
      <section className="py-20 bg-white">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl font-bold text-gray-900 text-center mb-12">
            よくある質問
          </h2>
          
          <div className="space-y-8">
            <div className="border-b border-gray-200 pb-6">
              <h3 className="text-lg font-semibold text-gray-900 mb-2">
                プロジェクトの期間はどのくらいですか？
              </h3>
              <p className="text-gray-600">
                プロジェクトの規模や内容によって異なりますが、
                小規模なWebサイトで2-4週間、中規模なアプリケーションで2-3ヶ月程度が目安です。
                詳細はお打ち合わせ時にご相談させていただきます。
              </p>
            </div>
            
            <div className="border-b border-gray-200 pb-6">
              <h3 className="text-lg font-semibold text-gray-900 mb-2">
                対応可能な技術スタックは？
              </h3>
              <p className="text-gray-600">
                React、Next.js、TypeScript、Vue.js、Node.js等のモダンな技術スタックに対応しています。
                お客様の要件に応じて最適な技術選定をご提案いたします。
              </p>
            </div>
            
            <div className="border-b border-gray-200 pb-6">
              <h3 className="text-lg font-semibold text-gray-900 mb-2">
                料金体系について教えてください
              </h3>
              <p className="text-gray-600">
                プロジェクトの内容と期間に応じて見積もりをいたします。
                時間単価制または固定価格制のいずれかで対応可能です。
                まずはお気軽にご相談ください。
              </p>
            </div>
            
            <div>
              <h3 className="text-lg font-semibold text-gray-900 mb-2">
                リモートワークは可能ですか？
              </h3>
              <p className="text-gray-600">
                はい、フルリモートでの作業が可能です。
                定期的なオンラインミーティングでプロジェクトの進捗を共有し、
                効率的なコミュニケーションを心がけています。
              </p>
            </div>
          </div>
        </div>
      </section>
    </Layout>
  )
}
