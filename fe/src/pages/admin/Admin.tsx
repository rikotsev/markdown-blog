import React from "react";
import CategoryManager from "../../components/admin/CategoryManager";
import styles from './Admin.module.css'
import ArticleManager from "../../components/admin/ArticleManager";


const Admin: React.FC = () => {
    return (
        <div className={styles.adminContainer}>
            <div className={styles.categorySection}>
                <CategoryManager/>
            </div>
            <div className={styles.otherSection}>
                <ArticleManager />
            </div>
        </div>
    )
}

export default Admin;