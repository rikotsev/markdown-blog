import React from "react";
import {Link, useParams} from "react-router-dom";
import ArticleApi from "../../services/ArticleApi";

const Category: React.FC = () => {

    const {category} = useParams<{category: string}>();
    const articles = ArticleApi.getInstance().getArticles(category!)

    return (
        <div className="main-content">
            <div className="container">
                <h1>{category}</h1>
            </div>
            <ul className="article-list">
                {articles.map((article) => (
                    <li key={article.id} className="article-item">
                        <Link to={`articles/${article.prettyId}`}>
                            <h2>{article.title}</h2>
                        </Link>
                        <p>{article.description}</p>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default Category