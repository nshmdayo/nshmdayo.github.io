*, *::before, *::after {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

:root {
    --bg: #0e0e0e;
    --text: #d4d4d4;
    --text-dim: #666;
    --text-dimmer: #444;
    --link: #d4d4d4;
    --link-hover: #fff;
    --border: #222;
    --accent: #888;
}

html {
    font-size: 16px;
}

body {
    background: var(--bg);
    color: var(--text);
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif;
    line-height: 1.6;
    -webkit-font-smoothing: antialiased;
}

.page {
    max-width: 680px;
    margin: 0 auto;
    padding: 0 24px;
}

/* Nav */
header {
    padding: 40px 0 0;
}

nav {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.nav-home {
    font-weight: 600;
    color: #fff;
    text-decoration: none;
    font-size: 0.95rem;
}

.nav-links {
    display: flex;
    gap: 1.5rem;
}

.nav-links a {
    font-size: 0.875rem;
    color: var(--text-dim);
    text-decoration: none;
    transition: color 0.15s;
}

.nav-links a:hover {
    color: var(--text);
}

/* Tabs */
.tabs {
    display: flex;
    gap: 0;
    margin-top: 32px;
    border-bottom: 1px solid var(--border);
}

.tab-btn {
    background: none;
    border: none;
    border-bottom: 2px solid transparent;
    margin-bottom: -1px;
    padding: 0.5rem 1rem;
    font-size: 0.875rem;
    color: var(--text-dim);
    cursor: pointer;
    transition: color 0.15s, border-color 0.15s;
}

.tab-btn:first-child {
    padding-left: 0;
}

.tab-btn:hover {
    color: var(--text);
}

.tab-btn.active {
    color: #fff;
    border-bottom-color: #fff;
}

/* Tab panels */
.tab-panel {
    display: none;
}

.tab-panel.active {
    display: block;
}

.empty-state {
    font-size: 0.875rem;
    color: var(--text-dimmer);
    padding: 1.5rem 0;
}

/* Main */
main {
    padding: 48px 0 80px;
    display: flex;
    flex-direction: column;
    gap: 60px;
}

/* Intro */
.intro h1 {
    font-size: 1.4rem;
    font-weight: 600;
    color: #fff;
    margin-bottom: 0.5rem;
    letter-spacing: -0.01em;
}

.intro p {
    font-size: 0.95rem;
    color: var(--text-dim);
    max-width: 520px;
}

.intro .about {
    margin-top: 1rem;
    color: var(--accent);
    line-height: 1.7;
}

/* Section headings */
section h2 {
    font-size: 0.75rem;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.1em;
    color: var(--text-dimmer);
    margin-bottom: 1.25rem;
}

/* Skills */
.skills {
    display: flex;
    flex-direction: column;
    gap: 0.6rem;
}

.skill-row {
    display: flex;
    gap: 1.5rem;
    font-size: 0.875rem;
}

.skill-category {
    color: var(--text-dim);
    min-width: 100px;
    flex-shrink: 0;
}

.skill-items {
    color: var(--text);
}

/* Projects */
.project-list {
    list-style: none;
    display: flex;
    flex-direction: column;
    gap: 0;
}

.project-item {
    padding: 1rem 0;
    border-top: 1px solid var(--border);
}

.project-item:last-child {
    border-bottom: 1px solid var(--border);
}

.project-title-row {
    display: flex;
    align-items: baseline;
    justify-content: space-between;
    gap: 1rem;
    margin-bottom: 0.35rem;
}

.project-name {
    font-size: 0.95rem;
    font-weight: 500;
    color: #fff;
}

.project-links {
    display: flex;
    gap: 1rem;
    flex-shrink: 0;
}

.project-links a {
    font-size: 0.8rem;
    color: var(--text-dim);
    text-decoration: none;
    transition: color 0.15s;
}

.project-links a:hover {
    color: var(--text);
}

.project-desc {
    font-size: 0.875rem;
    color: var(--text-dim);
    margin-bottom: 0.4rem;
    line-height: 1.6;
}

.project-tech {
    font-size: 0.8rem;
    color: var(--text-dimmer);
}

/* Writing / Posts */
.post-list {
    list-style: none;
    display: flex;
    flex-direction: column;
    gap: 0;
}

.post-item {
    display: flex;
    align-items: baseline;
    justify-content: space-between;
    gap: 1rem;
    padding: 0.6rem 0;
    border-top: 1px solid var(--border);
}

.post-item:last-child {
    border-bottom: 1px solid var(--border);
}

.post-item a {
    font-size: 0.9rem;
    color: var(--text);
    text-decoration: none;
    transition: color 0.15s;
}

.post-item a:hover {
    color: #fff;
}

.post-date {
    font-size: 0.8rem;
    color: var(--text-dimmer);
    white-space: nowrap;
    flex-shrink: 0;
}

/* Footer */
footer {
    padding: 32px 0 40px;
    border-top: 1px solid var(--border);
}

footer p {
    font-size: 0.8rem;
    color: var(--text-dimmer);
}

/* Blog post page */
.post-header {
    padding: 72px 0 0;
}

.post-back {
    display: inline-block;
    font-size: 0.875rem;
    color: var(--text-dim);
    text-decoration: none;
    margin-bottom: 2.5rem;
    transition: color 0.15s;
}

.post-back:hover {
    color: var(--text);
}

.post-header h1 {
    font-size: 1.6rem;
    font-weight: 600;
    color: #fff;
    letter-spacing: -0.02em;
    line-height: 1.3;
    margin-bottom: 0.5rem;
}

.post-header .meta {
    font-size: 0.85rem;
    color: var(--text-dimmer);
    margin-bottom: 3rem;
}

.post-body {
    padding: 0 0 80px;
    font-size: 0.95rem;
    color: var(--text);
    line-height: 1.8;
}

.post-body h2,
.post-body h3 {
    color: #fff;
    font-weight: 600;
    margin: 2.5rem 0 0.75rem;
    letter-spacing: -0.01em;
}

.post-body h2 { font-size: 1.2rem; }
.post-body h3 { font-size: 1rem; }

.post-body p {
    margin-bottom: 1.2rem;
    color: var(--text);
}

.post-body a {
    color: var(--text);
    text-decoration: underline;
    text-underline-offset: 3px;
    text-decoration-color: var(--text-dimmer);
    transition: text-decoration-color 0.15s;
}

.post-body a:hover {
    text-decoration-color: var(--text);
}

.post-body code {
    font-family: 'SFMono-Regular', Consolas, 'Liberation Mono', Menlo, monospace;
    font-size: 0.85em;
    background: #1a1a1a;
    color: #c9d1d9;
    padding: 0.15em 0.4em;
    border-radius: 3px;
}

.post-body pre {
    background: #111;
    border: 1px solid var(--border);
    border-radius: 4px;
    padding: 1.25rem;
    overflow-x: auto;
    margin: 1.5rem 0;
}

.post-body pre code {
    background: none;
    padding: 0;
    font-size: 0.875rem;
}

footer {
    margin-top: 6rem;
    padding: 2rem 0;
    border-top: 1px solid #e2e8f0;
    color: #64748b;
    font-size: 0.875rem;
}

@media (max-width: 768px) {
    .hero h2 {
        font-size: 1.8rem;
    }
    
    nav ul {
        gap: 1rem;

    }

    .skill-category {
        min-width: unset;
    }

    .project-title-row {
        flex-direction: column;
        gap: 0.2rem;
    }
}
