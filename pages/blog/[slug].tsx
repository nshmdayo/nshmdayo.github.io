import { GetStaticPaths, GetStaticProps } from 'next'
import Layout from '@/components/Layout'
import { getAllPosts, getPostBySlug } from '@/lib/blog'
import { Calendar, ArrowLeft } from 'lucide-react'
import Link from 'next/link'
import { MDXRemote, MDXRemoteSerializeResult } from 'next-mdx-remote'
import { serialize } from 'next-mdx-remote/serialize'

interface PostProps {
  post: {
    slug: string
    title: string
    date: string
    tags: string[]
    content: string
    mdxSource: MDXRemoteSerializeResult
  }
}

export default function BlogPost({ post }: PostProps) {
  return (
    <Layout>
      <article className="py-20 bg-white">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
          {/* Back Link */}
          <Link
            href="/blog"
            className="inline-flex items-center text-blue-600 hover:text-blue-800 mb-8 transition-colors duration-200"
          >
            <ArrowLeft size={16} className="mr-2" />
            ブログ一覧に戻る
          </Link>

          {/* Article Header */}
          <header className="mb-12">
            <h1 className="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
              {post.title}
            </h1>
            
            <div className="flex items-center text-gray-600 mb-6">
              <Calendar size={16} className="mr-2" />
              <span>{new Date(post.date).toLocaleDateString('ja-JP')}</span>
            </div>
            
            <div className="flex flex-wrap gap-2">
              {post.tags.map((tag) => (
                <span
                  key={tag}
                  className="px-3 py-1 bg-blue-100 text-blue-800 text-sm rounded-full"
                >
                  {tag}
                </span>
              ))}
            </div>
          </header>

          {/* Article Content */}
          <div className="prose prose-lg max-w-none">
            <MDXRemote {...post.mdxSource} />
          </div>
        </div>
      </article>
    </Layout>
  )
}

export const getStaticPaths: GetStaticPaths = async () => {
  const posts = getAllPosts()
  const paths = posts.map((post: any) => ({
    params: { slug: post.slug },
  }))

  return {
    paths,
    fallback: false,
  }
}

export const getStaticProps: GetStaticProps = async ({ params }) => {
  const post = getPostBySlug(params!.slug as string)
  const mdxSource = await serialize(post.content)

  return {
    props: {
      post: {
        ...post,
        mdxSource,
      },
    },
  }
}
