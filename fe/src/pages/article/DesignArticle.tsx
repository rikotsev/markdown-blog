import React, {useEffect, useState} from "react";
import styles from './DesignArticle.module.css'
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import {useCategoryApiCtx} from "../../services/CategoryApiContext";
import {useArticleApiCtx} from "../../services/ArticleApiContext";
import {useParams} from "react-router-dom";
import {ArticleCore} from "../../openapi";

const DesignArticle: React.FC = () => {
    const {categories} = useCategoryApiCtx();
    const {api} = useArticleApiCtx();
    const {id} = useParams<{ id: string }>();
    const [articleData, setArticleData] = useState<ArticleCore>({
        content: '',
        title: '',
        description: '',
        category: {
            entityType: 'category',
            id: ''
        }
    })


    useEffect(() => {
        if (id) {
            api.articleGet(id)
                .then((response) => {
                    setArticleData({
                        content: response.data.data!.content,
                        title: response.data.data!.title,
                        description: response.data.data!.description,
                        category: {
                            entityType: 'category',
                            id: response.data.data!.category.id
                        }
                    })
                })
                .catch((err) => {
                    console.error(err);
                })
        }
    }, []);

    const handleContentChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setArticleData((prev) => ({
            ...prev,
            ['content']: event.target.value
        }))
    }
    const handleTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setArticleData((prev) => ({
            ...prev,
            ['title']: event.target.value
        }))
    }
    const handleDescriptionChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setArticleData((prev) => ({
            ...prev,
            ['description']: event.target.value
        }))
    }

    function handleCategoryChange(event: React.ChangeEvent<HTMLSelectElement>) {
        setArticleData((prev) => ({
            ...prev,
            category: {
                ...prev.category,
                entityType: 'category',
                id: event.target.value
            }
        }))
    }

    const addArticle = async () => {
        api.articleCreate({
            title: articleData.title!,
            content: articleData.content!,
            description: articleData.description!,
            category: {
                entityType: 'category',
                id: articleData.category!.id
            }
        }).catch((err) => {
            console.error(err)
        })
    }


    const editArticle = async () => {
        api.articleEdit(id!, articleData)
            .catch((err) => {
                console.error(err)
            })

    }

    let actionButton
    if (id) {
        actionButton = <button className={styles['create-article-btn']} onClick={editArticle}>Save Changes</button>
    } else {
        actionButton = <button className={styles['create-article-btn']} onClick={addArticle}>Create Article</button>
    }

    return (
        <div className={styles['article-form']}>
            <div className={styles['top-controls']}>
                <input type="text" onChange={handleTitleChange} value={articleData.title}
                       className={styles['title-input']} placeholder="Article name"/>
                <select className={styles['category-select']} onChange={handleCategoryChange}
                        value={articleData.category?.id}>
                    <option value="">Select Category</option>
                    {categories.map((category) => (
                        <option key={category.id} value={category.id}>
                            {category.name}
                        </option>
                    ))}
                </select>
                {actionButton}
            </div>
            <textarea
                className={styles['short-description']}
                placeholder="Short description of the article..."
                onChange={handleDescriptionChange}
                value={articleData.description}
            />
            <div className={styles['design-article']}>
                <div className={styles['markdown-input']}>
                    <h2>Markdown Input</h2>
                    <textarea value={articleData.content}
                              onChange={handleContentChange}
                              placeholder={"Type your markdown here"}
                              className={styles['textarea']}/>
                </div>
                <div className={styles['article-output']}>
                    <h2>Article Output</h2>
                    <div className={styles['markdown-body']}>
                        <Markdown remarkPlugins={[remarkGfm]}>{articleData.content}</Markdown>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default DesignArticle