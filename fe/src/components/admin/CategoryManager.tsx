import React, {useState} from "react";
import styles from './CategoryManager.module.css'

const CategoryManager: React.FC = () => {
    const [categories, setCategories] = useState<Set<string>>(new Set());
    const [newCategory, setNewCategory] = useState<string>('');

    const handleAddCategory = () => {
        if (newCategory.trim() !== '') {
            setCategories((prevCategories) => new Set(prevCategories).add(newCategory.trim()))
            setNewCategory('')
        }
    };

    const handleRemoveCategory = (categoryToRemove: string) => {
        setCategories((prevCategories) => {
            const updatedCategories = new Set(prevCategories);
            updatedCategories.delete(categoryToRemove);
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
                    <li key={category}>
                        {category}
                        <button onClick={() => handleRemoveCategory(category)}>Remove</button>
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default CategoryManager