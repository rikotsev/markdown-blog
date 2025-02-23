import React, { useState } from "react";
import styles from "./CategoryManager.module.css";
import { useCategoryApiCtx } from "../../services/CategoryApiContext";

const CategoryManager: React.FC = () => {
  const { categories, add, remove } = useCategoryApiCtx();
  const [newCategory, setNewCategory] = useState<string>("");

  const handleAddCategory = async () => {
    if (newCategory.trim() === "") {
      return;
    }

    add({
      name: newCategory,
    }).then(() => {
      setNewCategory("");
    });
  };

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setNewCategory(event.target.value);
  };

  const handleKeyDown = async (
    event: React.KeyboardEvent<HTMLInputElement>,
  ) => {
    if (event.key === "Enter") {
      event.preventDefault();
      await handleAddCategory();
    }
  };

  return (
    <div className={styles["category-manager"]}>
      <h2>Category Manager</h2>
      <input
        type="text"
        value={newCategory}
        onChange={handleInputChange}
        onKeyDown={handleKeyDown}
        placeholder="Add a new category"
      />
      <button onClick={handleAddCategory}>Add category</button>
      <ul>
        {Array.from(categories).map((category) => (
          <li key={category.urlId}>
            <div className={styles["category-text"]}>
              <div className={styles["category-title"]}>{category.name}</div>
              <div className={styles["category-subtitle"]}>
                {category.urlId}
              </div>
            </div>
            <button onClick={() => remove(category.urlId)}>Remove</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default CategoryManager;
