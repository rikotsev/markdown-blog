import React from "react";
import CategoryManager from "../../components/admin/CategoryManager";
import styles from "./Admin.module.css";
import ArticleManager from "../../components/admin/ArticleManager";
import { useAuth0 } from "@auth0/auth0-react";
import PageManager from "../../components/admin/PageManager";

const Admin: React.FC = () => {
  const { isAuthenticated } = useAuth0();

  if (!isAuthenticated) {
    return (
      <div className={styles.adminContainer}>
        <span>You should not be here you naughty boy!</span>
      </div>
    );
  }

  return (
    <div className={styles.adminContainer}>
      <div className={styles.categorySection}>
        <CategoryManager />
      </div>
      <div className={styles.otherSection}>
        <ArticleManager />
        <PageManager />
      </div>
    </div>
  );
};

export default Admin;
