import React, {useEffect, useState} from "react";
import {useParams} from "react-router-dom";
import {usePageApiCtx} from "../../services/PageApiContext";
import {Page as PageModel} from "../../openapi";
import remarkGfm from "remark-gfm";
import Markdown from "react-markdown";

export const Page: React.FC = () => {
    const {id} = useParams<{ id: string }>();
    const {api} = usePageApiCtx()
    const [page, setPage] = useState<PageModel | undefined>(undefined)


    useEffect(() => {
        if (id) {
            api.pageGet(id)
                .then((result) => {
                    setPage(result.data.data)
                })
                .catch((err) => {
                    console.error('failed to get page', err)
                })
        }
    }, [api, id])

    if (!page) {
        return (
            <div className="main-content">
                <div className="container">
                    <h1>Hi there!</h1>
                    <p>Welcome to Markdown Blog</p>
                </div>
            </div>
        )
    }

    return (
        <div className="main-content">
            <div className="container">
                <Markdown remarkPlugins={[remarkGfm]}>{page.content}</Markdown>
            </div>
        </div>
    )
}