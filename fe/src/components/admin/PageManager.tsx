import React, {useEffect, useState} from "react";
import {usePageApiCtx} from "../../services/PageApiContext";
import {PageUrlIdAndTitle} from "../../openapi";

const PageManager: React.FC = () => {
    const {api} = usePageApiCtx();
    const [loading, setLoading] = useState(true)
    const [pages, setPages] = useState<PageUrlIdAndTitle[]>([])

    useEffect(() => {
        const fetchInitialData = async () => {
            try {
                let pageResponseList = await api.pageList()

                if (pageResponseList.status === 200) {
                    setPages(pageResponseList.data.data)
                    return;
                }

                console.error("failed to retrieve pages", pageResponseList)
            } catch (err) {
                console.error(err)
            } finally {
                setLoading(false)
            }
        }

        if (loading) {
            fetchInitialData()
        }
    }, [api, loading])

    return (
        <div></div>
    )
}