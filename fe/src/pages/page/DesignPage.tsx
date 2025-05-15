import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { usePageApiCtx } from "../../services/PageApiContext";
import styles from "../article/DesignArticle.module.css";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { PageCore } from "../../openapi";

const DesignPage: React.FC = () => {
  const { api, add, edit } = usePageApiCtx();
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  const [pageData, setPageData] = useState<PageCore>({
    content: "",
    title: "",
    position: 0,
  });

  useEffect(() => {
    if (id) {
      api
        .pageGet(id)
        .then((response) => {
          setPageData({
            content: response.data.data!.content,
            title: response.data.data!.title,
            position: response.data.data!.position,
          });
        })
        .catch((err) => {
          console.error(err);
        });
    }
  }, [id, api]);

  const handleContentChange = (
    event: React.ChangeEvent<HTMLTextAreaElement>,
  ) => {
    setPageData((prev) => ({
      ...prev,
      content: event.target.value,
    }));
  };
  const handleTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setPageData((prev) => ({
      ...prev,
      title: event.target.value,
    }));
  };
  const handlePositionChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setPageData((prev) => ({
      ...prev,
      position: Number(event.target.value),
    }));
  };

  const addPage = async () => {
    await add({
      title: pageData.title!,
      content: pageData.content!,
      position: pageData.position!,
    }).then((newId) => {
      if (newId) {
        navigate("/page/" + newId);
      }
    });
  };

  const editPage = async () => {
    await edit(id!, pageData).then((modifiedId) => {
      if (modifiedId) {
        navigate("/page/" + modifiedId);
      }
    });
  };

  let actionButton;
  if (id) {
    actionButton = (
      <button className={styles["create-page-btn"]} onClick={editPage}>
        Save Changes
      </button>
    );
  } else {
    actionButton = (
      <button className={styles["create-page-btn"]} onClick={addPage}>
        Create Page
      </button>
    );
  }

  return (
    <div className="main-content">
      <div className="container">
        <div className={styles["page-form"]}>
          <div className={styles["top-controls"]}>
            <input
              type="text"
              onChange={handleTitleChange}
              value={pageData.title}
              className={styles["title-input"]}
              placeholder="Page name"
            />
            <input
              type="number"
              onChange={handlePositionChange}
              value={pageData.position}
              className={styles["position-input"]}
              placeholder="Page position"
            />
            {actionButton}
          </div>
          <div className={styles["design-page"]}>
            <div className={styles["markdown-input"]}>
              <h2>Markdown Input</h2>
              <textarea
                value={pageData.content}
                onChange={handleContentChange}
                placeholder={"Type your markdown here"}
                className={styles["textarea"]}
              />
            </div>
            <div className={styles["page-output"]}>
              <h2>Page Output</h2>
              <div className={styles["markdown-body"]}>
                <Markdown remarkPlugins={[remarkGfm]}>
                  {pageData.content}
                </Markdown>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default DesignPage;
