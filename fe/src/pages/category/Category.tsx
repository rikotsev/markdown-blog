import React from "react";
import {Link, useParams} from "react-router-dom";

const Category: React.FC = () => {

    const {category} = useParams<{category: string}>();

    const articles: Article[] = [
        {
            id: "article_1",
            title: "Article 1",
            description: "This is an example article."
        },
        {
            id: "article_2",
            title: "Article 2",
            description: "This is another example article. Please read it."
        },
        {
            id: "article_3",
            title: "Article 3",
            description: "The final example article. Please read it."
        }
    ]

    return (
        <div className="main-content">
            <div className="container">
                <h1>{category}</h1>
            </div>
            <ul className="article-list">
                {articles.map((article) => (
                    <li key={article.id} className="article-item">
                        <Link to={`articles/${article.id}`}>
                            <h2>{article.title}</h2>
                        </Link>
                        <p>{article.description}</p>
                    </li>
                ))}
            </ul>
        </div>
    );
}

interface Article {
    id: string
    title: string
    description: string
}

export default Category