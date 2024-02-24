import { useState, ChangeEvent } from "react";
import StarImg from './assets/star.svg';

const NewRecipeModal = () => {

    const [titleValue, setTitleValue] = useState<string>('');
    const [descriptionValue, setDescriptionValue] = useState<string>('');
    const [ingredientsValue, setIngredientsValue] = useState<string>('');

    const handleTitleChange = (e: ChangeEvent<HTMLInputElement>) => {
        setTitleValue(e.target.value);
    }

    const handleDescriptionChange = (e: ChangeEvent<HTMLInputElement>) => {
        setDescriptionValue(e.target.value);
    }

    const handleIngredientsChange = (e: ChangeEvent<HTMLInputElement>) => {
        setIngredientsValue(e.target.value);
    }

    return <div className="recipe-modal-container">
        <div className="text-green-500 text-xl">New recipe</div>
        <div className="flex">
            <div>
                <div>
                    <label htmlFor="title">Title:</label>
                    <input
                        type="text"
                        id="title"
                        value={titleValue}
                        onChange={handleTitleChange}
                    />
                </div>
                <div>
                    <label htmlFor="description">Description:</label>
                    <input
                        type="text"
                        id="description"
                        value={descriptionValue}
                        onChange={handleDescriptionChange}
                    />
                </div>
                <div>
                    <label htmlFor="ingredients">Ingredients:</label>
                    <input
                        type="text"
                        id="ingredients"
                        value={ingredientsValue}
                        onChange={handleIngredientsChange}
                    />
                </div>
            </div>
            <div>Picture:</div>
        </div>
        <div>
            <div>
                <label htmlFor="ingredients">Preparation time:</label>
                <input
                    type="text"
                    id="preparation-h"
                // value={preparationHValue}
                // onChange={handlePreparationHChange}
                />
                Hours
                <input
                    type="text"
                    id="preparation-h"
                // value={preparationHValue}
                // onChange={handlePreparationMChange}
                />
                Mins
            </div>
            <div>
                <label htmlFor="ingredients">Cooking time:</label>
                <input
                    type="text"
                    id="preparation-h"
                // value={cookingHValue}
                // onChange={handleCookingHChange}
                />
                Hours
                <input
                    type="text"
                    id="preparation-h"
                // value={cookingMValue}
                // onChange={handleCookingMChange}
                />
                Mins
            </div>
        </div>
        <div className="flex">
            Difficulty:
            <img src={StarImg} alt="star" />
        </div>
    </div>
}

export default NewRecipeModal;