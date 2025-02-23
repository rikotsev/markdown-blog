import React, { useEffect, useState } from "react";
import styles from "./ArticleManager.module.css";
import { useNavigate } from "react-router-dom";
import { Article } from "../../openapi";
import { useCategoryApiCtx } from "../../services/CategoryApiContext";
import { useArticleApiCtx } from "../../services/ArticleApiContext";

const ArticleManager: React.FC = () => {
  const { categories } = useCategoryApiCtx();
  const { api } = useArticleApiCtx();
  const [loading, setLoading] = useState(true);
  const [articles, setArticles] = useState<Article[]>([]);
  const [searchTerm, setSearchTerm] = useState<string>("");
  const [selectedCategory, setSelectedCategory] = useState<string>("All");
  const navigate = useNavigate();

  useEffect(() => {
    const fetchInitialData = async () => {
      try {
        let resp = await api.articleList();

        if (resp.status === 200) {
          setArticles(resp.data.data);
          return;
        }

        console.error("failed to retrieve articles", resp);
      } catch (err) {
        console.error(err);
      } finally {
        setLoading(false);
      }
    };
    if (loading) {
      fetchInitialData();
    }
  }, []);

  const handleCreateArticle = () => {
    navigate("/article/create");
  };

  const handleDeleteArticle = (id: string) => {
    alert("delete article");
  };

  const handleEditArticle = (id: string) => {
    navigate("/article/" + id);
  };

  const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
  };

  const handleCategoryChange = (
    event: React.ChangeEvent<HTMLSelectElement>,
  ) => {
    setSelectedCategory(event.target.value);
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div className={styles["article-manager"]}>
      <h2>Article Manager</h2>
      <div className={styles.filters}>
        <input
          type="text"
          placeholder="Search article by title"
          onChange={handleSearchChange}
          value={searchTerm}
        />
        <select
          value={selectedCategory}
          onChange={handleCategoryChange}
          className={styles.categoryDropdown}
        >
          <option value="all">All</option>
          {categories.map((category) => (
            <option key={category.id} value={category.urlId}>
              {category.name}
            </option>
          ))}
        </select>
      </div>
      <button onClick={handleCreateArticle}>Add new article</button>
      <ul>
        {articles.map((article) => (
          <li key={article.id}>
            <div>
              <h3>{article.title}</h3>
              <p>{article.description}</p>
            </div>
            <div>
              <button
                onClick={() => {
                  handleEditArticle(article.urlId);
                }}
              >
                Edit
              </button>
              <button onClick={() => handleDeleteArticle(article.urlId)}>
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
