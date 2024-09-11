import React, {useState} from "react";
import styles from './CategoryManager.module.css'
import ArticleApi, {Category} from "../../services/ArticleApi";

const CategoryManager: React.FC = () => {

    const existingCategories = ArticleApi.getInstance().getCategories()
    const [categories, setCategories] = useState<Set<Category>>(new Set(existingCategories));
    const [newCategory, setNewCategory] = useState<string>('');

    const handleAddCategory = () => {
        if (newCategory.trim() !== '') {
            setCategories((prevCategories) => new Set(prevCategories).add({
                id: newCategory,
                prettyId: newCategory,
                title: newCategory
            }))
            setNewCategory('')
        }
    };

    const handleRemoveCategory = (categoryToRemove: string) => {
        setCategories((prevCategories) => {
            const updatedCategories = new Set<Category>()
            prevCategories.forEach(cat => {
                if (cat.id !== categoryToRemove) {
                    updatedCategories.add(cat)
                }
            })
            return updatedCategories
        })
    };

    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setNewCategory(event.target.value)
    }

    const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
        if (event.key === 'Enter') {
            event.preventDefault();
            handleAddCategory();
        }
    };

    return (
        <div className={styles['category-manager']}>
            <h2>Category Manager</h2>
            <input type="text"
                   value={newCategory}
                   onChange={handleInputChange}
                   onKeyDown={handleKeyDown}
                   placeholder="Add a new category" />
            <button onClick={handleAddCategory}>Add category</button>
            <ul>
                {Array.from(categories).map((category) => (
                    <li key={category.id}>
                        <div className={styles['category-text']}>
                            <div className={styles['category-title']}>{category.title}</div>
                            <div className={styles['category-subtitle']}>{category.prettyId}</div>
                        </div>
                        <button onClick={() => handleRemoveCategory(category.id)}>Remove</button>
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default CategoryManager