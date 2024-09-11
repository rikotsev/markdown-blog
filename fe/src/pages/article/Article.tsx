import React from "react";
import {useParams} from "react-router-dom";
import ArticleApi from "../../services/ArticleApi";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";


const Article: React.FC = () => {

    const {id} = useParams<{id: string}>();
    const article = ArticleApi.getInstance().getArticle(id!)

    return (
        <div className="main-content">
            <div className="container">
                <Markdown remarkPlugins={[remarkGfm]}>{article.content}</Markdown>
            </div>
        </div>
    )
}

export default Article;