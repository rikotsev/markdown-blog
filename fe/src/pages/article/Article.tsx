import React from "react";


const Article: React.FC = () => {

    const item: ArticleItem = {
        id: "article_1",
        title: "Exciting tech news",
        description: "This is an exciting update, please see it.",
        content: "This is a very long and elaborate article."
    }

    return (
        <div className="main-content">
            <div className="container">
                <h1>{item.title}</h1>
                <p>{item.content}</p>
            </div>
        </div>
    )
}

interface ArticleItem {
    id: string
    title: string
    description: string
    content: string
}

export default Article;