import React, { useState } from 'react';

const blogPosts = [
	{
		id: 1,
		title: 'Getting Started with Go Initializer',
		content: `Go Initializer helps you scaffold Go projects with best practices and modern frameworks. Select your project type, Go version, and framework to generate a ready-to-use starter kit.`,
	},
	{
		id: 2,
		title: 'Microservices Best Practices',
		content: `Learn how to structure your Go microservices for scalability and maintainability. Explore recommended folder structures, dependency management, and testing strategies.`,
	},
	{
		id: 3,
		title: 'Choosing a Go Web Framework',
		content: `Compare popular Go web frameworks like Gin, Echo, and Fiber. Understand their strengths, community support, and when to use each for your project.`,
	},
	{
		id: 4,
		title: 'Community Resources',
		content: `Find curated links to Go tutorials, open-source projects, and community forums to accelerate your learning and development.`,
	},
];

const Explore: React.FC<{ onBack: () => void }> = ({ onBack }) => {
	const [selectedId, setSelectedId] = useState(blogPosts[0].id);
	const selectedPost = blogPosts.find((post) => post.id === selectedId);

	return (
		<div style={{ width: '100%', padding: '2rem 0' }}>
			<div style={{ maxWidth: 1200, width: '100%', margin: '0 auto', display: 'flex', gap: 32 }}>
				{/* Sidebar: Titles */}
				<aside style={{ width: 300, minWidth: 220, background: 'var(--content-bg)', borderRadius: 16, boxShadow: '0 2px 12px 0 rgba(0,0,0,0.04)', padding: '2rem 0', display: 'flex', flexDirection: 'column', gap: 0, height: 'fit-content' }}>
					<h2 style={{ fontSize: 22, fontWeight: 800, margin: '0 0 1.5rem 2rem', color: 'var(--text)' }}>Explore</h2>
					{blogPosts.map((post) => (
						<button
							key={post.id}
							onClick={() => setSelectedId(post.id)}
							style={{
								background: selectedId === post.id ? '#ffd70022' : 'none',
								color: selectedId === post.id ? '#ffd700' : 'var(--text)',
								border: 'none',
								borderLeft: selectedId === post.id ? '4px solid #ffd700' : '4px solid transparent',
								textAlign: 'left',
								padding: '1.1rem 2rem',
								fontWeight: selectedId === post.id ? 700 : 500,
								fontSize: 18,
								cursor: 'pointer',
								outline: 'none',
								transition: 'background 0.2s, color 0.2s',
								width: '100%',
								borderRadius: 0,
							}}
						>
							{post.title}
						</button>
					))}
				</aside>
				{/* Main Content: Selected Post */}
				   <main style={{ flex: 1, height: '100%', background: 'var(--card-bg)', borderRadius: 16, boxShadow: '0 2px 12px 0 rgba(0,0,0,0.08)', padding: '2.5rem 2.5rem 2.5rem 4.5rem', display: 'flex', flexDirection: 'column', justifyContent: 'flex-start', minHeight: 400, position: 'relative' }}>
					   <h1 style={{ fontSize: 32, fontWeight: 800, marginBottom: 8, marginTop: 0 }}>{selectedPost?.title}</h1>
					   <div style={{ fontSize: 18, lineHeight: 1.7, color: 'var(--text)' }}>{selectedPost?.content}</div>
					   <div style={{ marginTop: 'auto', textAlign: 'center' }}>
						   <button onClick={onBack} style={{ color: '#ffd700', fontWeight: 700, fontSize: 17, textDecoration: 'none', border: '2px solid #ffd700', borderRadius: 8, padding: '0.6rem 1.7rem', background: 'var(--card-bg)', transition: 'background 0.2s, color 0.2s', cursor: 'pointer' }}>
							   ← Back to Generator
						   </button>
					   </div>
				   </main>
			</div>
		</div>
	);
};

export default Explore;
