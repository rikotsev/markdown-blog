import React, {useState} from "react";
import styles from './DesignArticle.module.css'
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";

const DesignArticle: React.FC = () => {

    const [markdownContent, setMarkdownContent] = useState<string>('')

    const handleTextareaChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setMarkdownContent(event.target.value)
    }

    return (
        <div className={styles['design-article']}>
            <div className={styles['markdown-input']}>
                <h2>Markdown Input</h2>
                <textarea value={markdownContent}
                          onChange={handleTextareaChange}
                          placeholder={"Type your markdown here"}
                          className={styles['textarea']}/>
            </div>
            <div className={styles['article-output']}>
                <h2>Article Output</h2>
                <div className={styles['markdown-body']}>
                    <Markdown remarkPlugins={[remarkGfm]}>{markdownContent}</Markdown>
                </div>
            </div>
        </div>
    )
}

export default DesignArticle