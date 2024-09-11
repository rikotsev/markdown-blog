import React, {useState} from "react";
import styles from './ArticleManager.module.css'
import {useNavigate} from "react-router-dom";
import ArticleApi, {Article} from "../../services/ArticleApi";


const ArticleManager: React.FC = () => {
    const initialArticles = ArticleApi.getInstance().getAllArticles()
    const initialCategories = ArticleApi.getInstance().getCategories()
    const [articles, setArticles] = useState<Article[]>(initialArticles)
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
                    <option value="all">All</option>
                    {initialCategories.map((category) => (
                        <option key={category.id} value={category.prettyId}>
                            {category.title}
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
                            <p>{article.description}</p>
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