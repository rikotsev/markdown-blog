import React from "react";
import { usePageApiCtx } from "../../services/PageApiContext";
import styles from "./ArticleManager.module.css";
import { useNavigate } from "react-router-dom";

const PageManager: React.FC = () => {
  const { pages, remove } = usePageApiCtx();
  const navigate = useNavigate();

  const handleCreatePage = () => {
    navigate("/page/create");
  };

  const handleDeletePage = async (id: string | undefined) => {
    await remove(id!)
      .then(() => window.alert("page removed"))
      .catch((err) => {
        console.error(err);
      });
  };

  const handleEditPage = (id: string | undefined) => {
    navigate("/page/design/" + id);
  };

  return (
    <div className={styles["page-manager"]}>
      <h2>Page Manager</h2>
      <button onClick={handleCreatePage}>Add new page</button>
      <ul>
        {pages.map((page) => (
          <li key={page.urlId}>
            <div>
              <h3>{page.title}</h3>
            </div>
            <div>
              <button
                onClick={() => {
                  handleEditPage(page.urlId);
                }}
              >
                Edit
              </button>
              <button onClick={() => handleDeletePage(page.urlId)}>
                Delete
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default PageManager;
