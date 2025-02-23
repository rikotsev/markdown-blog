import React, { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { useArticleApiCtx } from "../../services/ArticleApiContext";
import { Article } from "../../openapi";

const Category: React.FC = () => {
  const { category } = useParams<{ category: string }>();
  const { api } = useArticleApiCtx();
  const [articles, setArticles] = useState<Article[]>([]);
  const [categoryName, setCategoryName] = useState<string>("");

  useEffect(() => {
    api.articleList(category).then((result) => {
      setArticles(result.data.data);
      let includedItem = result.data.included[0];
      if (includedItem.entityType === "category") {
        setCategoryName(includedItem.name);
      }
    });
  }, []);

  return (
    <div className="main-content">
      <div className="container">
        <h1>{categoryName}</h1>
      </div>
      <ul className="article-list">
        {articles.map((article) => (
          <li key={article.id} className="article-item">
            <Link to={`articles/${article.urlId}`}>
              <h2>{article.title}</h2>
            </Link>
            <p>{article.description}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Category;
