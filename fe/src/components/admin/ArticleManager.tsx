import React, {useState} from "react";
import styles from './ArticleManager.module.css'
import {useNavigate} from "react-router-dom";


interface DummyArticle {
    id: string
    title: string
    content: string
}

const initialArticles: DummyArticle[] = [
    { id: 'article_1', title: 'Understanding React Components', content: 'Detailed content about React components...' },
    { id: 'article_2', title: 'JavaScript ES6 Features', content: 'Content about ES6 features in JavaScript...' },
    { id: 'article_3', title: 'TypeScript Basics', content: 'Content about getting started with TypeScript...' },
];

const categories = ['All', 'React', 'JavaScript', 'TypeScript'];

const ArticleManager: React.FC = () => {
    const [articles, setArticles] = useState<DummyArticle[]>(initialArticles)
    const [searchTerm, setSearchTerm] = useState<string>('')
    const [selectedCategory, setSelectedCategory] = useState<string>('All');
    const navigate = useNavigate()

    const handleCreateArticle = () => {
        navigate("/article/create")
    }

    const handleDeleteArticle = (id: string) => {
        alert('delete article')
    }

    const handleEditArticle = (id: string) => {
        navigate("/article/edit/" + id)
    }

    const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setSearchTerm(event.target.value)
    }

    const handleCategoryChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedCategory(event.target.value);
    };

    return (
        <div className={styles['article-manager']}>
            <h2>Article Manager</h2>
            <div className={styles.filters}>
                <input type="text"
                       placeholder="Search article by title"
                       onChange={handleSearchChange}
                       value={searchTerm} />
                <select value={selectedCategory} onChange={handleCategoryChange} className={styles.categoryDropdown}>
                    {categories.map((category) => (
                        <option key={category} value={category}>
                            {category}
                        </option>
                    ))}
                </select>
            </div>
            <button onClick={handleCreateArticle}>
                Add new article
            </button>
            <ul>
                {articles.map((article) => (
                    <li key={article.id}>
                        <div>
                            <h3>{article.title}</h3>
                            <p>{article.content}</p>
                        </div>
                        <div>
                            <button onClick={() => {handleEditArticle(article.id)}}>
                                Edit
                            </button>
                            <button onClick={() => handleDeleteArticle(article.id)}>
                                Delete
                            </button>
                        </div>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default ArticleManager;