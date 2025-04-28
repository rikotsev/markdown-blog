import React, { useEffect, useState } from "react";
import { useLocation, useParams } from "react-router-dom";
import { usePageApiCtx } from "../../services/PageApiContext";
import { Page as PageModel } from "../../openapi";
import remarkGfm from "remark-gfm";
import Markdown from "react-markdown";

export const Page: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const location = useLocation();
  const { api } = usePageApiCtx();
  const [page, setPage] = useState<PageModel | undefined>(undefined);

  useEffect(() => {
    if (id) {
      api
        .pageGet(id)
        .then((result) => {
          setPage(result.data.data);
        })
        .catch((err) => {
          console.error("failed to get page", err);
        });
    }
  }, [api, id]);

  useEffect(() => {
    if (location.pathname === "/") {
      api
        .pageList()
        .then((pages) => {
          api
            .pageGet(pages.data.data[0].urlId!)
            .then((result) => {
              setPage(result.data.data);
            })
            .catch((err) => {
              console.error("failed to get default page", err);
            });
        })
        .catch((err) => {
          console.error("failed to get page list", err);
        });
    }
  }, [api, location]);

  if (!page) {
    return (
      <div className="main-content">
        <div className="container">
          <h1>Hi there!</h1>
          <p>Welcome to Markdown Blog</p>
        </div>
      </div>
    );
  }

  return (
    <div className="main-content">
      <div className="container">
        <Markdown remarkPlugins={[remarkGfm]}>{page.content}</Markdown>
      </div>
    </div>
  );
};
