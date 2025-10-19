* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
    line-height: 1.6;
    color: #333;
    background: #f8f9fa;
}

.container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
}

nav {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

nav h1 {
    font-size: 1.5rem;
    color: #2563eb;
}

nav ul {
    display: flex;
    list-style: none;
    gap: 2rem;
}

nav a {
    color: #333;
    text-decoration: none;
    transition: color 0.3s;
}

nav a:hover {
    color: #2563eb;
}

.hero {
    text-align: center;
    padding: 4rem 0;
    margin-bottom: 3rem;
}

.hero h2 {
    font-size: 2.5rem;
    margin-bottom: 1rem;
}

.hero p {
    font-size: 1.2rem;
    opacity: 0.9;
}

section {
    margin-bottom: 4rem;
}

h2 {
    font-size: 2rem;
    margin-bottom: 1.5rem;
    color: #1e293b;
}

.grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 2rem;
}

.card {
    background: white;
    border-radius: 8px;
    padding: 1.5rem;
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
    transition: transform 0.3s, box-shadow 0.3s;
}

.card:hover {
    transform: translateY(-4px);
    box-shadow: 0 4px 12px rgba(0,0,0,0.15);
}

.card h3 {
    color: #2563eb;
    margin-bottom: 0.5rem;
}

.card .date {
    color: #64748b;
    font-size: 0.9rem;
    margin-bottom: 0.5rem;
}

.card p {
    color: #475569;
    margin-bottom: 1rem;
}

.tags {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
    margin-top: 1rem;
}

.tag {
    background: #e0e7ff;
    color: #3730a3;
    padding: 0.25rem 0.75rem;
    border-radius: 9999px;
    font-size: 0.875rem;
}

.btn {
    display: inline-block;
    padding: 0.5rem 1rem;
    background: #2563eb;
    color: white;
    text-decoration: none;
    border-radius: 4px;
    transition: background 0.3s;
}

.btn:hover {
    background: #1d4ed8;
}

.social-links {
    display: flex;
    gap: 1rem;
    justify-content: center;
    margin-top: 2rem;
}

.social-links a {
    color: #64748b;
    transition: color 0.3s;
}

.social-links a:hover {
    color: #2563eb;
}

article {
    background: white;
    padding: 3rem;
    border-radius: 8px;
    max-width: 800px;
    margin: 0 auto;
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

article h1 {
    color: #1e293b;
    margin-bottom: 0.5rem;
}

article img {
    max-width: 100%;
    height: auto;
    border-radius: 4px;
    margin: 1rem 0;
}

article pre {
    background: #1e293b;
    color: #e2e8f0;
    padding: 1rem;
    border-radius: 4px;
    overflow-x: auto;
    margin: 1rem 0;
}

article code {
    background: #f1f5f9;
    padding: 0.2rem 0.4rem;
    border-radius: 4px;
    font-family: 'Courier New', monospace;
}

article pre code {
    background: none;
    padding: 0;
}

@media (max-width: 768px) {
    .hero h2 {
        font-size: 1.8rem;
    }
    
    nav ul {
        gap: 1rem;
    }
    
    .grid {
        grid-template-columns: 1fr;
    }
    
    article {
        padding: 1.5rem;
    }
}