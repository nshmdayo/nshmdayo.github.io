import Layout from '@/components/Layout'
import Link from 'next/link'
import { ExternalLink, Github } from 'lucide-react'

export default function Portfolio() {
  const projects = [
    {
      id: 1,
      title: 'E-Commerce Platform',
      description: 'Next.js と Stripe を使用したフルスタック EC サイト。レスポンシブデザインと高いパフォーマンスを実現。',
      image: '/projects/ecommerce.jpg',
      technologies: ['Next.js', 'TypeScript', 'Tailwind CSS', 'Stripe', 'Prisma'],
      github: 'https://github.com/nshmdayo/ecommerce-platform',
      demo: 'https://ecommerce-demo.example.com',
      featured: true
    },
    {
      id: 2,
      title: 'Task Management App',
      description: 'React と Firebase を使用したリアルタイムタスク管理アプリケーション。チーム作業に最適化されています。',
      image: '/projects/taskapp.jpg',
      technologies: ['React', 'Firebase', 'Material-UI', 'Redux'],
      github: 'https://github.com/nshmdayo/task-management',
      demo: 'https://taskapp-demo.example.com',
      featured: true
    },
    {
      id: 3,
      title: 'Weather Dashboard',
      description: 'OpenWeather API を使用した天気予報ダッシュボード。データビジュアライゼーションに注力。',
      image: '/projects/weather.jpg',
      technologies: ['Vue.js', 'Chart.js', 'OpenWeather API', 'Vuetify'],
      github: 'https://github.com/nshmdayo/weather-dashboard',
      demo: 'https://weather-demo.example.com',
      featured: false
    },
    {
      id: 4,
      title: 'Blog Platform',
      description: 'MDX を使用したブログプラットフォーム。管理者機能とコメントシステムを実装。',
      image: '/projects/blog.jpg',
      technologies: ['Next.js', 'MDX', 'MongoDB', 'NextAuth.js'],
      github: 'https://github.com/nshmdayo/blog-platform',
      demo: 'https://blog-demo.example.com',
      featured: false
    }
  ]

  const featuredProjects = projects.filter(project => project.featured)
  const otherProjects = projects.filter(project => !project.featured)

  return (
    <Layout>
      {/* Hero Section */}
      <section className="bg-gradient-to-r from-purple-600 to-blue-600 text-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h1 className="text-4xl md:text-6xl font-bold mb-6">
            ポートフォリオ
          </h1>
          <p className="text-xl md:text-2xl text-purple-100">
            これまでに制作したプロジェクトをご紹介します
          </p>
        </div>
      </section>

      {/* Featured Projects */}
      <section className="py-20 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-16 text-center">
            主要プロジェクト
          </h2>
          
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-12">
            {featuredProjects.map((project) => (
              <div key={project.id} className="bg-white rounded-xl shadow-lg overflow-hidden">
                <div className="h-64 bg-gray-200 relative">
                  {/* Placeholder for project image */}
                  <div className="absolute inset-0 flex items-center justify-center text-gray-500">
                    <span>プロジェクト画像</span>
                  </div>
                </div>
                
                <div className="p-8">
                  <h3 className="text-2xl font-bold text-gray-900 mb-4">
                    {project.title}
                  </h3>
                  <p className="text-gray-600 mb-6">
                    {project.description}
                  </p>
                  
                  <div className="flex flex-wrap gap-2 mb-6">
                    {project.technologies.map((tech) => (
                      <span
                        key={tech}
                        className="px-3 py-1 bg-blue-100 text-blue-800 text-sm rounded-full"
                      >
                        {tech}
                      </span>
                    ))}
                  </div>
                  
                  <div className="flex space-x-4">
                    <a
                      href={project.github}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="flex items-center text-gray-600 hover:text-gray-900 transition-colors"
                    >
                      <Github size={20} className="mr-2" />
                      コード
                    </a>
                    <a
                      href={project.demo}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="flex items-center text-blue-600 hover:text-blue-800 transition-colors"
                    >
                      <ExternalLink size={20} className="mr-2" />
                      デモ
                    </a>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Other Projects */}
      <section className="py-20 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-16 text-center">
            その他のプロジェクト
          </h2>
          
          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            {otherProjects.map((project) => (
              <div key={project.id} className="bg-white rounded-lg shadow-md overflow-hidden">
                <div className="h-48 bg-gray-200 relative">
                  {/* Placeholder for project image */}
                  <div className="absolute inset-0 flex items-center justify-center text-gray-500">
                    <span>プロジェクト画像</span>
                  </div>
                </div>
                
                <div className="p-6">
                  <h3 className="text-xl font-bold text-gray-900 mb-3">
                    {project.title}
                  </h3>
                  <p className="text-gray-600 mb-4 text-sm">
                    {project.description}
                  </p>
                  
                  <div className="flex flex-wrap gap-2 mb-4">
                    {project.technologies.map((tech) => (
                      <span
                        key={tech}
                        className="px-2 py-1 bg-gray-100 text-gray-700 text-xs rounded-full"
                      >
                        {tech}
                      </span>
                    ))}
                  </div>
                  
                  <div className="flex space-x-4">
                    <a
                      href={project.github}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="flex items-center text-gray-600 hover:text-gray-900 transition-colors text-sm"
                    >
                      <Github size={16} className="mr-1" />
                      コード
                    </a>
                    <a
                      href={project.demo}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="flex items-center text-blue-600 hover:text-blue-800 transition-colors text-sm"
                    >
                      <ExternalLink size={16} className="mr-1" />
                      デモ
                    </a>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Skills Section */}
      <section className="py-20 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-16 text-center">
            スキル・技術
          </h2>
          
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div className="text-center">
              <h3 className="text-xl font-semibold mb-4 text-gray-900">フロントエンド</h3>
              <div className="space-y-2">
                {['React', 'Next.js', 'TypeScript', 'Vue.js', 'Tailwind CSS', 'SCSS'].map((skill) => (
                  <span key={skill} className="block text-gray-600">{skill}</span>
                ))}
              </div>
            </div>
            
            <div className="text-center">
              <h3 className="text-xl font-semibold mb-4 text-gray-900">バックエンド</h3>
              <div className="space-y-2">
                {['Node.js', 'Express', 'Python', 'Django', 'PostgreSQL', 'MongoDB'].map((skill) => (
                  <span key={skill} className="block text-gray-600">{skill}</span>
                ))}
              </div>
            </div>
            
            <div className="text-center">
              <h3 className="text-xl font-semibold mb-4 text-gray-900">ツール・その他</h3>
              <div className="space-y-2">
                {['Git', 'Docker', 'AWS', 'Figma', 'Jest', 'Cypress'].map((skill) => (
                  <span key={skill} className="block text-gray-600">{skill}</span>
                ))}
              </div>
            </div>
          </div>
        </div>
      </section>
    </Layout>
  )
}
