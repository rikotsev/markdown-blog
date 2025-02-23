import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { useArticleApiCtx } from "../../services/ArticleApiContext";
import { Article as ArticleModel } from "../../openapi";

const Article: React.FC = () => {
  const { api } = useArticleApiCtx();
  const { id } = useParams<{ id: string }>();
  const [article, setArticle] = useState<ArticleModel | undefined>(undefined);

  useEffect(() => {
    api
      .articleGet(id!)
      .then((resp) => {
        setArticle(resp.data.data);
      })
      .catch((err) => {
        console.error(err);
      });
  }, [id]);

  if (!article) {
    return <div>Loading...</div>;
  }

  return (
    <div className="main-content">
      <div className="container">
        <Markdown remarkPlugins={[remarkGfm]}>{article.content}</Markdown>
      </div>
    </div>
  );
};

export default Article;
